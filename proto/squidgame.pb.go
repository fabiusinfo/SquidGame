// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/squidgame.proto

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes1 string `protobuf:"bytes,1,opt,name=codes1,proto3" json:"codes1,omitempty"`
	Codes2 string `protobuf:"bytes,2,opt,name=codes2,proto3" json:"codes2,omitempty"`
	Codes3 string `protobuf:"bytes,3,opt,name=codes3,proto3" json:"codes3,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{3}
}

func (x *JoinReply) GetCodes1() string {
	if x != nil {
		return x.Codes1
	}
	return ""
}

func (x *JoinReply) GetCodes2() string {
	if x != nil {
		return x.Codes2
	}
	return ""
}

func (x *JoinReply) GetCodes3() string {
	if x != nil {
		return x.Codes3
	}
	return ""
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"` //id jugador
	Play   string `protobuf:"bytes,2,opt,name=play,proto3" json:"play,omitempty"`     //jugada
	Stage  string `protobuf:"bytes,3,opt,name=stage,proto3" json:"stage,omitempty"`   //etapa
	Round  int32  `protobuf:"varint,4,opt,name=round,proto3" json:"round,omitempty"`  //ronda
	Score  int32  `protobuf:"varint,5,opt,name=score,proto3" json:"score,omitempty"`  //puntaje
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{4}
}

func (x *SendRequest) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *SendRequest) GetPlay() string {
	if x != nil {
		return x.Play
	}
	return ""
}

func (x *SendRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *SendRequest) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *SendRequest) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type SendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage   string `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
	Alive   bool   `protobuf:"varint,2,opt,name=alive,proto3" json:"alive,omitempty"`
	Round   int32  `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	Started bool   `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
	Lround  int32  `protobuf:"varint,6,opt,name=lround,proto3" json:"lround,omitempty"`
}

func (x *SendReply) Reset() {
	*x = SendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReply) ProtoMessage() {}

func (x *SendReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReply.ProtoReflect.Descriptor instead.
func (*SendReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{5}
}

func (x *SendReply) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *SendReply) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *SendReply) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *SendReply) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *SendReply) GetLround() int32 {
	if x != nil {
		return x.Lround
	}
	return 0
}

type AmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AmountRequest) Reset() {
	*x = AmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountRequest) ProtoMessage() {}

func (x *AmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountRequest.ProtoReflect.Descriptor instead.
func (*AmountRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{6}
}

func (x *AmountRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AmountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto string `protobuf:"bytes,1,opt,name=monto,proto3" json:"monto,omitempty"`
}

func (x *AmountReply) Reset() {
	*x = AmountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountReply) ProtoMessage() {}

func (x *AmountReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountReply.ProtoReflect.Descriptor instead.
func (*AmountReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{7}
}

func (x *AmountReply) GetMonto() string {
	if x != nil {
		return x.Monto
	}
	return ""
}

type DeadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Stage  string `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *DeadRequest) Reset() {
	*x = DeadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadRequest) ProtoMessage() {}

func (x *DeadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadRequest.ProtoReflect.Descriptor instead.
func (*DeadRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{8}
}

func (x *DeadRequest) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *DeadRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type DeadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dead bool `protobuf:"varint,1,opt,name=dead,proto3" json:"dead,omitempty"`
}

func (x *DeadReply) Reset() {
	*x = DeadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadReply) ProtoMessage() {}

func (x *DeadReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadReply.ProtoReflect.Descriptor instead.
func (*DeadReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{9}
}

func (x *DeadReply) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{10}
}

func (x *StartRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Started bool `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
}

func (x *StartReply) Reset() {
	*x = StartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReply) ProtoMessage() {}

func (x *StartReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReply.ProtoReflect.Descriptor instead.
func (*StartReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{11}
}

func (x *StartReply) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

var File_proto_squidgame_proto protoreflect.FileDescriptor

var file_proto_squidgame_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x22, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x53, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x33, 0x22, 0x7b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x7f, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x44, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x22, 0x1f, 0x0a, 0x09, 0x44, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x61,
	0x64, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x32, 0xcc, 0x02, 0x0a, 0x10, 0x53, 0x71, 0x75, 0x69, 0x64, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x08,
	0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0b, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0b, 0x44, 0x65,
	0x61, 0x64, 0x4f, 0x72, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x61, 0x62, 0x69, 0x75, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x53, 0x71, 0x75, 0x69,
	0x64, 0x47, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_squidgame_proto_rawDescOnce sync.Once
	file_proto_squidgame_proto_rawDescData = file_proto_squidgame_proto_rawDesc
)

func file_proto_squidgame_proto_rawDescGZIP() []byte {
	file_proto_squidgame_proto_rawDescOnce.Do(func() {
		file_proto_squidgame_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_squidgame_proto_rawDescData)
	})
	return file_proto_squidgame_proto_rawDescData
}

var file_proto_squidgame_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_squidgame_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: grpc.HelloRequest
	(*HelloReply)(nil),    // 1: grpc.HelloReply
	(*JoinRequest)(nil),   // 2: grpc.JoinRequest
	(*JoinReply)(nil),     // 3: grpc.JoinReply
	(*SendRequest)(nil),   // 4: grpc.SendRequest
	(*SendReply)(nil),     // 5: grpc.SendReply
	(*AmountRequest)(nil), // 6: grpc.AmountRequest
	(*AmountReply)(nil),   // 7: grpc.AmountReply
	(*DeadRequest)(nil),   // 8: grpc.DeadRequest
	(*DeadReply)(nil),     // 9: grpc.DeadReply
	(*StartRequest)(nil),  // 10: grpc.StartRequest
	(*StartReply)(nil),    // 11: grpc.StartReply
}
var file_proto_squidgame_proto_depIdxs = []int32{
	0,  // 0: grpc.SquidGameService.SayHello:input_type -> grpc.HelloRequest
	2,  // 1: grpc.SquidGameService.JoinGame:input_type -> grpc.JoinRequest
	4,  // 2: grpc.SquidGameService.SendPlays:input_type -> grpc.SendRequest
	6,  // 3: grpc.SquidGameService.AmountCheck:input_type -> grpc.AmountRequest
	8,  // 4: grpc.SquidGameService.DeadOrAlive:input_type -> grpc.DeadRequest
	10, // 5: grpc.SquidGameService.Started:input_type -> grpc.StartRequest
	1,  // 6: grpc.SquidGameService.SayHello:output_type -> grpc.HelloReply
	3,  // 7: grpc.SquidGameService.JoinGame:output_type -> grpc.JoinReply
	5,  // 8: grpc.SquidGameService.SendPlays:output_type -> grpc.SendReply
	7,  // 9: grpc.SquidGameService.AmountCheck:output_type -> grpc.AmountReply
	9,  // 10: grpc.SquidGameService.DeadOrAlive:output_type -> grpc.DeadReply
	11, // 11: grpc.SquidGameService.Started:output_type -> grpc.StartReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_squidgame_proto_init() }
func file_proto_squidgame_proto_init() {
	if File_proto_squidgame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_squidgame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_squidgame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_proto_squidgame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_proto_squidgame_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_proto_squidgame_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_proto_squidgame_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReply); i {
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
		file_proto_squidgame_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountRequest); i {
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
		file_proto_squidgame_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountReply); i {
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
		file_proto_squidgame_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadRequest); i {
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
		file_proto_squidgame_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadReply); i {
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
		file_proto_squidgame_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_proto_squidgame_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReply); i {
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
			RawDescriptor: file_proto_squidgame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_squidgame_proto_goTypes,
		DependencyIndexes: file_proto_squidgame_proto_depIdxs,
		MessageInfos:      file_proto_squidgame_proto_msgTypes,
	}.Build()
	File_proto_squidgame_proto = out.File
	file_proto_squidgame_proto_rawDesc = nil
	file_proto_squidgame_proto_goTypes = nil
	file_proto_squidgame_proto_depIdxs = nil
}
