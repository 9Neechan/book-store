// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: genre.proto

package genre

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

// GenreV1Client is the client API for GenreV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenreV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type genreV1Client struct {
	cc grpc.ClientConnInterface
}

func NewGenreV1Client(cc grpc.ClientConnInterface) GenreV1Client {
	return &genreV1Client{cc}
}

func (c *genreV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.GenreV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.GenreV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenreV1Server is the server API for GenreV1 service.
// All implementations must embed UnimplementedGenreV1Server
// for forward compatibility
type GenreV1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedGenreV1Server()
}

// UnimplementedGenreV1Server must be embedded to have forward compatible implementations.
type UnimplementedGenreV1Server struct {
}

func (UnimplementedGenreV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGenreV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGenreV1Server) mustEmbedUnimplementedGenreV1Server() {}

// UnsafeGenreV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenreV1Server will
// result in compilation errors.
type UnsafeGenreV1Server interface {
	mustEmbedUnimplementedGenreV1Server()
}

func RegisterGenreV1Server(s grpc.ServiceRegistrar, srv GenreV1Server) {
	s.RegisterService(&GenreV1_ServiceDesc, srv)
}

func _GenreV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.GenreV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.GenreV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenreV1_ServiceDesc is the grpc.ServiceDesc for GenreV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenreV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.book_store.v1.GenreV1",
	HandlerType: (*GenreV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GenreV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GenreV1_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "genre.proto",
}
