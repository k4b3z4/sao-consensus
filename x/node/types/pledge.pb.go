// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/pledge.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Pledge struct {
	Creator    string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Pledged    types.Coin `protobuf:"bytes,2,opt,name=pledged,proto3" json:"pledged"`
	Reward     types.Coin `protobuf:"bytes,3,opt,name=reward,proto3" json:"reward"`
	RewardDebt types.Coin `protobuf:"bytes,4,opt,name=rewardDebt,proto3" json:"rewardDebt"`
}

func (m *Pledge) Reset()         { *m = Pledge{} }
func (m *Pledge) String() string { return proto.CompactTextString(m) }
func (*Pledge) ProtoMessage()    {}
func (*Pledge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44a3a94106aee56, []int{0}
}
func (m *Pledge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pledge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pledge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pledge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pledge.Merge(m, src)
}
func (m *Pledge) XXX_Size() int {
	return m.Size()
}
func (m *Pledge) XXX_DiscardUnknown() {
	xxx_messageInfo_Pledge.DiscardUnknown(m)
}

var xxx_messageInfo_Pledge proto.InternalMessageInfo

func (m *Pledge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Pledge) GetPledged() types.Coin {
	if m != nil {
		return m.Pledged
	}
	return types.Coin{}
}

func (m *Pledge) GetReward() types.Coin {
	if m != nil {
		return m.Reward
	}
	return types.Coin{}
}

func (m *Pledge) GetRewardDebt() types.Coin {
	if m != nil {
		return m.RewardDebt
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Pledge)(nil), "saonetwork.sao.node.Pledge")
}

func init() { proto.RegisterFile("sao/node/pledge.proto", fileDescriptor_c44a3a94106aee56) }

var fileDescriptor_c44a3a94106aee56 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0x2f, 0x5a, 0xae, 0x18, 0xb7, 0xa8, 0x10, 0x3b, 0xc4, 0xe2, 0x20, 0x9d, 0x12, 0xaa,
	0x83, 0x38, 0x29, 0xd5, 0x59, 0xa4, 0x6e, 0x6e, 0xc9, 0xdd, 0xc7, 0x79, 0x68, 0xef, 0x3b, 0x92,
	0x68, 0xf5, 0x5f, 0xf8, 0xb3, 0x3a, 0x76, 0xd3, 0x49, 0xe4, 0xee, 0x8f, 0x48, 0x2e, 0x2d, 0x74,
	0xec, 0xf6, 0x85, 0xf7, 0x7d, 0x9f, 0xc0, 0x43, 0x8f, 0x9c, 0x46, 0x55, 0x61, 0x0e, 0xaa, 0x7e,
	0x85, 0xbc, 0x00, 0x59, 0x5b, 0xf4, 0xc8, 0x0e, 0x9c, 0xc6, 0x0a, 0xfc, 0x1c, 0xed, 0x8b, 0x74,
	0x1a, 0x65, 0x68, 0x0c, 0x0e, 0x0b, 0x2c, 0xb0, 0xcb, 0x55, 0xb8, 0x62, 0x75, 0x20, 0x32, 0x74,
	0x33, 0x74, 0xca, 0x68, 0x07, 0xea, 0x7d, 0x6c, 0xc0, 0xeb, 0xb1, 0xca, 0xb0, 0xac, 0x62, 0x7e,
	0xfa, 0x4d, 0x68, 0xfa, 0xd0, 0xb1, 0x19, 0xa7, 0xfd, 0xcc, 0x82, 0xf6, 0x68, 0x39, 0x19, 0x92,
	0xd1, 0xde, 0x74, 0xfd, 0x64, 0x57, 0xb4, 0x1f, 0xff, 0xcf, 0xf9, 0xce, 0x90, 0x8c, 0xf6, 0xcf,
	0x8f, 0x65, 0xc4, 0xca, 0x80, 0x95, 0x2b, 0xac, 0xbc, 0xc5, 0xb2, 0x9a, 0xf4, 0x16, 0xbf, 0x27,
	0xc9, 0x74, 0xdd, 0x67, 0x97, 0x34, 0xb5, 0x30, 0xd7, 0x36, 0xe7, 0xbb, 0xdb, 0x2d, 0x57, 0x75,
	0x76, 0x4d, 0x69, 0xbc, 0xee, 0xc0, 0x78, 0xde, 0xdb, 0x6e, 0xbc, 0x31, 0x99, 0xdc, 0x2c, 0x1a,
	0x41, 0x96, 0x8d, 0x20, 0x7f, 0x8d, 0x20, 0x5f, 0xad, 0x48, 0x96, 0xad, 0x48, 0x7e, 0x5a, 0x91,
	0x3c, 0x9d, 0x15, 0xa5, 0x7f, 0x7e, 0x33, 0x32, 0xc3, 0x99, 0x7a, 0xd4, 0x78, 0x1f, 0x4d, 0xaa,
	0xe0, 0xfa, 0x23, 0xda, 0xf6, 0x9f, 0x35, 0x38, 0x93, 0x76, 0x8a, 0x2e, 0xfe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0xf1, 0x64, 0x57, 0x86, 0x01, 0x00, 0x00,
}

func (m *Pledge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pledge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pledge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RewardDebt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Pledged.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPledge(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPledge(dAtA []byte, offset int, v uint64) int {
	offset -= sovPledge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pledge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPledge(uint64(l))
	}
	l = m.Pledged.Size()
	n += 1 + l + sovPledge(uint64(l))
	l = m.Reward.Size()
	n += 1 + l + sovPledge(uint64(l))
	l = m.RewardDebt.Size()
	n += 1 + l + sovPledge(uint64(l))
	return n
}

func sovPledge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPledge(x uint64) (n int) {
	return sovPledge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pledge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPledge
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
			return fmt.Errorf("proto: Pledge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pledge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pledged", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pledged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDebt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPledge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPledge
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
func skipPledge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPledge
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
					return 0, ErrIntOverflowPledge
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
					return 0, ErrIntOverflowPledge
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
				return 0, ErrInvalidLengthPledge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPledge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPledge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPledge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPledge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPledge = fmt.Errorf("proto: unexpected end of group")
)