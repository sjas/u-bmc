// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bmc.proto

package bmc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Button int32

const (
	Button_BUTTON_UNSPEC Button = 0
	Button_BUTTON_POWER  Button = 1
	Button_BUTTON_RESET  Button = 2
)

var Button_name = map[int32]string{
	0: "BUTTON_UNSPEC",
	1: "BUTTON_POWER",
	2: "BUTTON_RESET",
}
var Button_value = map[string]int32{
	"BUTTON_UNSPEC": 0,
	"BUTTON_POWER":  1,
	"BUTTON_RESET":  2,
}

func (x Button) String() string {
	return proto.EnumName(Button_name, int32(x))
}
func (Button) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{0}
}

type FanMode int32

const (
	FanMode_FAN_MODE_UNSPEC     FanMode = 0
	FanMode_FAN_MODE_PERCENTAGE FanMode = 1
)

var FanMode_name = map[int32]string{
	0: "FAN_MODE_UNSPEC",
	1: "FAN_MODE_PERCENTAGE",
}
var FanMode_value = map[string]int32{
	"FAN_MODE_UNSPEC":     0,
	"FAN_MODE_PERCENTAGE": 1,
}

func (x FanMode) String() string {
	return proto.EnumName(FanMode_name, int32(x))
}
func (FanMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{1}
}

type ButtonPressRequest struct {
	// Required: which button to press
	Button Button `protobuf:"varint,1,opt,name=button,proto3,enum=bmc.Button" json:"button,omitempty"`
	// Required: duration in milliseconds for how long to press the button
	DurationMs           uint32   `protobuf:"varint,2,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButtonPressRequest) Reset()         { *m = ButtonPressRequest{} }
func (m *ButtonPressRequest) String() string { return proto.CompactTextString(m) }
func (*ButtonPressRequest) ProtoMessage()    {}
func (*ButtonPressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{0}
}
func (m *ButtonPressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButtonPressRequest.Unmarshal(m, b)
}
func (m *ButtonPressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButtonPressRequest.Marshal(b, m, deterministic)
}
func (dst *ButtonPressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButtonPressRequest.Merge(dst, src)
}
func (m *ButtonPressRequest) XXX_Size() int {
	return xxx_messageInfo_ButtonPressRequest.Size(m)
}
func (m *ButtonPressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ButtonPressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ButtonPressRequest proto.InternalMessageInfo

func (m *ButtonPressRequest) GetButton() Button {
	if m != nil {
		return m.Button
	}
	return Button_BUTTON_UNSPEC
}

func (m *ButtonPressRequest) GetDurationMs() uint32 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

type ButtonPressResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButtonPressResponse) Reset()         { *m = ButtonPressResponse{} }
func (m *ButtonPressResponse) String() string { return proto.CompactTextString(m) }
func (*ButtonPressResponse) ProtoMessage()    {}
func (*ButtonPressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{1}
}
func (m *ButtonPressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButtonPressResponse.Unmarshal(m, b)
}
func (m *ButtonPressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButtonPressResponse.Marshal(b, m, deterministic)
}
func (dst *ButtonPressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButtonPressResponse.Merge(dst, src)
}
func (m *ButtonPressResponse) XXX_Size() int {
	return xxx_messageInfo_ButtonPressResponse.Size(m)
}
func (m *ButtonPressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ButtonPressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ButtonPressResponse proto.InternalMessageInfo

type SetFanRequest struct {
	// Required: which fan to modify
	Fan uint32 `protobuf:"varint,1,opt,name=fan,proto3" json:"fan,omitempty"`
	// Required: which mode to put the fan in
	Mode FanMode `protobuf:"varint,2,opt,name=mode,proto3,enum=bmc.FanMode" json:"mode,omitempty"`
	// Required if mode is PERCENTAGE
	Percentage           uint32   `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFanRequest) Reset()         { *m = SetFanRequest{} }
func (m *SetFanRequest) String() string { return proto.CompactTextString(m) }
func (*SetFanRequest) ProtoMessage()    {}
func (*SetFanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{2}
}
func (m *SetFanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFanRequest.Unmarshal(m, b)
}
func (m *SetFanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFanRequest.Marshal(b, m, deterministic)
}
func (dst *SetFanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFanRequest.Merge(dst, src)
}
func (m *SetFanRequest) XXX_Size() int {
	return xxx_messageInfo_SetFanRequest.Size(m)
}
func (m *SetFanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFanRequest proto.InternalMessageInfo

func (m *SetFanRequest) GetFan() uint32 {
	if m != nil {
		return m.Fan
	}
	return 0
}

func (m *SetFanRequest) GetMode() FanMode {
	if m != nil {
		return m.Mode
	}
	return FanMode_FAN_MODE_UNSPEC
}

func (m *SetFanRequest) GetPercentage() uint32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type SetFanResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFanResponse) Reset()         { *m = SetFanResponse{} }
func (m *SetFanResponse) String() string { return proto.CompactTextString(m) }
func (*SetFanResponse) ProtoMessage()    {}
func (*SetFanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{3}
}
func (m *SetFanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFanResponse.Unmarshal(m, b)
}
func (m *SetFanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFanResponse.Marshal(b, m, deterministic)
}
func (dst *SetFanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFanResponse.Merge(dst, src)
}
func (m *SetFanResponse) XXX_Size() int {
	return xxx_messageInfo_SetFanResponse.Size(m)
}
func (m *SetFanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetFanResponse proto.InternalMessageInfo

type GetFansRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFansRequest) Reset()         { *m = GetFansRequest{} }
func (m *GetFansRequest) String() string { return proto.CompactTextString(m) }
func (*GetFansRequest) ProtoMessage()    {}
func (*GetFansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{4}
}
func (m *GetFansRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFansRequest.Unmarshal(m, b)
}
func (m *GetFansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFansRequest.Marshal(b, m, deterministic)
}
func (dst *GetFansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFansRequest.Merge(dst, src)
}
func (m *GetFansRequest) XXX_Size() int {
	return xxx_messageInfo_GetFansRequest.Size(m)
}
func (m *GetFansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFansRequest proto.InternalMessageInfo

type Fan struct {
	Fan                  uint32   `protobuf:"varint,1,opt,name=fan,proto3" json:"fan,omitempty"`
	Mode                 FanMode  `protobuf:"varint,2,opt,name=mode,proto3,enum=bmc.FanMode" json:"mode,omitempty"`
	Percentage           uint32   `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Rpm                  uint32   `protobuf:"varint,4,opt,name=rpm,proto3" json:"rpm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fan) Reset()         { *m = Fan{} }
func (m *Fan) String() string { return proto.CompactTextString(m) }
func (*Fan) ProtoMessage()    {}
func (*Fan) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{5}
}
func (m *Fan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fan.Unmarshal(m, b)
}
func (m *Fan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fan.Marshal(b, m, deterministic)
}
func (dst *Fan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fan.Merge(dst, src)
}
func (m *Fan) XXX_Size() int {
	return xxx_messageInfo_Fan.Size(m)
}
func (m *Fan) XXX_DiscardUnknown() {
	xxx_messageInfo_Fan.DiscardUnknown(m)
}

var xxx_messageInfo_Fan proto.InternalMessageInfo

func (m *Fan) GetFan() uint32 {
	if m != nil {
		return m.Fan
	}
	return 0
}

func (m *Fan) GetMode() FanMode {
	if m != nil {
		return m.Mode
	}
	return FanMode_FAN_MODE_UNSPEC
}

func (m *Fan) GetPercentage() uint32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *Fan) GetRpm() uint32 {
	if m != nil {
		return m.Rpm
	}
	return 0
}

type GetFansResponse struct {
	Fan                  []*Fan   `protobuf:"bytes,1,rep,name=fan,proto3" json:"fan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFansResponse) Reset()         { *m = GetFansResponse{} }
func (m *GetFansResponse) String() string { return proto.CompactTextString(m) }
func (*GetFansResponse) ProtoMessage()    {}
func (*GetFansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{6}
}
func (m *GetFansResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFansResponse.Unmarshal(m, b)
}
func (m *GetFansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFansResponse.Marshal(b, m, deterministic)
}
func (dst *GetFansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFansResponse.Merge(dst, src)
}
func (m *GetFansResponse) XXX_Size() int {
	return xxx_messageInfo_GetFansResponse.Size(m)
}
func (m *GetFansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFansResponse proto.InternalMessageInfo

func (m *GetFansResponse) GetFan() []*Fan {
	if m != nil {
		return m.Fan
	}
	return nil
}

type ConsoleData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsoleData) Reset()         { *m = ConsoleData{} }
func (m *ConsoleData) String() string { return proto.CompactTextString(m) }
func (*ConsoleData) ProtoMessage()    {}
func (*ConsoleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{7}
}
func (m *ConsoleData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsoleData.Unmarshal(m, b)
}
func (m *ConsoleData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsoleData.Marshal(b, m, deterministic)
}
func (dst *ConsoleData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsoleData.Merge(dst, src)
}
func (m *ConsoleData) XXX_Size() int {
	return xxx_messageInfo_ConsoleData.Size(m)
}
func (m *ConsoleData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsoleData.DiscardUnknown(m)
}

var xxx_messageInfo_ConsoleData proto.InternalMessageInfo

func (m *ConsoleData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{8}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(dst, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

type GetVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GitHash              string   `protobuf:"bytes,2,opt,name=git_hash,json=gitHash,proto3" json:"git_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionResponse) Reset()         { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()    {}
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bmc_38e3e9ddd88e3820, []int{9}
}
func (m *GetVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionResponse.Unmarshal(m, b)
}
func (m *GetVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionResponse.Marshal(b, m, deterministic)
}
func (dst *GetVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionResponse.Merge(dst, src)
}
func (m *GetVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetVersionResponse.Size(m)
}
func (m *GetVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionResponse proto.InternalMessageInfo

func (m *GetVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetVersionResponse) GetGitHash() string {
	if m != nil {
		return m.GitHash
	}
	return ""
}

func init() {
	proto.RegisterType((*ButtonPressRequest)(nil), "bmc.ButtonPressRequest")
	proto.RegisterType((*ButtonPressResponse)(nil), "bmc.ButtonPressResponse")
	proto.RegisterType((*SetFanRequest)(nil), "bmc.SetFanRequest")
	proto.RegisterType((*SetFanResponse)(nil), "bmc.SetFanResponse")
	proto.RegisterType((*GetFansRequest)(nil), "bmc.GetFansRequest")
	proto.RegisterType((*Fan)(nil), "bmc.Fan")
	proto.RegisterType((*GetFansResponse)(nil), "bmc.GetFansResponse")
	proto.RegisterType((*ConsoleData)(nil), "bmc.ConsoleData")
	proto.RegisterType((*GetVersionRequest)(nil), "bmc.GetVersionRequest")
	proto.RegisterType((*GetVersionResponse)(nil), "bmc.GetVersionResponse")
	proto.RegisterEnum("bmc.Button", Button_name, Button_value)
	proto.RegisterEnum("bmc.FanMode", FanMode_name, FanMode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagementServiceClient interface {
	PressButton(ctx context.Context, in *ButtonPressRequest, opts ...grpc.CallOption) (*ButtonPressResponse, error)
	SetFan(ctx context.Context, in *SetFanRequest, opts ...grpc.CallOption) (*SetFanResponse, error)
	GetFans(ctx context.Context, in *GetFansRequest, opts ...grpc.CallOption) (*GetFansResponse, error)
	StreamConsole(ctx context.Context, opts ...grpc.CallOption) (ManagementService_StreamConsoleClient, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type managementServiceClient struct {
	cc *grpc.ClientConn
}

func NewManagementServiceClient(cc *grpc.ClientConn) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) PressButton(ctx context.Context, in *ButtonPressRequest, opts ...grpc.CallOption) (*ButtonPressResponse, error) {
	out := new(ButtonPressResponse)
	err := c.cc.Invoke(ctx, "/bmc.ManagementService/PressButton", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) SetFan(ctx context.Context, in *SetFanRequest, opts ...grpc.CallOption) (*SetFanResponse, error) {
	out := new(SetFanResponse)
	err := c.cc.Invoke(ctx, "/bmc.ManagementService/SetFan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) GetFans(ctx context.Context, in *GetFansRequest, opts ...grpc.CallOption) (*GetFansResponse, error) {
	out := new(GetFansResponse)
	err := c.cc.Invoke(ctx, "/bmc.ManagementService/GetFans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) StreamConsole(ctx context.Context, opts ...grpc.CallOption) (ManagementService_StreamConsoleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ManagementService_serviceDesc.Streams[0], "/bmc.ManagementService/StreamConsole", opts...)
	if err != nil {
		return nil, err
	}
	x := &managementServiceStreamConsoleClient{stream}
	return x, nil
}

type ManagementService_StreamConsoleClient interface {
	Send(*ConsoleData) error
	Recv() (*ConsoleData, error)
	grpc.ClientStream
}

type managementServiceStreamConsoleClient struct {
	grpc.ClientStream
}

func (x *managementServiceStreamConsoleClient) Send(m *ConsoleData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managementServiceStreamConsoleClient) Recv() (*ConsoleData, error) {
	m := new(ConsoleData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managementServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/bmc.ManagementService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
type ManagementServiceServer interface {
	PressButton(context.Context, *ButtonPressRequest) (*ButtonPressResponse, error)
	SetFan(context.Context, *SetFanRequest) (*SetFanResponse, error)
	GetFans(context.Context, *GetFansRequest) (*GetFansResponse, error)
	StreamConsole(ManagementService_StreamConsoleServer) error
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
}

func RegisterManagementServiceServer(s *grpc.Server, srv ManagementServiceServer) {
	s.RegisterService(&_ManagementService_serviceDesc, srv)
}

func _ManagementService_PressButton_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ButtonPressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).PressButton(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bmc.ManagementService/PressButton",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).PressButton(ctx, req.(*ButtonPressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_SetFan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).SetFan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bmc.ManagementService/SetFan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).SetFan(ctx, req.(*SetFanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_GetFans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).GetFans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bmc.ManagementService/GetFans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).GetFans(ctx, req.(*GetFansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_StreamConsole_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagementServiceServer).StreamConsole(&managementServiceStreamConsoleServer{stream})
}

type ManagementService_StreamConsoleServer interface {
	Send(*ConsoleData) error
	Recv() (*ConsoleData, error)
	grpc.ServerStream
}

type managementServiceStreamConsoleServer struct {
	grpc.ServerStream
}

func (x *managementServiceStreamConsoleServer) Send(m *ConsoleData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managementServiceStreamConsoleServer) Recv() (*ConsoleData, error) {
	m := new(ConsoleData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ManagementService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bmc.ManagementService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bmc.ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PressButton",
			Handler:    _ManagementService_PressButton_Handler,
		},
		{
			MethodName: "SetFan",
			Handler:    _ManagementService_SetFan_Handler,
		},
		{
			MethodName: "GetFans",
			Handler:    _ManagementService_GetFans_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ManagementService_GetVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamConsole",
			Handler:       _ManagementService_StreamConsole_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bmc.proto",
}

func init() { proto.RegisterFile("bmc.proto", fileDescriptor_bmc_38e3e9ddd88e3820) }

var fileDescriptor_bmc_38e3e9ddd88e3820 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4d, 0x8f, 0xda, 0x3c,
	0x10, 0x80, 0x09, 0x20, 0x58, 0x86, 0x8f, 0x0d, 0xc3, 0xfb, 0x96, 0x34, 0x87, 0x96, 0xba, 0x17,
	0xb4, 0x52, 0x57, 0x15, 0x2b, 0xb5, 0xea, 0x69, 0xb5, 0xb0, 0x81, 0xf6, 0xc0, 0x87, 0x0c, 0xdb,
	0x4a, 0xbd, 0x20, 0x03, 0x2e, 0xa0, 0x36, 0x36, 0x8d, 0xcd, 0xfe, 0xbe, 0xfe, 0xb4, 0x0a, 0x27,
	0x06, 0x22, 0xce, 0xbd, 0x39, 0x8f, 0xc7, 0xf3, 0x8c, 0x3d, 0x13, 0x28, 0x2d, 0xc2, 0xe5, 0xed,
	0x2e, 0x92, 0x5a, 0x62, 0x6e, 0x11, 0x2e, 0xc9, 0x77, 0xc0, 0xee, 0x5e, 0x6b, 0x29, 0x26, 0x11,
	0x57, 0x8a, 0xf2, 0xdf, 0x7b, 0xae, 0x34, 0xbe, 0x85, 0xc2, 0xc2, 0x50, 0xcf, 0x69, 0x39, 0xed,
	0x5a, 0xa7, 0x7c, 0x7b, 0x38, 0x16, 0x07, 0xd2, 0x64, 0x0b, 0x5f, 0x43, 0x79, 0xb5, 0x8f, 0x98,
	0xde, 0x4a, 0x31, 0x0f, 0x95, 0x97, 0x6d, 0x39, 0xed, 0x2a, 0x05, 0x8b, 0x86, 0x8a, 0xfc, 0x0f,
	0x8d, 0x54, 0x6e, 0xb5, 0x93, 0x42, 0x71, 0xb2, 0x84, 0xea, 0x94, 0xeb, 0x3e, 0x13, 0xd6, 0xe6,
	0x42, 0xee, 0x07, 0x8b, 0x55, 0x55, 0x7a, 0x58, 0x62, 0x0b, 0xf2, 0xa1, 0x5c, 0x71, 0x93, 0xb3,
	0xd6, 0xa9, 0x18, 0x7b, 0x9f, 0x89, 0xa1, 0x5c, 0x71, 0x6a, 0x76, 0xf0, 0x15, 0xc0, 0x8e, 0x47,
	0x4b, 0x2e, 0x34, 0x5b, 0x73, 0x2f, 0x17, 0xbb, 0x4f, 0x84, 0xb8, 0x50, 0xb3, 0x92, 0x44, 0xeb,
	0x42, 0x6d, 0x60, 0x88, 0xbd, 0x25, 0xf9, 0x09, 0xb9, 0x3e, 0x13, 0xff, 0x42, 0x7f, 0xc8, 0x19,
	0xed, 0x42, 0x2f, 0x1f, 0xe7, 0x8c, 0x76, 0x21, 0x79, 0x07, 0xd7, 0x47, 0x7d, 0x5c, 0x11, 0xfa,
	0x56, 0x9c, 0x6b, 0x97, 0x3b, 0x57, 0xd6, 0x62, 0x4a, 0x20, 0x6f, 0xa0, 0xdc, 0x93, 0x42, 0xc9,
	0x5f, 0xfc, 0x91, 0x69, 0x86, 0x08, 0xf9, 0x15, 0xd3, 0xcc, 0x14, 0x59, 0xa1, 0x66, 0x4d, 0x1a,
	0x50, 0x1f, 0x70, 0xfd, 0x95, 0x47, 0x6a, 0x2b, 0xed, 0x5b, 0x92, 0x2f, 0x80, 0xe7, 0x30, 0x31,
	0x79, 0x50, 0x7c, 0x8e, 0x91, 0xc9, 0x50, 0xa2, 0xf6, 0x13, 0x5f, 0xc2, 0xd5, 0x7a, 0xab, 0xe7,
	0x1b, 0xa6, 0x36, 0xe6, 0xba, 0x25, 0x5a, 0x5c, 0x6f, 0xf5, 0x67, 0xa6, 0x36, 0x37, 0xf7, 0x50,
	0x88, 0xdb, 0x87, 0x75, 0xa8, 0x76, 0x9f, 0x66, 0xb3, 0xf1, 0x68, 0xfe, 0x34, 0x9a, 0x4e, 0x82,
	0x9e, 0x9b, 0x41, 0x17, 0x2a, 0x09, 0x9a, 0x8c, 0xbf, 0x05, 0xd4, 0x75, 0xce, 0x08, 0x0d, 0xa6,
	0xc1, 0xcc, 0xcd, 0xde, 0x7c, 0x84, 0x62, 0xf2, 0x6a, 0xd8, 0x80, 0xeb, 0xfe, 0xc3, 0x68, 0x3e,
	0x1c, 0x3f, 0x06, 0xa7, 0x1c, 0x4d, 0x68, 0x1c, 0xe1, 0x24, 0xa0, 0xbd, 0x60, 0x34, 0x7b, 0x18,
	0x04, 0xae, 0xd3, 0xf9, 0x93, 0x85, 0xfa, 0x90, 0x09, 0xb6, 0xe6, 0x21, 0x17, 0x7a, 0xca, 0xa3,
	0xe7, 0xed, 0x92, 0x63, 0x17, 0xca, 0x66, 0x90, 0x92, 0xa2, 0x9a, 0x67, 0x33, 0x79, 0x3e, 0xbc,
	0xbe, 0x77, 0xb9, 0x91, 0x8c, 0x40, 0x06, 0xef, 0xa0, 0x10, 0x8f, 0x05, 0xa2, 0x89, 0x4a, 0x0d,
	0xa2, 0xdf, 0x48, 0xb1, 0xe3, 0xa1, 0x0f, 0x50, 0x4c, 0x5a, 0x87, 0x71, 0x44, 0x7a, 0x8e, 0xfc,
	0xff, 0xd2, 0xf0, 0x78, 0xee, 0x13, 0x54, 0xa7, 0x3a, 0xe2, 0x2c, 0x4c, 0x3a, 0x89, 0xae, 0x09,
	0x3c, 0xeb, 0xab, 0x7f, 0x41, 0x48, 0xa6, 0xed, 0xbc, 0x77, 0xf0, 0x1e, 0xe0, 0xd4, 0x46, 0x7c,
	0x61, 0x05, 0xe9, 0x66, 0xfb, 0xcd, 0x0b, 0x6e, 0xdd, 0x8b, 0x82, 0xf9, 0xc7, 0xef, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x4c, 0xf4, 0x69, 0xf0, 0x03, 0x00, 0x00,
}
