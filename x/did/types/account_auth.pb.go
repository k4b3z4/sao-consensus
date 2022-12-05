// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/did/account_auth.proto

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

type AccountAuth struct {
	AccountDid           string `protobuf:"bytes,1,opt,name=accountDid,proto3" json:"accountDid,omitempty"`
	AccountEncryptedSeed string `protobuf:"bytes,2,opt,name=accountEncryptedSeed,proto3" json:"accountEncryptedSeed,omitempty"`
	SidEncryptedAccount  string `protobuf:"bytes,3,opt,name=sidEncryptedAccount,proto3" json:"sidEncryptedAccount,omitempty"`
}

func (m *AccountAuth) Reset()         { *m = AccountAuth{} }
func (m *AccountAuth) String() string { return proto.CompactTextString(m) }
func (*AccountAuth) ProtoMessage()    {}
func (*AccountAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ecf7a147823a3b0, []int{0}
}
func (m *AccountAuth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountAuth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAuth.Merge(m, src)
}
func (m *AccountAuth) XXX_Size() int {
	return m.Size()
}
func (m *AccountAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAuth.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAuth proto.InternalMessageInfo

func (m *AccountAuth) GetAccountDid() string {
	if m != nil {
		return m.AccountDid
	}
	return ""
}

func (m *AccountAuth) GetAccountEncryptedSeed() string {
	if m != nil {
		return m.AccountEncryptedSeed
	}
	return ""
}

func (m *AccountAuth) GetSidEncryptedAccount() string {
	if m != nil {
		return m.SidEncryptedAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountAuth)(nil), "saonetwork.sao.did.AccountAuth")
}

func init() { proto.RegisterFile("sao/did/account_auth.proto", fileDescriptor_5ecf7a147823a3b0) }

var fileDescriptor_5ecf7a147823a3b0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0xcc, 0xd7,
	0x4f, 0xc9, 0x4c, 0xd1, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0x4f, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x4e, 0xcc, 0xcf, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0xcc, 0xd7, 0x4b, 0xc9, 0x4c, 0x51, 0x9a, 0xcc, 0xc8, 0xc5, 0xed,
	0x08, 0x51, 0xea, 0x58, 0x5a, 0x92, 0x21, 0x24, 0xc7, 0xc5, 0x05, 0xd5, 0xe9, 0x92, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x64, 0xc4, 0x25, 0x02, 0xe5, 0xb9, 0xe6,
	0x25, 0x17, 0x55, 0x16, 0x94, 0xa4, 0xa6, 0x04, 0xa7, 0xa6, 0xa6, 0x48, 0x30, 0x81, 0x55, 0x62,
	0x95, 0x13, 0x32, 0xe0, 0x12, 0x2e, 0xce, 0x4c, 0x81, 0x8b, 0x41, 0xad, 0x93, 0x60, 0x06, 0x6b,
	0xc1, 0x26, 0xe5, 0x64, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xaa,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc1, 0x89, 0xf9, 0x7e, 0x10,
	0xef, 0xe8, 0x83, 0x7c, 0x5d, 0x01, 0xf6, 0x77, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8,
	0xc7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x90, 0xd3, 0xa4, 0x0f, 0x01, 0x00, 0x00,
}

func (m *AccountAuth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountAuth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountAuth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SidEncryptedAccount) > 0 {
		i -= len(m.SidEncryptedAccount)
		copy(dAtA[i:], m.SidEncryptedAccount)
		i = encodeVarintAccountAuth(dAtA, i, uint64(len(m.SidEncryptedAccount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccountEncryptedSeed) > 0 {
		i -= len(m.AccountEncryptedSeed)
		copy(dAtA[i:], m.AccountEncryptedSeed)
		i = encodeVarintAccountAuth(dAtA, i, uint64(len(m.AccountEncryptedSeed)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountDid) > 0 {
		i -= len(m.AccountDid)
		copy(dAtA[i:], m.AccountDid)
		i = encodeVarintAccountAuth(dAtA, i, uint64(len(m.AccountDid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountAuth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountDid)
	if l > 0 {
		n += 1 + l + sovAccountAuth(uint64(l))
	}
	l = len(m.AccountEncryptedSeed)
	if l > 0 {
		n += 1 + l + sovAccountAuth(uint64(l))
	}
	l = len(m.SidEncryptedAccount)
	if l > 0 {
		n += 1 + l + sovAccountAuth(uint64(l))
	}
	return n
}

func sovAccountAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountAuth(x uint64) (n int) {
	return sovAccountAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountAuth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountAuth
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
			return fmt.Errorf("proto: AccountAuth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountAuth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountAuth
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
				return ErrInvalidLengthAccountAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountEncryptedSeed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountAuth
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
				return ErrInvalidLengthAccountAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountEncryptedSeed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SidEncryptedAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountAuth
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
				return ErrInvalidLengthAccountAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SidEncryptedAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountAuth
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
func skipAccountAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountAuth
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
					return 0, ErrIntOverflowAccountAuth
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
					return 0, ErrIntOverflowAccountAuth
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
				return 0, ErrInvalidLengthAccountAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountAuth = fmt.Errorf("proto: unexpected end of group")
)