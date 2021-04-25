// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/GetImageCode.proto

package GetImageCode

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GetImageCode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GetImageCode_proto_msgTypes[0]
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
	return file_proto_GetImageCode_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string         `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string         `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Pix    []uint32       `protobuf:"varint,3,rep,packed,name=pix,proto3" json:"pix,omitempty"`
	Stride int64          `protobuf:"varint,4,opt,name=stride,proto3" json:"stride,omitempty"`
	Min    *ResponsePoint `protobuf:"bytes,5,opt,name=min,proto3" json:"min,omitempty"`
	Max    *ResponsePoint `protobuf:"bytes,6,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GetImageCode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GetImageCode_proto_msgTypes[1]
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
	return file_proto_GetImageCode_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetPix() []uint32 {
	if x != nil {
		return x.Pix
	}
	return nil
}

func (x *Response) GetStride() int64 {
	if x != nil {
		return x.Stride
	}
	return 0
}

func (x *Response) GetMin() *ResponsePoint {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *Response) GetMax() *ResponsePoint {
	if x != nil {
		return x.Max
	}
	return nil
}

type ResponsePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ResponsePoint) Reset() {
	*x = ResponsePoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GetImageCode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePoint) ProtoMessage() {}

func (x *ResponsePoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GetImageCode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePoint.ProtoReflect.Descriptor instead.
func (*ResponsePoint) Descriptor() ([]byte, []int) {
	return file_proto_GetImageCode_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResponsePoint) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ResponsePoint) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_proto_GetImageCode_proto protoreflect.FileDescriptor

var file_proto_GetImageCode_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0xe7, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d,
	0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x78, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x03, 0x70, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x03,
	0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x1a, 0x23, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x79, 0x32, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_GetImageCode_proto_rawDescOnce sync.Once
	file_proto_GetImageCode_proto_rawDescData = file_proto_GetImageCode_proto_rawDesc
)

func file_proto_GetImageCode_proto_rawDescGZIP() []byte {
	file_proto_GetImageCode_proto_rawDescOnce.Do(func() {
		file_proto_GetImageCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GetImageCode_proto_rawDescData)
	})
	return file_proto_GetImageCode_proto_rawDescData
}

var file_proto_GetImageCode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_GetImageCode_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: GetImageCode.Request
	(*Response)(nil),      // 1: GetImageCode.Response
	(*ResponsePoint)(nil), // 2: GetImageCode.Response.point
}
var file_proto_GetImageCode_proto_depIdxs = []int32{
	2, // 0: GetImageCode.Response.min:type_name -> GetImageCode.Response.point
	2, // 1: GetImageCode.Response.max:type_name -> GetImageCode.Response.point
	0, // 2: GetImageCode.GetImageCode.GetImageCode:input_type -> GetImageCode.Request
	1, // 3: GetImageCode.GetImageCode.GetImageCode:output_type -> GetImageCode.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_GetImageCode_proto_init() }
func file_proto_GetImageCode_proto_init() {
	if File_proto_GetImageCode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_GetImageCode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_GetImageCode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_GetImageCode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePoint); i {
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
			RawDescriptor: file_proto_GetImageCode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_GetImageCode_proto_goTypes,
		DependencyIndexes: file_proto_GetImageCode_proto_depIdxs,
		MessageInfos:      file_proto_GetImageCode_proto_msgTypes,
	}.Build()
	File_proto_GetImageCode_proto = out.File
	file_proto_GetImageCode_proto_rawDesc = nil
	file_proto_GetImageCode_proto_goTypes = nil
	file_proto_GetImageCode_proto_depIdxs = nil
}
