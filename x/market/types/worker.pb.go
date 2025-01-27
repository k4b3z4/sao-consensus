// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/market/worker.proto

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

type Worker struct {
	Workername      string        `protobuf:"bytes,1,opt,name=workername,proto3" json:"workername,omitempty"`
	Storage         uint64        `protobuf:"varint,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Debt            types.DecCoin `protobuf:"bytes,3,opt,name=debt,proto3" json:"debt"`
	Reward          types.DecCoin `protobuf:"bytes,4,opt,name=reward,proto3" json:"reward"`
	IncomePerSecond types.DecCoin `protobuf:"bytes,5,opt,name=incomePerSecond,proto3" json:"incomePerSecond"`
	TotalStorage    uint64        `protobuf:"varint,6,opt,name=total_storage,json=totalStorage,proto3" json:"total_storage,omitempty"`
	LastRewardAt    int64         `protobuf:"varint,7,opt,name=last_reward_at,json=lastRewardAt,proto3" json:"last_reward_at,omitempty"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_c067fca2be4b797f, []int{0}
}
func (m *Worker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return m.Size()
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetWorkername() string {
	if m != nil {
		return m.Workername
	}
	return ""
}

func (m *Worker) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Worker) GetDebt() types.DecCoin {
	if m != nil {
		return m.Debt
	}
	return types.DecCoin{}
}

func (m *Worker) GetReward() types.DecCoin {
	if m != nil {
		return m.Reward
	}
	return types.DecCoin{}
}

func (m *Worker) GetIncomePerSecond() types.DecCoin {
	if m != nil {
		return m.IncomePerSecond
	}
	return types.DecCoin{}
}

func (m *Worker) GetTotalStorage() uint64 {
	if m != nil {
		return m.TotalStorage
	}
	return 0
}

func (m *Worker) GetLastRewardAt() int64 {
	if m != nil {
		return m.LastRewardAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Worker)(nil), "saonetwork.sao.market.Worker")
}

func init() { proto.RegisterFile("sao/market/worker.proto", fileDescriptor_c067fca2be4b797f) }

var fileDescriptor_c067fca2be4b797f = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x9b, 0xad, 0xff, 0x8e, 0x7f, 0x9c, 0x0a, 0x45, 0xb1, 0x0c, 0x89, 0x45, 0x3d, 0xf4,
	0x94, 0x30, 0x05, 0x0f, 0xde, 0x9c, 0x1e, 0x45, 0xa4, 0x3b, 0x08, 0x5e, 0x46, 0xda, 0xbd, 0xd4,
	0xb2, 0xb5, 0xef, 0x48, 0xa2, 0xd3, 0x6f, 0xe1, 0xc7, 0xda, 0xcd, 0x1d, 0x3d, 0x89, 0x6c, 0x5f,
	0x44, 0xda, 0x6c, 0x20, 0x9e, 0x76, 0x4b, 0x7e, 0x79, 0xde, 0xf7, 0xc9, 0xc3, 0x43, 0x0f, 0xb4,
	0x44, 0x51, 0x48, 0x35, 0x02, 0x23, 0xa6, 0xa8, 0x46, 0xa0, 0xf8, 0x44, 0xa1, 0x41, 0x7f, 0x5f,
	0x4b, 0x2c, 0xc1, 0x54, 0x8c, 0x6b, 0x89, 0xdc, 0x6a, 0x3a, 0x7b, 0x19, 0x66, 0x58, 0x2b, 0x44,
	0x75, 0xb2, 0xe2, 0x0e, 0x4b, 0x51, 0x17, 0xa8, 0x45, 0x22, 0x35, 0x88, 0x97, 0x6e, 0x02, 0x46,
	0x76, 0x45, 0x8a, 0x79, 0x69, 0xdf, 0x8f, 0x3f, 0x1a, 0xd4, 0x7b, 0xa8, 0xb7, 0xfb, 0x8c, 0x52,
	0xeb, 0x53, 0xca, 0x02, 0x02, 0x12, 0x92, 0xe8, 0x7f, 0xfc, 0x8b, 0xf8, 0x01, 0x6d, 0x69, 0x83,
	0x4a, 0x66, 0x10, 0x34, 0x42, 0x12, 0xb9, 0xf1, 0xfa, 0xea, 0x5f, 0x50, 0x77, 0x08, 0x89, 0x09,
	0x9a, 0x21, 0x89, 0xb6, 0xce, 0x0e, 0xb9, 0xf5, 0xe4, 0x95, 0x27, 0x5f, 0x79, 0xf2, 0x1b, 0x48,
	0xaf, 0x31, 0x2f, 0x7b, 0xee, 0xec, 0xeb, 0xc8, 0x89, 0x6b, 0xbd, 0x7f, 0x49, 0x3d, 0x05, 0x53,
	0xa9, 0x86, 0x81, 0xbb, 0xf1, 0xe4, 0x6a, 0xc2, 0xbf, 0xa5, 0xbb, 0x79, 0x99, 0x62, 0x01, 0xf7,
	0xa0, 0xfa, 0x90, 0x62, 0x39, 0x0c, 0xfe, 0x6d, 0xbc, 0xe4, 0xef, 0xa8, 0x7f, 0x42, 0xb7, 0x0d,
	0x1a, 0x39, 0x1e, 0xac, 0x13, 0x7a, 0x75, 0xc2, 0x76, 0x0d, 0xfb, 0xab, 0x98, 0xa7, 0x74, 0x67,
	0x2c, 0xb5, 0x19, 0xd8, 0x1f, 0x0c, 0xa4, 0x09, 0x5a, 0x21, 0x89, 0x9a, 0x71, 0xbb, 0xa2, 0x71,
	0x0d, 0xaf, 0x4c, 0xaf, 0x37, 0x5b, 0x30, 0x32, 0x5f, 0x30, 0xf2, 0xbd, 0x60, 0xe4, 0x7d, 0xc9,
	0x9c, 0xf9, 0x92, 0x39, 0x9f, 0x4b, 0xe6, 0x3c, 0x46, 0x59, 0x6e, 0x9e, 0x9e, 0x13, 0x9e, 0x62,
	0x21, 0xfa, 0x12, 0xef, 0x6c, 0x87, 0xa2, 0xea, 0xf9, 0x75, 0xdd, 0xb4, 0x79, 0x9b, 0x80, 0x4e,
	0xbc, 0xba, 0x9c, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x5c, 0x1d, 0x65, 0x04, 0x02,
	0x00, 0x00,
}

func (m *Worker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Worker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Worker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastRewardAt != 0 {
		i = encodeVarintWorker(dAtA, i, uint64(m.LastRewardAt))
		i--
		dAtA[i] = 0x38
	}
	if m.TotalStorage != 0 {
		i = encodeVarintWorker(dAtA, i, uint64(m.TotalStorage))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.IncomePerSecond.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWorker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWorker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Debt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWorker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Storage != 0 {
		i = encodeVarintWorker(dAtA, i, uint64(m.Storage))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Workername) > 0 {
		i -= len(m.Workername)
		copy(dAtA[i:], m.Workername)
		i = encodeVarintWorker(dAtA, i, uint64(len(m.Workername)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorker(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Worker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Workername)
	if l > 0 {
		n += 1 + l + sovWorker(uint64(l))
	}
	if m.Storage != 0 {
		n += 1 + sovWorker(uint64(m.Storage))
	}
	l = m.Debt.Size()
	n += 1 + l + sovWorker(uint64(l))
	l = m.Reward.Size()
	n += 1 + l + sovWorker(uint64(l))
	l = m.IncomePerSecond.Size()
	n += 1 + l + sovWorker(uint64(l))
	if m.TotalStorage != 0 {
		n += 1 + sovWorker(uint64(m.TotalStorage))
	}
	if m.LastRewardAt != 0 {
		n += 1 + sovWorker(uint64(m.LastRewardAt))
	}
	return n
}

func sovWorker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorker(x uint64) (n int) {
	return sovWorker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Worker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorker
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
			return fmt.Errorf("proto: Worker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Worker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Workername", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
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
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Workername = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			m.Storage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Storage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
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
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Debt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
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
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncomePerSecond", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
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
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncomePerSecond.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStorage", wireType)
			}
			m.TotalStorage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalStorage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardAt", wireType)
			}
			m.LastRewardAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastRewardAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorker
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
func skipWorker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorker
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
					return 0, ErrIntOverflowWorker
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
					return 0, ErrIntOverflowWorker
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
				return 0, ErrInvalidLengthWorker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorker = fmt.Errorf("proto: unexpected end of group")
)
