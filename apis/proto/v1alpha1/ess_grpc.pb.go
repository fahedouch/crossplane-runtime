//
//Copyright 2023 The Crossplane Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/v1alpha1/ess.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExternalSecretStorePluginService_GetSecret_FullMethodName   = "/ess.proto.v1alpha1.ExternalSecretStorePluginService/GetSecret"
	ExternalSecretStorePluginService_ApplySecret_FullMethodName = "/ess.proto.v1alpha1.ExternalSecretStorePluginService/ApplySecret"
	ExternalSecretStorePluginService_DeleteKeys_FullMethodName  = "/ess.proto.v1alpha1.ExternalSecretStorePluginService/DeleteKeys"
)

// ExternalSecretStorePluginServiceClient is the client API for ExternalSecretStorePluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalSecretStorePluginServiceClient interface {
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	ApplySecret(ctx context.Context, in *ApplySecretRequest, opts ...grpc.CallOption) (*ApplySecretResponse, error)
	DeleteKeys(ctx context.Context, in *DeleteKeysRequest, opts ...grpc.CallOption) (*DeleteKeysResponse, error)
}

type externalSecretStorePluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalSecretStorePluginServiceClient(cc grpc.ClientConnInterface) ExternalSecretStorePluginServiceClient {
	return &externalSecretStorePluginServiceClient{cc}
}

func (c *externalSecretStorePluginServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, ExternalSecretStorePluginService_GetSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalSecretStorePluginServiceClient) ApplySecret(ctx context.Context, in *ApplySecretRequest, opts ...grpc.CallOption) (*ApplySecretResponse, error) {
	out := new(ApplySecretResponse)
	err := c.cc.Invoke(ctx, ExternalSecretStorePluginService_ApplySecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalSecretStorePluginServiceClient) DeleteKeys(ctx context.Context, in *DeleteKeysRequest, opts ...grpc.CallOption) (*DeleteKeysResponse, error) {
	out := new(DeleteKeysResponse)
	err := c.cc.Invoke(ctx, ExternalSecretStorePluginService_DeleteKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalSecretStorePluginServiceServer is the server API for ExternalSecretStorePluginService service.
// All implementations must embed UnimplementedExternalSecretStorePluginServiceServer
// for forward compatibility
type ExternalSecretStorePluginServiceServer interface {
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	ApplySecret(context.Context, *ApplySecretRequest) (*ApplySecretResponse, error)
	DeleteKeys(context.Context, *DeleteKeysRequest) (*DeleteKeysResponse, error)
	mustEmbedUnimplementedExternalSecretStorePluginServiceServer()
}

// UnimplementedExternalSecretStorePluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExternalSecretStorePluginServiceServer struct {
}

func (UnimplementedExternalSecretStorePluginServiceServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedExternalSecretStorePluginServiceServer) ApplySecret(context.Context, *ApplySecretRequest) (*ApplySecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySecret not implemented")
}
func (UnimplementedExternalSecretStorePluginServiceServer) DeleteKeys(context.Context, *DeleteKeysRequest) (*DeleteKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeys not implemented")
}
func (UnimplementedExternalSecretStorePluginServiceServer) mustEmbedUnimplementedExternalSecretStorePluginServiceServer() {
}

// UnsafeExternalSecretStorePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalSecretStorePluginServiceServer will
// result in compilation errors.
type UnsafeExternalSecretStorePluginServiceServer interface {
	mustEmbedUnimplementedExternalSecretStorePluginServiceServer()
}

func RegisterExternalSecretStorePluginServiceServer(s grpc.ServiceRegistrar, srv ExternalSecretStorePluginServiceServer) {
	s.RegisterService(&ExternalSecretStorePluginService_ServiceDesc, srv)
}

func _ExternalSecretStorePluginService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalSecretStorePluginServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalSecretStorePluginService_GetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalSecretStorePluginServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalSecretStorePluginService_ApplySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalSecretStorePluginServiceServer).ApplySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalSecretStorePluginService_ApplySecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalSecretStorePluginServiceServer).ApplySecret(ctx, req.(*ApplySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalSecretStorePluginService_DeleteKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalSecretStorePluginServiceServer).DeleteKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalSecretStorePluginService_DeleteKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalSecretStorePluginServiceServer).DeleteKeys(ctx, req.(*DeleteKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalSecretStorePluginService_ServiceDesc is the grpc.ServiceDesc for ExternalSecretStorePluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalSecretStorePluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ess.proto.v1alpha1.ExternalSecretStorePluginService",
	HandlerType: (*ExternalSecretStorePluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _ExternalSecretStorePluginService_GetSecret_Handler,
		},
		{
			MethodName: "ApplySecret",
			Handler:    _ExternalSecretStorePluginService_ApplySecret_Handler,
		},
		{
			MethodName: "DeleteKeys",
			Handler:    _ExternalSecretStorePluginService_DeleteKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/ess.proto",
}
