// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/zkproof/v1beta/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryProofRequest struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ClaimType string `protobuf:"bytes,2,opt,name=claim_type,json=claimType,proto3" json:"claim_type,omitempty"`
}

func (m *QueryProofRequest) Reset()         { *m = QueryProofRequest{} }
func (m *QueryProofRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProofRequest) ProtoMessage()    {}
func (*QueryProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee094e553bd11a3, []int{0}
}
func (m *QueryProofRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProofRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProofRequest.Merge(m, src)
}
func (m *QueryProofRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProofRequest proto.InternalMessageInfo

func (m *QueryProofRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QueryProofRequest) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

type QueryProofResponse struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ClaimType   string `protobuf:"bytes,2,opt,name=claim_type,json=claimType,proto3" json:"claim_type,omitempty"`
	ProofData   string `protobuf:"bytes,3,opt,name=proof_data,json=proofData,proto3" json:"proof_data,omitempty"`
	PublicInput string `protobuf:"bytes,4,opt,name=public_input,json=publicInput,proto3" json:"public_input,omitempty"`
}

func (m *QueryProofResponse) Reset()         { *m = QueryProofResponse{} }
func (m *QueryProofResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProofResponse) ProtoMessage()    {}
func (*QueryProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee094e553bd11a3, []int{1}
}
func (m *QueryProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProofResponse.Merge(m, src)
}
func (m *QueryProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProofResponse proto.InternalMessageInfo

func (m *QueryProofResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QueryProofResponse) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

func (m *QueryProofResponse) GetProofData() string {
	if m != nil {
		return m.ProofData
	}
	return ""
}

func (m *QueryProofResponse) GetPublicInput() string {
	if m != nil {
		return m.PublicInput
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryProofRequest)(nil), "ixo.zkproof.v1beta.QueryProofRequest")
	proto.RegisterType((*QueryProofResponse)(nil), "ixo.zkproof.v1beta.QueryProofResponse")
}

func init() { proto.RegisterFile("ixo/zkproof/v1beta/query.proto", fileDescriptor_bee094e553bd11a3) }

var fileDescriptor_bee094e553bd11a3 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x1b, 0xb5, 0x4a, 0xa3, 0x8b, 0xc1, 0xe1, 0x28, 0x18, 0xb4, 0xa0, 0xb8, 0x98, 0xa0,
	0xe2, 0xe2, 0x28, 0x2e, 0x82, 0x83, 0x96, 0xe2, 0xe0, 0x72, 0xe4, 0xd2, 0xb4, 0x0d, 0x6d, 0xef,
	0xa5, 0x77, 0x49, 0xb9, 0xfa, 0x25, 0xf4, 0x63, 0x39, 0x76, 0x74, 0x94, 0xbb, 0x2f, 0x22, 0xc9,
	0xa9, 0x08, 0x1d, 0x04, 0xb7, 0xbc, 0xff, 0xef, 0x9f, 0xc7, 0x7b, 0xff, 0x87, 0xa9, 0x2e, 0x80,
	0x3f, 0x8f, 0x4d, 0x06, 0x30, 0xe0, 0xf3, 0xb3, 0x44, 0x59, 0xc1, 0x67, 0x4e, 0x65, 0x0b, 0x66,
	0x32, 0xb0, 0x40, 0x88, 0x2e, 0x80, 0x7d, 0x71, 0x56, 0xf3, 0xf6, 0xde, 0x10, 0x86, 0x10, 0x30,
	0xf7, 0xaf, 0xda, 0xd9, 0xb9, 0xc3, 0xbb, 0x0f, 0xfe, 0xe3, 0xbd, 0xb7, 0x76, 0xd5, 0xcc, 0xa9,
	0xdc, 0x92, 0x08, 0x6f, 0xc9, 0x4c, 0x09, 0x0b, 0x59, 0x84, 0x0e, 0xd0, 0x49, 0xab, 0xfb, 0x5d,
	0x92, 0x7d, 0x8c, 0xe5, 0x44, 0xe8, 0x69, 0x6c, 0x17, 0x46, 0x45, 0x6b, 0x01, 0xb6, 0x82, 0xd2,
	0x5b, 0x18, 0xd5, 0x79, 0x41, 0x98, 0xfc, 0x6e, 0x97, 0x1b, 0x48, 0x73, 0xf5, 0xef, 0x7e, 0x1e,
	0x87, 0x1d, 0xe2, 0xbe, 0xb0, 0x22, 0x5a, 0xaf, 0x71, 0x50, 0x6e, 0x84, 0x15, 0xe4, 0x10, 0xef,
	0x18, 0x97, 0x4c, 0xb4, 0x8c, 0x75, 0x6a, 0x9c, 0x8d, 0x36, 0x82, 0x61, 0xbb, 0xd6, 0x6e, 0xbd,
	0x74, 0x1e, 0xe3, 0x66, 0x18, 0x88, 0x3c, 0xe2, 0x66, 0x18, 0x8a, 0x1c, 0xb1, 0xd5, 0x70, 0xd8,
	0x4a, 0x06, 0xed, 0xe3, 0xbf, 0x6c, 0xf5, 0x6e, 0xd7, 0xbd, 0xb7, 0x92, 0xa2, 0x65, 0x49, 0xd1,
	0x47, 0x49, 0xd1, 0x6b, 0x45, 0x1b, 0xcb, 0x8a, 0x36, 0xde, 0x2b, 0xda, 0x78, 0xba, 0x1a, 0x6a,
	0x3b, 0x72, 0x09, 0x93, 0x30, 0xe5, 0xba, 0x80, 0x01, 0xb8, 0xb4, 0x2f, 0xac, 0x86, 0xd4, 0x57,
	0xa7, 0xc9, 0x04, 0xe4, 0x58, 0x8e, 0x84, 0x4e, 0xf9, 0xfc, 0x92, 0x17, 0x3f, 0xd7, 0xf4, 0x41,
	0xe4, 0xc9, 0x66, 0xb8, 0xce, 0xc5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x9d, 0x8b, 0x1e,
	0xe9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Proof(ctx context.Context, in *QueryProofRequest, opts ...grpc.CallOption) (*QueryProofResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proof(ctx context.Context, in *QueryProofRequest, opts ...grpc.CallOption) (*QueryProofResponse, error) {
	out := new(QueryProofResponse)
	err := c.cc.Invoke(ctx, "/ixo.zkproof.v1beta.Query/Proof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Proof(context.Context, *QueryProofRequest) (*QueryProofResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proof(ctx context.Context, req *QueryProofRequest) (*QueryProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proof not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.zkproof.v1beta.Query/Proof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proof(ctx, req.(*QueryProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ixo.zkproof.v1beta.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proof",
			Handler:    _Query_Proof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ixo/zkproof/v1beta/query.proto",
}

func (m *QueryProofRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProofRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProofRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicInput) > 0 {
		i -= len(m.PublicInput)
		copy(dAtA[i:], m.PublicInput)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PublicInput)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProofData) > 0 {
		i -= len(m.ProofData)
		copy(dAtA[i:], m.ProofData)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ProofData)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProofRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ProofData)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PublicInput)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProofRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProofRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProofRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicInput", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicInput = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
