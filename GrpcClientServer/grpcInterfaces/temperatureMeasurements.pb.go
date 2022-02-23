// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: grpcInterfaces/temperatureMeasurements.proto

package grpcInterfaces

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration string `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_grpcInterfaces_temperatureMeasurements_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorId    int32 `protobuf:"varint,1,opt,name=sensorId,proto3" json:"sensorId,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_grpcInterfaces_temperatureMeasurements_proto_rawDescGZIP(), []int{1}
}

func (x *SetRequest) GetSensorId() int32 {
	if x != nil {
		return x.SensorId
	}
	return 0
}

func (x *SetRequest) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type MeasurementsDataResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MeasurementsDataResult) Reset() {
	*x = MeasurementsDataResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementsDataResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementsDataResult) ProtoMessage() {}

func (x *MeasurementsDataResult) ProtoReflect() protoreflect.Message {
	mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementsDataResult.ProtoReflect.Descriptor instead.
func (*MeasurementsDataResult) Descriptor() ([]byte, []int) {
	return file_grpcInterfaces_temperatureMeasurements_proto_rawDescGZIP(), []int{2}
}

func (x *MeasurementsDataResult) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpcInterfaces_temperatureMeasurements_proto_rawDescGZIP(), []int{3}
}

var File_grpcInterfaces_temperatureMeasurements_proto protoreflect.FileDescriptor

var file_grpcInterfaces_temperatureMeasurements_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x72, 0x70, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x8c, 0x01,
	0x0a, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x12, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcInterfaces_temperatureMeasurements_proto_rawDescOnce sync.Once
	file_grpcInterfaces_temperatureMeasurements_proto_rawDescData = file_grpcInterfaces_temperatureMeasurements_proto_rawDesc
)

func file_grpcInterfaces_temperatureMeasurements_proto_rawDescGZIP() []byte {
	file_grpcInterfaces_temperatureMeasurements_proto_rawDescOnce.Do(func() {
		file_grpcInterfaces_temperatureMeasurements_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcInterfaces_temperatureMeasurements_proto_rawDescData)
	})
	return file_grpcInterfaces_temperatureMeasurements_proto_rawDescData
}

var file_grpcInterfaces_temperatureMeasurements_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpcInterfaces_temperatureMeasurements_proto_goTypes = []interface{}{
	(*GetRequest)(nil),             // 0: GetRequest
	(*SetRequest)(nil),             // 1: SetRequest
	(*MeasurementsDataResult)(nil), // 2: MeasurementsDataResult
	(*Empty)(nil),                  // 3: Empty
}
var file_grpcInterfaces_temperatureMeasurements_proto_depIdxs = []int32{
	0, // 0: TemperatureMeasurementsService.getMeasurementsData:input_type -> GetRequest
	1, // 1: TemperatureMeasurementsService.setMeasurementData:input_type -> SetRequest
	2, // 2: TemperatureMeasurementsService.getMeasurementsData:output_type -> MeasurementsDataResult
	3, // 3: TemperatureMeasurementsService.setMeasurementData:output_type -> Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcInterfaces_temperatureMeasurements_proto_init() }
func file_grpcInterfaces_temperatureMeasurements_proto_init() {
	if File_grpcInterfaces_temperatureMeasurements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasurementsDataResult); i {
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
		file_grpcInterfaces_temperatureMeasurements_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpcInterfaces_temperatureMeasurements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcInterfaces_temperatureMeasurements_proto_goTypes,
		DependencyIndexes: file_grpcInterfaces_temperatureMeasurements_proto_depIdxs,
		MessageInfos:      file_grpcInterfaces_temperatureMeasurements_proto_msgTypes,
	}.Build()
	File_grpcInterfaces_temperatureMeasurements_proto = out.File
	file_grpcInterfaces_temperatureMeasurements_proto_rawDesc = nil
	file_grpcInterfaces_temperatureMeasurements_proto_goTypes = nil
	file_grpcInterfaces_temperatureMeasurements_proto_depIdxs = nil
}
