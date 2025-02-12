// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: geofence.proto

package geofence

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

// Geofence types.
type FenceType int32

const (
	FenceType_FENCE_TYPE_INCLUSION FenceType = 0 // Type representing an inclusion fence
	FenceType_FENCE_TYPE_EXCLUSION FenceType = 1 // Type representing an exclusion fence
)

// Enum value maps for FenceType.
var (
	FenceType_name = map[int32]string{
		0: "FENCE_TYPE_INCLUSION",
		1: "FENCE_TYPE_EXCLUSION",
	}
	FenceType_value = map[string]int32{
		"FENCE_TYPE_INCLUSION": 0,
		"FENCE_TYPE_EXCLUSION": 1,
	}
)

func (x FenceType) Enum() *FenceType {
	p := new(FenceType)
	*p = x
	return p
}

func (x FenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_geofence_proto_enumTypes[0].Descriptor()
}

func (FenceType) Type() protoreflect.EnumType {
	return &file_geofence_proto_enumTypes[0]
}

func (x FenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FenceType.Descriptor instead.
func (FenceType) EnumDescriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

// Possible results returned for geofence requests.
type GeofenceResult_Result int32

const (
	GeofenceResult_RESULT_UNKNOWN                 GeofenceResult_Result = 0 // Unknown result
	GeofenceResult_RESULT_SUCCESS                 GeofenceResult_Result = 1 // Request succeeded
	GeofenceResult_RESULT_ERROR                   GeofenceResult_Result = 2 // Error
	GeofenceResult_RESULT_TOO_MANY_GEOFENCE_ITEMS GeofenceResult_Result = 3 // Too many objects in the geofence
	GeofenceResult_RESULT_BUSY                    GeofenceResult_Result = 4 // Vehicle is busy
	GeofenceResult_RESULT_TIMEOUT                 GeofenceResult_Result = 5 // Request timed out
	GeofenceResult_RESULT_INVALID_ARGUMENT        GeofenceResult_Result = 6 // Invalid argument
	GeofenceResult_RESULT_NO_SYSTEM               GeofenceResult_Result = 7 // No system connected
)

// Enum value maps for GeofenceResult_Result.
var (
	GeofenceResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_ERROR",
		3: "RESULT_TOO_MANY_GEOFENCE_ITEMS",
		4: "RESULT_BUSY",
		5: "RESULT_TIMEOUT",
		6: "RESULT_INVALID_ARGUMENT",
		7: "RESULT_NO_SYSTEM",
	}
	GeofenceResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":                 0,
		"RESULT_SUCCESS":                 1,
		"RESULT_ERROR":                   2,
		"RESULT_TOO_MANY_GEOFENCE_ITEMS": 3,
		"RESULT_BUSY":                    4,
		"RESULT_TIMEOUT":                 5,
		"RESULT_INVALID_ARGUMENT":        6,
		"RESULT_NO_SYSTEM":               7,
	}
)

func (x GeofenceResult_Result) Enum() *GeofenceResult_Result {
	p := new(GeofenceResult_Result)
	*p = x
	return p
}

func (x GeofenceResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeofenceResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_geofence_proto_enumTypes[1].Descriptor()
}

func (GeofenceResult_Result) Type() protoreflect.EnumType {
	return &file_geofence_proto_enumTypes[1]
}

func (x GeofenceResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeofenceResult_Result.Descriptor instead.
func (GeofenceResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{8, 0}
}

// Point type.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatitudeDeg  float64 `protobuf:"fixed64,1,opt,name=latitude_deg,json=latitudeDeg,proto3" json:"latitude_deg,omitempty"`    // Latitude in degrees (range: -90 to +90)
	LongitudeDeg float64 `protobuf:"fixed64,2,opt,name=longitude_deg,json=longitudeDeg,proto3" json:"longitude_deg,omitempty"` // Longitude in degrees (range: -180 to +180)
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitudeDeg() float64 {
	if x != nil {
		return x.LatitudeDeg
	}
	return 0
}

func (x *Point) GetLongitudeDeg() float64 {
	if x != nil {
		return x.LongitudeDeg
	}
	return 0
}

// Polygon type.
type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points    []*Point  `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`                                                            // Points defining the polygon
	FenceType FenceType `protobuf:"varint,2,opt,name=fence_type,json=fenceType,proto3,enum=mavsdk.rpc.geofence.FenceType" json:"fence_type,omitempty"` // Fence type
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{1}
}

func (x *Polygon) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *Polygon) GetFenceType() FenceType {
	if x != nil {
		return x.FenceType
	}
	return FenceType_FENCE_TYPE_INCLUSION
}

// Circular type.
type Circle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point     *Point    `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`                                                              // Point defining the center
	Radius    float32   `protobuf:"fixed32,2,opt,name=radius,proto3" json:"radius,omitempty"`                                                          // Radius of the circular fence
	FenceType FenceType `protobuf:"varint,3,opt,name=fence_type,json=fenceType,proto3,enum=mavsdk.rpc.geofence.FenceType" json:"fence_type,omitempty"` // Fence type
}

func (x *Circle) Reset() {
	*x = Circle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{2}
}

func (x *Circle) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *Circle) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *Circle) GetFenceType() FenceType {
	if x != nil {
		return x.FenceType
	}
	return FenceType_FENCE_TYPE_INCLUSION
}

// Geofence data type.
type GeofenceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygons []*Polygon `protobuf:"bytes,1,rep,name=polygons,proto3" json:"polygons,omitempty"` // Polygon(s) representing the geofence(s)
	Circles  []*Circle  `protobuf:"bytes,2,rep,name=circles,proto3" json:"circles,omitempty"`   // Circle(s) representing the geofence(s)
}

func (x *GeofenceData) Reset() {
	*x = GeofenceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeofenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeofenceData) ProtoMessage() {}

func (x *GeofenceData) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeofenceData.ProtoReflect.Descriptor instead.
func (*GeofenceData) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{3}
}

func (x *GeofenceData) GetPolygons() []*Polygon {
	if x != nil {
		return x.Polygons
	}
	return nil
}

func (x *GeofenceData) GetCircles() []*Circle {
	if x != nil {
		return x.Circles
	}
	return nil
}

type UploadGeofenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeofenceData *GeofenceData `protobuf:"bytes,1,opt,name=geofence_data,json=geofenceData,proto3" json:"geofence_data,omitempty"` // Circle(s) and/or Polygon(s) representing the geofence(s)
}

func (x *UploadGeofenceRequest) Reset() {
	*x = UploadGeofenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadGeofenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadGeofenceRequest) ProtoMessage() {}

func (x *UploadGeofenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadGeofenceRequest.ProtoReflect.Descriptor instead.
func (*UploadGeofenceRequest) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{4}
}

func (x *UploadGeofenceRequest) GetGeofenceData() *GeofenceData {
	if x != nil {
		return x.GeofenceData
	}
	return nil
}

type UploadGeofenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeofenceResult *GeofenceResult `protobuf:"bytes,1,opt,name=geofence_result,json=geofenceResult,proto3" json:"geofence_result,omitempty"`
}

func (x *UploadGeofenceResponse) Reset() {
	*x = UploadGeofenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadGeofenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadGeofenceResponse) ProtoMessage() {}

func (x *UploadGeofenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadGeofenceResponse.ProtoReflect.Descriptor instead.
func (*UploadGeofenceResponse) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{5}
}

func (x *UploadGeofenceResponse) GetGeofenceResult() *GeofenceResult {
	if x != nil {
		return x.GeofenceResult
	}
	return nil
}

type ClearGeofenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearGeofenceRequest) Reset() {
	*x = ClearGeofenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearGeofenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearGeofenceRequest) ProtoMessage() {}

func (x *ClearGeofenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearGeofenceRequest.ProtoReflect.Descriptor instead.
func (*ClearGeofenceRequest) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{6}
}

type ClearGeofenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeofenceResult *GeofenceResult `protobuf:"bytes,1,opt,name=geofence_result,json=geofenceResult,proto3" json:"geofence_result,omitempty"`
}

func (x *ClearGeofenceResponse) Reset() {
	*x = ClearGeofenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearGeofenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearGeofenceResponse) ProtoMessage() {}

func (x *ClearGeofenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearGeofenceResponse.ProtoReflect.Descriptor instead.
func (*ClearGeofenceResponse) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{7}
}

func (x *ClearGeofenceResponse) GetGeofenceResult() *GeofenceResult {
	if x != nil {
		return x.GeofenceResult
	}
	return nil
}

// Result type.
type GeofenceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    GeofenceResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.geofence.GeofenceResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string                `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                          // Human-readable English string describing the result
}

func (x *GeofenceResult) Reset() {
	*x = GeofenceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeofenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeofenceResult) ProtoMessage() {}

func (x *GeofenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeofenceResult.ProtoReflect.Descriptor instead.
func (*GeofenceResult) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{8}
}

func (x *GeofenceResult) GetResult() GeofenceResult_Result {
	if x != nil {
		return x.Result
	}
	return GeofenceResult_RESULT_UNKNOWN
}

func (x *GeofenceResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_geofence_proto protoreflect.FileDescriptor

var file_geofence_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x05, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x5f, 0x64, 0x65, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x44, 0x65, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x44, 0x65, 0x67, 0x22, 0x7c, 0x0a, 0x07,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x43,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42, 0x07, 0x82, 0xb5, 0x18, 0x03, 0x4e, 0x61, 0x4e,
	0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7f, 0x0a, 0x0c, 0x47, 0x65, 0x6f, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x76, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x73, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52,
	0x07, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x46, 0x0a, 0x0d, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x67, 0x65, 0x6f,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x16, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61,
	0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0xb4, 0x02, 0x0a, 0x0e, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x22, 0xbe, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x47,
	0x45, 0x4f, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x53, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x07, 0x2a, 0x3f, 0x0a, 0x09, 0x46, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x43,
	0x4c, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x32, 0xe8, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x6f,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a,
	0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x65, 0x6f, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x61, 0x76,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x6d, 0x61, 0x76,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1b, 0x42, 0x0d, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0a, 0x2e, 0x3b, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geofence_proto_rawDescOnce sync.Once
	file_geofence_proto_rawDescData = file_geofence_proto_rawDesc
)

func file_geofence_proto_rawDescGZIP() []byte {
	file_geofence_proto_rawDescOnce.Do(func() {
		file_geofence_proto_rawDescData = protoimpl.X.CompressGZIP(file_geofence_proto_rawDescData)
	})
	return file_geofence_proto_rawDescData
}

var file_geofence_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_geofence_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_geofence_proto_goTypes = []interface{}{
	(FenceType)(0),                 // 0: mavsdk.rpc.geofence.FenceType
	(GeofenceResult_Result)(0),     // 1: mavsdk.rpc.geofence.GeofenceResult.Result
	(*Point)(nil),                  // 2: mavsdk.rpc.geofence.Point
	(*Polygon)(nil),                // 3: mavsdk.rpc.geofence.Polygon
	(*Circle)(nil),                 // 4: mavsdk.rpc.geofence.Circle
	(*GeofenceData)(nil),           // 5: mavsdk.rpc.geofence.GeofenceData
	(*UploadGeofenceRequest)(nil),  // 6: mavsdk.rpc.geofence.UploadGeofenceRequest
	(*UploadGeofenceResponse)(nil), // 7: mavsdk.rpc.geofence.UploadGeofenceResponse
	(*ClearGeofenceRequest)(nil),   // 8: mavsdk.rpc.geofence.ClearGeofenceRequest
	(*ClearGeofenceResponse)(nil),  // 9: mavsdk.rpc.geofence.ClearGeofenceResponse
	(*GeofenceResult)(nil),         // 10: mavsdk.rpc.geofence.GeofenceResult
}
var file_geofence_proto_depIdxs = []int32{
	2,  // 0: mavsdk.rpc.geofence.Polygon.points:type_name -> mavsdk.rpc.geofence.Point
	0,  // 1: mavsdk.rpc.geofence.Polygon.fence_type:type_name -> mavsdk.rpc.geofence.FenceType
	2,  // 2: mavsdk.rpc.geofence.Circle.point:type_name -> mavsdk.rpc.geofence.Point
	0,  // 3: mavsdk.rpc.geofence.Circle.fence_type:type_name -> mavsdk.rpc.geofence.FenceType
	3,  // 4: mavsdk.rpc.geofence.GeofenceData.polygons:type_name -> mavsdk.rpc.geofence.Polygon
	4,  // 5: mavsdk.rpc.geofence.GeofenceData.circles:type_name -> mavsdk.rpc.geofence.Circle
	5,  // 6: mavsdk.rpc.geofence.UploadGeofenceRequest.geofence_data:type_name -> mavsdk.rpc.geofence.GeofenceData
	10, // 7: mavsdk.rpc.geofence.UploadGeofenceResponse.geofence_result:type_name -> mavsdk.rpc.geofence.GeofenceResult
	10, // 8: mavsdk.rpc.geofence.ClearGeofenceResponse.geofence_result:type_name -> mavsdk.rpc.geofence.GeofenceResult
	1,  // 9: mavsdk.rpc.geofence.GeofenceResult.result:type_name -> mavsdk.rpc.geofence.GeofenceResult.Result
	6,  // 10: mavsdk.rpc.geofence.GeofenceService.UploadGeofence:input_type -> mavsdk.rpc.geofence.UploadGeofenceRequest
	8,  // 11: mavsdk.rpc.geofence.GeofenceService.ClearGeofence:input_type -> mavsdk.rpc.geofence.ClearGeofenceRequest
	7,  // 12: mavsdk.rpc.geofence.GeofenceService.UploadGeofence:output_type -> mavsdk.rpc.geofence.UploadGeofenceResponse
	9,  // 13: mavsdk.rpc.geofence.GeofenceService.ClearGeofence:output_type -> mavsdk.rpc.geofence.ClearGeofenceResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_geofence_proto_init() }
func file_geofence_proto_init() {
	if File_geofence_proto != nil {
		return
	}
	file_mavsdk_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_geofence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_geofence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
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
		file_geofence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circle); i {
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
		file_geofence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeofenceData); i {
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
		file_geofence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadGeofenceRequest); i {
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
		file_geofence_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadGeofenceResponse); i {
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
		file_geofence_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearGeofenceRequest); i {
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
		file_geofence_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearGeofenceResponse); i {
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
		file_geofence_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeofenceResult); i {
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
			RawDescriptor: file_geofence_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geofence_proto_goTypes,
		DependencyIndexes: file_geofence_proto_depIdxs,
		EnumInfos:         file_geofence_proto_enumTypes,
		MessageInfos:      file_geofence_proto_msgTypes,
	}.Build()
	File_geofence_proto = out.File
	file_geofence_proto_rawDesc = nil
	file_geofence_proto_goTypes = nil
	file_geofence_proto_depIdxs = nil
}
