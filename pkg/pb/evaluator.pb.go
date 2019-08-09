// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/evaluator.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type EvaluateRequest struct {
	// List of Matches to evaluate.
	Matches              []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateRequest) Reset()         { *m = EvaluateRequest{} }
func (m *EvaluateRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateRequest) ProtoMessage()    {}
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{0}
}

func (m *EvaluateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateRequest.Unmarshal(m, b)
}
func (m *EvaluateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateRequest.Merge(m, src)
}
func (m *EvaluateRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateRequest.Size(m)
}
func (m *EvaluateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateRequest proto.InternalMessageInfo

func (m *EvaluateRequest) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type EvaluateResponse struct {
	// Accepted list of Matches.
	Matches              []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateResponse) Reset()         { *m = EvaluateResponse{} }
func (m *EvaluateResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponse) ProtoMessage()    {}
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{1}
}

func (m *EvaluateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponse.Unmarshal(m, b)
}
func (m *EvaluateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponse.Merge(m, src)
}
func (m *EvaluateResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponse.Size(m)
}
func (m *EvaluateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponse proto.InternalMessageInfo

func (m *EvaluateResponse) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

func init() {
	proto.RegisterType((*EvaluateRequest)(nil), "api.EvaluateRequest")
	proto.RegisterType((*EvaluateResponse)(nil), "api.EvaluateResponse")
}

func init() { proto.RegisterFile("api/evaluator.proto", fileDescriptor_8c58cb7dff9acb0f) }

var fileDescriptor_8c58cb7dff9acb0f = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x36, 0x28, 0x05, 0x73, 0xa0, 0x32, 0x1f, 0xaa, 0x22, 0x84, 0x4c, 0xca, 0x01, 0x22,
	0xb2, 0x4e, 0x43, 0x24, 0x50, 0x10, 0x52, 0x0b, 0xe4, 0x50, 0xa9, 0x80, 0x14, 0x24, 0x0e, 0xdc,
	0xbc, 0xce, 0xe0, 0x35, 0x64, 0x3d, 0xc6, 0xe3, 0x4d, 0x39, 0xf3, 0x13, 0xe0, 0xc6, 0xdf, 0xe0,
	0xa7, 0x70, 0xe3, 0x8c, 0xf8, 0x1d, 0x68, 0x77, 0xbb, 0x4a, 0xd5, 0x72, 0xe8, 0xc9, 0xf2, 0x7b,
	0x6f, 0xe6, 0x8d, 0x9f, 0x87, 0x5d, 0x57, 0xde, 0x4a, 0x58, 0xab, 0x55, 0xa9, 0x22, 0x86, 0xd4,
	0x07, 0x8c, 0xc8, 0xbb, 0xca, 0xdb, 0x3e, 0xaf, 0x98, 0x02, 0x88, 0x94, 0x01, 0x6a, 0x88, 0xfe,
	0x6d, 0x83, 0x68, 0x56, 0x20, 0x2b, 0x4a, 0x39, 0x87, 0x51, 0x45, 0x8b, 0xae, 0x65, 0x1f, 0xd6,
	0x87, 0x1e, 0x19, 0x70, 0x23, 0x3a, 0x56, 0xc6, 0x40, 0x90, 0xe8, 0x6b, 0xc5, 0x79, 0xf5, 0xe0,
	0x31, 0xbb, 0x36, 0x6f, 0x7c, 0x61, 0x01, 0x9f, 0x4b, 0xa0, 0xc8, 0xef, 0xb1, 0xad, 0x42, 0x45,
	0x9d, 0x03, 0xed, 0x74, 0x44, 0xf7, 0xfe, 0xd5, 0x09, 0x4b, 0x95, 0xb7, 0xe9, 0xab, 0x0a, 0x5b,
	0xb4, 0xd4, 0xe0, 0x09, 0xdb, 0xde, 0x14, 0x92, 0x47, 0x47, 0x70, 0xb1, 0xca, 0x09, 0xb2, 0x2b,
	0xf3, 0xf6, 0xa9, 0x3c, 0x63, 0x97, 0xdb, 0x36, 0xfc, 0x46, 0xad, 0x3e, 0x33, 0x4e, 0xff, 0xe6,
	0x19, 0xb4, 0xf1, 0x1a, 0x3c, 0xf8, 0xfa, 0xeb, 0xcf, 0xf7, 0x64, 0x77, 0x70, 0x47, 0xae, 0xf7,
	0x36, 0xd1, 0xc9, 0x13, 0x93, 0xd9, 0x09, 0x02, 0xb3, 0xce, 0xf0, 0xf9, 0xdf, 0xe4, 0xdb, 0xc1,
	0xef, 0x84, 0xff, 0xec, 0x9c, 0x32, 0x1e, 0x1c, 0x32, 0xf6, 0xc6, 0x83, 0x13, 0xf5, 0x70, 0xfc,
	0x56, 0x1e, 0xa3, 0xa7, 0x99, 0x94, 0xe8, 0xc1, 0x8d, 0xea, 0x26, 0xe9, 0x12, 0xd6, 0xfd, 0xdd,
	0xcd, 0x7d, 0xb4, 0xb4, 0xa4, 0x4b, 0xa2, 0xfd, 0x26, 0x7d, 0x13, 0xb0, 0xf4, 0x94, 0x6a, 0x2c,
	0x86, 0xef, 0x18, 0x3f, 0xf0, 0x4a, 0xe7, 0x20, 0x26, 0xe9, 0x58, 0x1c, 0x59, 0x0d, 0x55, 0x18,
	0xfb, 0x6d, 0x4b, 0x63, 0x63, 0x5e, 0x66, 0x95, 0x52, 0x36, 0xa5, 0x1f, 0x30, 0x18, 0x55, 0x00,
	0x9d, 0x32, 0x93, 0xd9, 0x0a, 0x33, 0x59, 0x28, 0x8a, 0x10, 0xe4, 0xd1, 0xe1, 0x8b, 0xf9, 0xeb,
	0xb7, 0xf3, 0x49, 0x77, 0x2f, 0x1d, 0x0f, 0x93, 0x4e, 0x32, 0xd9, 0x56, 0xde, 0xaf, 0xac, 0xae,
	0x3f, 0x4e, 0x7e, 0x24, 0x74, 0xb3, 0x73, 0xc8, 0xe2, 0x29, 0xeb, 0x4e, 0xc7, 0x53, 0x3e, 0x65,
	0xc3, 0x05, 0xc4, 0x32, 0x38, 0x58, 0x8a, 0xe3, 0x1c, 0x9c, 0x88, 0x39, 0x88, 0x00, 0x84, 0x65,
	0xd0, 0x20, 0x96, 0x08, 0x24, 0x1c, 0x46, 0x01, 0x5f, 0x2c, 0xc5, 0x94, 0xf7, 0xd8, 0xa5, 0x1f,
	0x49, 0x67, 0x2b, 0x3c, 0x63, 0x3b, 0x9b, 0x30, 0xc4, 0x4b, 0xd4, 0x65, 0x01, 0xae, 0x59, 0x14,
	0x7e, 0xf7, 0xff, 0xd1, 0x48, 0xb2, 0x11, 0xe4, 0x12, 0x35, 0xc9, 0xf7, 0x3d, 0xff, 0xc9, 0x48,
	0x9f, 0x65, 0xbd, 0x7a, 0xa7, 0x1e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x87, 0x98, 0x9e, 0x87,
	0xcf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EvaluatorClient is the client API for Evaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluatorClient interface {
	// Evaluate accepts a list of proposed matches, evaluates them for quality,
	// collisions etc. and returns matches that should be accepted as results.
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error)
}

type evaluatorClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatorClient(cc *grpc.ClientConn) EvaluatorClient {
	return &evaluatorClient{cc}
}

func (c *evaluatorClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResponse, error) {
	out := new(EvaluateResponse)
	err := c.cc.Invoke(ctx, "/api.Evaluator/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluatorServer is the server API for Evaluator service.
type EvaluatorServer interface {
	// Evaluate accepts a list of proposed matches, evaluates them for quality,
	// collisions etc. and returns matches that should be accepted as results.
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResponse, error)
}

func RegisterEvaluatorServer(s *grpc.Server, srv EvaluatorServer) {
	s.RegisterService(&_Evaluator_serviceDesc, srv)
}

func _Evaluator_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Evaluator/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Evaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Evaluator",
	HandlerType: (*EvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _Evaluator_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/evaluator.proto",
}
