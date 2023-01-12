// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/did/account_id.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type AccountId struct {
	AccountDid string `protobuf:"bytes,1,opt,name=accountDid,proto3" json:"accountDid,omitempty"`
	AccountId  string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Creator    string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733f72293115692, []int{0}
}
func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(m, src)
}
func (m *AccountId) XXX_Size() int {
	return m.Size()
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetAccountDid() string {
	if m != nil {
		return m.AccountDid
	}
	return ""
}

func (m *AccountId) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccountId) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountId)(nil), "saonetwork.sao.did.AccountId")
}

func init() { proto.RegisterFile("sao/did/account_id.proto", fileDescriptor_1733f72293115692) }

var fileDescriptor_1733f72293115692 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4e, 0xcc, 0xd7,
	0x4f, 0xc9, 0x4c, 0xd1, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca,
	0xd6, 0x2b, 0x4e, 0xcc, 0xd7, 0x4b, 0xc9, 0x4c, 0x51, 0x4a, 0xe6, 0xe2, 0x74, 0x84, 0xa8, 0xf3,
	0x4c, 0x11, 0x92, 0xe3, 0xe2, 0x82, 0x6a, 0x72, 0xc9, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x42, 0x12, 0x11, 0x92, 0xe1, 0xe2, 0x4c, 0x84, 0x29, 0x96, 0x60, 0x02, 0x4b, 0x23, 0x04,
	0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x98, 0xc1, 0x72, 0x30,
	0xae, 0x93, 0xfd, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x07, 0x27, 0xe6, 0xfb, 0x41, 0x5c, 0xa7,
	0x0f, 0xf2, 0x42, 0x05, 0xd8, 0x13, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x0f, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xbc, 0x07, 0x31, 0xdc, 0x00, 0x00, 0x00,
}

func (m *AccountId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAccountId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountId) > 0 {
		i -= len(m.AccountId)
		copy(dAtA[i:], m.AccountId)
		i = encodeVarintAccountId(dAtA, i, uint64(len(m.AccountId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountDid) > 0 {
		i -= len(m.AccountDid)
		copy(dAtA[i:], m.AccountDid)
		i = encodeVarintAccountId(dAtA, i, uint64(len(m.AccountDid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountId(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountDid)
	if l > 0 {
		n += 1 + l + sovAccountId(uint64(l))
	}
	l = len(m.AccountId)
	if l > 0 {
		n += 1 + l + sovAccountId(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAccountId(uint64(l))
	}
	return n
}

func sovAccountId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountId(x uint64) (n int) {
	return sovAccountId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountId
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
			return fmt.Errorf("proto: AccountId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountId
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
				return ErrInvalidLengthAccountId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountId
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
				return ErrInvalidLengthAccountId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountId
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
				return ErrInvalidLengthAccountId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountId
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
func skipAccountId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountId
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
					return 0, ErrIntOverflowAccountId
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
					return 0, ErrIntOverflowAccountId
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
				return 0, ErrInvalidLengthAccountId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountId = fmt.Errorf("proto: unexpected end of group")
)
