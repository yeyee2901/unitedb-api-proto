// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: unitedb/v1/unitedb_service.proto

package unitepb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_unitedb_v1_unitedb_service_proto protoreflect.FileDescriptor

var file_unitedb_v1_unitedb_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69,
	0x74, 0x65, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1d,
	0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x82, 0x01, 0x0a, 0x0e,
	0x55, 0x6e, 0x69, 0x74, 0x65, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x20, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2d, 0x69, 0x74, 0x65, 0x6d,
	0x42, 0x8e, 0x01, 0x5a, 0x14, 0x2e, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76,
	0x31, 0x3b, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x70, 0x62, 0x92, 0x41, 0x75, 0x12, 0x73, 0x0a, 0x1b,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x20, 0x55, 0x6e, 0x69, 0x74, 0x65, 0x20, 0x44, 0x42,
	0x20, 0x2d, 0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x41, 0x50, 0x49, 0x22, 0x4f, 0x0a, 0x05, 0x59,
	0x65, 0x79, 0x65, 0x65, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x65, 0x79, 0x65, 0x65, 0x32, 0x39,
	0x30, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c, 0x37, 0x37, 0x37,
	0x73, 0x68, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_unitedb_v1_unitedb_service_proto_goTypes = []interface{}{
	(*GetBattleItemRequest)(nil),  // 0: unitedb.v1.GetBattleItemRequest
	(*GetBattleItemResponse)(nil), // 1: unitedb.v1.GetBattleItemResponse
}
var file_unitedb_v1_unitedb_service_proto_depIdxs = []int32{
	0, // 0: unitedb.v1.UniteDBService.GetBattleItem:input_type -> unitedb.v1.GetBattleItemRequest
	1, // 1: unitedb.v1.UniteDBService.GetBattleItem:output_type -> unitedb.v1.GetBattleItemResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_unitedb_v1_unitedb_service_proto_init() }
func file_unitedb_v1_unitedb_service_proto_init() {
	if File_unitedb_v1_unitedb_service_proto != nil {
		return
	}
	file_unitedb_v1_battle_items_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_unitedb_v1_unitedb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unitedb_v1_unitedb_service_proto_goTypes,
		DependencyIndexes: file_unitedb_v1_unitedb_service_proto_depIdxs,
	}.Build()
	File_unitedb_v1_unitedb_service_proto = out.File
	file_unitedb_v1_unitedb_service_proto_rawDesc = nil
	file_unitedb_v1_unitedb_service_proto_goTypes = nil
	file_unitedb_v1_unitedb_service_proto_depIdxs = nil
}