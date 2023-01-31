// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/energy.proto

package proto

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

type PowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PowerResponse) Reset() {
	*x = PowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_energy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerResponse) ProtoMessage() {}

func (x *PowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_energy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerResponse.ProtoReflect.Descriptor instead.
func (*PowerResponse) Descriptor() ([]byte, []int) {
	return file_proto_energy_proto_rawDescGZIP(), []int{0}
}

func (x *PowerResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PowerConsumptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year           int64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	ResponseAmount int64 `protobuf:"varint,2,opt,name=responseAmount,proto3" json:"responseAmount,omitempty"`
}

func (x *PowerConsumptionRequest) Reset() {
	*x = PowerConsumptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_energy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerConsumptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerConsumptionRequest) ProtoMessage() {}

func (x *PowerConsumptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_energy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerConsumptionRequest.ProtoReflect.Descriptor instead.
func (*PowerConsumptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_energy_proto_rawDescGZIP(), []int{1}
}

func (x *PowerConsumptionRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *PowerConsumptionRequest) GetResponseAmount() int64 {
	if x != nil {
		return x.ResponseAmount
	}
	return 0
}

type PowerFromHomes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Period    string  `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	Year      int64   `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Unit      string  `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Precision int64   `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
	Character string  `protobuf:"bytes,6,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *PowerFromHomes) Reset() {
	*x = PowerFromHomes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_energy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerFromHomes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerFromHomes) ProtoMessage() {}

func (x *PowerFromHomes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_energy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerFromHomes.ProtoReflect.Descriptor instead.
func (*PowerFromHomes) Descriptor() ([]byte, []int) {
	return file_proto_energy_proto_rawDescGZIP(), []int{2}
}

func (x *PowerFromHomes) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PowerFromHomes) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *PowerFromHomes) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *PowerFromHomes) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *PowerFromHomes) GetPrecision() int64 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *PowerFromHomes) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_energy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_energy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_proto_energy_proto_rawDescGZIP(), []int{3}
}

type YearList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Years []string `protobuf:"bytes,1,rep,name=years,proto3" json:"years,omitempty"`
}

func (x *YearList) Reset() {
	*x = YearList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_energy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YearList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YearList) ProtoMessage() {}

func (x *YearList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_energy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YearList.ProtoReflect.Descriptor instead.
func (*YearList) Descriptor() ([]byte, []int) {
	return file_proto_energy_proto_rawDescGZIP(), []int{4}
}

func (x *YearList) GetYears() []string {
	if x != nil {
		return x.Years
	}
	return nil
}

var File_proto_energy_proto protoreflect.FileDescriptor

var file_proto_energy_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x17, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x6f,
	0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x22, 0x20, 0x0a, 0x08, 0x59, 0x65, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x79, 0x65, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x79, 0x65, 0x61,
	0x72, 0x73, 0x32, 0xc2, 0x01, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x61, 0x72, 0x45,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e, 0x73,
	0x6f, 0x6c, 0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x25,
	0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x6f,
	0x6d, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_energy_proto_rawDescOnce sync.Once
	file_proto_energy_proto_rawDescData = file_proto_energy_proto_rawDesc
)

func file_proto_energy_proto_rawDescGZIP() []byte {
	file_proto_energy_proto_rawDescOnce.Do(func() {
		file_proto_energy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_energy_proto_rawDescData)
	})
	return file_proto_energy_proto_rawDescData
}

var file_proto_energy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_energy_proto_goTypes = []interface{}{
	(*PowerResponse)(nil),           // 0: solarservice.PowerResponse
	(*PowerConsumptionRequest)(nil), // 1: solarservice.PowerConsumptionRequest
	(*PowerFromHomes)(nil),          // 2: solarservice.PowerFromHomes
	(*NoParam)(nil),                 // 3: solarservice.NoParam
	(*YearList)(nil),                // 4: solarservice.YearList
}
var file_proto_energy_proto_depIdxs = []int32{
	3, // 0: solarservice.SolarService.GetSolarEnergy:input_type -> solarservice.NoParam
	1, // 1: solarservice.SolarService.GetSolarEnergyFromHomesByParams:input_type -> solarservice.PowerConsumptionRequest
	0, // 2: solarservice.SolarService.GetSolarEnergy:output_type -> solarservice.PowerResponse
	2, // 3: solarservice.SolarService.GetSolarEnergyFromHomesByParams:output_type -> solarservice.PowerFromHomes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_energy_proto_init() }
func file_proto_energy_proto_init() {
	if File_proto_energy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_energy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerResponse); i {
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
		file_proto_energy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerConsumptionRequest); i {
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
		file_proto_energy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerFromHomes); i {
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
		file_proto_energy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParam); i {
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
		file_proto_energy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YearList); i {
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
			RawDescriptor: file_proto_energy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_energy_proto_goTypes,
		DependencyIndexes: file_proto_energy_proto_depIdxs,
		MessageInfos:      file_proto_energy_proto_msgTypes,
	}.Build()
	File_proto_energy_proto = out.File
	file_proto_energy_proto_rawDesc = nil
	file_proto_energy_proto_goTypes = nil
	file_proto_energy_proto_depIdxs = nil
}