// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: commander_detach.proto

package dhctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommanderDetachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*CommanderDetachRequest_Start
	//	*CommanderDetachRequest_Cancel
	Message isCommanderDetachRequest_Message `protobuf_oneof:"message"`
}

func (x *CommanderDetachRequest) Reset() {
	*x = CommanderDetachRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachRequest) ProtoMessage() {}

func (x *CommanderDetachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachRequest.ProtoReflect.Descriptor instead.
func (*CommanderDetachRequest) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{0}
}

func (m *CommanderDetachRequest) GetMessage() isCommanderDetachRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *CommanderDetachRequest) GetStart() *CommanderDetachStart {
	if x, ok := x.GetMessage().(*CommanderDetachRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *CommanderDetachRequest) GetCancel() *CommanderDetachCancel {
	if x, ok := x.GetMessage().(*CommanderDetachRequest_Cancel); ok {
		return x.Cancel
	}
	return nil
}

type isCommanderDetachRequest_Message interface {
	isCommanderDetachRequest_Message()
}

type CommanderDetachRequest_Start struct {
	Start *CommanderDetachStart `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type CommanderDetachRequest_Cancel struct {
	Cancel *CommanderDetachCancel `protobuf:"bytes,2,opt,name=cancel,proto3,oneof"`
}

func (*CommanderDetachRequest_Start) isCommanderDetachRequest_Message() {}

func (*CommanderDetachRequest_Cancel) isCommanderDetachRequest_Message() {}

type CommanderDetachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*CommanderDetachResponse_Result
	//	*CommanderDetachResponse_Logs
	//	*CommanderDetachResponse_Progress
	Message isCommanderDetachResponse_Message `protobuf_oneof:"message"`
}

func (x *CommanderDetachResponse) Reset() {
	*x = CommanderDetachResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachResponse) ProtoMessage() {}

func (x *CommanderDetachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachResponse.ProtoReflect.Descriptor instead.
func (*CommanderDetachResponse) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{1}
}

func (m *CommanderDetachResponse) GetMessage() isCommanderDetachResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *CommanderDetachResponse) GetResult() *CommanderDetachResult {
	if x, ok := x.GetMessage().(*CommanderDetachResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *CommanderDetachResponse) GetLogs() *Logs {
	if x, ok := x.GetMessage().(*CommanderDetachResponse_Logs); ok {
		return x.Logs
	}
	return nil
}

func (x *CommanderDetachResponse) GetProgress() *Progress {
	if x, ok := x.GetMessage().(*CommanderDetachResponse_Progress); ok {
		return x.Progress
	}
	return nil
}

type isCommanderDetachResponse_Message interface {
	isCommanderDetachResponse_Message()
}

type CommanderDetachResponse_Result struct {
	Result *CommanderDetachResult `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type CommanderDetachResponse_Logs struct {
	Logs *Logs `protobuf:"bytes,2,opt,name=logs,proto3,oneof"`
}

type CommanderDetachResponse_Progress struct {
	Progress *Progress `protobuf:"bytes,4,opt,name=progress,proto3,oneof"`
}

func (*CommanderDetachResponse_Result) isCommanderDetachResponse_Message() {}

func (*CommanderDetachResponse_Logs) isCommanderDetachResponse_Message() {}

func (*CommanderDetachResponse_Progress) isCommanderDetachResponse_Message() {}

type CommanderDetachStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfig              string                       `protobuf:"bytes,1,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	ClusterConfig                 string                       `protobuf:"bytes,2,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	ProviderSpecificClusterConfig string                       `protobuf:"bytes,3,opt,name=provider_specific_cluster_config,json=providerSpecificClusterConfig,proto3" json:"provider_specific_cluster_config,omitempty"`
	State                         string                       `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	CreateResourcesTemplate       string                       `protobuf:"bytes,5,opt,name=create_resources_template,json=createResourcesTemplate,proto3" json:"create_resources_template,omitempty"`
	DeleteResourcesTemplate       string                       `protobuf:"bytes,6,opt,name=delete_resources_template,json=deleteResourcesTemplate,proto3" json:"delete_resources_template,omitempty"`
	CreateResourcesValues         *structpb.Struct             `protobuf:"bytes,7,opt,name=create_resources_values,json=createResourcesValues,proto3" json:"create_resources_values,omitempty"`
	DeleteResourcesValues         *structpb.Struct             `protobuf:"bytes,8,opt,name=delete_resources_values,json=deleteResourcesValues,proto3" json:"delete_resources_values,omitempty"`
	Options                       *CommanderDetachStartOptions `protobuf:"bytes,9,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CommanderDetachStart) Reset() {
	*x = CommanderDetachStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachStart) ProtoMessage() {}

func (x *CommanderDetachStart) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachStart.ProtoReflect.Descriptor instead.
func (*CommanderDetachStart) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{2}
}

func (x *CommanderDetachStart) GetConnectionConfig() string {
	if x != nil {
		return x.ConnectionConfig
	}
	return ""
}

func (x *CommanderDetachStart) GetClusterConfig() string {
	if x != nil {
		return x.ClusterConfig
	}
	return ""
}

func (x *CommanderDetachStart) GetProviderSpecificClusterConfig() string {
	if x != nil {
		return x.ProviderSpecificClusterConfig
	}
	return ""
}

func (x *CommanderDetachStart) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CommanderDetachStart) GetCreateResourcesTemplate() string {
	if x != nil {
		return x.CreateResourcesTemplate
	}
	return ""
}

func (x *CommanderDetachStart) GetDeleteResourcesTemplate() string {
	if x != nil {
		return x.DeleteResourcesTemplate
	}
	return ""
}

func (x *CommanderDetachStart) GetCreateResourcesValues() *structpb.Struct {
	if x != nil {
		return x.CreateResourcesValues
	}
	return nil
}

func (x *CommanderDetachStart) GetDeleteResourcesValues() *structpb.Struct {
	if x != nil {
		return x.DeleteResourcesValues
	}
	return nil
}

func (x *CommanderDetachStart) GetOptions() *CommanderDetachStartOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type CommanderDetachCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommanderDetachCancel) Reset() {
	*x = CommanderDetachCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachCancel) ProtoMessage() {}

func (x *CommanderDetachCancel) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachCancel.ProtoReflect.Descriptor instead.
func (*CommanderDetachCancel) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{3}
}

type CommanderDetachStartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommanderMode    bool                 `protobuf:"varint,1,opt,name=commander_mode,json=commanderMode,proto3" json:"commander_mode,omitempty"`
	CommanderUuid    string               `protobuf:"bytes,2,opt,name=commander_uuid,json=commanderUuid,proto3" json:"commander_uuid,omitempty"`
	LogWidth         int32                `protobuf:"varint,3,opt,name=log_width,json=logWidth,proto3" json:"log_width,omitempty"`
	ResourcesTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=resources_timeout,json=resourcesTimeout,proto3" json:"resources_timeout,omitempty"`
	DeckhouseTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=deckhouse_timeout,json=deckhouseTimeout,proto3" json:"deckhouse_timeout,omitempty"`
	CommonOptions    *OperationOptions    `protobuf:"bytes,10,opt,name=common_options,json=commonOptions,proto3" json:"common_options,omitempty"`
}

func (x *CommanderDetachStartOptions) Reset() {
	*x = CommanderDetachStartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachStartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachStartOptions) ProtoMessage() {}

func (x *CommanderDetachStartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachStartOptions.ProtoReflect.Descriptor instead.
func (*CommanderDetachStartOptions) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{4}
}

func (x *CommanderDetachStartOptions) GetCommanderMode() bool {
	if x != nil {
		return x.CommanderMode
	}
	return false
}

func (x *CommanderDetachStartOptions) GetCommanderUuid() string {
	if x != nil {
		return x.CommanderUuid
	}
	return ""
}

func (x *CommanderDetachStartOptions) GetLogWidth() int32 {
	if x != nil {
		return x.LogWidth
	}
	return 0
}

func (x *CommanderDetachStartOptions) GetResourcesTimeout() *durationpb.Duration {
	if x != nil {
		return x.ResourcesTimeout
	}
	return nil
}

func (x *CommanderDetachStartOptions) GetDeckhouseTimeout() *durationpb.Duration {
	if x != nil {
		return x.DeckhouseTimeout
	}
	return nil
}

func (x *CommanderDetachStartOptions) GetCommonOptions() *OperationOptions {
	if x != nil {
		return x.CommonOptions
	}
	return nil
}

type CommanderDetachResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *CommanderDetachResult) Reset() {
	*x = CommanderDetachResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commander_detach_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommanderDetachResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommanderDetachResult) ProtoMessage() {}

func (x *CommanderDetachResult) ProtoReflect() protoreflect.Message {
	mi := &file_commander_detach_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommanderDetachResult.ProtoReflect.Descriptor instead.
func (*CommanderDetachResult) Descriptor() ([]byte, []int) {
	return file_commander_detach_proto_rawDescGZIP(), []int{5}
}

func (x *CommanderDetachResult) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CommanderDetachResult) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_commander_detach_proto protoreflect.FileDescriptor

var file_commander_detach_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x16,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x68,
	0x63, 0x74, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xae,
	0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x68, 0x63,
	0x74, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x48, 0x00, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xa1, 0x04, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x63, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x20,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x15, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x15,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22, 0xd8, 0x02, 0x0a,
	0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f,
	0x67, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x6f, 0x67, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x46, 0x0a, 0x11, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x62, 0x2f, 0x64,
	0x68, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commander_detach_proto_rawDescOnce sync.Once
	file_commander_detach_proto_rawDescData = file_commander_detach_proto_rawDesc
)

func file_commander_detach_proto_rawDescGZIP() []byte {
	file_commander_detach_proto_rawDescOnce.Do(func() {
		file_commander_detach_proto_rawDescData = protoimpl.X.CompressGZIP(file_commander_detach_proto_rawDescData)
	})
	return file_commander_detach_proto_rawDescData
}

var file_commander_detach_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commander_detach_proto_goTypes = []interface{}{
	(*CommanderDetachRequest)(nil),      // 0: dhctl.CommanderDetachRequest
	(*CommanderDetachResponse)(nil),     // 1: dhctl.CommanderDetachResponse
	(*CommanderDetachStart)(nil),        // 2: dhctl.CommanderDetachStart
	(*CommanderDetachCancel)(nil),       // 3: dhctl.CommanderDetachCancel
	(*CommanderDetachStartOptions)(nil), // 4: dhctl.CommanderDetachStartOptions
	(*CommanderDetachResult)(nil),       // 5: dhctl.CommanderDetachResult
	(*Logs)(nil),                        // 6: dhctl.Logs
	(*Progress)(nil),                    // 7: dhctl.Progress
	(*structpb.Struct)(nil),             // 8: google.protobuf.Struct
	(*durationpb.Duration)(nil),         // 9: google.protobuf.Duration
	(*OperationOptions)(nil),            // 10: dhctl.OperationOptions
}
var file_commander_detach_proto_depIdxs = []int32{
	2,  // 0: dhctl.CommanderDetachRequest.start:type_name -> dhctl.CommanderDetachStart
	3,  // 1: dhctl.CommanderDetachRequest.cancel:type_name -> dhctl.CommanderDetachCancel
	5,  // 2: dhctl.CommanderDetachResponse.result:type_name -> dhctl.CommanderDetachResult
	6,  // 3: dhctl.CommanderDetachResponse.logs:type_name -> dhctl.Logs
	7,  // 4: dhctl.CommanderDetachResponse.progress:type_name -> dhctl.Progress
	8,  // 5: dhctl.CommanderDetachStart.create_resources_values:type_name -> google.protobuf.Struct
	8,  // 6: dhctl.CommanderDetachStart.delete_resources_values:type_name -> google.protobuf.Struct
	4,  // 7: dhctl.CommanderDetachStart.options:type_name -> dhctl.CommanderDetachStartOptions
	9,  // 8: dhctl.CommanderDetachStartOptions.resources_timeout:type_name -> google.protobuf.Duration
	9,  // 9: dhctl.CommanderDetachStartOptions.deckhouse_timeout:type_name -> google.protobuf.Duration
	10, // 10: dhctl.CommanderDetachStartOptions.common_options:type_name -> dhctl.OperationOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_commander_detach_proto_init() }
func file_commander_detach_proto_init() {
	if File_commander_detach_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commander_detach_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachRequest); i {
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
		file_commander_detach_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachResponse); i {
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
		file_commander_detach_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachStart); i {
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
		file_commander_detach_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachCancel); i {
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
		file_commander_detach_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachStartOptions); i {
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
		file_commander_detach_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommanderDetachResult); i {
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
	file_commander_detach_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommanderDetachRequest_Start)(nil),
		(*CommanderDetachRequest_Cancel)(nil),
	}
	file_commander_detach_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommanderDetachResponse_Result)(nil),
		(*CommanderDetachResponse_Logs)(nil),
		(*CommanderDetachResponse_Progress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commander_detach_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commander_detach_proto_goTypes,
		DependencyIndexes: file_commander_detach_proto_depIdxs,
		MessageInfos:      file_commander_detach_proto_msgTypes,
	}.Build()
	File_commander_detach_proto = out.File
	file_commander_detach_proto_rawDesc = nil
	file_commander_detach_proto_goTypes = nil
	file_commander_detach_proto_depIdxs = nil
}
