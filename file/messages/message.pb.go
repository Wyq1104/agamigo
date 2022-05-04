// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: message.proto

package messages

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

type ClimateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UTC_TIMESTAMP       int64   `protobuf:"varint,1,opt,name=UTC_TIMESTAMP,json=UTCTIMESTAMP,proto3" json:"UTC_TIMESTAMP,omitempty"`
	LONGITUDE           float32 `protobuf:"fixed32,2,opt,name=LONGITUDE,proto3" json:"LONGITUDE,omitempty"`
	LATITUDE            float32 `protobuf:"fixed32,3,opt,name=LATITUDE,proto3" json:"LATITUDE,omitempty"`
	AIR_TEMPERATURE     float32 `protobuf:"fixed32,4,opt,name=AIR_TEMPERATURE,json=AIRTEMPERATURE,proto3" json:"AIR_TEMPERATURE,omitempty"`
	PRECIPITATION       float32 `protobuf:"fixed32,5,opt,name=PRECIPITATION,proto3" json:"PRECIPITATION,omitempty"`
	SOLAR_RADIATION     float32 `protobuf:"fixed32,6,opt,name=SOLAR_RADIATION,json=SOLARRADIATION,proto3" json:"SOLAR_RADIATION,omitempty"`
	SURFACE_TEMPERATURE float32 `protobuf:"fixed32,7,opt,name=SURFACE_TEMPERATURE,json=SURFACETEMPERATURE,proto3" json:"SURFACE_TEMPERATURE,omitempty"`
	RELATIVE_HUMIDITY   int64   `protobuf:"varint,8,opt,name=RELATIVE_HUMIDITY,json=RELATIVEHUMIDITY,proto3" json:"RELATIVE_HUMIDITY,omitempty"`
}

func (x *ClimateData) Reset() {
	*x = ClimateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClimateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClimateData) ProtoMessage() {}

func (x *ClimateData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClimateData.ProtoReflect.Descriptor instead.
func (*ClimateData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *ClimateData) GetUTC_TIMESTAMP() int64 {
	if x != nil {
		return x.UTC_TIMESTAMP
	}
	return 0
}

func (x *ClimateData) GetLONGITUDE() float32 {
	if x != nil {
		return x.LONGITUDE
	}
	return 0
}

func (x *ClimateData) GetLATITUDE() float32 {
	if x != nil {
		return x.LATITUDE
	}
	return 0
}

func (x *ClimateData) GetAIR_TEMPERATURE() float32 {
	if x != nil {
		return x.AIR_TEMPERATURE
	}
	return 0
}

func (x *ClimateData) GetPRECIPITATION() float32 {
	if x != nil {
		return x.PRECIPITATION
	}
	return 0
}

func (x *ClimateData) GetSOLAR_RADIATION() float32 {
	if x != nil {
		return x.SOLAR_RADIATION
	}
	return 0
}

func (x *ClimateData) GetSURFACE_TEMPERATURE() float32 {
	if x != nil {
		return x.SURFACE_TEMPERATURE
	}
	return 0
}

func (x *ClimateData) GetRELATIVE_HUMIDITY() int64 {
	if x != nil {
		return x.RELATIVE_HUMIDITY
	}
	return 0
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*Wrapper_ClimateData
	Msg isWrapper_Msg `protobuf_oneof:"msg"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (m *Wrapper) GetMsg() isWrapper_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *Wrapper) GetClimateData() *ClimateData {
	if x, ok := x.GetMsg().(*Wrapper_ClimateData); ok {
		return x.ClimateData
	}
	return nil
}

type isWrapper_Msg interface {
	isWrapper_Msg()
}

type Wrapper_ClimateData struct {
	ClimateData *ClimateData `protobuf:"bytes,1,opt,name=climate_data,json=climateData,proto3,oneof"`
}

func (*Wrapper_ClimateData) isWrapper_Msg() {}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x02, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x23, 0x0a, 0x0d, 0x55, 0x54, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x55, 0x54, 0x43, 0x54, 0x49, 0x4d, 0x45, 0x53,
	0x54, 0x41, 0x4d, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x4f, 0x4e, 0x47, 0x49, 0x54, 0x55, 0x44,
	0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x4c, 0x4f, 0x4e, 0x47, 0x49, 0x54, 0x55,
	0x44, 0x45, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x41, 0x54, 0x49, 0x54, 0x55, 0x44, 0x45, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x4c, 0x41, 0x54, 0x49, 0x54, 0x55, 0x44, 0x45, 0x12, 0x27,
	0x0a, 0x0f, 0x41, 0x49, 0x52, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x41, 0x49, 0x52, 0x54, 0x45, 0x4d, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x52, 0x45, 0x43, 0x49,
	0x50, 0x49, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x50, 0x52, 0x45, 0x43, 0x49, 0x50, 0x49, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x27, 0x0a,
	0x0f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x52, 0x41, 0x44, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x52, 0x41, 0x44,
	0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x2f, 0x0a, 0x13, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43,
	0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x12, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x54, 0x45, 0x4d, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x12, 0x2b, 0x0a, 0x11, 0x52, 0x45, 0x4c, 0x41, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x48, 0x55, 0x4d, 0x49, 0x44, 0x49, 0x54, 0x59, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x48, 0x55, 0x4d, 0x49,
	0x44, 0x49, 0x54, 0x59, 0x22, 0x43, 0x0a, 0x07, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(*ClimateData)(nil), // 0: ClimateData
	(*Wrapper)(nil),     // 1: Wrapper
}
var file_message_proto_depIdxs = []int32{
	0, // 0: Wrapper.climate_data:type_name -> ClimateData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClimateData); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
	file_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Wrapper_ClimateData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
