// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.2
// source: tracking_server.proto

package tracking_server

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

// TrackingServerServiceClient is the client API for TrackingServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackingServerServiceClient interface {
	// Set/update the current point tracking status.
	SetTrackingPointStatus(ctx context.Context, in *SetTrackingPointStatusRequest, opts ...grpc.CallOption) (*SetTrackingPointStatusResponse, error)
	// Set/update the current rectangle tracking status.
	SetTrackingRectangleStatus(ctx context.Context, in *SetTrackingRectangleStatusRequest, opts ...grpc.CallOption) (*SetTrackingRectangleStatusResponse, error)
	// Set the current tracking status to off.
	SetTrackingOffStatus(ctx context.Context, in *SetTrackingOffStatusRequest, opts ...grpc.CallOption) (*SetTrackingOffStatusResponse, error)
	// Subscribe to incoming tracking point command.
	SubscribeTrackingPointCommand(ctx context.Context, in *SubscribeTrackingPointCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingPointCommandClient, error)
	// Subscribe to incoming tracking rectangle command.
	SubscribeTrackingRectangleCommand(ctx context.Context, in *SubscribeTrackingRectangleCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingRectangleCommandClient, error)
	// Subscribe to incoming tracking off command.
	SubscribeTrackingOffCommand(ctx context.Context, in *SubscribeTrackingOffCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingOffCommandClient, error)
	// Respond to an incoming tracking point command.
	RespondTrackingPointCommand(ctx context.Context, in *RespondTrackingPointCommandRequest, opts ...grpc.CallOption) (*RespondTrackingPointCommandResponse, error)
	// Respond to an incoming tracking rectangle command.
	RespondTrackingRectangleCommand(ctx context.Context, in *RespondTrackingRectangleCommandRequest, opts ...grpc.CallOption) (*RespondTrackingRectangleCommandResponse, error)
	// Respond to an incoming tracking off command.
	RespondTrackingOffCommand(ctx context.Context, in *RespondTrackingOffCommandRequest, opts ...grpc.CallOption) (*RespondTrackingOffCommandResponse, error)
}

type trackingServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingServerServiceClient(cc grpc.ClientConnInterface) TrackingServerServiceClient {
	return &trackingServerServiceClient{cc}
}

func (c *trackingServerServiceClient) SetTrackingPointStatus(ctx context.Context, in *SetTrackingPointStatusRequest, opts ...grpc.CallOption) (*SetTrackingPointStatusResponse, error) {
	out := new(SetTrackingPointStatusResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingPointStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServerServiceClient) SetTrackingRectangleStatus(ctx context.Context, in *SetTrackingRectangleStatusRequest, opts ...grpc.CallOption) (*SetTrackingRectangleStatusResponse, error) {
	out := new(SetTrackingRectangleStatusResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingRectangleStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServerServiceClient) SetTrackingOffStatus(ctx context.Context, in *SetTrackingOffStatusRequest, opts ...grpc.CallOption) (*SetTrackingOffStatusResponse, error) {
	out := new(SetTrackingOffStatusResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingOffStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServerServiceClient) SubscribeTrackingPointCommand(ctx context.Context, in *SubscribeTrackingPointCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingPointCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrackingServerService_ServiceDesc.Streams[0], "/mavsdk.rpc.tracking_server.TrackingServerService/SubscribeTrackingPointCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackingServerServiceSubscribeTrackingPointCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrackingServerService_SubscribeTrackingPointCommandClient interface {
	Recv() (*TrackingPointCommandResponse, error)
	grpc.ClientStream
}

type trackingServerServiceSubscribeTrackingPointCommandClient struct {
	grpc.ClientStream
}

func (x *trackingServerServiceSubscribeTrackingPointCommandClient) Recv() (*TrackingPointCommandResponse, error) {
	m := new(TrackingPointCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trackingServerServiceClient) SubscribeTrackingRectangleCommand(ctx context.Context, in *SubscribeTrackingRectangleCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingRectangleCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrackingServerService_ServiceDesc.Streams[1], "/mavsdk.rpc.tracking_server.TrackingServerService/SubscribeTrackingRectangleCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackingServerServiceSubscribeTrackingRectangleCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrackingServerService_SubscribeTrackingRectangleCommandClient interface {
	Recv() (*TrackingRectangleCommandResponse, error)
	grpc.ClientStream
}

type trackingServerServiceSubscribeTrackingRectangleCommandClient struct {
	grpc.ClientStream
}

func (x *trackingServerServiceSubscribeTrackingRectangleCommandClient) Recv() (*TrackingRectangleCommandResponse, error) {
	m := new(TrackingRectangleCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trackingServerServiceClient) SubscribeTrackingOffCommand(ctx context.Context, in *SubscribeTrackingOffCommandRequest, opts ...grpc.CallOption) (TrackingServerService_SubscribeTrackingOffCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrackingServerService_ServiceDesc.Streams[2], "/mavsdk.rpc.tracking_server.TrackingServerService/SubscribeTrackingOffCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackingServerServiceSubscribeTrackingOffCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrackingServerService_SubscribeTrackingOffCommandClient interface {
	Recv() (*TrackingOffCommandResponse, error)
	grpc.ClientStream
}

type trackingServerServiceSubscribeTrackingOffCommandClient struct {
	grpc.ClientStream
}

func (x *trackingServerServiceSubscribeTrackingOffCommandClient) Recv() (*TrackingOffCommandResponse, error) {
	m := new(TrackingOffCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trackingServerServiceClient) RespondTrackingPointCommand(ctx context.Context, in *RespondTrackingPointCommandRequest, opts ...grpc.CallOption) (*RespondTrackingPointCommandResponse, error) {
	out := new(RespondTrackingPointCommandResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingPointCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServerServiceClient) RespondTrackingRectangleCommand(ctx context.Context, in *RespondTrackingRectangleCommandRequest, opts ...grpc.CallOption) (*RespondTrackingRectangleCommandResponse, error) {
	out := new(RespondTrackingRectangleCommandResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingRectangleCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServerServiceClient) RespondTrackingOffCommand(ctx context.Context, in *RespondTrackingOffCommandRequest, opts ...grpc.CallOption) (*RespondTrackingOffCommandResponse, error) {
	out := new(RespondTrackingOffCommandResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingOffCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingServerServiceServer is the server API for TrackingServerService service.
// All implementations must embed UnimplementedTrackingServerServiceServer
// for forward compatibility
type TrackingServerServiceServer interface {
	// Set/update the current point tracking status.
	SetTrackingPointStatus(context.Context, *SetTrackingPointStatusRequest) (*SetTrackingPointStatusResponse, error)
	// Set/update the current rectangle tracking status.
	SetTrackingRectangleStatus(context.Context, *SetTrackingRectangleStatusRequest) (*SetTrackingRectangleStatusResponse, error)
	// Set the current tracking status to off.
	SetTrackingOffStatus(context.Context, *SetTrackingOffStatusRequest) (*SetTrackingOffStatusResponse, error)
	// Subscribe to incoming tracking point command.
	SubscribeTrackingPointCommand(*SubscribeTrackingPointCommandRequest, TrackingServerService_SubscribeTrackingPointCommandServer) error
	// Subscribe to incoming tracking rectangle command.
	SubscribeTrackingRectangleCommand(*SubscribeTrackingRectangleCommandRequest, TrackingServerService_SubscribeTrackingRectangleCommandServer) error
	// Subscribe to incoming tracking off command.
	SubscribeTrackingOffCommand(*SubscribeTrackingOffCommandRequest, TrackingServerService_SubscribeTrackingOffCommandServer) error
	// Respond to an incoming tracking point command.
	RespondTrackingPointCommand(context.Context, *RespondTrackingPointCommandRequest) (*RespondTrackingPointCommandResponse, error)
	// Respond to an incoming tracking rectangle command.
	RespondTrackingRectangleCommand(context.Context, *RespondTrackingRectangleCommandRequest) (*RespondTrackingRectangleCommandResponse, error)
	// Respond to an incoming tracking off command.
	RespondTrackingOffCommand(context.Context, *RespondTrackingOffCommandRequest) (*RespondTrackingOffCommandResponse, error)
	mustEmbedUnimplementedTrackingServerServiceServer()
}

// UnimplementedTrackingServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrackingServerServiceServer struct {
}

func (UnimplementedTrackingServerServiceServer) SetTrackingPointStatus(context.Context, *SetTrackingPointStatusRequest) (*SetTrackingPointStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrackingPointStatus not implemented")
}
func (UnimplementedTrackingServerServiceServer) SetTrackingRectangleStatus(context.Context, *SetTrackingRectangleStatusRequest) (*SetTrackingRectangleStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrackingRectangleStatus not implemented")
}
func (UnimplementedTrackingServerServiceServer) SetTrackingOffStatus(context.Context, *SetTrackingOffStatusRequest) (*SetTrackingOffStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrackingOffStatus not implemented")
}
func (UnimplementedTrackingServerServiceServer) SubscribeTrackingPointCommand(*SubscribeTrackingPointCommandRequest, TrackingServerService_SubscribeTrackingPointCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTrackingPointCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) SubscribeTrackingRectangleCommand(*SubscribeTrackingRectangleCommandRequest, TrackingServerService_SubscribeTrackingRectangleCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTrackingRectangleCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) SubscribeTrackingOffCommand(*SubscribeTrackingOffCommandRequest, TrackingServerService_SubscribeTrackingOffCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTrackingOffCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) RespondTrackingPointCommand(context.Context, *RespondTrackingPointCommandRequest) (*RespondTrackingPointCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondTrackingPointCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) RespondTrackingRectangleCommand(context.Context, *RespondTrackingRectangleCommandRequest) (*RespondTrackingRectangleCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondTrackingRectangleCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) RespondTrackingOffCommand(context.Context, *RespondTrackingOffCommandRequest) (*RespondTrackingOffCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondTrackingOffCommand not implemented")
}
func (UnimplementedTrackingServerServiceServer) mustEmbedUnimplementedTrackingServerServiceServer() {}

// UnsafeTrackingServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackingServerServiceServer will
// result in compilation errors.
type UnsafeTrackingServerServiceServer interface {
	mustEmbedUnimplementedTrackingServerServiceServer()
}

func RegisterTrackingServerServiceServer(s grpc.ServiceRegistrar, srv TrackingServerServiceServer) {
	s.RegisterService(&TrackingServerService_ServiceDesc, srv)
}

func _TrackingServerService_SetTrackingPointStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTrackingPointStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).SetTrackingPointStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingPointStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).SetTrackingPointStatus(ctx, req.(*SetTrackingPointStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingServerService_SetTrackingRectangleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTrackingRectangleStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).SetTrackingRectangleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingRectangleStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).SetTrackingRectangleStatus(ctx, req.(*SetTrackingRectangleStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingServerService_SetTrackingOffStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTrackingOffStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).SetTrackingOffStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/SetTrackingOffStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).SetTrackingOffStatus(ctx, req.(*SetTrackingOffStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingServerService_SubscribeTrackingPointCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTrackingPointCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackingServerServiceServer).SubscribeTrackingPointCommand(m, &trackingServerServiceSubscribeTrackingPointCommandServer{stream})
}

type TrackingServerService_SubscribeTrackingPointCommandServer interface {
	Send(*TrackingPointCommandResponse) error
	grpc.ServerStream
}

type trackingServerServiceSubscribeTrackingPointCommandServer struct {
	grpc.ServerStream
}

func (x *trackingServerServiceSubscribeTrackingPointCommandServer) Send(m *TrackingPointCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TrackingServerService_SubscribeTrackingRectangleCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTrackingRectangleCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackingServerServiceServer).SubscribeTrackingRectangleCommand(m, &trackingServerServiceSubscribeTrackingRectangleCommandServer{stream})
}

type TrackingServerService_SubscribeTrackingRectangleCommandServer interface {
	Send(*TrackingRectangleCommandResponse) error
	grpc.ServerStream
}

type trackingServerServiceSubscribeTrackingRectangleCommandServer struct {
	grpc.ServerStream
}

func (x *trackingServerServiceSubscribeTrackingRectangleCommandServer) Send(m *TrackingRectangleCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TrackingServerService_SubscribeTrackingOffCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTrackingOffCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackingServerServiceServer).SubscribeTrackingOffCommand(m, &trackingServerServiceSubscribeTrackingOffCommandServer{stream})
}

type TrackingServerService_SubscribeTrackingOffCommandServer interface {
	Send(*TrackingOffCommandResponse) error
	grpc.ServerStream
}

type trackingServerServiceSubscribeTrackingOffCommandServer struct {
	grpc.ServerStream
}

func (x *trackingServerServiceSubscribeTrackingOffCommandServer) Send(m *TrackingOffCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TrackingServerService_RespondTrackingPointCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondTrackingPointCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).RespondTrackingPointCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingPointCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).RespondTrackingPointCommand(ctx, req.(*RespondTrackingPointCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingServerService_RespondTrackingRectangleCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondTrackingRectangleCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).RespondTrackingRectangleCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingRectangleCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).RespondTrackingRectangleCommand(ctx, req.(*RespondTrackingRectangleCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingServerService_RespondTrackingOffCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondTrackingOffCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServerServiceServer).RespondTrackingOffCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.tracking_server.TrackingServerService/RespondTrackingOffCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServerServiceServer).RespondTrackingOffCommand(ctx, req.(*RespondTrackingOffCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackingServerService_ServiceDesc is the grpc.ServiceDesc for TrackingServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackingServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.tracking_server.TrackingServerService",
	HandlerType: (*TrackingServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTrackingPointStatus",
			Handler:    _TrackingServerService_SetTrackingPointStatus_Handler,
		},
		{
			MethodName: "SetTrackingRectangleStatus",
			Handler:    _TrackingServerService_SetTrackingRectangleStatus_Handler,
		},
		{
			MethodName: "SetTrackingOffStatus",
			Handler:    _TrackingServerService_SetTrackingOffStatus_Handler,
		},
		{
			MethodName: "RespondTrackingPointCommand",
			Handler:    _TrackingServerService_RespondTrackingPointCommand_Handler,
		},
		{
			MethodName: "RespondTrackingRectangleCommand",
			Handler:    _TrackingServerService_RespondTrackingRectangleCommand_Handler,
		},
		{
			MethodName: "RespondTrackingOffCommand",
			Handler:    _TrackingServerService_RespondTrackingOffCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeTrackingPointCommand",
			Handler:       _TrackingServerService_SubscribeTrackingPointCommand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeTrackingRectangleCommand",
			Handler:       _TrackingServerService_SubscribeTrackingRectangleCommand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeTrackingOffCommand",
			Handler:       _TrackingServerService_SubscribeTrackingOffCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tracking_server.proto",
}