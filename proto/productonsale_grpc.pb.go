// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/productonsale.proto

package product

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
	ProductService_CreateProductOnSale_FullMethodName = "/productpnsale.ProductService/CreateProductOnSale"
	ProductService_GetProductsOnSale_FullMethodName   = "/productpnsale.ProductService/GetProductsOnSale"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProductOnSale(ctx context.Context, in *ProductOnSale, opts ...grpc.CallOption) (*CreateProductOnSaleResponse, error)
	GetProductsOnSale(ctx context.Context, in *GetProductsOnSaleRequest, opts ...grpc.CallOption) (*GetProductsOnSaleResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProductOnSale(ctx context.Context, in *ProductOnSale, opts ...grpc.CallOption) (*CreateProductOnSaleResponse, error) {
	out := new(CreateProductOnSaleResponse)
	err := c.cc.Invoke(ctx, ProductService_CreateProductOnSale_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductsOnSale(ctx context.Context, in *GetProductsOnSaleRequest, opts ...grpc.CallOption) (*GetProductsOnSaleResponse, error) {
	out := new(GetProductsOnSaleResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProductsOnSale_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProductOnSale(context.Context, *ProductOnSale) (*CreateProductOnSaleResponse, error)
	GetProductsOnSale(context.Context, *GetProductsOnSaleRequest) (*GetProductsOnSaleResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProductOnSale(context.Context, *ProductOnSale) (*CreateProductOnSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductOnSale not implemented")
}
func (UnimplementedProductServiceServer) GetProductsOnSale(context.Context, *GetProductsOnSaleRequest) (*GetProductsOnSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsOnSale not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProductOnSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductOnSale)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProductOnSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CreateProductOnSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProductOnSale(ctx, req.(*ProductOnSale))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductsOnSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsOnSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductsOnSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductsOnSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductsOnSale(ctx, req.(*GetProductsOnSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "productpnsale.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductOnSale",
			Handler:    _ProductService_CreateProductOnSale_Handler,
		},
		{
			MethodName: "GetProductsOnSale",
			Handler:    _ProductService_GetProductsOnSale_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/productonsale.proto",
}
