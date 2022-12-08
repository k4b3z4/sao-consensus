// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/sao/proposal.proto

package types

import (
	fmt "fmt"
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

type Proposal struct {
	Owner         string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider      string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	GroupId       string   `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Duration      uint64   `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Replica       int32    `protobuf:"varint,5,opt,name=replica,proto3" json:"replica,omitempty"`
	Timeout       int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Alias         string   `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
	DataId        string   `protobuf:"bytes,8,opt,name=dataId,proto3" json:"dataId,omitempty"`
	CommitId      string   `protobuf:"bytes,9,opt,name=commitId,proto3" json:"commitId,omitempty"`
	Tags          []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Cid           string   `protobuf:"bytes,11,opt,name=cid,proto3" json:"cid,omitempty"`
	Rule          string   `protobuf:"bytes,12,opt,name=rule,proto3" json:"rule,omitempty"`
	ExtendInfo    string   `protobuf:"bytes,13,opt,name=extendInfo,proto3" json:"extendInfo,omitempty"`
	Size_         uint64   `protobuf:"varint,14,opt,name=size,proto3" json:"size,omitempty"`
	Operation     uint32   `protobuf:"varint,15,opt,name=operation,proto3" json:"operation,omitempty"`
	ReadonlyDids  []string `protobuf:"bytes,16,rep,name=readonlyDids,proto3" json:"readonlyDids,omitempty"`
	ReadwriteDids []string `protobuf:"bytes,17,rep,name=readwriteDids,proto3" json:"readwriteDids,omitempty"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e976f5d1907af06, []int{0}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Proposal) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Proposal) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Proposal) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Proposal) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Proposal) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Proposal) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Proposal) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *Proposal) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *Proposal) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Proposal) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Proposal) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *Proposal) GetExtendInfo() string {
	if m != nil {
		return m.ExtendInfo
	}
	return ""
}

func (m *Proposal) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Proposal) GetOperation() uint32 {
	if m != nil {
		return m.Operation
	}
	return 0
}

func (m *Proposal) GetReadonlyDids() []string {
	if m != nil {
		return m.ReadonlyDids
	}
	return nil
}

func (m *Proposal) GetReadwriteDids() []string {
	if m != nil {
		return m.ReadwriteDids
	}
	return nil
}

func init() {
	proto.RegisterType((*Proposal)(nil), "saonetwork.sao.sao.Proposal")
}

func init() { proto.RegisterFile("sao/sao/proposal.proto", fileDescriptor_8e976f5d1907af06) }

var fileDescriptor_8e976f5d1907af06 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xd7, 0x64, 0x77, 0xbb, 0x31, 0x5d, 0x28, 0x56, 0x55, 0x8d, 0x2a, 0x14, 0x45, 0x15,
	0x48, 0x39, 0xed, 0x1e, 0x78, 0x00, 0x24, 0xc4, 0x25, 0x17, 0x84, 0xc2, 0x8d, 0x9b, 0x1b, 0x9b,
	0x60, 0x91, 0x64, 0x2c, 0xdb, 0x61, 0x5b, 0x9e, 0x82, 0xa7, 0x42, 0x1c, 0x7b, 0xe4, 0x88, 0x76,
	0x5f, 0x04, 0x79, 0x92, 0x2d, 0xed, 0xc1, 0xd1, 0xff, 0xfd, 0xff, 0x3f, 0xca, 0x24, 0x32, 0xbf,
	0xf0, 0x12, 0xb7, 0xf1, 0x58, 0x87, 0x16, 0xbd, 0x6c, 0x37, 0xd6, 0x61, 0x40, 0x21, 0xbc, 0xc4,
	0x5e, 0x87, 0x1d, 0xba, 0x6f, 0x1b, 0x2f, 0x31, 0x9e, 0xcb, 0xf3, 0x06, 0x1b, 0xa4, 0x78, 0x1b,
	0xd5, 0xd8, 0xbc, 0xfa, 0x95, 0xf0, 0xd5, 0xc7, 0x69, 0x58, 0x9c, 0xf3, 0x05, 0xee, 0x7a, 0xed,
	0x80, 0xe5, 0xac, 0x48, 0xab, 0x11, 0xc4, 0x25, 0x5f, 0x59, 0x87, 0xdf, 0x8d, 0xd2, 0x0e, 0x9e,
	0x50, 0x70, 0xcf, 0x02, 0xf8, 0x49, 0xe3, 0x70, 0xb0, 0xa5, 0x82, 0x84, 0xa2, 0x23, 0xc6, 0x29,
	0x35, 0x38, 0x19, 0x0c, 0xf6, 0x30, 0xcf, 0x59, 0x31, 0xaf, 0xee, 0x39, 0x4e, 0x39, 0x6d, 0x5b,
	0x53, 0x4b, 0x58, 0xe4, 0xac, 0x58, 0x54, 0x47, 0x8c, 0x49, 0x30, 0x9d, 0xc6, 0x21, 0xc0, 0x72,
	0x4c, 0x26, 0x8c, 0xbb, 0xc9, 0xd6, 0x48, 0x0f, 0x27, 0xe3, 0x6e, 0x04, 0xe2, 0x82, 0x2f, 0x95,
	0x0c, 0xb2, 0x54, 0xb0, 0x22, 0x7b, 0xa2, 0xf8, 0xf6, 0x1a, 0xbb, 0xce, 0x84, 0x52, 0x41, 0x3a,
	0xee, 0x7c, 0x64, 0x21, 0xf8, 0x3c, 0xc8, 0xc6, 0x03, 0xcf, 0x93, 0x22, 0xad, 0x48, 0x8b, 0x33,
	0x9e, 0xd4, 0x46, 0xc1, 0x53, 0xaa, 0x46, 0x19, 0x5b, 0x6e, 0x68, 0x35, 0x9c, 0x92, 0x45, 0x5a,
	0x64, 0x9c, 0xeb, 0x9b, 0xa0, 0x7b, 0x55, 0xf6, 0x5f, 0x10, 0xd6, 0x94, 0x3c, 0x70, 0xe2, 0x8c,
	0x37, 0x3f, 0x34, 0x3c, 0xa3, 0xef, 0x25, 0x2d, 0x5e, 0xf2, 0x14, 0xad, 0x9e, 0x7e, 0xc4, 0xf3,
	0x9c, 0x15, 0xeb, 0xea, 0xbf, 0x21, 0xae, 0xf8, 0xa9, 0xd3, 0x52, 0x61, 0xdf, 0xde, 0xbe, 0x37,
	0xca, 0xc3, 0x19, 0xed, 0xf4, 0xc8, 0x13, 0xaf, 0xf8, 0x3a, 0xf2, 0xce, 0x99, 0xa0, 0xa9, 0xf4,
	0x82, 0x4a, 0x8f, 0xcd, 0x77, 0x6f, 0x7f, 0xef, 0x33, 0x76, 0xb7, 0xcf, 0xd8, 0xdf, 0x7d, 0xc6,
	0x7e, 0x1e, 0xb2, 0xd9, 0xdd, 0x21, 0x9b, 0xfd, 0x39, 0x64, 0xb3, 0xcf, 0xaf, 0x1b, 0x13, 0xbe,
	0x0e, 0xd7, 0x9b, 0x1a, 0xbb, 0xed, 0x27, 0x89, 0x1f, 0xc6, 0x7b, 0x41, 0xd7, 0xe6, 0x86, 0x9e,
	0xe1, 0xd6, 0x6a, 0x7f, 0xbd, 0xa4, 0x0b, 0xf1, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x8c, 0x57, 0x72, 0x54, 0x02, 0x00, 0x00,
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReadwriteDids) > 0 {
		for iNdEx := len(m.ReadwriteDids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReadwriteDids[iNdEx])
			copy(dAtA[i:], m.ReadwriteDids[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.ReadwriteDids[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.ReadonlyDids) > 0 {
		for iNdEx := len(m.ReadonlyDids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReadonlyDids[iNdEx])
			copy(dAtA[i:], m.ReadonlyDids[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.ReadonlyDids[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if m.Operation != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Operation))
		i--
		dAtA[i] = 0x78
	}
	if m.Size_ != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x70
	}
	if len(m.ExtendInfo) > 0 {
		i -= len(m.ExtendInfo)
		copy(dAtA[i:], m.ExtendInfo)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ExtendInfo)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Rule) > 0 {
		i -= len(m.Rule)
		copy(dAtA[i:], m.Rule)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Rule)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.CommitId) > 0 {
		i -= len(m.CommitId)
		copy(dAtA[i:], m.CommitId)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.CommitId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DataId) > 0 {
		i -= len(m.DataId)
		copy(dAtA[i:], m.DataId)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.DataId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Alias) > 0 {
		i -= len(m.Alias)
		copy(dAtA[i:], m.Alias)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Alias)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timeout != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x30
	}
	if m.Replica != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Replica))
		i--
		dAtA[i] = 0x28
	}
	if m.Duration != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovProposal(uint64(m.Duration))
	}
	if m.Replica != 0 {
		n += 1 + sovProposal(uint64(m.Replica))
	}
	if m.Timeout != 0 {
		n += 1 + sovProposal(uint64(m.Timeout))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.DataId)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.CommitId)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Rule)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ExtendInfo)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovProposal(uint64(m.Size_))
	}
	if m.Operation != 0 {
		n += 1 + sovProposal(uint64(m.Operation))
	}
	if len(m.ReadonlyDids) > 0 {
		for _, s := range m.ReadonlyDids {
			l = len(s)
			n += 2 + l + sovProposal(uint64(l))
		}
	}
	if len(m.ReadwriteDids) > 0 {
		for _, s := range m.ReadwriteDids {
			l = len(s)
			n += 2 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			m.Replica = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Replica |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtendInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadonlyDids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReadonlyDids = append(m.ReadonlyDids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadwriteDids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReadwriteDids = append(m.ReadwriteDids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
