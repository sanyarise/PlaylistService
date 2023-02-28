// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: proto/playlist.proto

package playlist

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

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint32 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *Song) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Song) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Song) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type CreateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint32 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *CreateSongRequest) Reset() {
	*x = CreateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongRequest) ProtoMessage() {}

func (x *CreateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongRequest.ProtoReflect.Descriptor instead.
func (*CreateSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSongRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSongRequest) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type CreateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSongResponse) Reset() {
	*x = CreateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongResponse) ProtoMessage() {}

func (x *CreateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongResponse.ProtoReflect.Descriptor instead.
func (*CreateSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSongResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSongRequest) Reset() {
	*x = GetSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongRequest) ProtoMessage() {}

func (x *GetSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongRequest.ProtoReflect.Descriptor instead.
func (*GetSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{3}
}

func (x *GetSongRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *GetSongResponse) Reset() {
	*x = GetSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongResponse) ProtoMessage() {}

func (x *GetSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongResponse.ProtoReflect.Descriptor instead.
func (*GetSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{4}
}

func (x *GetSongResponse) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type UpdateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint32 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *UpdateSongRequest) Reset() {
	*x = UpdateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongRequest) ProtoMessage() {}

func (x *UpdateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongRequest.ProtoReflect.Descriptor instead.
func (*UpdateSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSongRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSongRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateSongRequest) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type UpdateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSongResponse) Reset() {
	*x = UpdateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongResponse) ProtoMessage() {}

func (x *UpdateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongResponse.ProtoReflect.Descriptor instead.
func (*UpdateSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{6}
}

type DeleteSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSongRequest) Reset() {
	*x = DeleteSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongRequest) ProtoMessage() {}

func (x *DeleteSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongRequest.ProtoReflect.Descriptor instead.
func (*DeleteSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSongRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSongResponse) Reset() {
	*x = DeleteSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongResponse) ProtoMessage() {}

func (x *DeleteSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{8}
}

type PlaySongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaySongRequest) Reset() {
	*x = PlaySongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaySongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaySongRequest) ProtoMessage() {}

func (x *PlaySongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaySongRequest.ProtoReflect.Descriptor instead.
func (*PlaySongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{9}
}

type PlaySongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaySongResponse) Reset() {
	*x = PlaySongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaySongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaySongResponse) ProtoMessage() {}

func (x *PlaySongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaySongResponse.ProtoReflect.Descriptor instead.
func (*PlaySongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{10}
}

type PauseSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PauseSongRequest) Reset() {
	*x = PauseSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseSongRequest) ProtoMessage() {}

func (x *PauseSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseSongRequest.ProtoReflect.Descriptor instead.
func (*PauseSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{11}
}

type PauseSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PauseSongResponse) Reset() {
	*x = PauseSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseSongResponse) ProtoMessage() {}

func (x *PauseSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseSongResponse.ProtoReflect.Descriptor instead.
func (*PauseSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{12}
}

type NextSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextSongRequest) Reset() {
	*x = NextSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextSongRequest) ProtoMessage() {}

func (x *NextSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextSongRequest.ProtoReflect.Descriptor instead.
func (*NextSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{13}
}

type NextSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextSongResponse) Reset() {
	*x = NextSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextSongResponse) ProtoMessage() {}

func (x *NextSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextSongResponse.ProtoReflect.Descriptor instead.
func (*NextSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{14}
}

type PrevSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrevSongRequest) Reset() {
	*x = PrevSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrevSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrevSongRequest) ProtoMessage() {}

func (x *PrevSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrevSongRequest.ProtoReflect.Descriptor instead.
func (*PrevSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{15}
}

type PrevSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrevSongResponse) Reset() {
	*x = PrevSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrevSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrevSongResponse) ProtoMessage() {}

func (x *PrevSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrevSongResponse.ProtoReflect.Descriptor instead.
func (*PrevSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{16}
}

type AddSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *AddSongRequest) Reset() {
	*x = AddSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongRequest) ProtoMessage() {}

func (x *AddSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongRequest.ProtoReflect.Descriptor instead.
func (*AddSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{17}
}

func (x *AddSongRequest) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type AddSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddSongResponse) Reset() {
	*x = AddSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongResponse) ProtoMessage() {}

func (x *AddSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongResponse.ProtoReflect.Descriptor instead.
func (*AddSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{18}
}

var File_proto_playlist_proto protoreflect.FileDescriptor

var file_proto_playlist_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x48, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04,
	0x73, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67,
	0x22, 0x55, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x50,
	0x6c, 0x61, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x12, 0x0a, 0x10, 0x50, 0x61, 0x75, 0x73, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x61, 0x75, 0x73, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x65, 0x78, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4e,
	0x65, 0x78, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x76, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x76, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22, 0x11, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x86, 0x05, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x18, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x50, 0x61, 0x75, 0x73, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08,
	0x4e, 0x65, 0x78, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4e,
	0x65, 0x78, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x76, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e,
	0x67, 0x12, 0x18, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_playlist_proto_rawDescOnce sync.Once
	file_proto_playlist_proto_rawDescData = file_proto_playlist_proto_rawDesc
)

func file_proto_playlist_proto_rawDescGZIP() []byte {
	file_proto_playlist_proto_rawDescOnce.Do(func() {
		file_proto_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_playlist_proto_rawDescData)
	})
	return file_proto_playlist_proto_rawDescData
}

var file_proto_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_proto_playlist_proto_goTypes = []interface{}{
	(*Song)(nil),               // 0: playlist.Song
	(*CreateSongRequest)(nil),  // 1: playlist.CreateSongRequest
	(*CreateSongResponse)(nil), // 2: playlist.CreateSongResponse
	(*GetSongRequest)(nil),     // 3: playlist.GetSongRequest
	(*GetSongResponse)(nil),    // 4: playlist.GetSongResponse
	(*UpdateSongRequest)(nil),  // 5: playlist.UpdateSongRequest
	(*UpdateSongResponse)(nil), // 6: playlist.UpdateSongResponse
	(*DeleteSongRequest)(nil),  // 7: playlist.DeleteSongRequest
	(*DeleteSongResponse)(nil), // 8: playlist.DeleteSongResponse
	(*PlaySongRequest)(nil),    // 9: playlist.PlaySongRequest
	(*PlaySongResponse)(nil),   // 10: playlist.PlaySongResponse
	(*PauseSongRequest)(nil),   // 11: playlist.PauseSongRequest
	(*PauseSongResponse)(nil),  // 12: playlist.PauseSongResponse
	(*NextSongRequest)(nil),    // 13: playlist.NextSongRequest
	(*NextSongResponse)(nil),   // 14: playlist.NextSongResponse
	(*PrevSongRequest)(nil),    // 15: playlist.PrevSongRequest
	(*PrevSongResponse)(nil),   // 16: playlist.PrevSongResponse
	(*AddSongRequest)(nil),     // 17: playlist.AddSongRequest
	(*AddSongResponse)(nil),    // 18: playlist.AddSongResponse
}
var file_proto_playlist_proto_depIdxs = []int32{
	0,  // 0: playlist.GetSongResponse.song:type_name -> playlist.Song
	0,  // 1: playlist.AddSongRequest.song:type_name -> playlist.Song
	1,  // 2: playlist.Playlist.CreateSong:input_type -> playlist.CreateSongRequest
	3,  // 3: playlist.Playlist.GetSong:input_type -> playlist.GetSongRequest
	5,  // 4: playlist.Playlist.UpdateSong:input_type -> playlist.UpdateSongRequest
	7,  // 5: playlist.Playlist.DeleteSong:input_type -> playlist.DeleteSongRequest
	9,  // 6: playlist.Playlist.PlaySong:input_type -> playlist.PlaySongRequest
	11, // 7: playlist.Playlist.PauseSong:input_type -> playlist.PauseSongRequest
	13, // 8: playlist.Playlist.NextSong:input_type -> playlist.NextSongRequest
	15, // 9: playlist.Playlist.PrevSong:input_type -> playlist.PrevSongRequest
	17, // 10: playlist.Playlist.AddSong:input_type -> playlist.AddSongRequest
	2,  // 11: playlist.Playlist.CreateSong:output_type -> playlist.CreateSongResponse
	4,  // 12: playlist.Playlist.GetSong:output_type -> playlist.GetSongResponse
	6,  // 13: playlist.Playlist.UpdateSong:output_type -> playlist.UpdateSongResponse
	8,  // 14: playlist.Playlist.DeleteSong:output_type -> playlist.DeleteSongResponse
	10, // 15: playlist.Playlist.PlaySong:output_type -> playlist.PlaySongResponse
	12, // 16: playlist.Playlist.PauseSong:output_type -> playlist.PauseSongResponse
	14, // 17: playlist.Playlist.NextSong:output_type -> playlist.NextSongResponse
	16, // 18: playlist.Playlist.PrevSong:output_type -> playlist.PrevSongResponse
	18, // 19: playlist.Playlist.AddSong:output_type -> playlist.AddSongResponse
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_playlist_proto_init() }
func file_proto_playlist_proto_init() {
	if File_proto_playlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_playlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_proto_playlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaySongRequest); i {
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
		file_proto_playlist_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaySongResponse); i {
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
		file_proto_playlist_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrevSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrevSongResponse); i {
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
		file_proto_playlist_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSongRequest); i {
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
		file_proto_playlist_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSongResponse); i {
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
			RawDescriptor: file_proto_playlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_playlist_proto_goTypes,
		DependencyIndexes: file_proto_playlist_proto_depIdxs,
		MessageInfos:      file_proto_playlist_proto_msgTypes,
	}.Build()
	File_proto_playlist_proto = out.File
	file_proto_playlist_proto_rawDesc = nil
	file_proto_playlist_proto_goTypes = nil
	file_proto_playlist_proto_depIdxs = nil
}
