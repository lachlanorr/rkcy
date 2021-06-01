// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package edge

import (
	context "context"
	concerns "github.com/lachlanorr/rocketcycle/examples/rpg/concerns"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RpgServiceClient is the client API for RpgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpgServiceClient interface {
	ReadPlayer(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*concerns.Player, error)
	CreatePlayer(ctx context.Context, in *concerns.Player, opts ...grpc.CallOption) (*concerns.Player, error)
	UpdatePlayer(ctx context.Context, in *concerns.Player, opts ...grpc.CallOption) (*concerns.Player, error)
	DeletePlayer(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*RpgResponse, error)
	ReadCharacter(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*concerns.Character, error)
	CreateCharacter(ctx context.Context, in *concerns.Character, opts ...grpc.CallOption) (*concerns.Character, error)
	UpdateCharacter(ctx context.Context, in *concerns.Character, opts ...grpc.CallOption) (*concerns.Character, error)
	DeleteCharacter(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*RpgResponse, error)
	FundCharacter(ctx context.Context, in *concerns.FundingRequest, opts ...grpc.CallOption) (*concerns.Character, error)
}

type rpgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpgServiceClient(cc grpc.ClientConnInterface) RpgServiceClient {
	return &rpgServiceClient{cc}
}

func (c *rpgServiceClient) ReadPlayer(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*concerns.Player, error) {
	out := new(concerns.Player)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/ReadPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) CreatePlayer(ctx context.Context, in *concerns.Player, opts ...grpc.CallOption) (*concerns.Player, error) {
	out := new(concerns.Player)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) UpdatePlayer(ctx context.Context, in *concerns.Player, opts ...grpc.CallOption) (*concerns.Player, error) {
	out := new(concerns.Player)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/UpdatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) DeletePlayer(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*RpgResponse, error) {
	out := new(RpgResponse)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/DeletePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) ReadCharacter(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*concerns.Character, error) {
	out := new(concerns.Character)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/ReadCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) CreateCharacter(ctx context.Context, in *concerns.Character, opts ...grpc.CallOption) (*concerns.Character, error) {
	out := new(concerns.Character)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/CreateCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) UpdateCharacter(ctx context.Context, in *concerns.Character, opts ...grpc.CallOption) (*concerns.Character, error) {
	out := new(concerns.Character)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/UpdateCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) DeleteCharacter(ctx context.Context, in *RpgRequest, opts ...grpc.CallOption) (*RpgResponse, error) {
	out := new(RpgResponse)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/DeleteCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpgServiceClient) FundCharacter(ctx context.Context, in *concerns.FundingRequest, opts ...grpc.CallOption) (*concerns.Character, error) {
	out := new(concerns.Character)
	err := c.cc.Invoke(ctx, "/rocketcycle.examples.rpg.edge.RpgService/FundCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpgServiceServer is the server API for RpgService service.
// All implementations must embed UnimplementedRpgServiceServer
// for forward compatibility
type RpgServiceServer interface {
	ReadPlayer(context.Context, *RpgRequest) (*concerns.Player, error)
	CreatePlayer(context.Context, *concerns.Player) (*concerns.Player, error)
	UpdatePlayer(context.Context, *concerns.Player) (*concerns.Player, error)
	DeletePlayer(context.Context, *RpgRequest) (*RpgResponse, error)
	ReadCharacter(context.Context, *RpgRequest) (*concerns.Character, error)
	CreateCharacter(context.Context, *concerns.Character) (*concerns.Character, error)
	UpdateCharacter(context.Context, *concerns.Character) (*concerns.Character, error)
	DeleteCharacter(context.Context, *RpgRequest) (*RpgResponse, error)
	FundCharacter(context.Context, *concerns.FundingRequest) (*concerns.Character, error)
	mustEmbedUnimplementedRpgServiceServer()
}

// UnimplementedRpgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRpgServiceServer struct {
}

func (UnimplementedRpgServiceServer) ReadPlayer(context.Context, *RpgRequest) (*concerns.Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPlayer not implemented")
}
func (UnimplementedRpgServiceServer) CreatePlayer(context.Context, *concerns.Player) (*concerns.Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}
func (UnimplementedRpgServiceServer) UpdatePlayer(context.Context, *concerns.Player) (*concerns.Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayer not implemented")
}
func (UnimplementedRpgServiceServer) DeletePlayer(context.Context, *RpgRequest) (*RpgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlayer not implemented")
}
func (UnimplementedRpgServiceServer) ReadCharacter(context.Context, *RpgRequest) (*concerns.Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCharacter not implemented")
}
func (UnimplementedRpgServiceServer) CreateCharacter(context.Context, *concerns.Character) (*concerns.Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacter not implemented")
}
func (UnimplementedRpgServiceServer) UpdateCharacter(context.Context, *concerns.Character) (*concerns.Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCharacter not implemented")
}
func (UnimplementedRpgServiceServer) DeleteCharacter(context.Context, *RpgRequest) (*RpgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCharacter not implemented")
}
func (UnimplementedRpgServiceServer) FundCharacter(context.Context, *concerns.FundingRequest) (*concerns.Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCharacter not implemented")
}
func (UnimplementedRpgServiceServer) mustEmbedUnimplementedRpgServiceServer() {}

// UnsafeRpgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpgServiceServer will
// result in compilation errors.
type UnsafeRpgServiceServer interface {
	mustEmbedUnimplementedRpgServiceServer()
}

func RegisterRpgServiceServer(s grpc.ServiceRegistrar, srv RpgServiceServer) {
	s.RegisterService(&_RpgService_serviceDesc, srv)
}

func _RpgService_ReadPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).ReadPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/ReadPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).ReadPlayer(ctx, req.(*RpgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(concerns.Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).CreatePlayer(ctx, req.(*concerns.Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_UpdatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(concerns.Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).UpdatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/UpdatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).UpdatePlayer(ctx, req.(*concerns.Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_DeletePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).DeletePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/DeletePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).DeletePlayer(ctx, req.(*RpgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_ReadCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).ReadCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/ReadCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).ReadCharacter(ctx, req.(*RpgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_CreateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(concerns.Character)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).CreateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/CreateCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).CreateCharacter(ctx, req.(*concerns.Character))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_UpdateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(concerns.Character)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).UpdateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/UpdateCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).UpdateCharacter(ctx, req.(*concerns.Character))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_DeleteCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).DeleteCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/DeleteCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).DeleteCharacter(ctx, req.(*RpgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpgService_FundCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(concerns.FundingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpgServiceServer).FundCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocketcycle.examples.rpg.edge.RpgService/FundCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpgServiceServer).FundCharacter(ctx, req.(*concerns.FundingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rocketcycle.examples.rpg.edge.RpgService",
	HandlerType: (*RpgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadPlayer",
			Handler:    _RpgService_ReadPlayer_Handler,
		},
		{
			MethodName: "CreatePlayer",
			Handler:    _RpgService_CreatePlayer_Handler,
		},
		{
			MethodName: "UpdatePlayer",
			Handler:    _RpgService_UpdatePlayer_Handler,
		},
		{
			MethodName: "DeletePlayer",
			Handler:    _RpgService_DeletePlayer_Handler,
		},
		{
			MethodName: "ReadCharacter",
			Handler:    _RpgService_ReadCharacter_Handler,
		},
		{
			MethodName: "CreateCharacter",
			Handler:    _RpgService_CreateCharacter_Handler,
		},
		{
			MethodName: "UpdateCharacter",
			Handler:    _RpgService_UpdateCharacter_Handler,
		},
		{
			MethodName: "DeleteCharacter",
			Handler:    _RpgService_DeleteCharacter_Handler,
		},
		{
			MethodName: "FundCharacter",
			Handler:    _RpgService_FundCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edge.proto",
}
