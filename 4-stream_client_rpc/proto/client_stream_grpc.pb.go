// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// StreamClientClient is the client API for StreamClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClientClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error)
}

type streamClientClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClientClient(cc grpc.ClientConnInterface) StreamClientClient {
	return &streamClientClient{cc}
}

func (c *streamClientClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.StreamClient/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClientClient) RouteList(ctx context.Context, opts ...grpc.CallOption) (StreamClient_RouteListClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamClient_ServiceDesc.Streams[0], "/proto.StreamClient/RouteList", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClientRouteListClient{stream}
	return x, nil
}

type StreamClient_RouteListClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type streamClientRouteListClient struct {
	grpc.ClientStream
}

func (x *streamClientRouteListClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClientRouteListClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClientServer is the server API for StreamClient service.
// All implementations must embed UnimplementedStreamClientServer
// for forward compatibility
type StreamClientServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
	RouteList(StreamClient_RouteListServer) error
	mustEmbedUnimplementedStreamClientServer()
}

// UnimplementedStreamClientServer must be embedded to have forward compatible implementations.
type UnimplementedStreamClientServer struct {
}

func (UnimplementedStreamClientServer) Route(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedStreamClientServer) RouteList(StreamClient_RouteListServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteList not implemented")
}
func (UnimplementedStreamClientServer) mustEmbedUnimplementedStreamClientServer() {}

// UnsafeStreamClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamClientServer will
// result in compilation errors.
type UnsafeStreamClientServer interface {
	mustEmbedUnimplementedStreamClientServer()
}

func RegisterStreamClientServer(s grpc.ServiceRegistrar, srv StreamClientServer) {
	s.RegisterService(&StreamClient_ServiceDesc, srv)
}

func _StreamClient_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamClientServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StreamClient/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamClientServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamClient_RouteList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamClientServer).RouteList(&streamClientRouteListServer{stream})
}

type StreamClient_RouteListServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamClientRouteListServer struct {
	grpc.ServerStream
}

func (x *streamClientRouteListServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamClientRouteListServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClient_ServiceDesc is the grpc.ServiceDesc for StreamClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamClient",
	HandlerType: (*StreamClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _StreamClient_Route_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteList",
			Handler:       _StreamClient_RouteList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/client_stream.proto",
}
