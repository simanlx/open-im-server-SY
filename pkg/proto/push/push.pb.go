// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: push/push.proto

package pbPush

import (
	sdk_ws "Open_IM/pkg/proto/sdk_ws"
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

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID  string          `protobuf:"bytes,1,opt,name=operationID,proto3" json:"operationID,omitempty"`
	MsgData      *sdk_ws.MsgData `protobuf:"bytes,2,opt,name=msgData,proto3" json:"msgData,omitempty"`
	PushToUserID string          `protobuf:"bytes,3,opt,name=pushToUserID,proto3" json:"pushToUserID,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *PushMsgReq) GetMsgData() *sdk_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *PushMsgReq) GetPushToUserID() string {
	if x != nil {
		return x.PushToUserID
	}
	return ""
}

type PushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int32 `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
}

func (x *PushMsgResp) Reset() {
	*x = PushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgResp) ProtoMessage() {}

func (x *PushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgResp.ProtoReflect.Descriptor instead.
func (*PushMsgResp) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{1}
}

func (x *PushMsgResp) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

type DelUserPushTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID string `protobuf:"bytes,1,opt,name=operationID,proto3" json:"operationID,omitempty"`
	UserID      string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	PlatformID  int32  `protobuf:"varint,3,opt,name=platformID,proto3" json:"platformID,omitempty"`
}

func (x *DelUserPushTokenReq) Reset() {
	*x = DelUserPushTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserPushTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserPushTokenReq) ProtoMessage() {}

func (x *DelUserPushTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserPushTokenReq.ProtoReflect.Descriptor instead.
func (*DelUserPushTokenReq) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{2}
}

func (x *DelUserPushTokenReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *DelUserPushTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *DelUserPushTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

type DelUserPushTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *DelUserPushTokenResp) Reset() {
	*x = DelUserPushTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserPushTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserPushTokenResp) ProtoMessage() {}

func (x *DelUserPushTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserPushTokenResp.ProtoReflect.Descriptor instead.
func (*DelUserPushTokenResp) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{3}
}

func (x *DelUserPushTokenResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *DelUserPushTokenResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_push_push_proto protoreflect.FileDescriptor

var file_push_push_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x70, 0x75, 0x73, 0x68, 0x1a, 0x28, 0x4f, 0x70, 0x65, 0x6e, 0x2d, 0x49, 0x4d,
	0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x5f, 0x77, 0x73, 0x2f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x0b,
	0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x6f, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x32, 0x8b, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e,
	0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x1f, 0x5a, 0x1d, 0x4f, 0x70, 0x65, 0x6e, 0x5f, 0x49, 0x4d, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x3b, 0x70,
	0x62, 0x50, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_push_push_proto_rawDescOnce sync.Once
	file_push_push_proto_rawDescData = file_push_push_proto_rawDesc
)

func file_push_push_proto_rawDescGZIP() []byte {
	file_push_push_proto_rawDescOnce.Do(func() {
		file_push_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_push_proto_rawDescData)
	})
	return file_push_push_proto_rawDescData
}

var file_push_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_push_push_proto_goTypes = []interface{}{
	(*PushMsgReq)(nil),           // 0: push.PushMsgReq
	(*PushMsgResp)(nil),          // 1: push.PushMsgResp
	(*DelUserPushTokenReq)(nil),  // 2: push.DelUserPushTokenReq
	(*DelUserPushTokenResp)(nil), // 3: push.DelUserPushTokenResp
	(*sdk_ws.MsgData)(nil),       // 4: server_api_params.MsgData
}
var file_push_push_proto_depIdxs = []int32{
	4, // 0: push.PushMsgReq.msgData:type_name -> server_api_params.MsgData
	0, // 1: push.PushMsgService.PushMsg:input_type -> push.PushMsgReq
	2, // 2: push.PushMsgService.DelUserPushToken:input_type -> push.DelUserPushTokenReq
	1, // 3: push.PushMsgService.PushMsg:output_type -> push.PushMsgResp
	3, // 4: push.PushMsgService.DelUserPushToken:output_type -> push.DelUserPushTokenResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_push_push_proto_init() }
func file_push_push_proto_init() {
	if File_push_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_push_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgResp); i {
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
		file_push_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserPushTokenReq); i {
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
		file_push_push_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserPushTokenResp); i {
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
			RawDescriptor: file_push_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_push_proto_goTypes,
		DependencyIndexes: file_push_push_proto_depIdxs,
		MessageInfos:      file_push_push_proto_msgTypes,
	}.Build()
	File_push_push_proto = out.File
	file_push_push_proto_rawDesc = nil
	file_push_push_proto_goTypes = nil
	file_push_push_proto_depIdxs = nil
}
