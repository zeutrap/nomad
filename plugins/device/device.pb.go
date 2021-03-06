// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/hashicorp/nomad/plugins/device/device.proto

package device

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DetectedDevices is the set of devices that the device plugin has
// detected and is exposing
type DetectedDevices struct {
	// vendor is the name of the vendor of the device
	Vendor string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// device_type is the type of the device (gpu, fpga, etc).
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// device_name is the name of the device.
	DeviceName string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// devices is the set of devices detected by the plugin.
	Devices []*DetectedDevice `protobuf:"bytes,4,rep,name=devices,proto3" json:"devices,omitempty"`
	// node_attributes allows adding node attributes to be used for
	// constraints or affinities.
	NodeAttributes       map[string]string `protobuf:"bytes,5,rep,name=node_attributes,json=nodeAttributes,proto3" json:"node_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DetectedDevices) Reset()         { *m = DetectedDevices{} }
func (m *DetectedDevices) String() string { return proto.CompactTextString(m) }
func (*DetectedDevices) ProtoMessage()    {}
func (*DetectedDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{0}
}
func (m *DetectedDevices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedDevices.Unmarshal(m, b)
}
func (m *DetectedDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedDevices.Marshal(b, m, deterministic)
}
func (dst *DetectedDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedDevices.Merge(dst, src)
}
func (m *DetectedDevices) XXX_Size() int {
	return xxx_messageInfo_DetectedDevices.Size(m)
}
func (m *DetectedDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedDevices.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedDevices proto.InternalMessageInfo

func (m *DetectedDevices) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DetectedDevices) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DetectedDevices) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *DetectedDevices) GetDevices() []*DetectedDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *DetectedDevices) GetNodeAttributes() map[string]string {
	if m != nil {
		return m.NodeAttributes
	}
	return nil
}

// DetectedDevice is a single detected device.
type DetectedDevice struct {
	// ID is the ID of the device. This ID is used during
	// allocation.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Health of the device.
	Healthy bool `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
	// health_description allows the device plugin to optionally
	// annotate the health field with a human readable reason.
	HealthDescription string `protobuf:"bytes,3,opt,name=health_description,json=healthDescription,proto3" json:"health_description,omitempty"`
	// pci_bus_id is the PCI bus ID for the device. If reported, it
	// allows Nomad to make NUMA aware optimizations.
	PciBusId             string   `protobuf:"bytes,4,opt,name=pci_bus_id,json=pciBusId,proto3" json:"pci_bus_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectedDevice) Reset()         { *m = DetectedDevice{} }
func (m *DetectedDevice) String() string { return proto.CompactTextString(m) }
func (*DetectedDevice) ProtoMessage()    {}
func (*DetectedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{1}
}
func (m *DetectedDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedDevice.Unmarshal(m, b)
}
func (m *DetectedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedDevice.Marshal(b, m, deterministic)
}
func (dst *DetectedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedDevice.Merge(dst, src)
}
func (m *DetectedDevice) XXX_Size() int {
	return xxx_messageInfo_DetectedDevice.Size(m)
}
func (m *DetectedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedDevice proto.InternalMessageInfo

func (m *DetectedDevice) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DetectedDevice) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

func (m *DetectedDevice) GetHealthDescription() string {
	if m != nil {
		return m.HealthDescription
	}
	return ""
}

func (m *DetectedDevice) GetPciBusId() string {
	if m != nil {
		return m.PciBusId
	}
	return ""
}

// ReserveRequest is used to ask the device driver for information on
// how to allocate the requested devices.
type ReserveRequest struct {
	// device_ids are the requested devices.
	DeviceIds            []string `protobuf:"bytes,1,rep,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveRequest) Reset()         { *m = ReserveRequest{} }
func (m *ReserveRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveRequest) ProtoMessage()    {}
func (*ReserveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{2}
}
func (m *ReserveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveRequest.Unmarshal(m, b)
}
func (m *ReserveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveRequest.Marshal(b, m, deterministic)
}
func (dst *ReserveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveRequest.Merge(dst, src)
}
func (m *ReserveRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveRequest.Size(m)
}
func (m *ReserveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveRequest proto.InternalMessageInfo

func (m *ReserveRequest) GetDeviceIds() []string {
	if m != nil {
		return m.DeviceIds
	}
	return nil
}

// ReserveResponse informs Nomad how to expose the requested devices
// to the the task.
type ReserveResponse struct {
	// container_res contains information on how to mount the device
	// into a task isolated using container technologies (where the
	// host is shared)
	ContainerRes         *ContainerReservation `protobuf:"bytes,1,opt,name=container_res,json=containerRes,proto3" json:"container_res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ReserveResponse) Reset()         { *m = ReserveResponse{} }
func (m *ReserveResponse) String() string { return proto.CompactTextString(m) }
func (*ReserveResponse) ProtoMessage()    {}
func (*ReserveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{3}
}
func (m *ReserveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveResponse.Unmarshal(m, b)
}
func (m *ReserveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveResponse.Marshal(b, m, deterministic)
}
func (dst *ReserveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveResponse.Merge(dst, src)
}
func (m *ReserveResponse) XXX_Size() int {
	return xxx_messageInfo_ReserveResponse.Size(m)
}
func (m *ReserveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveResponse proto.InternalMessageInfo

func (m *ReserveResponse) GetContainerRes() *ContainerReservation {
	if m != nil {
		return m.ContainerRes
	}
	return nil
}

// ContainerReservation returns how to mount the device into a
// container that shares the host OS.
type ContainerReservation struct {
	// List of environment variable to be set
	Envs map[string]string `protobuf:"bytes,1,rep,name=envs,proto3" json:"envs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Mounts for the task.
	Mounts []*Mount `protobuf:"bytes,2,rep,name=mounts,proto3" json:"mounts,omitempty"`
	// Devices for the task.
	Devices              []*DeviceSpec `protobuf:"bytes,3,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContainerReservation) Reset()         { *m = ContainerReservation{} }
func (m *ContainerReservation) String() string { return proto.CompactTextString(m) }
func (*ContainerReservation) ProtoMessage()    {}
func (*ContainerReservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{4}
}
func (m *ContainerReservation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerReservation.Unmarshal(m, b)
}
func (m *ContainerReservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerReservation.Marshal(b, m, deterministic)
}
func (dst *ContainerReservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerReservation.Merge(dst, src)
}
func (m *ContainerReservation) XXX_Size() int {
	return xxx_messageInfo_ContainerReservation.Size(m)
}
func (m *ContainerReservation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerReservation.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerReservation proto.InternalMessageInfo

func (m *ContainerReservation) GetEnvs() map[string]string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *ContainerReservation) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *ContainerReservation) GetDevices() []*DeviceSpec {
	if m != nil {
		return m.Devices
	}
	return nil
}

// Mount specifies a host volume to mount into a task.
// where device library or tools are installed on host and task
type Mount struct {
	// Path of the mount within the task.
	TaskPath string `protobuf:"bytes,1,opt,name=task_path,json=taskPath,proto3" json:"task_path,omitempty"`
	// Path of the mount on the host.
	HostPath string `protobuf:"bytes,2,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	// If set, the mount is read-only.
	ReadOnly             bool     `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mount) Reset()         { *m = Mount{} }
func (m *Mount) String() string { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()    {}
func (*Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{5}
}
func (m *Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mount.Unmarshal(m, b)
}
func (m *Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mount.Marshal(b, m, deterministic)
}
func (dst *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(dst, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func (m *Mount) GetTaskPath() string {
	if m != nil {
		return m.TaskPath
	}
	return ""
}

func (m *Mount) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *Mount) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

// DeviceSpec specifies a host device to mount into a task.
type DeviceSpec struct {
	// Path of the device within the task.
	TaskPath string `protobuf:"bytes,1,opt,name=task_path,json=taskPath,proto3" json:"task_path,omitempty"`
	// Path of the device on the host.
	HostPath string `protobuf:"bytes,2,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	// Cgroups permissions of the device, candidates are one or more of
	// * r - allows task to read from the specified device.
	// * w - allows task to write to the specified device.
	// * m - allows task to create device files that do not yet exist
	Permissions          string   `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSpec) Reset()         { *m = DeviceSpec{} }
func (m *DeviceSpec) String() string { return proto.CompactTextString(m) }
func (*DeviceSpec) ProtoMessage()    {}
func (*DeviceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_13acb8ec0117c3b0, []int{6}
}
func (m *DeviceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSpec.Unmarshal(m, b)
}
func (m *DeviceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSpec.Marshal(b, m, deterministic)
}
func (dst *DeviceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSpec.Merge(dst, src)
}
func (m *DeviceSpec) XXX_Size() int {
	return xxx_messageInfo_DeviceSpec.Size(m)
}
func (m *DeviceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSpec proto.InternalMessageInfo

func (m *DeviceSpec) GetTaskPath() string {
	if m != nil {
		return m.TaskPath
	}
	return ""
}

func (m *DeviceSpec) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *DeviceSpec) GetPermissions() string {
	if m != nil {
		return m.Permissions
	}
	return ""
}

func init() {
	proto.RegisterType((*DetectedDevices)(nil), "hashicorp.nomad.plugins.device.DetectedDevices")
	proto.RegisterMapType((map[string]string)(nil), "hashicorp.nomad.plugins.device.DetectedDevices.NodeAttributesEntry")
	proto.RegisterType((*DetectedDevice)(nil), "hashicorp.nomad.plugins.device.DetectedDevice")
	proto.RegisterType((*ReserveRequest)(nil), "hashicorp.nomad.plugins.device.ReserveRequest")
	proto.RegisterType((*ReserveResponse)(nil), "hashicorp.nomad.plugins.device.ReserveResponse")
	proto.RegisterType((*ContainerReservation)(nil), "hashicorp.nomad.plugins.device.ContainerReservation")
	proto.RegisterMapType((map[string]string)(nil), "hashicorp.nomad.plugins.device.ContainerReservation.EnvsEntry")
	proto.RegisterType((*Mount)(nil), "hashicorp.nomad.plugins.device.Mount")
	proto.RegisterType((*DeviceSpec)(nil), "hashicorp.nomad.plugins.device.DeviceSpec")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DevicePluginClient is the client API for DevicePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicePluginClient interface {
	// Fingerprint allows the device plugin to return a set of
	// detected devices and provide a mechanism to update the state of
	// the device.
	Fingerprint(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DevicePlugin_FingerprintClient, error)
	// Reserve is called by the client before starting an allocation
	// that requires access to the plugin’s devices. The plugin can use
	// this to run any setup steps and provides the mounting details to
	// the Nomad client
	Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error)
}

type devicePluginClient struct {
	cc *grpc.ClientConn
}

func NewDevicePluginClient(cc *grpc.ClientConn) DevicePluginClient {
	return &devicePluginClient{cc}
}

func (c *devicePluginClient) Fingerprint(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DevicePlugin_FingerprintClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DevicePlugin_serviceDesc.Streams[0], "/hashicorp.nomad.plugins.device.DevicePlugin/Fingerprint", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicePluginFingerprintClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DevicePlugin_FingerprintClient interface {
	Recv() (*DetectedDevices, error)
	grpc.ClientStream
}

type devicePluginFingerprintClient struct {
	grpc.ClientStream
}

func (x *devicePluginFingerprintClient) Recv() (*DetectedDevices, error) {
	m := new(DetectedDevices)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *devicePluginClient) Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error) {
	out := new(ReserveResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.device.DevicePlugin/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicePluginServer is the server API for DevicePlugin service.
type DevicePluginServer interface {
	// Fingerprint allows the device plugin to return a set of
	// detected devices and provide a mechanism to update the state of
	// the device.
	Fingerprint(*empty.Empty, DevicePlugin_FingerprintServer) error
	// Reserve is called by the client before starting an allocation
	// that requires access to the plugin’s devices. The plugin can use
	// this to run any setup steps and provides the mounting details to
	// the Nomad client
	Reserve(context.Context, *ReserveRequest) (*ReserveResponse, error)
}

func RegisterDevicePluginServer(s *grpc.Server, srv DevicePluginServer) {
	s.RegisterService(&_DevicePlugin_serviceDesc, srv)
}

func _DevicePlugin_Fingerprint_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DevicePluginServer).Fingerprint(m, &devicePluginFingerprintServer{stream})
}

type DevicePlugin_FingerprintServer interface {
	Send(*DetectedDevices) error
	grpc.ServerStream
}

type devicePluginFingerprintServer struct {
	grpc.ServerStream
}

func (x *devicePluginFingerprintServer) Send(m *DetectedDevices) error {
	return x.ServerStream.SendMsg(m)
}

func _DevicePlugin_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicePluginServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.device.DevicePlugin/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicePluginServer).Reserve(ctx, req.(*ReserveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevicePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.device.DevicePlugin",
	HandlerType: (*DevicePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reserve",
			Handler:    _DevicePlugin_Reserve_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fingerprint",
			Handler:       _DevicePlugin_Fingerprint_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/hashicorp/nomad/plugins/device/device.proto",
}

func init() {
	proto.RegisterFile("github.com/hashicorp/nomad/plugins/device/device.proto", fileDescriptor_device_13acb8ec0117c3b0)
}

var fileDescriptor_device_13acb8ec0117c3b0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0x92, 0x36, 0x75, 0x26, 0xfd, 0xd2, 0x8f, 0xa5, 0xaa, 0xac, 0x94, 0x9f, 0xc8, 0x12,
	0x52, 0x85, 0x84, 0x8d, 0x02, 0x02, 0x84, 0x04, 0x52, 0xdb, 0x14, 0x91, 0x0b, 0x4a, 0x65, 0xb8,
	0xa1, 0x17, 0x58, 0x8e, 0x3d, 0xc4, 0xab, 0xda, 0xbb, 0x8b, 0x77, 0x1d, 0xc9, 0x4f, 0xc0, 0xab,
	0xf0, 0x48, 0x3c, 0x01, 0xcf, 0x81, 0xbc, 0xeb, 0xa4, 0x09, 0xaa, 0x88, 0x0a, 0x57, 0xf6, 0xce,
	0x9c, 0x73, 0xe6, 0x78, 0x3c, 0x3b, 0xf0, 0x6c, 0x4a, 0x55, 0x52, 0x4c, 0xdc, 0x88, 0x67, 0x5e,
	0x12, 0xca, 0x84, 0x46, 0x3c, 0x17, 0x1e, 0xe3, 0x59, 0x18, 0x7b, 0x22, 0x2d, 0xa6, 0x94, 0x49,
	0x2f, 0xc6, 0x19, 0x8d, 0xb0, 0x7e, 0xb8, 0x22, 0xe7, 0x8a, 0x93, 0x7b, 0x0b, 0xb0, 0xab, 0xc1,
	0x6e, 0x0d, 0x76, 0x0d, 0xaa, 0x7f, 0x30, 0xe5, 0x7c, 0x9a, 0xa2, 0xa7, 0xd1, 0x93, 0xe2, 0x8b,
	0x87, 0x99, 0x50, 0xa5, 0x21, 0x3b, 0x3f, 0x9b, 0xb0, 0x3b, 0x42, 0x85, 0x91, 0xc2, 0x78, 0xa4,
	0xf1, 0x92, 0xec, 0x43, 0x7b, 0x86, 0x2c, 0xe6, 0xb9, 0xdd, 0x18, 0x34, 0x0e, 0x3b, 0x7e, 0x7d,
	0x22, 0xf7, 0xa1, 0x6b, 0x24, 0x03, 0x55, 0x0a, 0xb4, 0x9b, 0x3a, 0x09, 0x26, 0xf4, 0xb1, 0x14,
	0xb8, 0x04, 0x60, 0x61, 0x86, 0x76, 0x6b, 0x19, 0x70, 0x16, 0x66, 0x48, 0xde, 0xc2, 0xb6, 0x39,
	0x49, 0x7b, 0x73, 0xd0, 0x3a, 0xec, 0x0e, 0x5d, 0xf7, 0xcf, 0xe6, 0xdd, 0x55, 0x6f, 0xfe, 0x9c,
	0x4e, 0x52, 0xd8, 0x65, 0x3c, 0xc6, 0x20, 0x54, 0x2a, 0xa7, 0x93, 0x42, 0xa1, 0xb4, 0xb7, 0xb4,
	0xe2, 0xc9, 0xcd, 0x14, 0xa5, 0x7b, 0xc6, 0x63, 0x3c, 0x5a, 0xa8, 0x9c, 0x32, 0x95, 0x97, 0x7e,
	0x8f, 0xad, 0x04, 0xfb, 0x47, 0x70, 0xfb, 0x1a, 0x18, 0xf9, 0x1f, 0x5a, 0x97, 0x58, 0xd6, 0x5d,
	0xaa, 0x5e, 0xc9, 0x1e, 0x6c, 0xcd, 0xc2, 0xb4, 0x98, 0x37, 0xc7, 0x1c, 0x5e, 0x36, 0x5f, 0x34,
	0x9c, 0x6f, 0x0d, 0xe8, 0xad, 0x96, 0x26, 0x3d, 0x68, 0x8e, 0x47, 0x35, 0xbb, 0x39, 0x1e, 0x11,
	0x1b, 0xb6, 0x13, 0x0c, 0x53, 0x95, 0x94, 0x9a, 0x6e, 0xf9, 0xf3, 0x23, 0x79, 0x04, 0xc4, 0xbc,
	0x06, 0x31, 0xca, 0x28, 0xa7, 0x42, 0x51, 0xce, 0xea, 0xfe, 0xde, 0x32, 0x99, 0xd1, 0x55, 0x82,
	0xdc, 0x01, 0x10, 0x11, 0x0d, 0x26, 0x85, 0x0c, 0x68, 0x6c, 0x6f, 0x6a, 0x98, 0x25, 0x22, 0x7a,
	0x5c, 0xc8, 0x71, 0xec, 0x78, 0xd0, 0xf3, 0x51, 0x62, 0x3e, 0x43, 0x1f, 0xbf, 0x16, 0x28, 0x15,
	0xb9, 0x0b, 0xf5, 0x4f, 0x0a, 0x68, 0x2c, 0xed, 0xc6, 0xa0, 0x75, 0xd8, 0xf1, 0x3b, 0x26, 0x32,
	0x8e, 0xa5, 0x93, 0xc2, 0xee, 0x82, 0x20, 0x05, 0x67, 0x12, 0xc9, 0x27, 0xf8, 0x2f, 0xe2, 0x4c,
	0x85, 0x94, 0x61, 0x1e, 0xe4, 0x28, 0xf5, 0x57, 0x74, 0x87, 0x4f, 0xd7, 0x35, 0xff, 0x64, 0x4e,
	0x32, 0x82, 0x61, 0x65, 0xd7, 0xdf, 0x89, 0x96, 0xa2, 0xce, 0xf7, 0x26, 0xec, 0x5d, 0x07, 0x23,
	0x3e, 0x6c, 0x22, 0x9b, 0x19, 0x7f, 0xdd, 0xe1, 0xeb, 0xbf, 0x29, 0xe5, 0x9e, 0xb2, 0x59, 0xfd,
	0x8b, 0xb5, 0x16, 0x79, 0x05, 0xed, 0x8c, 0x17, 0x4c, 0x49, 0xbb, 0xa9, 0x55, 0x1f, 0xac, 0x53,
	0x7d, 0x57, 0xa1, 0xfd, 0x9a, 0x44, 0x46, 0x57, 0xf3, 0xdc, 0xd2, 0xfc, 0x87, 0xeb, 0xa7, 0xaf,
	0x7a, 0x7c, 0x10, 0x18, 0x2d, 0x66, 0xb9, 0xff, 0x1c, 0x3a, 0x0b, 0x5f, 0x37, 0x9a, 0xa9, 0xcf,
	0xb0, 0xa5, 0xfd, 0x90, 0x03, 0xe8, 0xa8, 0x50, 0x5e, 0x06, 0x22, 0x54, 0x49, 0x4d, 0xb5, 0xaa,
	0xc0, 0x79, 0xa8, 0x92, 0x2a, 0x99, 0x70, 0xa9, 0x4c, 0xd2, 0x68, 0x58, 0x55, 0x60, 0x9e, 0xcc,
	0x31, 0x8c, 0x03, 0xce, 0xd2, 0x52, 0x0f, 0x94, 0xe5, 0x5b, 0x55, 0xe0, 0x3d, 0x4b, 0x4b, 0x27,
	0x01, 0xb8, 0xf2, 0xfb, 0x0f, 0x45, 0x06, 0xd0, 0x15, 0x98, 0x67, 0x54, 0x4a, 0xca, 0x99, 0xac,
	0xe7, 0x76, 0x39, 0x34, 0xfc, 0xd1, 0x80, 0x1d, 0x53, 0xea, 0x5c, 0xf7, 0x8b, 0x5c, 0x40, 0xf7,
	0x0d, 0x65, 0x53, 0xcc, 0x45, 0x4e, 0x99, 0x22, 0xfb, 0xae, 0x59, 0x62, 0xee, 0x7c, 0x89, 0xb9,
	0xa7, 0xd5, 0x12, 0xeb, 0x7b, 0x37, 0xbc, 0xed, 0xce, 0xc6, 0xe3, 0x06, 0x49, 0x61, 0xbb, 0x9e,
	0x67, 0xb2, 0x76, 0xff, 0xac, 0xde, 0x94, 0xf5, 0xf5, 0x7e, 0xbb, 0x28, 0xce, 0xc6, 0xb1, 0x75,
	0xd1, 0x36, 0xb9, 0x49, 0x5b, 0x9b, 0x7f, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x15, 0x03, 0xfe,
	0x52, 0xe9, 0x05, 0x00, 0x00,
}
