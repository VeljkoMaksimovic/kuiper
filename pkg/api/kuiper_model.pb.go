// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: kuiper_model.proto

package api

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

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Param) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NamedParamSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamSet []*Param `protobuf:"bytes,2,rep,name=paramSet,proto3" json:"paramSet,omitempty"`
}

func (x *NamedParamSet) Reset() {
	*x = NamedParamSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedParamSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedParamSet) ProtoMessage() {}

func (x *NamedParamSet) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedParamSet.ProtoReflect.Descriptor instead.
func (*NamedParamSet) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{1}
}

func (x *NamedParamSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedParamSet) GetParamSet() []*Param {
	if x != nil {
		return x.ParamSet
	}
	return nil
}

type NewStandaloneConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string   `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version      string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ParamSet     []*Param `protobuf:"bytes,4,rep,name=paramSet,proto3" json:"paramSet,omitempty"`
}

func (x *NewStandaloneConfig) Reset() {
	*x = NewStandaloneConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewStandaloneConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewStandaloneConfig) ProtoMessage() {}

func (x *NewStandaloneConfig) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewStandaloneConfig.ProtoReflect.Descriptor instead.
func (*NewStandaloneConfig) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{2}
}

func (x *NewStandaloneConfig) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *NewStandaloneConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewStandaloneConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NewStandaloneConfig) GetParamSet() []*Param {
	if x != nil {
		return x.ParamSet
	}
	return nil
}

type StandaloneConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string   `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version      string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt    string   `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ParamSet     []*Param `protobuf:"bytes,5,rep,name=paramSet,proto3" json:"paramSet,omitempty"`
}

func (x *StandaloneConfig) Reset() {
	*x = StandaloneConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandaloneConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandaloneConfig) ProtoMessage() {}

func (x *StandaloneConfig) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandaloneConfig.ProtoReflect.Descriptor instead.
func (*StandaloneConfig) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{3}
}

func (x *StandaloneConfig) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *StandaloneConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StandaloneConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StandaloneConfig) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *StandaloneConfig) GetParamSet() []*Param {
	if x != nil {
		return x.ParamSet
	}
	return nil
}

type NewConfigGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string           `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version      string           `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ParamSets    []*NamedParamSet `protobuf:"bytes,4,rep,name=paramSets,proto3" json:"paramSets,omitempty"`
}

func (x *NewConfigGroup) Reset() {
	*x = NewConfigGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewConfigGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewConfigGroup) ProtoMessage() {}

func (x *NewConfigGroup) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewConfigGroup.ProtoReflect.Descriptor instead.
func (*NewConfigGroup) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{4}
}

func (x *NewConfigGroup) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *NewConfigGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewConfigGroup) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NewConfigGroup) GetParamSets() []*NamedParamSet {
	if x != nil {
		return x.ParamSets
	}
	return nil
}

type ConfigGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string           `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version      string           `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt    string           `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ParamSets    []*NamedParamSet `protobuf:"bytes,5,rep,name=paramSets,proto3" json:"paramSets,omitempty"`
}

func (x *ConfigGroup) Reset() {
	*x = ConfigGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigGroup) ProtoMessage() {}

func (x *ConfigGroup) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigGroup.ProtoReflect.Descriptor instead.
func (*ConfigGroup) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigGroup) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ConfigGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigGroup) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ConfigGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ConfigGroup) GetParamSets() []*NamedParamSet {
	if x != nil {
		return x.ParamSets
	}
	return nil
}

type ConfigId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConfigId) Reset() {
	*x = ConfigId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigId) ProtoMessage() {}

func (x *ConfigId) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigId.ProtoReflect.Descriptor instead.
func (*ConfigId) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigId) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ConfigId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigId) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type PlacementTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Node       string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status     string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	AcceptedAt string `protobuf:"bytes,5,opt,name=acceptedAt,proto3" json:"acceptedAt,omitempty"`
	ResolvedAt string `protobuf:"bytes,6,opt,name=resolvedAt,proto3" json:"resolvedAt,omitempty"`
}

func (x *PlacementTask) Reset() {
	*x = PlacementTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementTask) ProtoMessage() {}

func (x *PlacementTask) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementTask.ProtoReflect.Descriptor instead.
func (*PlacementTask) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{7}
}

func (x *PlacementTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlacementTask) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PlacementTask) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PlacementTask) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PlacementTask) GetAcceptedAt() string {
	if x != nil {
		return x.AcceptedAt
	}
	return ""
}

func (x *PlacementTask) GetResolvedAt() string {
	if x != nil {
		return x.ResolvedAt
	}
	return ""
}

type Diff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Diff map[string]string `protobuf:"bytes,2,rep,name=diff,proto3" json:"diff,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Diff) Reset() {
	*x = Diff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diff) ProtoMessage() {}

func (x *Diff) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diff.ProtoReflect.Descriptor instead.
func (*Diff) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{8}
}

func (x *Diff) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Diff) GetDiff() map[string]string {
	if x != nil {
		return x.Diff
	}
	return nil
}

type Diffs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diffs []*Diff `protobuf:"bytes,1,rep,name=diffs,proto3" json:"diffs,omitempty"`
}

func (x *Diffs) Reset() {
	*x = Diffs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diffs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diffs) ProtoMessage() {}

func (x *Diffs) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diffs.ProtoReflect.Descriptor instead.
func (*Diffs) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{9}
}

func (x *Diffs) GetDiffs() []*Diff {
	if x != nil {
		return x.Diffs
	}
	return nil
}

type ApplyStandaloneConfigCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *StandaloneConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	TaskId    string            `protobuf:"bytes,2,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Namespace string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ApplyStandaloneConfigCommand) Reset() {
	*x = ApplyStandaloneConfigCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyStandaloneConfigCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyStandaloneConfigCommand) ProtoMessage() {}

func (x *ApplyStandaloneConfigCommand) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyStandaloneConfigCommand.ProtoReflect.Descriptor instead.
func (*ApplyStandaloneConfigCommand) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{10}
}

func (x *ApplyStandaloneConfigCommand) GetConfig() *StandaloneConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ApplyStandaloneConfigCommand) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ApplyStandaloneConfigCommand) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ApplyConfigGroupCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group     *ConfigGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	TaskId    string       `protobuf:"bytes,2,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Namespace string       `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ApplyConfigGroupCommand) Reset() {
	*x = ApplyConfigGroupCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuiper_model_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyConfigGroupCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyConfigGroupCommand) ProtoMessage() {}

func (x *ApplyConfigGroupCommand) ProtoReflect() protoreflect.Message {
	mi := &file_kuiper_model_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyConfigGroupCommand.ProtoReflect.Descriptor instead.
func (*ApplyConfigGroupCommand) Descriptor() ([]byte, []int) {
	return file_kuiper_model_proto_rawDescGZIP(), []int{11}
}

func (x *ApplyConfigGroupCommand) GetGroup() *ConfigGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *ApplyConfigGroupCommand) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ApplyConfigGroupCommand) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_kuiper_model_proto protoreflect.FileDescriptor

var file_kuiper_model_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6b, 0x75, 0x69, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x05, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x0d,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x13,
	0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x22,
	0xac, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x22, 0x96,
	0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x53, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74,
	0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7e, 0x0a, 0x04, 0x44, 0x69, 0x66, 0x66, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x66,
	0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x1a, 0x37, 0x0a, 0x09,
	0x44, 0x69, 0x66, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x05, 0x44, 0x69, 0x66, 0x66, 0x73, 0x12, 0x21,
	0x0a, 0x05, 0x64, 0x69, 0x66, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x52, 0x05, 0x64, 0x69, 0x66, 0x66,
	0x73, 0x22, 0x85, 0x01, 0x0a, 0x1c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x79, 0x0a, 0x17, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x31, 0x32, 0x73, 0x2f, 0x6b, 0x75, 0x69, 0x70, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kuiper_model_proto_rawDescOnce sync.Once
	file_kuiper_model_proto_rawDescData = file_kuiper_model_proto_rawDesc
)

func file_kuiper_model_proto_rawDescGZIP() []byte {
	file_kuiper_model_proto_rawDescOnce.Do(func() {
		file_kuiper_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuiper_model_proto_rawDescData)
	})
	return file_kuiper_model_proto_rawDescData
}

var file_kuiper_model_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_kuiper_model_proto_goTypes = []interface{}{
	(*Param)(nil),                        // 0: proto.Param
	(*NamedParamSet)(nil),                // 1: proto.NamedParamSet
	(*NewStandaloneConfig)(nil),          // 2: proto.NewStandaloneConfig
	(*StandaloneConfig)(nil),             // 3: proto.StandaloneConfig
	(*NewConfigGroup)(nil),               // 4: proto.NewConfigGroup
	(*ConfigGroup)(nil),                  // 5: proto.ConfigGroup
	(*ConfigId)(nil),                     // 6: proto.ConfigId
	(*PlacementTask)(nil),                // 7: proto.PlacementTask
	(*Diff)(nil),                         // 8: proto.Diff
	(*Diffs)(nil),                        // 9: proto.Diffs
	(*ApplyStandaloneConfigCommand)(nil), // 10: proto.ApplyStandaloneConfigCommand
	(*ApplyConfigGroupCommand)(nil),      // 11: proto.ApplyConfigGroupCommand
	nil,                                  // 12: proto.Diff.DiffEntry
}
var file_kuiper_model_proto_depIdxs = []int32{
	0,  // 0: proto.NamedParamSet.paramSet:type_name -> proto.Param
	0,  // 1: proto.NewStandaloneConfig.paramSet:type_name -> proto.Param
	0,  // 2: proto.StandaloneConfig.paramSet:type_name -> proto.Param
	1,  // 3: proto.NewConfigGroup.paramSets:type_name -> proto.NamedParamSet
	1,  // 4: proto.ConfigGroup.paramSets:type_name -> proto.NamedParamSet
	12, // 5: proto.Diff.diff:type_name -> proto.Diff.DiffEntry
	8,  // 6: proto.Diffs.diffs:type_name -> proto.Diff
	3,  // 7: proto.ApplyStandaloneConfigCommand.config:type_name -> proto.StandaloneConfig
	5,  // 8: proto.ApplyConfigGroupCommand.group:type_name -> proto.ConfigGroup
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_kuiper_model_proto_init() }
func file_kuiper_model_proto_init() {
	if File_kuiper_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuiper_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_kuiper_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedParamSet); i {
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
		file_kuiper_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewStandaloneConfig); i {
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
		file_kuiper_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandaloneConfig); i {
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
		file_kuiper_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewConfigGroup); i {
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
		file_kuiper_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigGroup); i {
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
		file_kuiper_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigId); i {
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
		file_kuiper_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementTask); i {
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
		file_kuiper_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diff); i {
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
		file_kuiper_model_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diffs); i {
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
		file_kuiper_model_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyStandaloneConfigCommand); i {
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
		file_kuiper_model_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyConfigGroupCommand); i {
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
			RawDescriptor: file_kuiper_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuiper_model_proto_goTypes,
		DependencyIndexes: file_kuiper_model_proto_depIdxs,
		MessageInfos:      file_kuiper_model_proto_msgTypes,
	}.Build()
	File_kuiper_model_proto = out.File
	file_kuiper_model_proto_rawDesc = nil
	file_kuiper_model_proto_goTypes = nil
	file_kuiper_model_proto_depIdxs = nil
}
