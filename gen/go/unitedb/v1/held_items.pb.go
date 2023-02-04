// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: unitedb/v1/held_items.proto

package unitepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HeldItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// effect of the item
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// tier of the item (at the last update)
	Tier string `protobuf:"bytes,3,opt,name=tier,proto3" json:"tier,omitempty"`
	// which main stats that is affected (Atk, SpAtk, etc)
	StatsBonus string `protobuf:"bytes,4,opt,name=stats_bonus,proto3" json:"stats_bonus,omitempty"`
	// stats at level 1
	StatsAtLevel_1 string `protobuf:"bytes,5,opt,name=stats_at_level_1,proto3" json:"stats_at_level_1,omitempty"`
	// stats at level 10
	StatsAtLevel_10 string `protobuf:"bytes,6,opt,name=stats_at_level_10,proto3" json:"stats_at_level_10,omitempty"`
	// stats at level 20
	StatsAtLevel_20 string `protobuf:"bytes,7,opt,name=stats_at_level_20,proto3" json:"stats_at_level_20,omitempty"`
	// when was this item last fetched from the website
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=last_updated,proto3" json:"last_updated,omitempty"`
}

func (x *HeldItem) Reset() {
	*x = HeldItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitedb_v1_held_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeldItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeldItem) ProtoMessage() {}

func (x *HeldItem) ProtoReflect() protoreflect.Message {
	mi := &file_unitedb_v1_held_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeldItem.ProtoReflect.Descriptor instead.
func (*HeldItem) Descriptor() ([]byte, []int) {
	return file_unitedb_v1_held_items_proto_rawDescGZIP(), []int{0}
}

func (x *HeldItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeldItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HeldItem) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *HeldItem) GetStatsBonus() string {
	if x != nil {
		return x.StatsBonus
	}
	return ""
}

func (x *HeldItem) GetStatsAtLevel_1() string {
	if x != nil {
		return x.StatsAtLevel_1
	}
	return ""
}

func (x *HeldItem) GetStatsAtLevel_10() string {
	if x != nil {
		return x.StatsAtLevel_10
	}
	return ""
}

func (x *HeldItem) GetStatsAtLevel_20() string {
	if x != nil {
		return x.StatsAtLevel_20
	}
	return ""
}

func (x *HeldItem) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type GetHeldItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item name
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty" validate:"required_without_all,omitempty,max=30"`  
	// item tier
	Tier *string `protobuf:"bytes,2,opt,name=tier,proto3,oneof" json:"tier,omitempty" validate:"required_without_all,omitempty,max=4"`  
	// stats bonus category of the item (such as Atk, HP, etc)
	StatsBonus *string `protobuf:"bytes,3,opt,name=statsBonus,proto3,oneof" json:"statsBonus,omitempty" validate:"required_without_all,omitempty,max=10"`  
}

func (x *GetHeldItemRequest) Reset() {
	*x = GetHeldItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitedb_v1_held_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeldItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeldItemRequest) ProtoMessage() {}

func (x *GetHeldItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unitedb_v1_held_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeldItemRequest.ProtoReflect.Descriptor instead.
func (*GetHeldItemRequest) Descriptor() ([]byte, []int) {
	return file_unitedb_v1_held_items_proto_rawDescGZIP(), []int{1}
}

func (x *GetHeldItemRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetHeldItemRequest) GetTier() string {
	if x != nil && x.Tier != nil {
		return *x.Tier
	}
	return ""
}

func (x *GetHeldItemRequest) GetStatsBonus() string {
	if x != nil && x.StatsBonus != nil {
		return *x.StatsBonus
	}
	return ""
}

type GetHeldItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response data
	Data []*HeldItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// length of the item found
	Total uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetHeldItemResponse) Reset() {
	*x = GetHeldItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitedb_v1_held_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeldItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeldItemResponse) ProtoMessage() {}

func (x *GetHeldItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unitedb_v1_held_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeldItemResponse.ProtoReflect.Descriptor instead.
func (*GetHeldItemResponse) Descriptor() ([]byte, []int) {
	return file_unitedb_v1_held_items_proto_rawDescGZIP(), []int{2}
}

func (x *GetHeldItemResponse) GetData() []*HeldItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetHeldItemResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_unitedb_v1_held_items_proto protoreflect.FileDescriptor

var file_unitedb_v1_held_items_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c,
	0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75,
	0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x48,
	0x65, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x31, 0x12,
	0x2c, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x31, 0x30, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x31, 0x30, 0x12, 0x2c, 0x0a,
	0x11, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x32, 0x30, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x32, 0x30, 0x12, 0x3e, 0x0a, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x48, 0x65, 0x6c, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76,
	0x31, 0x3b, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_unitedb_v1_held_items_proto_rawDescOnce sync.Once
	file_unitedb_v1_held_items_proto_rawDescData = file_unitedb_v1_held_items_proto_rawDesc
)

func file_unitedb_v1_held_items_proto_rawDescGZIP() []byte {
	file_unitedb_v1_held_items_proto_rawDescOnce.Do(func() {
		file_unitedb_v1_held_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_unitedb_v1_held_items_proto_rawDescData)
	})
	return file_unitedb_v1_held_items_proto_rawDescData
}

var file_unitedb_v1_held_items_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_unitedb_v1_held_items_proto_goTypes = []interface{}{
	(*HeldItem)(nil),              // 0: unitedb.v1.HeldItem
	(*GetHeldItemRequest)(nil),    // 1: unitedb.v1.GetHeldItemRequest
	(*GetHeldItemResponse)(nil),   // 2: unitedb.v1.GetHeldItemResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_unitedb_v1_held_items_proto_depIdxs = []int32{
	3, // 0: unitedb.v1.HeldItem.last_updated:type_name -> google.protobuf.Timestamp
	0, // 1: unitedb.v1.GetHeldItemResponse.data:type_name -> unitedb.v1.HeldItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_unitedb_v1_held_items_proto_init() }
func file_unitedb_v1_held_items_proto_init() {
	if File_unitedb_v1_held_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unitedb_v1_held_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeldItem); i {
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
		file_unitedb_v1_held_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeldItemRequest); i {
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
		file_unitedb_v1_held_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeldItemResponse); i {
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
	file_unitedb_v1_held_items_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_unitedb_v1_held_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_unitedb_v1_held_items_proto_goTypes,
		DependencyIndexes: file_unitedb_v1_held_items_proto_depIdxs,
		MessageInfos:      file_unitedb_v1_held_items_proto_msgTypes,
	}.Build()
	File_unitedb_v1_held_items_proto = out.File
	file_unitedb_v1_held_items_proto_rawDesc = nil
	file_unitedb_v1_held_items_proto_goTypes = nil
	file_unitedb_v1_held_items_proto_depIdxs = nil
}
