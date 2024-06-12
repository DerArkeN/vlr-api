// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: vlr/api/api_service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// *
// Returns a list of match ids that match the given criteria
// You can further limit the number by filling one or more optional *options* fields, look at the comments for the *options* message for more information
//
// Timestamps are in UTC
// If the status is STATUS_LIVE, the *from* and *to* fields are ignored
// If the status is STATUS_UPCOMING and no *from* field is provided, the *from* field is set to the current time, if no *to* field is provided, the *to* field is set to the *from* field +24h
// If the status is STATUS_COMPLETED and no *from* field is provided, the *from* field is is set to the *to* field -24h, if no *to* field is provided, the *to* field is set to the current time
type GetMatchIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status                      `protobuf:"varint,1,opt,name=status,proto3,enum=vlr.api.Status" json:"status,omitempty"`
	From    *timestamppb.Timestamp      `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To      *timestamppb.Timestamp      `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Options *GetMatchIdsRequest_Options `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *GetMatchIdsRequest) Reset() {
	*x = GetMatchIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_api_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchIdsRequest) ProtoMessage() {}

func (x *GetMatchIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_api_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchIdsRequest.ProtoReflect.Descriptor instead.
func (*GetMatchIdsRequest) Descriptor() ([]byte, []int) {
	return file_vlr_api_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMatchIdsRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *GetMatchIdsRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetMatchIdsRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *GetMatchIdsRequest) GetOptions() *GetMatchIdsRequest_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type GetMatchIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchIds []string `protobuf:"bytes,1,rep,name=match_ids,json=matchIds,proto3" json:"match_ids,omitempty"`
}

func (x *GetMatchIdsResponse) Reset() {
	*x = GetMatchIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_api_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchIdsResponse) ProtoMessage() {}

func (x *GetMatchIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_api_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchIdsResponse.ProtoReflect.Descriptor instead.
func (*GetMatchIdsResponse) Descriptor() ([]byte, []int) {
	return file_vlr_api_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMatchIdsResponse) GetMatchIds() []string {
	if x != nil {
		return x.MatchIds
	}
	return nil
}

// Returns a match by its id
type GetMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
}

func (x *GetMatchRequest) Reset() {
	*x = GetMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_api_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchRequest) ProtoMessage() {}

func (x *GetMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_api_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchRequest.ProtoReflect.Descriptor instead.
func (*GetMatchRequest) Descriptor() ([]byte, []int) {
	return file_vlr_api_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMatchRequest) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

type GetMatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match *Match `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
}

func (x *GetMatchResponse) Reset() {
	*x = GetMatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_api_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchResponse) ProtoMessage() {}

func (x *GetMatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_api_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchResponse.ProtoReflect.Descriptor instead.
func (*GetMatchResponse) Descriptor() ([]byte, []int) {
	return file_vlr_api_api_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMatchResponse) GetMatch() *Match {
	if x != nil {
		return x.Match
	}
	return nil
}

type GetMatchIdsRequest_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If an event id is provided, *from* and *to* are ignored
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetMatchIdsRequest_Options) Reset() {
	*x = GetMatchIdsRequest_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_api_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchIdsRequest_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchIdsRequest_Options) ProtoMessage() {}

func (x *GetMatchIdsRequest_Options) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_api_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchIdsRequest_Options.ProtoReflect.Descriptor instead.
func (*GetMatchIdsRequest_Options) Descriptor() ([]byte, []int) {
	return file_vlr_api_api_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetMatchIdsRequest_Options) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

var File_vlr_api_api_service_proto protoreflect.FileDescriptor

var file_vlr_api_api_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x6c, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x6c, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x11, 0x76, 0x6c, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x76, 0x6c, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x3d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x24, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76,
	0x6c, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x32, 0xbf, 0x01, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x61, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x76, 0x6c, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a,
	0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x73, 0x12, 0x55,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x76, 0x6c, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x72, 0x61, 0x72, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x6c, 0x72,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vlr_api_api_service_proto_rawDescOnce sync.Once
	file_vlr_api_api_service_proto_rawDescData = file_vlr_api_api_service_proto_rawDesc
)

func file_vlr_api_api_service_proto_rawDescGZIP() []byte {
	file_vlr_api_api_service_proto_rawDescOnce.Do(func() {
		file_vlr_api_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_vlr_api_api_service_proto_rawDescData)
	})
	return file_vlr_api_api_service_proto_rawDescData
}

var file_vlr_api_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vlr_api_api_service_proto_goTypes = []interface{}{
	(*GetMatchIdsRequest)(nil),         // 0: vlr.api.GetMatchIdsRequest
	(*GetMatchIdsResponse)(nil),        // 1: vlr.api.GetMatchIdsResponse
	(*GetMatchRequest)(nil),            // 2: vlr.api.GetMatchRequest
	(*GetMatchResponse)(nil),           // 3: vlr.api.GetMatchResponse
	(*GetMatchIdsRequest_Options)(nil), // 4: vlr.api.GetMatchIdsRequest.Options
	(Status)(0),                        // 5: vlr.api.Status
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
	(*Match)(nil),                      // 7: vlr.api.Match
}
var file_vlr_api_api_service_proto_depIdxs = []int32{
	5, // 0: vlr.api.GetMatchIdsRequest.status:type_name -> vlr.api.Status
	6, // 1: vlr.api.GetMatchIdsRequest.from:type_name -> google.protobuf.Timestamp
	6, // 2: vlr.api.GetMatchIdsRequest.to:type_name -> google.protobuf.Timestamp
	4, // 3: vlr.api.GetMatchIdsRequest.options:type_name -> vlr.api.GetMatchIdsRequest.Options
	7, // 4: vlr.api.GetMatchResponse.match:type_name -> vlr.api.Match
	0, // 5: vlr.api.Api.GetMatchIds:input_type -> vlr.api.GetMatchIdsRequest
	2, // 6: vlr.api.Api.GetMatch:input_type -> vlr.api.GetMatchRequest
	1, // 7: vlr.api.Api.GetMatchIds:output_type -> vlr.api.GetMatchIdsResponse
	3, // 8: vlr.api.Api.GetMatch:output_type -> vlr.api.GetMatchResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vlr_api_api_service_proto_init() }
func file_vlr_api_api_service_proto_init() {
	if File_vlr_api_api_service_proto != nil {
		return
	}
	file_vlr_api_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vlr_api_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchIdsRequest); i {
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
		file_vlr_api_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchIdsResponse); i {
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
		file_vlr_api_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchRequest); i {
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
		file_vlr_api_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchResponse); i {
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
		file_vlr_api_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchIdsRequest_Options); i {
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
			RawDescriptor: file_vlr_api_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vlr_api_api_service_proto_goTypes,
		DependencyIndexes: file_vlr_api_api_service_proto_depIdxs,
		MessageInfos:      file_vlr_api_api_service_proto_msgTypes,
	}.Build()
	File_vlr_api_api_service_proto = out.File
	file_vlr_api_api_service_proto_rawDesc = nil
	file_vlr_api_api_service_proto_goTypes = nil
	file_vlr_api_api_service_proto_depIdxs = nil
}
