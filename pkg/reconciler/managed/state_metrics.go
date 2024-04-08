/*
Copyright 2024 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package managed

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
)

// A StateRecorder records the state of given GroupVersionKind.
type StateRecorder interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)

	Record(ctx context.Context, gvk schema.GroupVersionKind)
	Run(ctx context.Context, gvk schema.GroupVersionKind)
}

// A MRStateRecorder records the state of managed resources.
type MRStateRecorder struct {
	client    client.Client
	log       logging.Logger
	frequency time.Duration

	mrExists *prometheus.GaugeVec
	mrReady  *prometheus.GaugeVec
	mrSynced *prometheus.GaugeVec
}

// NewMRStateRecorder returns a new MRStateRecorder which records the state of managed resources.
func NewMRStateRecorder(client client.Client, log logging.Logger, o ...StateRecorderOption) *MRStateRecorder {
	r := &MRStateRecorder{
		client: client,
		log:    log,

		mrExists: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Subsystem: subSystem,
			Name:      "managed_resource_exists",
			Help:      "The number of managed resources that exist",
		}, []string{"gvk"}),
		mrReady: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Subsystem: subSystem,
			Name:      "managed_resource_ready",
			Help:      "The number of managed resources in Ready=True state",
		}, []string{"gvk"}),
		mrSynced: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Subsystem: subSystem,
			Name:      "managed_resource_synced",
			Help:      "The number of managed resources in Synced=True state",
		}, []string{"gvk"}),
	}

	for _, ro := range o {
		ro(r)
	}

	return r
}

// A StateRecorderOption configures a MRStateRecorder.
type StateRecorderOption func(*MRStateRecorder)

// WithFrequency configures the frequency at which the MRStateRecorder
// will record.
func WithFrequency(f time.Duration) StateRecorderOption {
	return func(r *MRStateRecorder) {
		r.frequency = f
	}
}

// Describe sends the super-set of all possible descriptors of metrics
// collected by this Collector to the provided channel and returns once
// the last descriptor has been sent.
func (r *MRStateRecorder) Describe(ch chan<- *prometheus.Desc) {
	r.mrExists.Describe(ch)
	r.mrReady.Describe(ch)
	r.mrSynced.Describe(ch)
}

// Collect is called by the Prometheus registry when collecting
// metrics. The implementation sends each collected metric via the
// provided channel and returns once the last metric has been sent.
func (r *MRStateRecorder) Collect(ch chan<- prometheus.Metric) {
	r.mrExists.Collect(ch)
	r.mrReady.Collect(ch)
	r.mrSynced.Collect(ch)
}

// Record records the state of managed resources.
func (r *MRStateRecorder) Record(ctx context.Context, gvk schema.GroupVersionKind) {
	l := &unstructured.UnstructuredList{}
	l.SetGroupVersionKind(gvk)
	err := r.client.List(ctx, l)
	if err != nil {
		r.log.Info("Failed to list managed resources", "error", err)
		return
	}

	label := gvk.String()
	r.mrExists.WithLabelValues(label).Set(float64(len(l.Items)))

	var numReady, numSynced float64 = 0, 0
	for _, o := range l.Items {
		conditioned := xpv1.ConditionedStatus{}
		err := fieldpath.Pave(o.Object).GetValueInto("status", &conditioned)
		if err != nil {
			r.log.Info("Failed to get conditions of managed resource", "error", err)
			continue
		}

		for _, condition := range conditioned.Conditions {
			if condition.Type == xpv1.TypeReady && condition.Status == corev1.ConditionTrue {
				numReady++
			} else if condition.Type == xpv1.TypeSynced && condition.Status == corev1.ConditionTrue {
				numSynced++
			}
		}
	}

	r.mrReady.WithLabelValues(label).Set(numReady)
	r.mrSynced.WithLabelValues(label).Set(numSynced)
}

// Run records state of managed resources with given frequency.
func (r *MRStateRecorder) Run(ctx context.Context, gvk schema.GroupVersionKind) {
	ticker := time.NewTicker(r.frequency)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				r.Record(ctx, gvk)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// A NopStateRecorder does nothing.
type NopStateRecorder struct{}

// NewNopStateRecorder returns a NopStateRecorder that does nothing.
func NewNopStateRecorder() *NopStateRecorder {
	return &NopStateRecorder{}
}

// Describe does nothing.
func (r *NopStateRecorder) Describe(_ chan<- *prometheus.Desc) {}

// Collect does nothing.
func (r *NopStateRecorder) Collect(_ chan<- prometheus.Metric) {}

// Record does nothing.
func (r *NopStateRecorder) Record(_ context.Context, _ schema.GroupVersionKind) {}

// Run does nothing.
func (r *NopStateRecorder) Run(_ context.Context, _ schema.GroupVersionKind) {}
