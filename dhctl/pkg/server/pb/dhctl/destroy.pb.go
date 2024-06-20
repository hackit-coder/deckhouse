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
// source: destroy.proto

package dhctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DestroyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*DestroyRequest_Start
	//	*DestroyRequest_Continue
	Message isDestroyRequest_Message `protobuf_oneof:"message"`
}

func (x *DestroyRequest) Reset() {
	*x = DestroyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyRequest) ProtoMessage() {}

func (x *DestroyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyRequest.ProtoReflect.Descriptor instead.
func (*DestroyRequest) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{0}
}

func (m *DestroyRequest) GetMessage() isDestroyRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *DestroyRequest) GetStart() *DestroyStart {
	if x, ok := x.GetMessage().(*DestroyRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *DestroyRequest) GetContinue() *DestroyContinue {
	if x, ok := x.GetMessage().(*DestroyRequest_Continue); ok {
		return x.Continue
	}
	return nil
}

type isDestroyRequest_Message interface {
	isDestroyRequest_Message()
}

type DestroyRequest_Start struct {
	Start *DestroyStart `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type DestroyRequest_Continue struct {
	Continue *DestroyContinue `protobuf:"bytes,2,opt,name=continue,proto3,oneof"`
}

func (*DestroyRequest_Start) isDestroyRequest_Message() {}

func (*DestroyRequest_Continue) isDestroyRequest_Message() {}

type DestroyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*DestroyResponse_Result
	//	*DestroyResponse_PhaseEnd
	//	*DestroyResponse_Logs
	Message isDestroyResponse_Message `protobuf_oneof:"message"`
}

func (x *DestroyResponse) Reset() {
	*x = DestroyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyResponse) ProtoMessage() {}

func (x *DestroyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyResponse.ProtoReflect.Descriptor instead.
func (*DestroyResponse) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{1}
}

func (m *DestroyResponse) GetMessage() isDestroyResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *DestroyResponse) GetResult() *DestroyResult {
	if x, ok := x.GetMessage().(*DestroyResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *DestroyResponse) GetPhaseEnd() *DestroyPhaseEnd {
	if x, ok := x.GetMessage().(*DestroyResponse_PhaseEnd); ok {
		return x.PhaseEnd
	}
	return nil
}

func (x *DestroyResponse) GetLogs() *Logs {
	if x, ok := x.GetMessage().(*DestroyResponse_Logs); ok {
		return x.Logs
	}
	return nil
}

type isDestroyResponse_Message interface {
	isDestroyResponse_Message()
}

type DestroyResponse_Result struct {
	Result *DestroyResult `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type DestroyResponse_PhaseEnd struct {
	PhaseEnd *DestroyPhaseEnd `protobuf:"bytes,2,opt,name=phase_end,json=phaseEnd,proto3,oneof"`
}

type DestroyResponse_Logs struct {
	Logs *Logs `protobuf:"bytes,3,opt,name=logs,proto3,oneof"`
}

func (*DestroyResponse_Result) isDestroyResponse_Message() {}

func (*DestroyResponse_PhaseEnd) isDestroyResponse_Message() {}

func (*DestroyResponse_Logs) isDestroyResponse_Message() {}

type DestroyStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfig              string               `protobuf:"bytes,1,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	InitConfig                    string               `protobuf:"bytes,2,opt,name=init_config,json=initConfig,proto3" json:"init_config,omitempty"`
	ClusterConfig                 string               `protobuf:"bytes,3,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	ProviderSpecificClusterConfig string               `protobuf:"bytes,4,opt,name=provider_specific_cluster_config,json=providerSpecificClusterConfig,proto3" json:"provider_specific_cluster_config,omitempty"`
	State                         string               `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Options                       *DestroyStartOptions `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *DestroyStart) Reset() {
	*x = DestroyStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyStart) ProtoMessage() {}

func (x *DestroyStart) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyStart.ProtoReflect.Descriptor instead.
func (*DestroyStart) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{2}
}

func (x *DestroyStart) GetConnectionConfig() string {
	if x != nil {
		return x.ConnectionConfig
	}
	return ""
}

func (x *DestroyStart) GetInitConfig() string {
	if x != nil {
		return x.InitConfig
	}
	return ""
}

func (x *DestroyStart) GetClusterConfig() string {
	if x != nil {
		return x.ClusterConfig
	}
	return ""
}

func (x *DestroyStart) GetProviderSpecificClusterConfig() string {
	if x != nil {
		return x.ProviderSpecificClusterConfig
	}
	return ""
}

func (x *DestroyStart) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *DestroyStart) GetOptions() *DestroyStartOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type DestroyPhaseEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompletedPhase      string            `protobuf:"bytes,1,opt,name=completed_phase,json=completedPhase,proto3" json:"completed_phase,omitempty"`
	CompletedPhaseState map[string][]byte `protobuf:"bytes,2,rep,name=completed_phase_state,json=completedPhaseState,proto3" json:"completed_phase_state,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NextPhase           string            `protobuf:"bytes,3,opt,name=next_phase,json=nextPhase,proto3" json:"next_phase,omitempty"`
	NextPhaseCritical   bool              `protobuf:"varint,4,opt,name=next_phase_critical,json=nextPhaseCritical,proto3" json:"next_phase_critical,omitempty"`
}

func (x *DestroyPhaseEnd) Reset() {
	*x = DestroyPhaseEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyPhaseEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyPhaseEnd) ProtoMessage() {}

func (x *DestroyPhaseEnd) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyPhaseEnd.ProtoReflect.Descriptor instead.
func (*DestroyPhaseEnd) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{3}
}

func (x *DestroyPhaseEnd) GetCompletedPhase() string {
	if x != nil {
		return x.CompletedPhase
	}
	return ""
}

func (x *DestroyPhaseEnd) GetCompletedPhaseState() map[string][]byte {
	if x != nil {
		return x.CompletedPhaseState
	}
	return nil
}

func (x *DestroyPhaseEnd) GetNextPhase() string {
	if x != nil {
		return x.NextPhase
	}
	return ""
}

func (x *DestroyPhaseEnd) GetNextPhaseCritical() bool {
	if x != nil {
		return x.NextPhaseCritical
	}
	return false
}

type DestroyContinue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Continue Continue `protobuf:"varint,1,opt,name=continue,proto3,enum=dhctl.Continue" json:"continue,omitempty"`
	Err      string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DestroyContinue) Reset() {
	*x = DestroyContinue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyContinue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyContinue) ProtoMessage() {}

func (x *DestroyContinue) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyContinue.ProtoReflect.Descriptor instead.
func (*DestroyContinue) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{4}
}

func (x *DestroyContinue) GetContinue() Continue {
	if x != nil {
		return x.Continue
	}
	return Continue_CONTINUE_UNSPECIFIED
}

func (x *DestroyContinue) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type DestroyStartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommanderMode    bool                 `protobuf:"varint,1,opt,name=commander_mode,json=commanderMode,proto3" json:"commander_mode,omitempty"`
	CommanderUuid    string               `protobuf:"bytes,2,opt,name=commander_uuid,json=commanderUuid,proto3" json:"commander_uuid,omitempty"`
	LogWidth         int32                `protobuf:"varint,3,opt,name=log_width,json=logWidth,proto3" json:"log_width,omitempty"`
	ResourcesTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=resources_timeout,json=resourcesTimeout,proto3" json:"resources_timeout,omitempty"`
	DeckhouseTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=deckhouse_timeout,json=deckhouseTimeout,proto3" json:"deckhouse_timeout,omitempty"`
}

func (x *DestroyStartOptions) Reset() {
	*x = DestroyStartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyStartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyStartOptions) ProtoMessage() {}

func (x *DestroyStartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyStartOptions.ProtoReflect.Descriptor instead.
func (*DestroyStartOptions) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{5}
}

func (x *DestroyStartOptions) GetCommanderMode() bool {
	if x != nil {
		return x.CommanderMode
	}
	return false
}

func (x *DestroyStartOptions) GetCommanderUuid() string {
	if x != nil {
		return x.CommanderUuid
	}
	return ""
}

func (x *DestroyStartOptions) GetLogWidth() int32 {
	if x != nil {
		return x.LogWidth
	}
	return 0
}

func (x *DestroyStartOptions) GetResourcesTimeout() *durationpb.Duration {
	if x != nil {
		return x.ResourcesTimeout
	}
	return nil
}

func (x *DestroyStartOptions) GetDeckhouseTimeout() *durationpb.Duration {
	if x != nil {
		return x.DeckhouseTimeout
	}
	return nil
}

type DestroyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DestroyResult) Reset() {
	*x = DestroyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_destroy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyResult) ProtoMessage() {}

func (x *DestroyResult) ProtoReflect() protoreflect.Message {
	mi := &file_destroy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyResult.ProtoReflect.Descriptor instead.
func (*DestroyResult) Descriptor() ([]byte, []int) {
	return file_destroy_proto_rawDescGZIP(), []int{6}
}

func (x *DestroyResult) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *DestroyResult) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_destroy_proto protoreflect.FileDescriptor

var file_destroy_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c,
	0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x68,
	0x63, 0x74, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x45, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x70, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x98, 0x02,
	0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2b,
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb6, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x1a, 0x46, 0x0a, 0x18, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x50, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x22, 0x90, 0x02, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x46,
	0x0a, 0x11, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42,
	0x0a, 0x5a, 0x08, 0x70, 0x62, 0x2f, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_destroy_proto_rawDescOnce sync.Once
	file_destroy_proto_rawDescData = file_destroy_proto_rawDesc
)

func file_destroy_proto_rawDescGZIP() []byte {
	file_destroy_proto_rawDescOnce.Do(func() {
		file_destroy_proto_rawDescData = protoimpl.X.CompressGZIP(file_destroy_proto_rawDescData)
	})
	return file_destroy_proto_rawDescData
}

var file_destroy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_destroy_proto_goTypes = []interface{}{
	(*DestroyRequest)(nil),      // 0: dhctl.DestroyRequest
	(*DestroyResponse)(nil),     // 1: dhctl.DestroyResponse
	(*DestroyStart)(nil),        // 2: dhctl.DestroyStart
	(*DestroyPhaseEnd)(nil),     // 3: dhctl.DestroyPhaseEnd
	(*DestroyContinue)(nil),     // 4: dhctl.DestroyContinue
	(*DestroyStartOptions)(nil), // 5: dhctl.DestroyStartOptions
	(*DestroyResult)(nil),       // 6: dhctl.DestroyResult
	nil,                         // 7: dhctl.DestroyPhaseEnd.CompletedPhaseStateEntry
	(*Logs)(nil),                // 8: dhctl.Logs
	(Continue)(0),               // 9: dhctl.Continue
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_destroy_proto_depIdxs = []int32{
	2,  // 0: dhctl.DestroyRequest.start:type_name -> dhctl.DestroyStart
	4,  // 1: dhctl.DestroyRequest.continue:type_name -> dhctl.DestroyContinue
	6,  // 2: dhctl.DestroyResponse.result:type_name -> dhctl.DestroyResult
	3,  // 3: dhctl.DestroyResponse.phase_end:type_name -> dhctl.DestroyPhaseEnd
	8,  // 4: dhctl.DestroyResponse.logs:type_name -> dhctl.Logs
	5,  // 5: dhctl.DestroyStart.options:type_name -> dhctl.DestroyStartOptions
	7,  // 6: dhctl.DestroyPhaseEnd.completed_phase_state:type_name -> dhctl.DestroyPhaseEnd.CompletedPhaseStateEntry
	9,  // 7: dhctl.DestroyContinue.continue:type_name -> dhctl.Continue
	10, // 8: dhctl.DestroyStartOptions.resources_timeout:type_name -> google.protobuf.Duration
	10, // 9: dhctl.DestroyStartOptions.deckhouse_timeout:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_destroy_proto_init() }
func file_destroy_proto_init() {
	if File_destroy_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_destroy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyRequest); i {
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
		file_destroy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyResponse); i {
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
		file_destroy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyStart); i {
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
		file_destroy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyPhaseEnd); i {
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
		file_destroy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyContinue); i {
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
		file_destroy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyStartOptions); i {
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
		file_destroy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyResult); i {
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
	file_destroy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DestroyRequest_Start)(nil),
		(*DestroyRequest_Continue)(nil),
	}
	file_destroy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DestroyResponse_Result)(nil),
		(*DestroyResponse_PhaseEnd)(nil),
		(*DestroyResponse_Logs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_destroy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_destroy_proto_goTypes,
		DependencyIndexes: file_destroy_proto_depIdxs,
		MessageInfos:      file_destroy_proto_msgTypes,
	}.Build()
	File_destroy_proto = out.File
	file_destroy_proto_rawDesc = nil
	file_destroy_proto_goTypes = nil
	file_destroy_proto_depIdxs = nil
}
