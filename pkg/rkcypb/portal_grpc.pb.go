// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: portal.proto

package rkcypb

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

// PortalServiceClient is the client API for PortalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalServiceClient interface {
	PlatformDef(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PlatformDef, error)
	ConfigRead(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ConfigReadResponse, error)
	DecodeInstance(ctx context.Context, in *DecodeInstanceArgs, opts ...grpc.CallOption) (*DecodeResponse, error)
	DecodeArgPayload(ctx context.Context, in *DecodePayloadArgs, opts ...grpc.CallOption) (*DecodeResponse, error)
	DecodeResultPayload(ctx context.Context, in *DecodePayloadArgs, opts ...grpc.CallOption) (*DecodeResponse, error)
	CancelTxn(ctx context.Context, in *CancelTxnRequest, opts ...grpc.CallOption) (*Void, error)
}

type portalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalServiceClient(cc grpc.ClientConnInterface) PortalServiceClient {
	return &portalServiceClient{cc}
}

func (c *portalServiceClient) PlatformDef(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PlatformDef, error) {
	out := new(PlatformDef)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/PlatformDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalServiceClient) ConfigRead(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ConfigReadResponse, error) {
	out := new(ConfigReadResponse)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/ConfigRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalServiceClient) DecodeInstance(ctx context.Context, in *DecodeInstanceArgs, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/DecodeInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalServiceClient) DecodeArgPayload(ctx context.Context, in *DecodePayloadArgs, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/DecodeArgPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalServiceClient) DecodeResultPayload(ctx context.Context, in *DecodePayloadArgs, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/DecodeResultPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalServiceClient) CancelTxn(ctx context.Context, in *CancelTxnRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rkcy.PortalService/CancelTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalServiceServer is the server API for PortalService service.
// All implementations must embed UnimplementedPortalServiceServer
// for forward compatibility
type PortalServiceServer interface {
	PlatformDef(context.Context, *Void) (*PlatformDef, error)
	ConfigRead(context.Context, *Void) (*ConfigReadResponse, error)
	DecodeInstance(context.Context, *DecodeInstanceArgs) (*DecodeResponse, error)
	DecodeArgPayload(context.Context, *DecodePayloadArgs) (*DecodeResponse, error)
	DecodeResultPayload(context.Context, *DecodePayloadArgs) (*DecodeResponse, error)
	CancelTxn(context.Context, *CancelTxnRequest) (*Void, error)
	mustEmbedUnimplementedPortalServiceServer()
}

// UnimplementedPortalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortalServiceServer struct {
}

func (UnimplementedPortalServiceServer) PlatformDef(context.Context, *Void) (*PlatformDef, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlatformDef not implemented")
}
func (UnimplementedPortalServiceServer) ConfigRead(context.Context, *Void) (*ConfigReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigRead not implemented")
}
func (UnimplementedPortalServiceServer) DecodeInstance(context.Context, *DecodeInstanceArgs) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeInstance not implemented")
}
func (UnimplementedPortalServiceServer) DecodeArgPayload(context.Context, *DecodePayloadArgs) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeArgPayload not implemented")
}
func (UnimplementedPortalServiceServer) DecodeResultPayload(context.Context, *DecodePayloadArgs) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeResultPayload not implemented")
}
func (UnimplementedPortalServiceServer) CancelTxn(context.Context, *CancelTxnRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTxn not implemented")
}
func (UnimplementedPortalServiceServer) mustEmbedUnimplementedPortalServiceServer() {}

// UnsafePortalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalServiceServer will
// result in compilation errors.
type UnsafePortalServiceServer interface {
	mustEmbedUnimplementedPortalServiceServer()
}

func RegisterPortalServiceServer(s grpc.ServiceRegistrar, srv PortalServiceServer) {
	s.RegisterService(&PortalService_ServiceDesc, srv)
}

func _PortalService_PlatformDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).PlatformDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/PlatformDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).PlatformDef(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalService_ConfigRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).ConfigRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/ConfigRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).ConfigRead(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalService_DecodeInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeInstanceArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).DecodeInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/DecodeInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).DecodeInstance(ctx, req.(*DecodeInstanceArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalService_DecodeArgPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodePayloadArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).DecodeArgPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/DecodeArgPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).DecodeArgPayload(ctx, req.(*DecodePayloadArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalService_DecodeResultPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodePayloadArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).DecodeResultPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/DecodeResultPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).DecodeResultPayload(ctx, req.(*DecodePayloadArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalService_CancelTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServiceServer).CancelTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rkcy.PortalService/CancelTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServiceServer).CancelTxn(ctx, req.(*CancelTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortalService_ServiceDesc is the grpc.ServiceDesc for PortalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rkcy.PortalService",
	HandlerType: (*PortalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlatformDef",
			Handler:    _PortalService_PlatformDef_Handler,
		},
		{
			MethodName: "ConfigRead",
			Handler:    _PortalService_ConfigRead_Handler,
		},
		{
			MethodName: "DecodeInstance",
			Handler:    _PortalService_DecodeInstance_Handler,
		},
		{
			MethodName: "DecodeArgPayload",
			Handler:    _PortalService_DecodeArgPayload_Handler,
		},
		{
			MethodName: "DecodeResultPayload",
			Handler:    _PortalService_DecodeResultPayload_Handler,
		},
		{
			MethodName: "CancelTxn",
			Handler:    _PortalService_CancelTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal.proto",
}
