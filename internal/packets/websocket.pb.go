// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: websocket.proto

package packets

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

type ActionType int32

const (
	ActionType_UNKNOWN ActionType = 0
	ActionType_LOGIN   ActionType = 1
	ActionType_LOGOUT  ActionType = 2
	ActionType_MOVE    ActionType = 3
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LOGIN",
		2: "LOGOUT",
		3: "MOVE",
	}
	ActionType_value = map[string]int32{
		"UNKNOWN": 0,
		"LOGIN":   1,
		"LOGOUT":  2,
		"MOVE":    3,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_websocket_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_websocket_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

// Wrapper for all possible messages
type WSMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*WSMessage_TextMessage
	//	*WSMessage_GameInvite
	Payload       isWSMessage_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WSMessage) Reset() {
	*x = WSMessage{}
	mi := &file_websocket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSMessage) ProtoMessage() {}

func (x *WSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSMessage.ProtoReflect.Descriptor instead.
func (*WSMessage) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

func (x *WSMessage) GetPayload() isWSMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *WSMessage) GetTextMessage() *TextMessage {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_TextMessage); ok {
			return x.TextMessage
		}
	}
	return nil
}

func (x *WSMessage) GetGameInvite() *GameInvite {
	if x != nil {
		if x, ok := x.Payload.(*WSMessage_GameInvite); ok {
			return x.GameInvite
		}
	}
	return nil
}

type isWSMessage_Payload interface {
	isWSMessage_Payload()
}

type WSMessage_TextMessage struct {
	TextMessage *TextMessage `protobuf:"bytes,1,opt,name=text_message,json=textMessage,proto3,oneof"`
}

type WSMessage_GameInvite struct {
	GameInvite *GameInvite `protobuf:"bytes,2,opt,name=game_invite,json=gameInvite,proto3,oneof"`
}

func (*WSMessage_TextMessage) isWSMessage_Payload() {}

func (*WSMessage_GameInvite) isWSMessage_Payload() {}

type TextMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	mi := &file_websocket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{1}
}

func (x *TextMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GameInvite struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GameId        string                 `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	InviterId     string                 `protobuf:"bytes,2,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	InviteeId     string                 `protobuf:"bytes,3,opt,name=invitee_id,json=inviteeId,proto3" json:"invitee_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameInvite) Reset() {
	*x = GameInvite{}
	mi := &file_websocket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInvite) ProtoMessage() {}

func (x *GameInvite) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInvite.ProtoReflect.Descriptor instead.
func (*GameInvite) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{2}
}

func (x *GameInvite) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GameInvite) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *GameInvite) GetInviteeId() string {
	if x != nil {
		return x.InviteeId
	}
	return ""
}

var File_websocket_proto protoreflect.FileDescriptor

var file_websocket_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x8d, 0x01, 0x0a,
	0x09, 0x57, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77,
	0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x27, 0x0a, 0x0b,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x63, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x49, 0x64, 0x2a, 0x3a, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x74,
	0x6f, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websocket_proto_rawDescOnce sync.Once
	file_websocket_proto_rawDescData = file_websocket_proto_rawDesc
)

func file_websocket_proto_rawDescGZIP() []byte {
	file_websocket_proto_rawDescOnce.Do(func() {
		file_websocket_proto_rawDescData = protoimpl.X.CompressGZIP(file_websocket_proto_rawDescData)
	})
	return file_websocket_proto_rawDescData
}

var file_websocket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_websocket_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_websocket_proto_goTypes = []any{
	(ActionType)(0),     // 0: websocket.ActionType
	(*WSMessage)(nil),   // 1: websocket.WSMessage
	(*TextMessage)(nil), // 2: websocket.TextMessage
	(*GameInvite)(nil),  // 3: websocket.GameInvite
}
var file_websocket_proto_depIdxs = []int32{
	2, // 0: websocket.WSMessage.text_message:type_name -> websocket.TextMessage
	3, // 1: websocket.WSMessage.game_invite:type_name -> websocket.GameInvite
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_websocket_proto_init() }
func file_websocket_proto_init() {
	if File_websocket_proto != nil {
		return
	}
	file_websocket_proto_msgTypes[0].OneofWrappers = []any{
		(*WSMessage_TextMessage)(nil),
		(*WSMessage_GameInvite)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_websocket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_websocket_proto_goTypes,
		DependencyIndexes: file_websocket_proto_depIdxs,
		EnumInfos:         file_websocket_proto_enumTypes,
		MessageInfos:      file_websocket_proto_msgTypes,
	}.Build()
	File_websocket_proto = out.File
	file_websocket_proto_rawDesc = nil
	file_websocket_proto_goTypes = nil
	file_websocket_proto_depIdxs = nil
}
