// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: envoy/extensions/compression/zstd/decompressor/v3/zstd.proto

package decompressorv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Zstd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dictionaries for decompression. Zstd offers dictionary compression, which greatly improves
	// efficiency on small files and messages. It is necessary to ensure that the dictionary used for
	// decompression is the same as the compression dictionary. Multiple dictionaries can be set, and the
	// dictionary will be automatically selected for decompression according to the dictionary ID in the
	// source content.
	// Please refer to `zstd manual <https://github.com/facebook/zstd/blob/dev/programs/zstd.1.md#dictionary-builder>`_
	// to train specific dictionaries for decompression.
	Dictionaries []*v3.DataSource `protobuf:"bytes,1,rep,name=dictionaries,proto3" json:"dictionaries,omitempty"`
	// Value for decompressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *Zstd) Reset() {
	*x = Zstd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zstd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zstd) ProtoMessage() {}

func (x *Zstd) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zstd.ProtoReflect.Descriptor instead.
func (*Zstd) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescGZIP(), []int{0}
}

func (x *Zstd) GetDictionaries() []*v3.DataSource {
	if x != nil {
		return x.Dictionaries
	}
	return nil
}

func (x *Zstd) GetChunkSize() *wrappers.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

var File_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto protoreflect.FileDescriptor

var file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7a,
	0x73, 0x74, 0x64, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x2f, 0x76, 0x33, 0x2f, 0x7a, 0x73, 0x74, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x7a, 0x73, 0x74,
	0x64, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x04, 0x5a,
	0x73, 0x74, 0x64, 0x12, 0x44, 0x0a, 0x0c, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09,
	0x2a, 0x07, 0x18, 0x80, 0x80, 0x04, 0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x42, 0xbf, 0x01, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x7a, 0x73, 0x74, 0x64, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x5a, 0x73, 0x74, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x73, 0x74, 0x64, 0x2f,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescOnce sync.Once
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescData = file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDesc
)

func file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescGZIP() []byte {
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescData)
	})
	return file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDescData
}

var file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_goTypes = []interface{}{
	(*Zstd)(nil),                 // 0: envoy.extensions.compression.zstd.decompressor.v3.Zstd
	(*v3.DataSource)(nil),        // 1: envoy.config.core.v3.DataSource
	(*wrappers.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
}
var file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.compression.zstd.decompressor.v3.Zstd.dictionaries:type_name -> envoy.config.core.v3.DataSource
	2, // 1: envoy.extensions.compression.zstd.decompressor.v3.Zstd.chunk_size:type_name -> google.protobuf.UInt32Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_init() }
func file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_init() {
	if File_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zstd); i {
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
			RawDescriptor: file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_msgTypes,
	}.Build()
	File_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto = out.File
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_rawDesc = nil
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_goTypes = nil
	file_envoy_extensions_compression_zstd_decompressor_v3_zstd_proto_depIdxs = nil
}
