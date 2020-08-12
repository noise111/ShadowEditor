// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: point_feature.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PointFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bbox       []float64                `protobuf:"fixed64,2,rep,packed,name=bbox,proto3" json:"bbox,omitempty"`
	Geometry   *PointGeometry           `protobuf:"bytes,3,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Properties map[string]*VariantValue `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PointFeature) Reset() {
	*x = PointFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointFeature) ProtoMessage() {}

func (x *PointFeature) ProtoReflect() protoreflect.Message {
	mi := &file_point_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointFeature.ProtoReflect.Descriptor instead.
func (*PointFeature) Descriptor() ([]byte, []int) {
	return file_point_feature_proto_rawDescGZIP(), []int{0}
}

func (x *PointFeature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PointFeature) GetBbox() []float64 {
	if x != nil {
		return x.Bbox
	}
	return nil
}

func (x *PointFeature) GetGeometry() *PointGeometry {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *PointFeature) GetProperties() map[string]*VariantValue {
	if x != nil {
		return x.Properties
	}
	return nil
}

type PointGeometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,2,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *PointGeometry) Reset() {
	*x = PointGeometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointGeometry) ProtoMessage() {}

func (x *PointGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_point_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointGeometry.ProtoReflect.Descriptor instead.
func (*PointGeometry) Descriptor() ([]byte, []int) {
	return file_point_feature_proto_rawDescGZIP(), []int{1}
}

func (x *PointGeometry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PointGeometry) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_point_feature_proto protoreflect.FileDescriptor

var file_point_feature_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4,
	0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x04, 0x62,
	0x62, 0x6f, 0x78, 0x12, 0x2d, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x1a, 0x4f, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x0d, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_point_feature_proto_rawDescOnce sync.Once
	file_point_feature_proto_rawDescData = file_point_feature_proto_rawDesc
)

func file_point_feature_proto_rawDescGZIP() []byte {
	file_point_feature_proto_rawDescOnce.Do(func() {
		file_point_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_point_feature_proto_rawDescData)
	})
	return file_point_feature_proto_rawDescData
}

var file_point_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_point_feature_proto_goTypes = []interface{}{
	(*PointFeature)(nil),  // 0: pb.PointFeature
	(*PointGeometry)(nil), // 1: pb.PointGeometry
	nil,                   // 2: pb.PointFeature.PropertiesEntry
	(*VariantValue)(nil),  // 3: pb.VariantValue
}
var file_point_feature_proto_depIdxs = []int32{
	1, // 0: pb.PointFeature.geometry:type_name -> pb.PointGeometry
	2, // 1: pb.PointFeature.properties:type_name -> pb.PointFeature.PropertiesEntry
	3, // 2: pb.PointFeature.PropertiesEntry.value:type_name -> pb.VariantValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_point_feature_proto_init() }
func file_point_feature_proto_init() {
	if File_point_feature_proto != nil {
		return
	}
	file_variant_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_point_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointFeature); i {
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
		file_point_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointGeometry); i {
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
			RawDescriptor: file_point_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_point_feature_proto_goTypes,
		DependencyIndexes: file_point_feature_proto_depIdxs,
		MessageInfos:      file_point_feature_proto_msgTypes,
	}.Build()
	File_point_feature_proto = out.File
	file_point_feature_proto_rawDesc = nil
	file_point_feature_proto_goTypes = nil
	file_point_feature_proto_depIdxs = nil
}