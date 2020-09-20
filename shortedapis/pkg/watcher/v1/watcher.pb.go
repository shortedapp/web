// // Copyright (c) 2020 Shorted Ltd Pty.
// //
// // Permission is hereby granted, free of charge, to any person obtaining a copy
// // of this software and associated documentation files (the "Software"), to deal
// // in the Software without restriction, including without limitation the rights
// // to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// // copies of the Software, and to permit persons to whom the Software is
// // furnished to do so, subject to the following conditions:
// //
// // The above copyright notice and this permission notice shall be included in
// // all copies or substantial portions of the Software.
// //
// // THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// // IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// // FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// // AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// // LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// // OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// // THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: watcher/v1/watcher.proto

package watcher

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// SyncStatus specifies the state of a given synchronisation request.
// Where a synchronisation state will correspond to a specfic source
type SyncStatus int32

const (
	SyncStatus_SYNC_STATUS_UNSPECIFIED SyncStatus = 0
	SyncStatus_SYNC_STATUS_PENDING     SyncStatus = 1
	SyncStatus_SYNC_STATUS_SUCCESS     SyncStatus = 2
	SyncStatus_SYNC_STATUS_FAILURE     SyncStatus = 3
)

// Enum value maps for SyncStatus.
var (
	SyncStatus_name = map[int32]string{
		0: "SYNC_STATUS_UNSPECIFIED",
		1: "SYNC_STATUS_PENDING",
		2: "SYNC_STATUS_SUCCESS",
		3: "SYNC_STATUS_FAILURE",
	}
	SyncStatus_value = map[string]int32{
		"SYNC_STATUS_UNSPECIFIED": 0,
		"SYNC_STATUS_PENDING":     1,
		"SYNC_STATUS_SUCCESS":     2,
		"SYNC_STATUS_FAILURE":     3,
	}
)

func (x SyncStatus) Enum() *SyncStatus {
	p := new(SyncStatus)
	*p = x
	return p
}

func (x SyncStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_watcher_v1_watcher_proto_enumTypes[0].Descriptor()
}

func (SyncStatus) Type() protoreflect.EnumType {
	return &file_watcher_v1_watcher_proto_enumTypes[0]
}

func (x SyncStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SyncStatus.Descriptor instead.
func (SyncStatus) EnumDescriptor() ([]byte, []int) {
	return file_watcher_v1_watcher_proto_rawDescGZIP(), []int{0}
}

// WatchDetails contains the set of information for a given watch.
type WatchDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source *Source    `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Index  *Index     `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
	Status SyncStatus `protobuf:"varint,5,opt,name=status,proto3,enum=shorted.watcher.v1.SyncStatus" json:"status,omitempty"`
}

func (x *WatchDetails) Reset() {
	*x = WatchDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_v1_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchDetails) ProtoMessage() {}

func (x *WatchDetails) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_v1_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchDetails.ProtoReflect.Descriptor instead.
func (*WatchDetails) Descriptor() ([]byte, []int) {
	return file_watcher_v1_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *WatchDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WatchDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WatchDetails) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *WatchDetails) GetIndex() *Index {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *WatchDetails) GetStatus() SyncStatus {
	if x != nil {
		return x.Status
	}
	return SyncStatus_SYNC_STATUS_UNSPECIFIED
}

// Source is a definition of a target that we want to "watch" and discover content from
type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL is the base URL for the givne source
	Url      string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Adapter  string             `protobuf:"bytes,2,opt,name=adapter,proto3" json:"adapter,omitempty"`
	Interval *duration.Duration `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_v1_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_v1_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_watcher_v1_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *Source) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Source) GetAdapter() string {
	if x != nil {
		return x.Adapter
	}
	return ""
}

func (x *Source) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

// Index defines where the discovered content of a Source is stored
type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//url where index is stored (i.e gs://<project-name>-index/index.json)
	Url           string               `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DocumentCount int64                `protobuf:"varint,2,opt,name=document_count,json=documentCount,proto3" json:"document_count,omitempty"`
	LastUpdated   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_v1_watcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_v1_watcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_watcher_v1_watcher_proto_rawDescGZIP(), []int{2}
}

func (x *Index) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Index) GetDocumentCount() int64 {
	if x != nil {
		return x.DocumentCount
	}
	return 0
}

func (x *Index) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type SyncDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status SyncStatus `protobuf:"varint,3,opt,name=status,proto3,enum=shorted.watcher.v1.SyncStatus" json:"status,omitempty"`
}

func (x *SyncDetails) Reset() {
	*x = SyncDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_v1_watcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncDetails) ProtoMessage() {}

func (x *SyncDetails) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_v1_watcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncDetails.ProtoReflect.Descriptor instead.
func (*SyncDetails) Descriptor() ([]byte, []int) {
	return file_watcher_v1_watcher_proto_rawDescGZIP(), []int{3}
}

func (x *SyncDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SyncDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncDetails) GetStatus() SyncStatus {
	if x != nil {
		return x.Status
	}
	return SyncStatus_SYNC_STATUS_UNSPECIFIED
}

var File_watcher_v1_watcher_proto protoreflect.FileDescriptor

var file_watcher_v1_watcher_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x05, 0x18, 0x0a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x40, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x88,
	0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90, 0x01, 0x01,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x52, 0x04, 0x61,
	0x73, 0x69, 0x63, 0x52, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0x98, 0xdf, 0x1f, 0x01, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x88, 0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x05,
	0x18, 0x0a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x74, 0x0a, 0x0a, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x59, 0x4e, 0x43,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03,
	0x42, 0x6c, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0xa2, 0x02,
	0x03, 0x53, 0x57, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x5c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watcher_v1_watcher_proto_rawDescOnce sync.Once
	file_watcher_v1_watcher_proto_rawDescData = file_watcher_v1_watcher_proto_rawDesc
)

func file_watcher_v1_watcher_proto_rawDescGZIP() []byte {
	file_watcher_v1_watcher_proto_rawDescOnce.Do(func() {
		file_watcher_v1_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_watcher_v1_watcher_proto_rawDescData)
	})
	return file_watcher_v1_watcher_proto_rawDescData
}

var file_watcher_v1_watcher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_watcher_v1_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_watcher_v1_watcher_proto_goTypes = []interface{}{
	(SyncStatus)(0),             // 0: shorted.watcher.v1.SyncStatus
	(*WatchDetails)(nil),        // 1: shorted.watcher.v1.WatchDetails
	(*Source)(nil),              // 2: shorted.watcher.v1.Source
	(*Index)(nil),               // 3: shorted.watcher.v1.Index
	(*SyncDetails)(nil),         // 4: shorted.watcher.v1.SyncDetails
	(*duration.Duration)(nil),   // 5: google.protobuf.Duration
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_watcher_v1_watcher_proto_depIdxs = []int32{
	2, // 0: shorted.watcher.v1.WatchDetails.source:type_name -> shorted.watcher.v1.Source
	3, // 1: shorted.watcher.v1.WatchDetails.index:type_name -> shorted.watcher.v1.Index
	0, // 2: shorted.watcher.v1.WatchDetails.status:type_name -> shorted.watcher.v1.SyncStatus
	5, // 3: shorted.watcher.v1.Source.interval:type_name -> google.protobuf.Duration
	6, // 4: shorted.watcher.v1.Index.last_updated:type_name -> google.protobuf.Timestamp
	0, // 5: shorted.watcher.v1.SyncDetails.status:type_name -> shorted.watcher.v1.SyncStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_watcher_v1_watcher_proto_init() }
func file_watcher_v1_watcher_proto_init() {
	if File_watcher_v1_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_watcher_v1_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchDetails); i {
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
		file_watcher_v1_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_watcher_v1_watcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_watcher_v1_watcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncDetails); i {
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
			RawDescriptor: file_watcher_v1_watcher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_watcher_v1_watcher_proto_goTypes,
		DependencyIndexes: file_watcher_v1_watcher_proto_depIdxs,
		EnumInfos:         file_watcher_v1_watcher_proto_enumTypes,
		MessageInfos:      file_watcher_v1_watcher_proto_msgTypes,
	}.Build()
	File_watcher_v1_watcher_proto = out.File
	file_watcher_v1_watcher_proto_rawDesc = nil
	file_watcher_v1_watcher_proto_goTypes = nil
	file_watcher_v1_watcher_proto_depIdxs = nil
}
