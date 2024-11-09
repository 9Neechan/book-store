// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: book.proto

package book

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

// BookV1Client is the client API for BookV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	ChangePrice(ctx context.Context, in *ChangePriceRequest, opts ...grpc.CallOption) (*ChangePriceResponse, error)
	ChangeAmount(ctx context.Context, in *ChangeAmountRequest, opts ...grpc.CallOption) (*ChangeAmountResponse, error)
}

type bookV1Client struct {
	cc grpc.ClientConnInterface
}

func NewBookV1Client(cc grpc.ClientConnInterface) BookV1Client {
	return &bookV1Client{cc}
}

func (c *bookV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.BookV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.BookV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookV1Client) ChangePrice(ctx context.Context, in *ChangePriceRequest, opts ...grpc.CallOption) (*ChangePriceResponse, error) {
	out := new(ChangePriceResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.BookV1/ChangePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookV1Client) ChangeAmount(ctx context.Context, in *ChangeAmountRequest, opts ...grpc.CallOption) (*ChangeAmountResponse, error) {
	out := new(ChangeAmountResponse)
	err := c.cc.Invoke(ctx, "/api.book_store.v1.BookV1/ChangeAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookV1Server is the server API for BookV1 service.
// All implementations must embed UnimplementedBookV1Server
// for forward compatibility
type BookV1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	ChangePrice(context.Context, *ChangePriceRequest) (*ChangePriceResponse, error)
	ChangeAmount(context.Context, *ChangeAmountRequest) (*ChangeAmountResponse, error)
	mustEmbedUnimplementedBookV1Server()
}

// UnimplementedBookV1Server must be embedded to have forward compatible implementations.
type UnimplementedBookV1Server struct {
}

func (UnimplementedBookV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookV1Server) ChangePrice(context.Context, *ChangePriceRequest) (*ChangePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePrice not implemented")
}
func (UnimplementedBookV1Server) ChangeAmount(context.Context, *ChangeAmountRequest) (*ChangeAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAmount not implemented")
}
func (UnimplementedBookV1Server) mustEmbedUnimplementedBookV1Server() {}

// UnsafeBookV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookV1Server will
// result in compilation errors.
type UnsafeBookV1Server interface {
	mustEmbedUnimplementedBookV1Server()
}

func RegisterBookV1Server(s grpc.ServiceRegistrar, srv BookV1Server) {
	s.RegisterService(&BookV1_ServiceDesc, srv)
}

func _BookV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.BookV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.BookV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookV1_ChangePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookV1Server).ChangePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.BookV1/ChangePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookV1Server).ChangePrice(ctx, req.(*ChangePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookV1_ChangeAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookV1Server).ChangeAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.book_store.v1.BookV1/ChangeAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookV1Server).ChangeAmount(ctx, req.(*ChangeAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookV1_ServiceDesc is the grpc.ServiceDesc for BookV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.book_store.v1.BookV1",
	HandlerType: (*BookV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookV1_Get_Handler,
		},
		{
			MethodName: "ChangePrice",
			Handler:    _BookV1_ChangePrice_Handler,
		},
		{
			MethodName: "ChangeAmount",
			Handler:    _BookV1_ChangeAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
