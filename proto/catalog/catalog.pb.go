// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/catalog/catalog.proto

package catalog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_proto_catalog_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Catalog) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Catalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCatalogRequest) Reset() {
	*x = CreateCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogRequest) ProtoMessage() {}

func (x *CreateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogRequest.ProtoReflect.Descriptor instead.
func (*CreateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCatalogResponse) Reset() {
	*x = CreateCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogResponse) ProtoMessage() {}

func (x *CreateCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogResponse.ProtoReflect.Descriptor instead.
func (*CreateCatalogResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCatalogResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCatalogResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_proto_catalog_catalog_proto_rawDescGZIP(), []int{3}
}

type GetCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog []*Catalog `protobuf:"bytes,1,rep,name=catalog,proto3" json:"catalog,omitempty"`
}

func (x *GetCatalogResponse) Reset() {
	*x = GetCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_catalog_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogResponse) ProtoMessage() {}

func (x *GetCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_catalog_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogResponse.ProtoReflect.Descriptor instead.
func (*GetCatalogResponse) Descriptor() ([]byte, []int) {
	return file_proto_catalog_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *GetCatalogResponse) GetCatalog() []*Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

var File_proto_catalog_catalog_proto protoreflect.FileDescriptor

var file_proto_catalog_catalog_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x07, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x32, 0x9d, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_catalog_catalog_proto_rawDescOnce sync.Once
	file_proto_catalog_catalog_proto_rawDescData = file_proto_catalog_catalog_proto_rawDesc
)

func file_proto_catalog_catalog_proto_rawDescGZIP() []byte {
	file_proto_catalog_catalog_proto_rawDescOnce.Do(func() {
		file_proto_catalog_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_catalog_catalog_proto_rawDescData)
	})
	return file_proto_catalog_catalog_proto_rawDescData
}

var file_proto_catalog_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_catalog_catalog_proto_goTypes = []interface{}{
	(*Catalog)(nil),               // 0: catalog.Catalog
	(*CreateCatalogRequest)(nil),  // 1: catalog.CreateCatalogRequest
	(*CreateCatalogResponse)(nil), // 2: catalog.CreateCatalogResponse
	(*GetCatalogRequest)(nil),     // 3: catalog.GetCatalogRequest
	(*GetCatalogResponse)(nil),    // 4: catalog.GetCatalogResponse
}
var file_proto_catalog_catalog_proto_depIdxs = []int32{
	0, // 0: catalog.GetCatalogResponse.catalog:type_name -> catalog.Catalog
	1, // 1: catalog.CatalogService.CreateCatalog:input_type -> catalog.CreateCatalogRequest
	3, // 2: catalog.CatalogService.GetCatalog:input_type -> catalog.GetCatalogRequest
	0, // 3: catalog.CatalogService.CreateCatalog:output_type -> catalog.Catalog
	4, // 4: catalog.CatalogService.GetCatalog:output_type -> catalog.GetCatalogResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_catalog_catalog_proto_init() }
func file_proto_catalog_catalog_proto_init() {
	if File_proto_catalog_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_catalog_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatalogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatalogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_catalog_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_catalog_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_catalog_catalog_proto_goTypes,
		DependencyIndexes: file_proto_catalog_catalog_proto_depIdxs,
		MessageInfos:      file_proto_catalog_catalog_proto_msgTypes,
	}.Build()
	File_proto_catalog_catalog_proto = out.File
	file_proto_catalog_catalog_proto_rawDesc = nil
	file_proto_catalog_catalog_proto_goTypes = nil
	file_proto_catalog_catalog_proto_depIdxs = nil
}