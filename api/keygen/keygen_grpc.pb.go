// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: keygen.proto

package keygen

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
	Keygen_GenerateKey_FullMethodName = "/keygen.Keygen/GenerateKey"
)

// KeygenClient is the client API for Keygen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeygenClient interface {
	GenerateKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error)
}

type keygenClient struct {
	cc grpc.ClientConnInterface
}

func NewKeygenClient(cc grpc.ClientConnInterface) KeygenClient {
	return &keygenClient{cc}
}

func (c *keygenClient) GenerateKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Key)
	err := c.cc.Invoke(ctx, Keygen_GenerateKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeygenServer is the server API for Keygen service.
// All implementations must embed UnimplementedKeygenServer
// for forward compatibility.
type KeygenServer interface {
	GenerateKey(context.Context, *Empty) (*Key, error)
	mustEmbedUnimplementedKeygenServer()
}

// UnimplementedKeygenServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeygenServer struct{}

func (UnimplementedKeygenServer) GenerateKey(context.Context, *Empty) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKey not implemented")
}
func (UnimplementedKeygenServer) mustEmbedUnimplementedKeygenServer() {}
func (UnimplementedKeygenServer) testEmbeddedByValue()                {}

// UnsafeKeygenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeygenServer will
// result in compilation errors.
type UnsafeKeygenServer interface {
	mustEmbedUnimplementedKeygenServer()
}

func RegisterKeygenServer(s grpc.ServiceRegistrar, srv KeygenServer) {
	// If the following call pancis, it indicates UnimplementedKeygenServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Keygen_ServiceDesc, srv)
}

func _Keygen_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeygenServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keygen_GenerateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeygenServer).GenerateKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Keygen_ServiceDesc is the grpc.ServiceDesc for Keygen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keygen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keygen.Keygen",
	HandlerType: (*KeygenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKey",
			Handler:    _Keygen_GenerateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keygen.proto",
}
