// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/v1/url.proto

package grpc

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

type OriginalUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *OriginalUrlRequest) Reset() {
	*x = OriginalUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalUrlRequest) ProtoMessage() {}

func (x *OriginalUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalUrlRequest.ProtoReflect.Descriptor instead.
func (*OriginalUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_url_proto_rawDescGZIP(), []int{0}
}

func (x *OriginalUrlRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type ShortUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *ShortUrlResponse) Reset() {
	*x = ShortUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUrlResponse) ProtoMessage() {}

func (x *ShortUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortUrlResponse.ProtoReflect.Descriptor instead.
func (*ShortUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortUrlResponse) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type OriginalUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *OriginalUrlResponse) Reset() {
	*x = OriginalUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalUrlResponse) ProtoMessage() {}

func (x *OriginalUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalUrlResponse.ProtoReflect.Descriptor instead.
func (*OriginalUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_url_proto_rawDescGZIP(), []int{2}
}

func (x *OriginalUrlResponse) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type ShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *ShortUrlRequest) Reset() {
	*x = ShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_url_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUrlRequest) ProtoMessage() {}

func (x *ShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_url_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortUrlRequest.ProtoReflect.Descriptor instead.
func (*ShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_url_proto_rawDescGZIP(), []int{3}
}

func (x *ShortUrlRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

var File_proto_v1_url_proto protoreflect.FileDescriptor

var file_proto_v1_url_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x26, 0x0a, 0x12, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52,
	0x4c, 0x22, 0x24, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x27, 0x0a, 0x13, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c,
	0x22, 0x23, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x52, 0x4c, 0x32, 0x97, 0x01, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x2e,
	0x75, 0x72, 0x6c, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_v1_url_proto_rawDescOnce sync.Once
	file_proto_v1_url_proto_rawDescData = file_proto_v1_url_proto_rawDesc
)

func file_proto_v1_url_proto_rawDescGZIP() []byte {
	file_proto_v1_url_proto_rawDescOnce.Do(func() {
		file_proto_v1_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_url_proto_rawDescData)
	})
	return file_proto_v1_url_proto_rawDescData
}

var file_proto_v1_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_url_proto_goTypes = []interface{}{
	(*OriginalUrlRequest)(nil),  // 0: url.OriginalUrlRequest
	(*ShortUrlResponse)(nil),    // 1: url.ShortUrlResponse
	(*OriginalUrlResponse)(nil), // 2: url.OriginalUrlResponse
	(*ShortUrlRequest)(nil),     // 3: url.ShortUrlRequest
}
var file_proto_v1_url_proto_depIdxs = []int32{
	0, // 0: url.UrlShortenService.CreateShortUrl:input_type -> url.OriginalUrlRequest
	3, // 1: url.UrlShortenService.GetOriginalUrl:input_type -> url.ShortUrlRequest
	1, // 2: url.UrlShortenService.CreateShortUrl:output_type -> url.ShortUrlResponse
	2, // 3: url.UrlShortenService.GetOriginalUrl:output_type -> url.OriginalUrlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_url_proto_init() }
func file_proto_v1_url_proto_init() {
	if File_proto_v1_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalUrlRequest); i {
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
		file_proto_v1_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortUrlResponse); i {
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
		file_proto_v1_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalUrlResponse); i {
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
		file_proto_v1_url_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortUrlRequest); i {
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
			RawDescriptor: file_proto_v1_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_url_proto_goTypes,
		DependencyIndexes: file_proto_v1_url_proto_depIdxs,
		MessageInfos:      file_proto_v1_url_proto_msgTypes,
	}.Build()
	File_proto_v1_url_proto = out.File
	file_proto_v1_url_proto_rawDesc = nil
	file_proto_v1_url_proto_goTypes = nil
	file_proto_v1_url_proto_depIdxs = nil
}