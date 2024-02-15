// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/loadbalancerconfig.proto

package nbcontractv1

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

type LoadBalancerSku int32

const (
	LoadBalancerSku_LOAD_BALANCER_SKU_UNSPECIFIED LoadBalancerSku = 0
	LoadBalancerSku_LOAD_BALANCER_SKU_BASIC       LoadBalancerSku = 1 // to be confirmed
	LoadBalancerSku_LOAD_BALANCER_SKU_STANDARD    LoadBalancerSku = 2
)

// Enum value maps for LoadBalancerSku.
var (
	LoadBalancerSku_name = map[int32]string{
		0: "LOAD_BALANCER_SKU_UNSPECIFIED",
		1: "LOAD_BALANCER_SKU_BASIC",
		2: "LOAD_BALANCER_SKU_STANDARD",
	}
	LoadBalancerSku_value = map[string]int32{
		"LOAD_BALANCER_SKU_UNSPECIFIED": 0,
		"LOAD_BALANCER_SKU_BASIC":       1,
		"LOAD_BALANCER_SKU_STANDARD":    2,
	}
)

func (x LoadBalancerSku) Enum() *LoadBalancerSku {
	p := new(LoadBalancerSku)
	*p = x
	return p
}

func (x LoadBalancerSku) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadBalancerSku) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_enumTypes[0].Descriptor()
}

func (LoadBalancerSku) Type() protoreflect.EnumType {
	return &file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_enumTypes[0]
}

func (x LoadBalancerSku) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadBalancerSku.Descriptor instead.
func (LoadBalancerSku) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescGZIP(), []int{0}
}

type LoadBalancerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadBalancerSku                       LoadBalancerSku `protobuf:"varint,1,opt,name=load_balancer_sku,json=loadBalancerSku,proto3,enum=nbcontract.v1.LoadBalancerSku" json:"load_balancer_sku,omitempty"`
	ExcludeMasterFromStandardLoadBalancer bool            `protobuf:"varint,2,opt,name=exclude_master_from_standard_load_balancer,json=excludeMasterFromStandardLoadBalancer,proto3" json:"exclude_master_from_standard_load_balancer,omitempty"`
	MaxLoadBalancerRuleCount              int32           `protobuf:"varint,3,opt,name=max_load_balancer_rule_count,json=maxLoadBalancerRuleCount,proto3" json:"max_load_balancer_rule_count,omitempty"`
}

func (x *LoadBalancerConfig) Reset() {
	*x = LoadBalancerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig) ProtoMessage() {}

func (x *LoadBalancerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescGZIP(), []int{0}
}

func (x *LoadBalancerConfig) GetLoadBalancerSku() LoadBalancerSku {
	if x != nil {
		return x.LoadBalancerSku
	}
	return LoadBalancerSku_LOAD_BALANCER_SKU_UNSPECIFIED
}

func (x *LoadBalancerConfig) GetExcludeMasterFromStandardLoadBalancer() bool {
	if x != nil {
		return x.ExcludeMasterFromStandardLoadBalancer
	}
	return false
}

func (x *LoadBalancerConfig) GetMaxLoadBalancerRuleCount() int32 {
	if x != nil {
		return x.MaxLoadBalancerRuleCount
	}
	return 0
}

var File_pkg_proto_nbcontract_v1_loadbalancerconfig_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0xfb, 0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x11, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x53, 0x6b, 0x75, 0x52, 0x0f, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x53, 0x6b, 0x75, 0x12, 0x59, 0x0a, 0x2a, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x25, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12,
	0x3e, 0x0a, 0x1c, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a,
	0x71, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53,
	0x6b, 0x75, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e,
	0x43, 0x45, 0x52, 0x5f, 0x53, 0x4b, 0x55, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41,
	0x4c, 0x41, 0x4e, 0x43, 0x45, 0x52, 0x5f, 0x53, 0x4b, 0x55, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e,
	0x43, 0x45, 0x52, 0x5f, 0x53, 0x4b, 0x55, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44,
	0x10, 0x02, 0x42, 0xc3, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x4c, 0x6f, 0x61, 0x64, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e, 0x62,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescData = file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_goTypes = []interface{}{
	(LoadBalancerSku)(0),       // 0: nbcontract.v1.LoadBalancerSku
	(*LoadBalancerConfig)(nil), // 1: nbcontract.v1.LoadBalancerConfig
}
var file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_depIdxs = []int32{
	0, // 0: nbcontract.v1.LoadBalancerConfig.load_balancer_sku:type_name -> nbcontract.v1.LoadBalancerSku
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_init() }
func file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_init() {
	if File_pkg_proto_nbcontract_v1_loadbalancerconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig); i {
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
			RawDescriptor: file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_depIdxs,
		EnumInfos:         file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_enumTypes,
		MessageInfos:      file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_loadbalancerconfig_proto = out.File
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_loadbalancerconfig_proto_depIdxs = nil
}
