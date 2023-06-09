// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: area.proto

package server

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

type Level int32

const (
	// 未知
	Level_UNKNOW Level = 0
	// 国家
	Level_COUNTRY Level = 1
	// 省份（直辖市会在province显示）
	Level_PROVINCE Level = 2
	// 市（直辖市会在province显示）
	Level_CITY Level = 3
	// 区县
	Level_DISTRICT Level = 4
	// 街道
	Level_STREET Level = 5
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "UNKNOW",
		1: "COUNTRY",
		2: "PROVINCE",
		3: "CITY",
		4: "DISTRICT",
		5: "STREET",
	}
	Level_value = map[string]int32{
		"UNKNOW":   0,
		"COUNTRY":  1,
		"PROVINCE": 2,
		"CITY":     3,
		"DISTRICT": 4,
		"STREET":   5,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_area_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_area_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[0]
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
	return file_area_proto_rawDescGZIP(), []int{0}
}

type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 行政区名称
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// 行政区划级别
	Level Level `protobuf:"varint,2,opt,name=Level,proto3,enum=area.Level" json:"Level,omitempty"`
	// 区域编码
	Code string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	// 区域中心点
	Center string `protobuf:"bytes,4,opt,name=Center,proto3" json:"Center,omitempty"`
	// 下级行政区列表
	Areas []*Area `protobuf:"bytes,5,rep,name=Areas,proto3" json:"Areas,omitempty"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{1}
}

func (x *Area) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Area) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_UNKNOW
}

func (x *Area) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Area) GetCenter() string {
	if x != nil {
		return x.Center
	}
	return ""
}

func (x *Area) GetAreas() []*Area {
	if x != nil {
		return x.Areas
	}
	return nil
}

type GetAreaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区域编码
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *GetAreaReq) Reset() {
	*x = GetAreaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAreaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAreaReq) ProtoMessage() {}

func (x *GetAreaReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAreaReq.ProtoReflect.Descriptor instead.
func (*GetAreaReq) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{2}
}

func (x *GetAreaReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetAreaAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区域信息
	Area *Area `protobuf:"bytes,1,opt,name=Area,proto3" json:"Area,omitempty"`
}

func (x *GetAreaAck) Reset() {
	*x = GetAreaAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAreaAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAreaAck) ProtoMessage() {}

func (x *GetAreaAck) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAreaAck.ProtoReflect.Descriptor instead.
func (*GetAreaAck) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{3}
}

func (x *GetAreaAck) GetArea() *Area {
	if x != nil {
		return x.Area
	}
	return nil
}

type GetProvincesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProvincesReq) Reset() {
	*x = GetProvincesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvincesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvincesReq) ProtoMessage() {}

func (x *GetProvincesReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvincesReq.ProtoReflect.Descriptor instead.
func (*GetProvincesReq) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{4}
}

type GetProvincesAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 省份列表
	Provinces []*Area `protobuf:"bytes,1,rep,name=Provinces,proto3" json:"Provinces,omitempty"`
}

func (x *GetProvincesAck) Reset() {
	*x = GetProvincesAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvincesAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvincesAck) ProtoMessage() {}

func (x *GetProvincesAck) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvincesAck.ProtoReflect.Descriptor instead.
func (*GetProvincesAck) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{5}
}

func (x *GetProvincesAck) GetProvinces() []*Area {
	if x != nil {
		return x.Provinces
	}
	return nil
}

type GetCitiesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 省份区域编码
	ProvinceCode string `protobuf:"bytes,1,opt,name=ProvinceCode,proto3" json:"ProvinceCode,omitempty"`
}

func (x *GetCitiesReq) Reset() {
	*x = GetCitiesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesReq) ProtoMessage() {}

func (x *GetCitiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesReq.ProtoReflect.Descriptor instead.
func (*GetCitiesReq) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{6}
}

func (x *GetCitiesReq) GetProvinceCode() string {
	if x != nil {
		return x.ProvinceCode
	}
	return ""
}

type GetCitiesAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 城市列表
	Cities []*Area `protobuf:"bytes,1,rep,name=Cities,proto3" json:"Cities,omitempty"`
}

func (x *GetCitiesAck) Reset() {
	*x = GetCitiesAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCitiesAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCitiesAck) ProtoMessage() {}

func (x *GetCitiesAck) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCitiesAck.ProtoReflect.Descriptor instead.
func (*GetCitiesAck) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{7}
}

func (x *GetCitiesAck) GetCities() []*Area {
	if x != nil {
		return x.Cities
	}
	return nil
}

type GetDistrictsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 城市区域编码
	CityCode string `protobuf:"bytes,1,opt,name=CityCode,proto3" json:"CityCode,omitempty"`
}

func (x *GetDistrictsReq) Reset() {
	*x = GetDistrictsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictsReq) ProtoMessage() {}

func (x *GetDistrictsReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictsReq.ProtoReflect.Descriptor instead.
func (*GetDistrictsReq) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{8}
}

func (x *GetDistrictsReq) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

type GetDistrictsAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区县列表
	Districts []*Area `protobuf:"bytes,1,rep,name=Districts,proto3" json:"Districts,omitempty"`
}

func (x *GetDistrictsAck) Reset() {
	*x = GetDistrictsAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDistrictsAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDistrictsAck) ProtoMessage() {}

func (x *GetDistrictsAck) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDistrictsAck.ProtoReflect.Descriptor instead.
func (*GetDistrictsAck) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{9}
}

func (x *GetDistrictsAck) GetDistricts() []*Area {
	if x != nil {
		return x.Districts
	}
	return nil
}

type GetStreetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区县区域编码
	DistrictCode string `protobuf:"bytes,1,opt,name=DistrictCode,proto3" json:"DistrictCode,omitempty"`
}

func (x *GetStreetsReq) Reset() {
	*x = GetStreetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreetsReq) ProtoMessage() {}

func (x *GetStreetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreetsReq.ProtoReflect.Descriptor instead.
func (*GetStreetsReq) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{10}
}

func (x *GetStreetsReq) GetDistrictCode() string {
	if x != nil {
		return x.DistrictCode
	}
	return ""
}

type GetStreetsAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 街道列表
	Streets []*Area `protobuf:"bytes,1,rep,name=Streets,proto3" json:"Streets,omitempty"`
}

func (x *GetStreetsAck) Reset() {
	*x = GetStreetsAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreetsAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreetsAck) ProtoMessage() {}

func (x *GetStreetsAck) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreetsAck.ProtoReflect.Descriptor instead.
func (*GetStreetsAck) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{11}
}

func (x *GetStreetsAck) GetStreets() []*Area {
	if x != nil {
		return x.Streets
	}
	return nil
}

var File_area_proto protoreflect.FileDescriptor

var file_area_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x8b, 0x01, 0x0a, 0x04,
	0x41, 0x72, 0x65, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x61, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x05, 0x41, 0x72, 0x65, 0x61, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x41, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x04, 0x41, 0x72, 0x65,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x22, 0x3b, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x12,
	0x28, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x09,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x12, 0x22, 0x0a,
	0x06, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x06, 0x43, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73,
	0x41, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x09, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x22, 0x33, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x22,
	0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73,
	0x41, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73, 0x2a, 0x52, 0x0a, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x49, 0x54,
	0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x45, 0x45, 0x54, 0x10, 0x05, 0x32, 0xd9, 0x02,
	0x0a, 0x0b, 0x41, 0x72, 0x65, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x0a, 0x46, 0x72, 0x65, 0x73, 0x68, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x0b, 0x2e, 0x61, 0x72,
	0x65, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x65, 0x61, 0x12, 0x10, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x65,
	0x61, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x65, 0x61, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x72, 0x65, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12,
	0x15, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73, 0x12, 0x13, 0x2e,
	0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x73, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_area_proto_rawDescOnce sync.Once
	file_area_proto_rawDescData = file_area_proto_rawDesc
)

func file_area_proto_rawDescGZIP() []byte {
	file_area_proto_rawDescOnce.Do(func() {
		file_area_proto_rawDescData = protoimpl.X.CompressGZIP(file_area_proto_rawDescData)
	})
	return file_area_proto_rawDescData
}

var file_area_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_area_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_area_proto_goTypes = []interface{}{
	(Level)(0),              // 0: area.Level
	(*Empty)(nil),           // 1: area.Empty
	(*Area)(nil),            // 2: area.Area
	(*GetAreaReq)(nil),      // 3: area.GetAreaReq
	(*GetAreaAck)(nil),      // 4: area.GetAreaAck
	(*GetProvincesReq)(nil), // 5: area.GetProvincesReq
	(*GetProvincesAck)(nil), // 6: area.GetProvincesAck
	(*GetCitiesReq)(nil),    // 7: area.GetCitiesReq
	(*GetCitiesAck)(nil),    // 8: area.GetCitiesAck
	(*GetDistrictsReq)(nil), // 9: area.GetDistrictsReq
	(*GetDistrictsAck)(nil), // 10: area.GetDistrictsAck
	(*GetStreetsReq)(nil),   // 11: area.GetStreetsReq
	(*GetStreetsAck)(nil),   // 12: area.GetStreetsAck
}
var file_area_proto_depIdxs = []int32{
	0,  // 0: area.Area.Level:type_name -> area.Level
	2,  // 1: area.Area.Areas:type_name -> area.Area
	2,  // 2: area.GetAreaAck.Area:type_name -> area.Area
	2,  // 3: area.GetProvincesAck.Provinces:type_name -> area.Area
	2,  // 4: area.GetCitiesAck.Cities:type_name -> area.Area
	2,  // 5: area.GetDistrictsAck.Districts:type_name -> area.Area
	2,  // 6: area.GetStreetsAck.Streets:type_name -> area.Area
	1,  // 7: area.AreaService.FreshAreas:input_type -> area.Empty
	3,  // 8: area.AreaService.GetArea:input_type -> area.GetAreaReq
	5,  // 9: area.AreaService.GetProvinces:input_type -> area.GetProvincesReq
	7,  // 10: area.AreaService.GetCities:input_type -> area.GetCitiesReq
	9,  // 11: area.AreaService.GetDistricts:input_type -> area.GetDistrictsReq
	11, // 12: area.AreaService.GetStreets:input_type -> area.GetStreetsReq
	1,  // 13: area.AreaService.FreshAreas:output_type -> area.Empty
	4,  // 14: area.AreaService.GetArea:output_type -> area.GetAreaAck
	6,  // 15: area.AreaService.GetProvinces:output_type -> area.GetProvincesAck
	8,  // 16: area.AreaService.GetCities:output_type -> area.GetCitiesAck
	10, // 17: area.AreaService.GetDistricts:output_type -> area.GetDistrictsAck
	12, // 18: area.AreaService.GetStreets:output_type -> area.GetStreetsAck
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_area_proto_init() }
func file_area_proto_init() {
	if File_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_area_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		file_area_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAreaReq); i {
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
		file_area_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAreaAck); i {
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
		file_area_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvincesReq); i {
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
		file_area_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvincesAck); i {
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
		file_area_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCitiesReq); i {
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
		file_area_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCitiesAck); i {
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
		file_area_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictsReq); i {
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
		file_area_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDistrictsAck); i {
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
		file_area_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreetsReq); i {
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
		file_area_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreetsAck); i {
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
			RawDescriptor: file_area_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_area_proto_goTypes,
		DependencyIndexes: file_area_proto_depIdxs,
		EnumInfos:         file_area_proto_enumTypes,
		MessageInfos:      file_area_proto_msgTypes,
	}.Build()
	File_area_proto = out.File
	file_area_proto_rawDesc = nil
	file_area_proto_goTypes = nil
	file_area_proto_depIdxs = nil
}
