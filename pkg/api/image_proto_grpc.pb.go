// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: image_proto.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Saver_Save_FullMethodName = "/Saver/Save"
	Saver_Get_FullMethodName  = "/Saver/Get"
)

// SaverClient is the client API for Saver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaverClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveStatus, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type saverClient struct {
	cc grpc.ClientConnInterface
}

func NewSaverClient(cc grpc.ClientConnInterface) SaverClient {
	return &saverClient{cc}
}

func (c *saverClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveStatus)
	err := c.cc.Invoke(ctx, Saver_Save_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saverClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Saver_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaverServer is the server API for Saver service.
// All implementations must embed UnimplementedSaverServer
// for forward compatibility.
type SaverServer interface {
	Save(context.Context, *SaveRequest) (*SaveStatus, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedSaverServer()
}

// UnimplementedSaverServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSaverServer struct{}

func (UnimplementedSaverServer) Save(context.Context, *SaveRequest) (*SaveStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedSaverServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSaverServer) mustEmbedUnimplementedSaverServer() {}
func (UnimplementedSaverServer) testEmbeddedByValue()               {}

// UnsafeSaverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaverServer will
// result in compilation errors.
type UnsafeSaverServer interface {
	mustEmbedUnimplementedSaverServer()
}

func RegisterSaverServer(s grpc.ServiceRegistrar, srv SaverServer) {
	// If the following call pancis, it indicates UnimplementedSaverServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Saver_ServiceDesc, srv)
}

func _Saver_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Saver_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Saver_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Saver_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Saver_ServiceDesc is the grpc.ServiceDesc for Saver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Saver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Saver",
	HandlerType: (*SaverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Saver_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Saver_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image_proto.proto",
}