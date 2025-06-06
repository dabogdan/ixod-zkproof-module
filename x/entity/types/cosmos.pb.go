// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/entity/v1beta1/cosmos.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Grant gives permissions to execute
// the provide method with expiration time.
type Grant struct {
	Authorization *types.Any `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    time.Time  `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bb875cebaa6563e, []int{0}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetAuthorization() *types.Any {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *Grant) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Grant)(nil), "ixo.entity.v1beta1.Grant")
}

func init() { proto.RegisterFile("ixo/entity/v1beta1/cosmos.proto", fileDescriptor_3bb875cebaa6563e) }

var fileDescriptor_3bb875cebaa6563e = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x63, 0x24, 0x10, 0x0a, 0xea, 0x40, 0xd4, 0x01, 0x32, 0x38, 0x88, 0x89, 0xa5, 0xb6,
	0x4a, 0x27, 0xc6, 0x46, 0x48, 0x4c, 0x2c, 0x85, 0x89, 0x05, 0x39, 0xc1, 0x4d, 0x2c, 0x1a, 0xff,
	0x51, 0xf2, 0xa7, 0x4a, 0x38, 0x45, 0x0f, 0xc0, 0x31, 0x38, 0x44, 0xc5, 0xd4, 0x91, 0x09, 0x50,
	0x72, 0x11, 0xd4, 0x98, 0x48, 0x81, 0x6e, 0x7e, 0x7a, 0xdf, 0xf7, 0x6c, 0xc9, 0xb6, 0xa7, 0x4a,
	0xe0, 0x52, 0xa3, 0xc2, 0x8a, 0x2f, 0xc7, 0x81, 0x44, 0x31, 0xe6, 0x21, 0xe4, 0x09, 0xe4, 0x2c,
	0xcd, 0x00, 0xc1, 0x71, 0x54, 0x09, 0xcc, 0x00, 0xec, 0x17, 0x70, 0x87, 0x11, 0x44, 0xd0, 0xd6,
	0x7c, 0x7b, 0x32, 0xa4, 0x7b, 0x6a, 0xbc, 0x47, 0x53, 0xf4, 0x47, 0x5c, 0x2f, 0x02, 0x88, 0x16,
	0x92, 0xb7, 0x29, 0x28, 0xe6, 0x1c, 0x55, 0x22, 0x73, 0x14, 0x49, 0xda, 0xb9, 0xff, 0x01, 0xa1,
	0x2b, 0x53, 0x9d, 0xbf, 0x12, 0x7b, 0xff, 0x26, 0x13, 0x1a, 0x9d, 0x5b, 0x7b, 0x20, 0x0a, 0x8c,
	0x21, 0x53, 0x2f, 0x02, 0x15, 0xe8, 0x13, 0x72, 0x46, 0x2e, 0x8e, 0x2e, 0x87, 0xcc, 0xc8, 0xac,
	0x93, 0xd9, 0x54, 0x57, 0xfe, 0xf1, 0xfb, 0xdb, 0x68, 0x30, 0xed, 0xe3, 0xb3, 0xbf, 0xb6, 0x73,
	0x6d, 0xdb, 0xb2, 0x4c, 0x55, 0x66, 0xb6, 0xf6, 0xda, 0x2d, 0x77, 0x67, 0xeb, 0xbe, 0x7b, 0xa9,
	0x7f, 0xb8, 0xfe, 0xf4, 0xac, 0xd5, 0x97, 0x47, 0x66, 0x3d, 0xcf, 0xbf, 0x5b, 0xd7, 0x94, 0x6c,
	0x6a, 0x4a, 0xbe, 0x6b, 0x4a, 0x56, 0x0d, 0xb5, 0x36, 0x0d, 0xb5, 0x3e, 0x1a, 0x6a, 0x3d, 0x5c,
	0x45, 0x0a, 0xe3, 0x22, 0x60, 0x21, 0x24, 0x5c, 0x95, 0x30, 0x87, 0x42, 0x3f, 0xb5, 0xce, 0x36,
	0x8d, 0x82, 0x05, 0x84, 0xcf, 0x61, 0x2c, 0x94, 0xe6, 0xcb, 0x09, 0x2f, 0xbb, 0x3f, 0xc0, 0x2a,
	0x95, 0x79, 0x70, 0xd0, 0x5e, 0x3f, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x39, 0x92, 0x8e,
	0x9e, 0x01, 0x00, 0x00,
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCosmos(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCosmos(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCosmos(dAtA []byte, offset int, v uint64) int {
	offset -= sovCosmos(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovCosmos(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovCosmos(uint64(l))
	return n
}

func sovCosmos(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCosmos(x uint64) (n int) {
	return sovCosmos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCosmos
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCosmos
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
				return ErrInvalidLengthCosmos
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCosmos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &types.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCosmos
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
				return ErrInvalidLengthCosmos
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCosmos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCosmos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCosmos
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
func skipCosmos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCosmos
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
					return 0, ErrIntOverflowCosmos
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
					return 0, ErrIntOverflowCosmos
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
				return 0, ErrInvalidLengthCosmos
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCosmos
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCosmos
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCosmos        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCosmos          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCosmos = fmt.Errorf("proto: unexpected end of group")
)
