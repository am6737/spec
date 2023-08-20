// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: staticvolume/staticvolume.proto

package staticvolume

import (
	_ "github.com/container-storage-interface/spec/lib/go/csi"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace        string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	VolumeHandle     string            `protobuf:"bytes,3,opt,name=volume_handle,json=volumeHandle,proto3" json:"volume_handle,omitempty"`
	CapacityRange    *CapacityRange    `protobuf:"bytes,4,opt,name=capacity_range,json=capacityRange,proto3" json:"capacity_range,omitempty"`
	VolumeAttributes map[string]string `protobuf:"bytes,5,rep,name=volume_attributes,json=volumeAttributes,proto3" json:"volume_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AccessMode       []string          `protobuf:"bytes,6,rep,name=access_mode,json=accessMode,proto3" json:"access_mode,omitempty"`
	ReclaimPolicy    string            `protobuf:"bytes,7,opt,name=reclaim_policy,json=reclaimPolicy,proto3" json:"reclaim_policy,omitempty"`
}

func (x *PV) Reset() {
	*x = PV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PV) ProtoMessage() {}

func (x *PV) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PV.ProtoReflect.Descriptor instead.
func (*PV) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{0}
}

func (x *PV) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PV) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PV) GetVolumeHandle() string {
	if x != nil {
		return x.VolumeHandle
	}
	return ""
}

func (x *PV) GetCapacityRange() *CapacityRange {
	if x != nil {
		return x.CapacityRange
	}
	return nil
}

func (x *PV) GetVolumeAttributes() map[string]string {
	if x != nil {
		return x.VolumeAttributes
	}
	return nil
}

func (x *PV) GetAccessMode() []string {
	if x != nil {
		return x.AccessMode
	}
	return nil
}

func (x *PV) GetReclaimPolicy() string {
	if x != nil {
		return x.ReclaimPolicy
	}
	return ""
}

type PVC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string         `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CapacityRange *CapacityRange `protobuf:"bytes,3,opt,name=capacity_range,json=capacityRange,proto3" json:"capacity_range,omitempty"`
	AccessMode    []string       `protobuf:"bytes,4,rep,name=access_mode,json=accessMode,proto3" json:"access_mode,omitempty"`
	ReadOnly      bool           `protobuf:"varint,5,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *PVC) Reset() {
	*x = PVC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PVC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PVC) ProtoMessage() {}

func (x *PVC) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PVC.ProtoReflect.Descriptor instead.
func (*PVC) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{1}
}

func (x *PVC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PVC) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PVC) GetCapacityRange() *CapacityRange {
	if x != nil {
		return x.CapacityRange
	}
	return nil
}

func (x *PVC) GetAccessMode() []string {
	if x != nil {
		return x.AccessMode
	}
	return nil
}

func (x *PVC) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

type SC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Provisioner string            `protobuf:"bytes,2,opt,name=provisioner,proto3" json:"provisioner,omitempty"`
	ClusterId   string            `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Parameters  map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SC) Reset() {
	*x = SC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC) ProtoMessage() {}

func (x *SC) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC.ProtoReflect.Descriptor instead.
func (*SC) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{2}
}

func (x *SC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SC) GetProvisioner() string {
	if x != nil {
		return x.Provisioner
	}
	return ""
}

func (x *SC) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *SC) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type CreateVolumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CapacityRange *CapacityRange    `protobuf:"bytes,2,opt,name=capacity_range,json=capacityRange,proto3" json:"capacity_range,omitempty"`
	Parameters    map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Secrets       map[string]string `protobuf:"bytes,4,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pv            *PV               `protobuf:"bytes,5,opt,name=pv,proto3" json:"pv,omitempty"`
	Pvc           *PVC              `protobuf:"bytes,6,opt,name=pvc,proto3" json:"pvc,omitempty"`
	Sc            *SC               `protobuf:"bytes,7,opt,name=sc,proto3" json:"sc,omitempty"`
}

func (x *CreateVolumeRequest) Reset() {
	*x = CreateVolumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeRequest) ProtoMessage() {}

func (x *CreateVolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeRequest.ProtoReflect.Descriptor instead.
func (*CreateVolumeRequest) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVolumeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVolumeRequest) GetCapacityRange() *CapacityRange {
	if x != nil {
		return x.CapacityRange
	}
	return nil
}

func (x *CreateVolumeRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *CreateVolumeRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *CreateVolumeRequest) GetPv() *PV {
	if x != nil {
		return x.Pv
	}
	return nil
}

func (x *CreateVolumeRequest) GetPvc() *PVC {
	if x != nil {
		return x.Pvc
	}
	return nil
}

func (x *CreateVolumeRequest) GetSc() *SC {
	if x != nil {
		return x.Sc
	}
	return nil
}

type CreateVolumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeId string `protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
}

func (x *CreateVolumeResponse) Reset() {
	*x = CreateVolumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeResponse) ProtoMessage() {}

func (x *CreateVolumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeResponse.ProtoReflect.Descriptor instead.
func (*CreateVolumeResponse) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVolumeResponse) GetVolumeId() string {
	if x != nil {
		return x.VolumeId
	}
	return ""
}

type DeleteVolumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the volume to be deprovisioned.
	// This field is REQUIRED.
	VolumeId string `protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	// Secrets required by plugin to complete volume deletion request.
	// This field is OPTIONAL. Refer to the `Secrets Requirements`
	// section on how to use this field.
	Secrets map[string]string `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeleteVolumeRequest) Reset() {
	*x = DeleteVolumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVolumeRequest) ProtoMessage() {}

func (x *DeleteVolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVolumeRequest.ProtoReflect.Descriptor instead.
func (*DeleteVolumeRequest) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteVolumeRequest) GetVolumeId() string {
	if x != nil {
		return x.VolumeId
	}
	return ""
}

func (x *DeleteVolumeRequest) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type DeleteVolumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteVolumeResponse) Reset() {
	*x = DeleteVolumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVolumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVolumeResponse) ProtoMessage() {}

func (x *DeleteVolumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVolumeResponse.ProtoReflect.Descriptor instead.
func (*DeleteVolumeResponse) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{6}
}

// The capacity of the storage space in bytes. To specify an exact size,
// `required_bytes` and `limit_bytes` SHALL be set to the same value. At
// least one of the these fields MUST be specified.
type CapacityRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Volume MUST be at least this big. This field is OPTIONAL.
	// A value of 0 is equal to an unspecified field value.
	// The value of this field MUST NOT be negative.
	RequiredBytes int64 `protobuf:"varint,1,opt,name=required_bytes,json=requiredBytes,proto3" json:"required_bytes,omitempty"`
	// Volume MUST not be bigger than this. This field is OPTIONAL.
	// A value of 0 is equal to an unspecified field value.
	// The value of this field MUST NOT be negative.
	LimitBytes int64 `protobuf:"varint,2,opt,name=limit_bytes,json=limitBytes,proto3" json:"limit_bytes,omitempty"`
}

func (x *CapacityRange) Reset() {
	*x = CapacityRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staticvolume_staticvolume_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapacityRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapacityRange) ProtoMessage() {}

func (x *CapacityRange) ProtoReflect() protoreflect.Message {
	mi := &file_staticvolume_staticvolume_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapacityRange.ProtoReflect.Descriptor instead.
func (*CapacityRange) Descriptor() ([]byte, []int) {
	return file_staticvolume_staticvolume_proto_rawDescGZIP(), []int{7}
}

func (x *CapacityRange) GetRequiredBytes() int64 {
	if x != nil {
		return x.RequiredBytes
	}
	return 0
}

func (x *CapacityRange) GetLimitBytes() int64 {
	if x != nil {
		return x.LimitBytes
	}
	return 0
}

var File_staticvolume_staticvolume_proto protoreflect.FileDescriptor

var file_staticvolume_staticvolume_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x73, 0x69, 0x2f, 0x63, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x02, 0x50, 0x56, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x42, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x50,
	0x56, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x1a, 0x43, 0x0a, 0x15, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x03, 0x50, 0x56, 0x43, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xda, 0x01, 0x0a, 0x02, 0x53, 0x43, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x53, 0x43, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf4, 0x03, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x48, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0d, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x48,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x02, 0x70, 0x76, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x50, 0x56, 0x52, 0x02, 0x70, 0x76, 0x12, 0x23, 0x0a, 0x03, 0x70, 0x76,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x50, 0x56, 0x43, 0x52, 0x03, 0x70, 0x76, 0x63, 0x12,
	0x20, 0x0a, 0x02, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x53, 0x43, 0x52, 0x02, 0x73,
	0x63, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49,
	0x64, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0x98, 0x42, 0x01, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x55, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staticvolume_staticvolume_proto_rawDescOnce sync.Once
	file_staticvolume_staticvolume_proto_rawDescData = file_staticvolume_staticvolume_proto_rawDesc
)

func file_staticvolume_staticvolume_proto_rawDescGZIP() []byte {
	file_staticvolume_staticvolume_proto_rawDescOnce.Do(func() {
		file_staticvolume_staticvolume_proto_rawDescData = protoimpl.X.CompressGZIP(file_staticvolume_staticvolume_proto_rawDescData)
	})
	return file_staticvolume_staticvolume_proto_rawDescData
}

var file_staticvolume_staticvolume_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_staticvolume_staticvolume_proto_goTypes = []interface{}{
	(*PV)(nil),                   // 0: staticvolume.PV
	(*PVC)(nil),                  // 1: staticvolume.PVC
	(*SC)(nil),                   // 2: staticvolume.SC
	(*CreateVolumeRequest)(nil),  // 3: staticvolume.CreateVolumeRequest
	(*CreateVolumeResponse)(nil), // 4: staticvolume.CreateVolumeResponse
	(*DeleteVolumeRequest)(nil),  // 5: staticvolume.DeleteVolumeRequest
	(*DeleteVolumeResponse)(nil), // 6: staticvolume.DeleteVolumeResponse
	(*CapacityRange)(nil),        // 7: staticvolume.CapacityRange
	nil,                          // 8: staticvolume.PV.VolumeAttributesEntry
	nil,                          // 9: staticvolume.SC.ParametersEntry
	nil,                          // 10: staticvolume.CreateVolumeRequest.ParametersEntry
	nil,                          // 11: staticvolume.CreateVolumeRequest.SecretsEntry
	nil,                          // 12: staticvolume.DeleteVolumeRequest.SecretsEntry
}
var file_staticvolume_staticvolume_proto_depIdxs = []int32{
	7,  // 0: staticvolume.PV.capacity_range:type_name -> staticvolume.CapacityRange
	8,  // 1: staticvolume.PV.volume_attributes:type_name -> staticvolume.PV.VolumeAttributesEntry
	7,  // 2: staticvolume.PVC.capacity_range:type_name -> staticvolume.CapacityRange
	9,  // 3: staticvolume.SC.parameters:type_name -> staticvolume.SC.ParametersEntry
	7,  // 4: staticvolume.CreateVolumeRequest.capacity_range:type_name -> staticvolume.CapacityRange
	10, // 5: staticvolume.CreateVolumeRequest.parameters:type_name -> staticvolume.CreateVolumeRequest.ParametersEntry
	11, // 6: staticvolume.CreateVolumeRequest.secrets:type_name -> staticvolume.CreateVolumeRequest.SecretsEntry
	0,  // 7: staticvolume.CreateVolumeRequest.pv:type_name -> staticvolume.PV
	1,  // 8: staticvolume.CreateVolumeRequest.pvc:type_name -> staticvolume.PVC
	2,  // 9: staticvolume.CreateVolumeRequest.sc:type_name -> staticvolume.SC
	12, // 10: staticvolume.DeleteVolumeRequest.secrets:type_name -> staticvolume.DeleteVolumeRequest.SecretsEntry
	3,  // 11: staticvolume.Controller.CreateVolume:input_type -> staticvolume.CreateVolumeRequest
	5,  // 12: staticvolume.Controller.DeleteVolume:input_type -> staticvolume.DeleteVolumeRequest
	4,  // 13: staticvolume.Controller.CreateVolume:output_type -> staticvolume.CreateVolumeResponse
	6,  // 14: staticvolume.Controller.DeleteVolume:output_type -> staticvolume.DeleteVolumeResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_staticvolume_staticvolume_proto_init() }
func file_staticvolume_staticvolume_proto_init() {
	if File_staticvolume_staticvolume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staticvolume_staticvolume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PV); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PVC); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeRequest); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeResponse); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVolumeRequest); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVolumeResponse); i {
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
		file_staticvolume_staticvolume_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapacityRange); i {
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
			RawDescriptor: file_staticvolume_staticvolume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staticvolume_staticvolume_proto_goTypes,
		DependencyIndexes: file_staticvolume_staticvolume_proto_depIdxs,
		MessageInfos:      file_staticvolume_staticvolume_proto_msgTypes,
	}.Build()
	File_staticvolume_staticvolume_proto = out.File
	file_staticvolume_staticvolume_proto_rawDesc = nil
	file_staticvolume_staticvolume_proto_goTypes = nil
	file_staticvolume_staticvolume_proto_depIdxs = nil
}
