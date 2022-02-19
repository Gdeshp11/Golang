// Proto file for movie info service. Note this is gRPC proto syntax (not Go)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: movieapi/movieapi.proto

package movieapi

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

// The request message containing movie name
type MovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *MovieRequest) Reset() {
	*x = MovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieapi_movieapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRequest) ProtoMessage() {}

func (x *MovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movieapi_movieapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRequest.ProtoReflect.Descriptor instead.
func (*MovieRequest) Descriptor() ([]byte, []int) {
	return file_movieapi_movieapi_proto_rawDescGZIP(), []int{0}
}

func (x *MovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// The response message containining movie info
type MovieReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year     int32    `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Director string   `protobuf:"bytes,2,opt,name=director,proto3" json:"director,omitempty"`
	Cast     []string `protobuf:"bytes,3,rep,name=cast,proto3" json:"cast,omitempty"`
}

func (x *MovieReply) Reset() {
	*x = MovieReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieapi_movieapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieReply) ProtoMessage() {}

func (x *MovieReply) ProtoReflect() protoreflect.Message {
	mi := &file_movieapi_movieapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieReply.ProtoReflect.Descriptor instead.
func (*MovieReply) Descriptor() ([]byte, []int) {
	return file_movieapi_movieapi_proto_rawDescGZIP(), []int{1}
}

func (x *MovieReply) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *MovieReply) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieReply) GetCast() []string {
	if x != nil {
		return x.Cast
	}
	return nil
}

// request message to set movie info
type MovieData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year     int32    `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Director string   `protobuf:"bytes,3,opt,name=director,proto3" json:"director,omitempty"`
	Cast     []string `protobuf:"bytes,4,rep,name=cast,proto3" json:"cast,omitempty"`
}

func (x *MovieData) Reset() {
	*x = MovieData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieapi_movieapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieData) ProtoMessage() {}

func (x *MovieData) ProtoReflect() protoreflect.Message {
	mi := &file_movieapi_movieapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieData.ProtoReflect.Descriptor instead.
func (*MovieData) Descriptor() ([]byte, []int) {
	return file_movieapi_movieapi_proto_rawDescGZIP(), []int{2}
}

func (x *MovieData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieData) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *MovieData) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieData) GetCast() []string {
	if x != nil {
		return x.Cast
	}
	return nil
}

// The response message containing status of SetMovieInfo rpc
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieapi_movieapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_movieapi_movieapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_movieapi_movieapi_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_movieapi_movieapi_proto protoreflect.FileDescriptor

var file_movieapi_movieapi_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x61, 0x70, 0x69, 0x22, 0x24, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x50, 0x0a, 0x0a, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x09, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61,
	0x73, 0x74, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x84, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x75, 0x6e, 0x72, 0x61, 0x76, 0x69, 0x6e, 0x64,
	0x72, 0x61, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x35, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movieapi_movieapi_proto_rawDescOnce sync.Once
	file_movieapi_movieapi_proto_rawDescData = file_movieapi_movieapi_proto_rawDesc
)

func file_movieapi_movieapi_proto_rawDescGZIP() []byte {
	file_movieapi_movieapi_proto_rawDescOnce.Do(func() {
		file_movieapi_movieapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_movieapi_movieapi_proto_rawDescData)
	})
	return file_movieapi_movieapi_proto_rawDescData
}

var file_movieapi_movieapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_movieapi_movieapi_proto_goTypes = []interface{}{
	(*MovieRequest)(nil), // 0: movieapi.MovieRequest
	(*MovieReply)(nil),   // 1: movieapi.MovieReply
	(*MovieData)(nil),    // 2: movieapi.MovieData
	(*Status)(nil),       // 3: movieapi.Status
}
var file_movieapi_movieapi_proto_depIdxs = []int32{
	0, // 0: movieapi.MovieInfo.GetMovieInfo:input_type -> movieapi.MovieRequest
	2, // 1: movieapi.MovieInfo.SetMovieInfo:input_type -> movieapi.MovieData
	1, // 2: movieapi.MovieInfo.GetMovieInfo:output_type -> movieapi.MovieReply
	3, // 3: movieapi.MovieInfo.SetMovieInfo:output_type -> movieapi.Status
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_movieapi_movieapi_proto_init() }
func file_movieapi_movieapi_proto_init() {
	if File_movieapi_movieapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movieapi_movieapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRequest); i {
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
		file_movieapi_movieapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieReply); i {
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
		file_movieapi_movieapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieData); i {
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
		file_movieapi_movieapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_movieapi_movieapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movieapi_movieapi_proto_goTypes,
		DependencyIndexes: file_movieapi_movieapi_proto_depIdxs,
		MessageInfos:      file_movieapi_movieapi_proto_msgTypes,
	}.Build()
	File_movieapi_movieapi_proto = out.File
	file_movieapi_movieapi_proto_rawDesc = nil
	file_movieapi_movieapi_proto_goTypes = nil
	file_movieapi_movieapi_proto_depIdxs = nil
}
