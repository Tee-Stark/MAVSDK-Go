// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: geofence.proto

package geofence

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

// GeofenceServiceClient is the client API for GeofenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeofenceServiceClient interface {
	// Upload geofences.
	//
	// Polygon and Circular geofences are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error)
	// Clear all geofences saved on the vehicle.
	ClearGeofence(ctx context.Context, in *ClearGeofenceRequest, opts ...grpc.CallOption) (*ClearGeofenceResponse, error)
}

type geofenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeofenceServiceClient(cc grpc.ClientConnInterface) GeofenceServiceClient {
	return &geofenceServiceClient{cc}
}

func (c *geofenceServiceClient) UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error) {
	out := new(UploadGeofenceResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geofenceServiceClient) ClearGeofence(ctx context.Context, in *ClearGeofenceRequest, opts ...grpc.CallOption) (*ClearGeofenceResponse, error) {
	out := new(ClearGeofenceResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.geofence.GeofenceService/ClearGeofence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeofenceServiceServer is the server API for GeofenceService service.
// All implementations must embed UnimplementedGeofenceServiceServer
// for forward compatibility
type GeofenceServiceServer interface {
	// Upload geofences.
	//
	// Polygon and Circular geofences are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(context.Context, *UploadGeofenceRequest) (*UploadGeofenceResponse, error)
	// Clear all geofences saved on the vehicle.
	ClearGeofence(context.Context, *ClearGeofenceRequest) (*ClearGeofenceResponse, error)
	mustEmbedUnimplementedGeofenceServiceServer()
}

// UnimplementedGeofenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeofenceServiceServer struct {
}

func (UnimplementedGeofenceServiceServer) UploadGeofence(context.Context, *UploadGeofenceRequest) (*UploadGeofenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGeofence not implemented")
}
func (UnimplementedGeofenceServiceServer) ClearGeofence(context.Context, *ClearGeofenceRequest) (*ClearGeofenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearGeofence not implemented")
}
func (UnimplementedGeofenceServiceServer) mustEmbedUnimplementedGeofenceServiceServer() {}

// UnsafeGeofenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeofenceServiceServer will
// result in compilation errors.
type UnsafeGeofenceServiceServer interface {
	mustEmbedUnimplementedGeofenceServiceServer()
}

func RegisterGeofenceServiceServer(s grpc.ServiceRegistrar, srv GeofenceServiceServer) {
	s.RegisterService(&GeofenceService_ServiceDesc, srv)
}

func _GeofenceService_UploadGeofence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadGeofenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, req.(*UploadGeofenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeofenceService_ClearGeofence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearGeofenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeofenceServiceServer).ClearGeofence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.geofence.GeofenceService/ClearGeofence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeofenceServiceServer).ClearGeofence(ctx, req.(*ClearGeofenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeofenceService_ServiceDesc is the grpc.ServiceDesc for GeofenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeofenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.geofence.GeofenceService",
	HandlerType: (*GeofenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadGeofence",
			Handler:    _GeofenceService_UploadGeofence_Handler,
		},
		{
			MethodName: "ClearGeofence",
			Handler:    _GeofenceService_ClearGeofence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geofence.proto",
}
