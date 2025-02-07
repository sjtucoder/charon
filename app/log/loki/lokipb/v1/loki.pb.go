// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: app/log/loki/lokipb/v1/loki.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Streams       []*Stream              `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_app_log_loki_lokipb_v1_loki_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetStreams() []*Stream {
	if x != nil {
		return x.Streams
	}
	return nil
}

type Stream struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Labels        string                 `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	Entries       []*Entry               `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	Hash          uint64                 `protobuf:"varint,3,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Stream) Reset() {
	*x = Stream{}
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_app_log_loki_lokipb_v1_loki_proto_rawDescGZIP(), []int{1}
}

func (x *Stream) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *Stream) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *Stream) GetHash() uint64 {
	if x != nil {
		return x.Hash
	}
	return 0
}

type Entry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Line          string                 `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Entry) Reset() {
	*x = Entry{}
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_app_log_loki_lokipb_v1_loki_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_app_log_loki_lokipb_v1_loki_proto_rawDescGZIP(), []int{2}
}

func (x *Entry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Entry) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

var File_app_log_loki_lokipb_v1_loki_proto protoreflect.FileDescriptor

var file_app_log_loki_lokipb_v1_loki_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x6b, 0x69, 0x2f, 0x6c,
	0x6f, 0x6b, 0x69, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x6b, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x6c, 0x6f, 0x6b,
	0x69, 0x2e, 0x6c, 0x6f, 0x6b, 0x69, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0b,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x6c, 0x6f, 0x6b, 0x69, 0x2e, 0x6c, 0x6f, 0x6b, 0x69,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x6d, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x6c, 0x6f, 0x6b, 0x69, 0x2e, 0x6c, 0x6f, 0x6b, 0x69, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x55, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x6f, 0x6c, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x6b, 0x69, 0x2f, 0x6c, 0x6f, 0x6b, 0x69, 0x70, 0x62,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_app_log_loki_lokipb_v1_loki_proto_rawDescOnce sync.Once
	file_app_log_loki_lokipb_v1_loki_proto_rawDescData []byte
)

func file_app_log_loki_lokipb_v1_loki_proto_rawDescGZIP() []byte {
	file_app_log_loki_lokipb_v1_loki_proto_rawDescOnce.Do(func() {
		file_app_log_loki_lokipb_v1_loki_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_log_loki_lokipb_v1_loki_proto_rawDesc), len(file_app_log_loki_lokipb_v1_loki_proto_rawDesc)))
	})
	return file_app_log_loki_lokipb_v1_loki_proto_rawDescData
}

var file_app_log_loki_lokipb_v1_loki_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_log_loki_lokipb_v1_loki_proto_goTypes = []any{
	(*PushRequest)(nil),           // 0: app.log.loki.lokipb.v1.PushRequest
	(*Stream)(nil),                // 1: app.log.loki.lokipb.v1.Stream
	(*Entry)(nil),                 // 2: app.log.loki.lokipb.v1.Entry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_app_log_loki_lokipb_v1_loki_proto_depIdxs = []int32{
	1, // 0: app.log.loki.lokipb.v1.PushRequest.streams:type_name -> app.log.loki.lokipb.v1.Stream
	2, // 1: app.log.loki.lokipb.v1.Stream.entries:type_name -> app.log.loki.lokipb.v1.Entry
	3, // 2: app.log.loki.lokipb.v1.Entry.timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_log_loki_lokipb_v1_loki_proto_init() }
func file_app_log_loki_lokipb_v1_loki_proto_init() {
	if File_app_log_loki_lokipb_v1_loki_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_log_loki_lokipb_v1_loki_proto_rawDesc), len(file_app_log_loki_lokipb_v1_loki_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_log_loki_lokipb_v1_loki_proto_goTypes,
		DependencyIndexes: file_app_log_loki_lokipb_v1_loki_proto_depIdxs,
		MessageInfos:      file_app_log_loki_lokipb_v1_loki_proto_msgTypes,
	}.Build()
	File_app_log_loki_lokipb_v1_loki_proto = out.File
	file_app_log_loki_lokipb_v1_loki_proto_goTypes = nil
	file_app_log_loki_lokipb_v1_loki_proto_depIdxs = nil
}
