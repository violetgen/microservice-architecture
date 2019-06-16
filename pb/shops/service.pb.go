// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetShopCountRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopCountRequest) Reset()         { *m = GetShopCountRequest{} }
func (m *GetShopCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetShopCountRequest) ProtoMessage()    {}
func (*GetShopCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetShopCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopCountRequest.Unmarshal(m, b)
}
func (m *GetShopCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopCountRequest.Marshal(b, m, deterministic)
}
func (m *GetShopCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopCountRequest.Merge(m, src)
}
func (m *GetShopCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetShopCountRequest.Size(m)
}
func (m *GetShopCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopCountRequest proto.InternalMessageInfo

type GetShopCountResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopCountResponse) Reset()         { *m = GetShopCountResponse{} }
func (m *GetShopCountResponse) String() string { return proto.CompactTextString(m) }
func (*GetShopCountResponse) ProtoMessage()    {}
func (*GetShopCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetShopCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopCountResponse.Unmarshal(m, b)
}
func (m *GetShopCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopCountResponse.Marshal(b, m, deterministic)
}
func (m *GetShopCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopCountResponse.Merge(m, src)
}
func (m *GetShopCountResponse) XXX_Size() int {
	return xxx_messageInfo_GetShopCountResponse.Size(m)
}
func (m *GetShopCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopCountResponse proto.InternalMessageInfo

func (m *GetShopCountResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetShopByIDRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopByIDRequest) Reset()         { *m = GetShopByIDRequest{} }
func (m *GetShopByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetShopByIDRequest) ProtoMessage()    {}
func (*GetShopByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetShopByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopByIDRequest.Unmarshal(m, b)
}
func (m *GetShopByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetShopByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopByIDRequest.Merge(m, src)
}
func (m *GetShopByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetShopByIDRequest.Size(m)
}
func (m *GetShopByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopByIDRequest proto.InternalMessageInfo

func (m *GetShopByIDRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetShopResponse struct {
	Shop                 *Shop    `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopResponse) Reset()         { *m = GetShopResponse{} }
func (m *GetShopResponse) String() string { return proto.CompactTextString(m) }
func (*GetShopResponse) ProtoMessage()    {}
func (*GetShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *GetShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopResponse.Unmarshal(m, b)
}
func (m *GetShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopResponse.Marshal(b, m, deterministic)
}
func (m *GetShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopResponse.Merge(m, src)
}
func (m *GetShopResponse) XXX_Size() int {
	return xxx_messageInfo_GetShopResponse.Size(m)
}
func (m *GetShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopResponse proto.InternalMessageInfo

func (m *GetShopResponse) GetShop() *Shop {
	if m != nil {
		return m.Shop
	}
	return nil
}

type Shop struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ShopDomain           string               `protobuf:"bytes,2,opt,name=shop_domain,json=shopDomain,proto3" json:"shop_domain,omitempty"`
	Currency             string               `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Shop) Reset()         { *m = Shop{} }
func (m *Shop) String() string { return proto.CompactTextString(m) }
func (*Shop) ProtoMessage()    {}
func (*Shop) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *Shop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shop.Unmarshal(m, b)
}
func (m *Shop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shop.Marshal(b, m, deterministic)
}
func (m *Shop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shop.Merge(m, src)
}
func (m *Shop) XXX_Size() int {
	return xxx_messageInfo_Shop.Size(m)
}
func (m *Shop) XXX_DiscardUnknown() {
	xxx_messageInfo_Shop.DiscardUnknown(m)
}

var xxx_messageInfo_Shop proto.InternalMessageInfo

func (m *Shop) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shop) GetShopDomain() string {
	if m != nil {
		return m.ShopDomain
	}
	return ""
}

func (m *Shop) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Shop) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Shop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*GetShopCountRequest)(nil), "boldapi.pre.GetShopCountRequest")
	proto.RegisterType((*GetShopCountResponse)(nil), "boldapi.pre.GetShopCountResponse")
	proto.RegisterType((*GetShopByIDRequest)(nil), "boldapi.pre.GetShopByIDRequest")
	proto.RegisterType((*GetShopResponse)(nil), "boldapi.pre.GetShopResponse")
	proto.RegisterType((*Shop)(nil), "boldapi.pre.Shop")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0x9a, 0x8a, 0x3d, 0x51, 0x6b, 0xa7, 0x15, 0x43, 0x28, 0xa4, 0x06, 0x85, 0x2e,
	0x74, 0x02, 0x75, 0xe3, 0x4e, 0x5a, 0x0b, 0xe2, 0x36, 0x75, 0xe5, 0xa6, 0x4c, 0x92, 0x31, 0x1d,
	0x68, 0x32, 0x63, 0x66, 0x22, 0x14, 0x71, 0xe3, 0x13, 0x08, 0x3e, 0x9a, 0xb8, 0x76, 0xa1, 0x0f,
	0x72, 0x99, 0x49, 0xda, 0x9b, 0x5e, 0xca, 0xe5, 0xee, 0xf2, 0x9f, 0xf3, 0xcd, 0x7f, 0x32, 0xff,
	0x19, 0x78, 0x28, 0x69, 0xf5, 0x95, 0xa5, 0x14, 0x8b, 0x8a, 0x2b, 0x8e, 0xdc, 0x84, 0xef, 0x33,
	0x22, 0x18, 0x16, 0x15, 0xf5, 0xa7, 0x39, 0xe7, 0xf9, 0x9e, 0x46, 0x44, 0xb0, 0x88, 0x94, 0x25,
	0x57, 0x44, 0x31, 0x5e, 0xca, 0x06, 0xf5, 0x83, 0xb6, 0x6b, 0x54, 0x52, 0x7f, 0x8e, 0x14, 0x2b,
	0xa8, 0x54, 0xa4, 0x10, 0x2d, 0xf0, 0x2a, 0x67, 0x6a, 0x57, 0x27, 0x38, 0xe5, 0x45, 0x94, 0xf3,
	0x9c, 0x5f, 0x93, 0x5a, 0x19, 0x61, 0xbe, 0x1a, 0x3c, 0x7c, 0x02, 0xe3, 0xf7, 0x54, 0x6d, 0x76,
	0x5c, 0xbc, 0xe3, 0x75, 0xa9, 0x62, 0xfa, 0xa5, 0xa6, 0x52, 0x85, 0x2f, 0x61, 0x72, 0x5e, 0x96,
	0x82, 0x97, 0x92, 0xa2, 0x09, 0xf4, 0x53, 0x5d, 0xf0, 0xac, 0x99, 0x35, 0xef, 0xc7, 0x8d, 0x08,
	0x9f, 0x03, 0x6a, 0xe9, 0xd5, 0xe1, 0xc3, 0xba, 0xf5, 0x40, 0x8f, 0xc0, 0x66, 0x99, 0x01, 0x7b,
	0xb1, 0xcd, 0xb2, 0xf0, 0x0d, 0x0c, 0x5b, 0xea, 0x64, 0xf7, 0x02, 0x1c, 0xb9, 0xe3, 0xc2, 0x40,
	0xee, 0x62, 0x84, 0x3b, 0x39, 0x60, 0x03, 0x9a, 0x76, 0xf8, 0xc7, 0x02, 0x47, 0xcb, 0x9b, 0x96,
	0x28, 0x00, 0x57, 0x03, 0xdb, 0x8c, 0x17, 0x84, 0x95, 0x9e, 0x3d, 0xb3, 0xe6, 0x83, 0x18, 0x74,
	0x69, 0x6d, 0x2a, 0xc8, 0x87, 0xfb, 0x69, 0x5d, 0x55, 0xb4, 0x4c, 0x0f, 0x5e, 0xcf, 0x74, 0x4f,
	0x1a, 0xbd, 0x05, 0x48, 0x2b, 0x4a, 0x14, 0xcd, 0xb6, 0x44, 0x79, 0x63, 0xf3, 0x0b, 0x3e, 0x6e,
	0xf2, 0xc5, 0xc7, 0xd4, 0xf0, 0xc7, 0x63, 0xbe, 0x2b, 0xe7, 0xe7, 0xdf, 0xc0, 0x8a, 0x07, 0xed,
	0x99, 0xa5, 0xd2, 0x06, 0xb5, 0xc8, 0x8e, 0x06, 0x93, 0xbb, 0x1a, 0xb4, 0x67, 0x96, 0x6a, 0xf1,
	0xcf, 0x02, 0x57, 0xdf, 0x6b, 0xd3, 0xbc, 0x06, 0x54, 0xc0, 0x83, 0x6e, 0xea, 0x68, 0x76, 0x16,
	0xc8, 0x85, 0x3d, 0xf9, 0xcf, 0x6e, 0x21, 0x9a, 0x8c, 0x43, 0xef, 0xc7, 0xef, 0xff, 0xbf, 0x6c,
	0x84, 0x1e, 0x9b, 0x17, 0xa5, 0xb3, 0x91, 0x91, 0x59, 0x1b, 0xa2, 0xe0, 0x76, 0xd6, 0x86, 0x82,
	0x4b, 0x5e, 0x9d, 0x85, 0xfa, 0xd3, 0x4b, 0xc0, 0x69, 0xce, 0x53, 0x33, 0x67, 0x84, 0x86, 0x9d,
	0x39, 0xdf, 0x58, 0xf6, 0x7d, 0xe5, 0x7c, 0xb2, 0x45, 0x92, 0xdc, 0x33, 0x81, 0xbc, 0xbe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x40, 0x8e, 0xcb, 0xfb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	GetShopCount(ctx context.Context, in *GetShopCountRequest, opts ...grpc.CallOption) (*GetShopCountResponse, error)
	GetShopByID(ctx context.Context, in *GetShopByIDRequest, opts ...grpc.CallOption) (*GetShopResponse, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShopCount(ctx context.Context, in *GetShopCountRequest, opts ...grpc.CallOption) (*GetShopCountResponse, error) {
	out := new(GetShopCountResponse)
	err := c.cc.Invoke(ctx, "/boldapi.pre.ShopService/GetShopCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetShopByID(ctx context.Context, in *GetShopByIDRequest, opts ...grpc.CallOption) (*GetShopResponse, error) {
	out := new(GetShopResponse)
	err := c.cc.Invoke(ctx, "/boldapi.pre.ShopService/GetShopByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	GetShopCount(context.Context, *GetShopCountRequest) (*GetShopCountResponse, error)
	GetShopByID(context.Context, *GetShopByIDRequest) (*GetShopResponse, error)
}

// UnimplementedShopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (*UnimplementedShopServiceServer) GetShopCount(ctx context.Context, req *GetShopCountRequest) (*GetShopCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopCount not implemented")
}
func (*UnimplementedShopServiceServer) GetShopByID(ctx context.Context, req *GetShopByIDRequest) (*GetShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopByID not implemented")
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShopCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShopCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boldapi.pre.ShopService/GetShopCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShopCount(ctx, req.(*GetShopCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetShopByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShopByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boldapi.pre.ShopService/GetShopByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShopByID(ctx, req.(*GetShopByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "boldapi.pre.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShopCount",
			Handler:    _ShopService_GetShopCount_Handler,
		},
		{
			MethodName: "GetShopByID",
			Handler:    _ShopService_GetShopByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}