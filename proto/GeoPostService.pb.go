// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GeoPostService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddGeoPostParameters struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddGeoPostParameters) Reset()         { *m = AddGeoPostParameters{} }
func (m *AddGeoPostParameters) String() string { return proto.CompactTextString(m) }
func (*AddGeoPostParameters) ProtoMessage()    {}
func (*AddGeoPostParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_59ee5a670483ada5, []int{0}
}

func (m *AddGeoPostParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddGeoPostParameters.Unmarshal(m, b)
}
func (m *AddGeoPostParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddGeoPostParameters.Marshal(b, m, deterministic)
}
func (m *AddGeoPostParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddGeoPostParameters.Merge(m, src)
}
func (m *AddGeoPostParameters) XXX_Size() int {
	return xxx_messageInfo_AddGeoPostParameters.Size(m)
}
func (m *AddGeoPostParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_AddGeoPostParameters.DiscardUnknown(m)
}

var xxx_messageInfo_AddGeoPostParameters proto.InternalMessageInfo

type AddGeoPostReturn struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddGeoPostReturn) Reset()         { *m = AddGeoPostReturn{} }
func (m *AddGeoPostReturn) String() string { return proto.CompactTextString(m) }
func (*AddGeoPostReturn) ProtoMessage()    {}
func (*AddGeoPostReturn) Descriptor() ([]byte, []int) {
	return fileDescriptor_59ee5a670483ada5, []int{1}
}

func (m *AddGeoPostReturn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddGeoPostReturn.Unmarshal(m, b)
}
func (m *AddGeoPostReturn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddGeoPostReturn.Marshal(b, m, deterministic)
}
func (m *AddGeoPostReturn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddGeoPostReturn.Merge(m, src)
}
func (m *AddGeoPostReturn) XXX_Size() int {
	return xxx_messageInfo_AddGeoPostReturn.Size(m)
}
func (m *AddGeoPostReturn) XXX_DiscardUnknown() {
	xxx_messageInfo_AddGeoPostReturn.DiscardUnknown(m)
}

var xxx_messageInfo_AddGeoPostReturn proto.InternalMessageInfo

func (m *AddGeoPostReturn) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*AddGeoPostParameters)(nil), "proto.AddGeoPostParameters")
	proto.RegisterType((*AddGeoPostReturn)(nil), "proto.AddGeoPostReturn")
}

func init() { proto.RegisterFile("GeoPostService.proto", fileDescriptor_59ee5a670483ada5) }

var fileDescriptor_59ee5a670483ada5 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0x4f, 0xcd, 0x0f,
	0xc8, 0x2f, 0x2e, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x4a, 0x62, 0x5c, 0x22, 0x8e, 0x29, 0x29, 0x50, 0x15, 0x01, 0x89, 0x45,
	0x89, 0xb9, 0xa9, 0x25, 0xa9, 0x45, 0xc5, 0x4a, 0x3a, 0x5c, 0x02, 0x08, 0xf1, 0xa0, 0xd4, 0x92,
	0xd2, 0xa2, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0xd7, 0x28, 0x8c, 0x8b, 0x0f, 0xd5, 0x12, 0x21, 0x17, 0x2e,
	0x2e, 0x84, 0x7e, 0x21, 0x69, 0x88, 0xa5, 0x7a, 0xd8, 0xac, 0x92, 0x12, 0xc7, 0x90, 0x84, 0xd8,
	0xa7, 0xc4, 0x90, 0xc4, 0x06, 0x96, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x09, 0x66,
	0xa3, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoPostServiceClient is the client API for GeoPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoPostServiceClient interface {
	AddGeoPost(ctx context.Context, in *AddGeoPostParameters, opts ...grpc.CallOption) (*AddGeoPostReturn, error)
}

type geoPostServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeoPostServiceClient(cc *grpc.ClientConn) GeoPostServiceClient {
	return &geoPostServiceClient{cc}
}

func (c *geoPostServiceClient) AddGeoPost(ctx context.Context, in *AddGeoPostParameters, opts ...grpc.CallOption) (*AddGeoPostReturn, error) {
	out := new(AddGeoPostReturn)
	err := c.cc.Invoke(ctx, "/proto.GeoPostService/AddGeoPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoPostServiceServer is the server API for GeoPostService service.
type GeoPostServiceServer interface {
	AddGeoPost(context.Context, *AddGeoPostParameters) (*AddGeoPostReturn, error)
}

// UnimplementedGeoPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeoPostServiceServer struct {
}

func (*UnimplementedGeoPostServiceServer) AddGeoPost(ctx context.Context, req *AddGeoPostParameters) (*AddGeoPostReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGeoPost not implemented")
}

func RegisterGeoPostServiceServer(s *grpc.Server, srv GeoPostServiceServer) {
	s.RegisterService(&_GeoPostService_serviceDesc, srv)
}

func _GeoPostService_AddGeoPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGeoPostParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoPostServiceServer).AddGeoPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GeoPostService/AddGeoPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoPostServiceServer).AddGeoPost(ctx, req.(*AddGeoPostParameters))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeoPostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GeoPostService",
	HandlerType: (*GeoPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGeoPost",
			Handler:    _GeoPostService_AddGeoPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GeoPostService.proto",
}
