// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: gmystcachepb.proto

package gmystcachepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmystcachepb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_gmystcachepb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_gmystcachepb_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmystcachepb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_gmystcachepb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_gmystcachepb_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_gmystcachepb_proto protoreflect.FileDescriptor

var file_gmystcachepb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x70, 0x62, 0x22, 0x31, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x42, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x67,
	0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67,
	0x6d, 0x79, 0x73, 0x74, 0x2f, 0x67, 0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f,
	0x67, 0x6d, 0x79, 0x73, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gmystcachepb_proto_rawDescOnce sync.Once
	file_gmystcachepb_proto_rawDescData = file_gmystcachepb_proto_rawDesc
)

func file_gmystcachepb_proto_rawDescGZIP() []byte {
	file_gmystcachepb_proto_rawDescOnce.Do(func() {
		file_gmystcachepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_gmystcachepb_proto_rawDescData)
	})
	return file_gmystcachepb_proto_rawDescData
}

var file_gmystcachepb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gmystcachepb_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: gmystcachepb.Request
	(*Response)(nil), // 1: gmystcachepb.Response
}
var file_gmystcachepb_proto_depIdxs = []int32{
	0, // 0: gmystcachepb.GroupCache.Get:input_type -> gmystcachepb.Request
	1, // 1: gmystcachepb.GroupCache.Get:output_type -> gmystcachepb.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gmystcachepb_proto_init() }
func file_gmystcachepb_proto_init() {
	if File_gmystcachepb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gmystcachepb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_gmystcachepb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_gmystcachepb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gmystcachepb_proto_goTypes,
		DependencyIndexes: file_gmystcachepb_proto_depIdxs,
		MessageInfos:      file_gmystcachepb_proto_msgTypes,
	}.Build()
	File_gmystcachepb_proto = out.File
	file_gmystcachepb_proto_rawDesc = nil
	file_gmystcachepb_proto_goTypes = nil
	file_gmystcachepb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupCacheClient is the client API for GroupCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupCacheClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type groupCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupCacheClient(cc grpc.ClientConnInterface) GroupCacheClient {
	return &groupCacheClient{cc}
}

func (c *groupCacheClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/gmystcachepb.GroupCache/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupCacheServer is the server API for GroupCache service.
type GroupCacheServer interface {
	Get(context.Context, *Request) (*Response, error)
}

// UnimplementedGroupCacheServer can be embedded to have forward compatible implementations.
type UnimplementedGroupCacheServer struct {
}

func (*UnimplementedGroupCacheServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterGroupCacheServer(s *grpc.Server, srv GroupCacheServer) {
	s.RegisterService(&_GroupCache_serviceDesc, srv)
}

func _GroupCache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupCacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmystcachepb.GroupCache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupCacheServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gmystcachepb.GroupCache",
	HandlerType: (*GroupCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GroupCache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gmystcachepb.proto",
}
