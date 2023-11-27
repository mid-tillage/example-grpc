// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/catalog/catalog.proto

package catalog

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
	CatalogService_CreateCatalog_FullMethodName = "/catalog.CatalogService/CreateCatalog"
	CatalogService_GetCatalog_FullMethodName    = "/catalog.CatalogService/GetCatalog"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*Catalog, error)
	GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogResponse, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*Catalog, error) {
	out := new(Catalog)
	err := c.cc.Invoke(ctx, CatalogService_CreateCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogResponse, error) {
	out := new(GetCatalogResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetCatalog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations must embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	CreateCatalog(context.Context, *CreateCatalogRequest) (*Catalog, error)
	GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogResponse, error)
	mustEmbedUnimplementedCatalogServiceServer()
}

// UnimplementedCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (UnimplementedCatalogServiceServer) CreateCatalog(context.Context, *CreateCatalogRequest) (*Catalog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalog not implemented")
}
func (UnimplementedCatalogServiceServer) GetCatalog(context.Context, *GetCatalogRequest) (*GetCatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalog not implemented")
}
func (UnimplementedCatalogServiceServer) mustEmbedUnimplementedCatalogServiceServer() {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_CreateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCatalog(ctx, req.(*CreateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalog(ctx, req.(*GetCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCatalog",
			Handler:    _CatalogService_CreateCatalog_Handler,
		},
		{
			MethodName: "GetCatalog",
			Handler:    _CatalogService_GetCatalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/catalog/catalog.proto",
}