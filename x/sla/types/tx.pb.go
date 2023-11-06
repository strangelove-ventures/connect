// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/sla/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgAddSLAs defines the Msg/AddSLAs request type. It contains the
// SLAs to be added to the store.
type MsgAddSLAs struct {
	// SLAs defines the SLAs to be added to the store.
	SLAs []PriceFeedSLA `protobuf:"bytes,1,rep,name=slas,proto3" json:"slas"`
	// Authority defines the authority that is adding the SLAs.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgAddSLAs) Reset()         { *m = MsgAddSLAs{} }
func (m *MsgAddSLAs) String() string { return proto.CompactTextString(m) }
func (*MsgAddSLAs) ProtoMessage()    {}
func (*MsgAddSLAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{0}
}
func (m *MsgAddSLAs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddSLAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddSLAs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddSLAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddSLAs.Merge(m, src)
}
func (m *MsgAddSLAs) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddSLAs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddSLAs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddSLAs proto.InternalMessageInfo

func (m *MsgAddSLAs) GetSLAs() []PriceFeedSLA {
	if m != nil {
		return m.SLAs
	}
	return nil
}

func (m *MsgAddSLAs) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgAddSLAsResponse defines the Msg/AddSLAs response type.
type MsgAddSLAsResponse struct {
}

func (m *MsgAddSLAsResponse) Reset()         { *m = MsgAddSLAsResponse{} }
func (m *MsgAddSLAsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddSLAsResponse) ProtoMessage()    {}
func (*MsgAddSLAsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{1}
}
func (m *MsgAddSLAsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddSLAsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddSLAsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddSLAsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddSLAsResponse.Merge(m, src)
}
func (m *MsgAddSLAsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddSLAsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddSLAsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddSLAsResponse proto.InternalMessageInfo

// MsgRemoveSLAs defines the Msg/RemoveSLAs request type. It contains the
// IDs of the SLAs to be removed from the store.
type MsgRemoveSLAs struct {
	// IDs defines the IDs of the SLAs to be removed from the store.
	IDs []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// Authority defines the authority that is removing the SLAs.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgRemoveSLAs) Reset()         { *m = MsgRemoveSLAs{} }
func (m *MsgRemoveSLAs) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveSLAs) ProtoMessage()    {}
func (*MsgRemoveSLAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{2}
}
func (m *MsgRemoveSLAs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveSLAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveSLAs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveSLAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveSLAs.Merge(m, src)
}
func (m *MsgRemoveSLAs) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveSLAs) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveSLAs.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveSLAs proto.InternalMessageInfo

func (m *MsgRemoveSLAs) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *MsgRemoveSLAs) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgRemoveSLAsResponse defines the Msg/RemoveSLAs response type.
type MsgRemoveSLAsResponse struct {
}

func (m *MsgRemoveSLAsResponse) Reset()         { *m = MsgRemoveSLAsResponse{} }
func (m *MsgRemoveSLAsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveSLAsResponse) ProtoMessage()    {}
func (*MsgRemoveSLAsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{3}
}
func (m *MsgRemoveSLAsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveSLAsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveSLAsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveSLAsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveSLAsResponse.Merge(m, src)
}
func (m *MsgRemoveSLAsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveSLAsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveSLAsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveSLAsResponse proto.InternalMessageInfo

// MsgParams defines the Msg/Params request type. It contains the
// new parameters for the SLA module.
type MsgParams struct {
	// Params defines the new parameters for the SLA module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// Authority defines the authority that is updating the SLA module parameters.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgParams) Reset()         { *m = MsgParams{} }
func (m *MsgParams) String() string { return proto.CompactTextString(m) }
func (*MsgParams) ProtoMessage()    {}
func (*MsgParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{4}
}
func (m *MsgParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParams.Merge(m, src)
}
func (m *MsgParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParams proto.InternalMessageInfo

func (m *MsgParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *MsgParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgParamsResponse defines the Msg/Params response type.
type MsgParamsResponse struct {
}

func (m *MsgParamsResponse) Reset()         { *m = MsgParamsResponse{} }
func (m *MsgParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgParamsResponse) ProtoMessage()    {}
func (*MsgParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e35178383738b0, []int{5}
}
func (m *MsgParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParamsResponse.Merge(m, src)
}
func (m *MsgParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddSLAs)(nil), "slinky.sla.v1.MsgAddSLAs")
	proto.RegisterType((*MsgAddSLAsResponse)(nil), "slinky.sla.v1.MsgAddSLAsResponse")
	proto.RegisterType((*MsgRemoveSLAs)(nil), "slinky.sla.v1.MsgRemoveSLAs")
	proto.RegisterType((*MsgRemoveSLAsResponse)(nil), "slinky.sla.v1.MsgRemoveSLAsResponse")
	proto.RegisterType((*MsgParams)(nil), "slinky.sla.v1.MsgParams")
	proto.RegisterType((*MsgParamsResponse)(nil), "slinky.sla.v1.MsgParamsResponse")
}

func init() { proto.RegisterFile("slinky/sla/v1/tx.proto", fileDescriptor_92e35178383738b0) }

var fileDescriptor_92e35178383738b0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6b, 0xdb, 0x40,
	0x14, 0xc7, 0x7d, 0xb5, 0xeb, 0xe0, 0x97, 0x64, 0x88, 0xea, 0x34, 0xb2, 0x52, 0x64, 0xd7, 0x94,
	0x62, 0x02, 0x91, 0x88, 0x03, 0x1d, 0x0a, 0x85, 0x5a, 0x84, 0x96, 0x42, 0x0d, 0x46, 0xd9, 0xba,
	0x04, 0xc5, 0xba, 0x5e, 0x8e, 0xf8, 0x7c, 0x42, 0x4f, 0x36, 0xf1, 0x56, 0xba, 0x76, 0xe9, 0xd4,
	0xcf, 0x91, 0xa1, 0x1f, 0x22, 0x63, 0xe8, 0xd4, 0xc9, 0x14, 0x79, 0xc8, 0xd4, 0xef, 0x50, 0xa4,
	0x93, 0xad, 0x3a, 0x21, 0x99, 0xbc, 0xe9, 0xbd, 0xff, 0xff, 0xff, 0xf4, 0xe3, 0x1d, 0x0f, 0x9e,
	0xe2, 0x80, 0x0f, 0xcf, 0x27, 0x36, 0x0e, 0x3c, 0x7b, 0x7c, 0x60, 0x47, 0x17, 0x56, 0x10, 0xca,
	0x48, 0x6a, 0x9b, 0xaa, 0x6f, 0xe1, 0xc0, 0xb3, 0xc6, 0x07, 0xc6, 0xee, 0xb2, 0x8d, 0xd1, 0x21,
	0x45, 0x8e, 0xca, 0x6b, 0xd4, 0xfa, 0x12, 0x85, 0xc4, 0x93, 0xb4, 0xb2, 0x55, 0x91, 0x49, 0x3b,
	0xaa, 0xb2, 0x05, 0xb2, 0x24, 0x27, 0x90, 0x65, 0x42, 0x95, 0x49, 0x26, 0x55, 0x20, 0xf9, 0x52,
	0xdd, 0xe6, 0x0f, 0x02, 0xd0, 0x45, 0xd6, 0xf1, 0xfd, 0xe3, 0x8f, 0x1d, 0xd4, 0xde, 0x40, 0x09,
	0x07, 0x1e, 0xea, 0xa4, 0x51, 0x6c, 0xad, 0xb7, 0x77, 0xad, 0x25, 0x26, 0xab, 0x17, 0xf2, 0x3e,
	0x7d, 0x47, 0x69, 0xe2, 0x75, 0x36, 0xae, 0xa6, 0xf5, 0x42, 0x3c, 0xad, 0x97, 0x92, 0xa0, 0x9b,
	0xc6, 0xb4, 0x57, 0x50, 0xf1, 0x46, 0xd1, 0x99, 0x0c, 0x79, 0x34, 0xd1, 0x1f, 0x35, 0x48, 0xab,
	0xe2, 0xe8, 0xbf, 0x7e, 0xee, 0x57, 0x33, 0xc2, 0x8e, 0xef, 0x87, 0x14, 0xf1, 0x38, 0x0a, 0xf9,
	0x90, 0xb9, 0xb9, 0xf5, 0xf5, 0xd6, 0xd7, 0x9b, 0xcb, 0xbd, 0x8d, 0xcf, 0xa1, 0x14, 0x27, 0x9e,
	0xf2, 0x34, 0xab, 0xa0, 0xe5, 0x5c, 0x2e, 0xc5, 0x40, 0x0e, 0x91, 0x36, 0x47, 0xb0, 0xd9, 0x45,
	0xe6, 0x52, 0x21, 0xc7, 0x34, 0x05, 0xae, 0x41, 0x91, 0xfb, 0x8a, 0xb7, 0xe2, 0xac, 0xc5, 0xd3,
	0x7a, 0xf1, 0xc3, 0x11, 0xba, 0x49, 0x6f, 0x95, 0x30, 0x3b, 0xb0, 0xbd, 0xf4, 0xdb, 0x05, 0xcf,
	0x37, 0x02, 0x95, 0x2e, 0xb2, 0x9e, 0x17, 0x7a, 0x02, 0xb5, 0x43, 0x28, 0x07, 0xe9, 0x97, 0x4e,
	0x1a, 0xa4, 0xb5, 0xde, 0xde, 0xbe, 0xbd, 0xbf, 0x54, 0x74, 0x4a, 0xc9, 0xe6, 0xdc, 0xcc, 0xba,
	0x4a, 0xcc, 0x27, 0xb0, 0xb5, 0x80, 0x99, 0x23, 0xb6, 0xff, 0x12, 0x28, 0x76, 0x91, 0x69, 0xef,
	0x61, 0x6d, 0xfe, 0xca, 0xb5, 0x5b, 0x5c, 0xf9, 0xa2, 0x8d, 0xe7, 0xf7, 0x4a, 0xf3, 0x81, 0x5a,
	0x0f, 0xe0, 0xbf, 0x07, 0x78, 0x76, 0x37, 0x90, 0xab, 0xc6, 0x8b, 0x87, 0xd4, 0xc5, 0xc4, 0x23,
	0x28, 0x67, 0x1b, 0xd4, 0xef, 0xfa, 0x95, 0x62, 0x34, 0xee, 0x53, 0xe6, 0x53, 0x8c, 0xc7, 0x5f,
	0x6e, 0x2e, 0xf7, 0x88, 0xf3, 0xf6, 0x2a, 0x36, 0xc9, 0x75, 0x6c, 0x92, 0x3f, 0xb1, 0x49, 0xbe,
	0xcf, 0xcc, 0xc2, 0xf5, 0xcc, 0x2c, 0xfc, 0x9e, 0x99, 0x85, 0x4f, 0x2f, 0x19, 0x8f, 0xce, 0x46,
	0xa7, 0x56, 0x5f, 0x0a, 0x1b, 0xcf, 0x79, 0xb0, 0x2f, 0xe8, 0xd8, 0xce, 0xce, 0xec, 0x22, 0x3d,
	0xb4, 0x68, 0x12, 0x50, 0x3c, 0x2d, 0xa7, 0xa7, 0x71, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x2c, 0x80, 0x9c, 0xaa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// AddSLA defines a method for adding a new SLAs to the store. Note, this will
	// overwrite any existing SLA with the same ID.
	AddSLAs(ctx context.Context, in *MsgAddSLAs, opts ...grpc.CallOption) (*MsgAddSLAsResponse, error)
	// RemoveSLA defines a method for removing existing SLAs from the store. Note, this
	// will not panic if the SLA does not exist.
	RemoveSLAs(ctx context.Context, in *MsgRemoveSLAs, opts ...grpc.CallOption) (*MsgRemoveSLAsResponse, error)
	// Params defines a method for updating the SLA module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddSLAs(ctx context.Context, in *MsgAddSLAs, opts ...grpc.CallOption) (*MsgAddSLAsResponse, error) {
	out := new(MsgAddSLAsResponse)
	err := c.cc.Invoke(ctx, "/slinky.sla.v1.Msg/AddSLAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveSLAs(ctx context.Context, in *MsgRemoveSLAs, opts ...grpc.CallOption) (*MsgRemoveSLAsResponse, error) {
	out := new(MsgRemoveSLAsResponse)
	err := c.cc.Invoke(ctx, "/slinky.sla.v1.Msg/RemoveSLAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, "/slinky.sla.v1.Msg/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AddSLA defines a method for adding a new SLAs to the store. Note, this will
	// overwrite any existing SLA with the same ID.
	AddSLAs(context.Context, *MsgAddSLAs) (*MsgAddSLAsResponse, error)
	// RemoveSLA defines a method for removing existing SLAs from the store. Note, this
	// will not panic if the SLA does not exist.
	RemoveSLAs(context.Context, *MsgRemoveSLAs) (*MsgRemoveSLAsResponse, error)
	// Params defines a method for updating the SLA module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddSLAs(ctx context.Context, req *MsgAddSLAs) (*MsgAddSLAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSLAs not implemented")
}
func (*UnimplementedMsgServer) RemoveSLAs(ctx context.Context, req *MsgRemoveSLAs) (*MsgRemoveSLAsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSLAs not implemented")
}
func (*UnimplementedMsgServer) Params(ctx context.Context, req *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddSLAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddSLAs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddSLAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.sla.v1.Msg/AddSLAs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddSLAs(ctx, req.(*MsgAddSLAs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveSLAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveSLAs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveSLAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.sla.v1.Msg/RemoveSLAs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveSLAs(ctx, req.(*MsgRemoveSLAs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.sla.v1.Msg/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.sla.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSLAs",
			Handler:    _Msg_AddSLAs_Handler,
		},
		{
			MethodName: "RemoveSLAs",
			Handler:    _Msg_RemoveSLAs_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/sla/v1/tx.proto",
}

func (m *MsgAddSLAs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddSLAs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddSLAs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SLAs) > 0 {
		for iNdEx := len(m.SLAs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SLAs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddSLAsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddSLAsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddSLAsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveSLAs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveSLAs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveSLAs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IDs) > 0 {
		for iNdEx := len(m.IDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IDs[iNdEx])
			copy(dAtA[i:], m.IDs[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.IDs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveSLAsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveSLAsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveSLAsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddSLAs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SLAs) > 0 {
		for _, e := range m.SLAs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddSLAsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveSLAs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRemoveSLAsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddSLAs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddSLAs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddSLAs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SLAs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SLAs = append(m.SLAs, PriceFeedSLA{})
			if err := m.SLAs[len(m.SLAs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddSLAsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddSLAsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddSLAsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveSLAs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveSLAs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveSLAs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDs = append(m.IDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveSLAsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveSLAsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveSLAsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
