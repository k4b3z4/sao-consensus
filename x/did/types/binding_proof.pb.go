// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/did/binding_proof.proto

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

type BindingProof struct {
	Version   int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Account   string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Did       string `protobuf:"bytes,5,opt,name=did,proto3" json:"did,omitempty"`
	Timestamp uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Address   string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	ProofType string `protobuf:"bytes,8,opt,name=proofType,proto3" json:"proofType,omitempty"`
	ChainID   int32  `protobuf:"varint,9,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (m *BindingProof) Reset()         { *m = BindingProof{} }
func (m *BindingProof) String() string { return proto.CompactTextString(m) }
func (*BindingProof) ProtoMessage()    {}
func (*BindingProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f59c20c74c278a, []int{0}
}
func (m *BindingProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BindingProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BindingProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BindingProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingProof.Merge(m, src)
}
func (m *BindingProof) XXX_Size() int {
	return m.Size()
}
func (m *BindingProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingProof.DiscardUnknown(m)
}

var xxx_messageInfo_BindingProof proto.InternalMessageInfo

func (m *BindingProof) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BindingProof) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BindingProof) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *BindingProof) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *BindingProof) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *BindingProof) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BindingProof) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BindingProof) GetProofType() string {
	if m != nil {
		return m.ProofType
	}
	return ""
}

func (m *BindingProof) GetChainID() int32 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func init() {
	proto.RegisterType((*BindingProof)(nil), "saonetwork.sao.did.BindingProof")
}

func init() { proto.RegisterFile("sao/did/binding_proof.proto", fileDescriptor_11f59c20c74c278a) }

var fileDescriptor_11f59c20c74c278a = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x27, 0xce, 0x9f, 0x53, 0x5c, 0x48, 0x57, 0x01, 0x25, 0x14, 0x41, 0xe8, 0xaa, 0x5d,
	0xf8, 0x00, 0xc2, 0xe0, 0xc6, 0x8d, 0x48, 0x75, 0xe5, 0x46, 0xd2, 0x26, 0x76, 0x82, 0x34, 0xb7,
	0xe4, 0xa6, 0xea, 0xbc, 0x85, 0x8f, 0xe5, 0x72, 0x96, 0x2e, 0xa5, 0x7d, 0x0d, 0x17, 0x92, 0xa4,
	0xe3, 0xec, 0xf2, 0x9d, 0x7b, 0xcf, 0x21, 0xf7, 0x44, 0x67, 0xc8, 0x21, 0x17, 0x4a, 0xe4, 0xa5,
	0xd2, 0x42, 0xe9, 0xfa, 0xb9, 0x35, 0x00, 0x2f, 0x59, 0x6b, 0xc0, 0x42, 0x1c, 0x23, 0x07, 0x2d,
	0xed, 0x3b, 0x98, 0xd7, 0x0c, 0x39, 0x64, 0x42, 0x89, 0x8b, 0x5f, 0x12, 0x9d, 0xac, 0xc3, 0xee,
	0xbd, 0x5b, 0x8d, 0x69, 0xb4, 0x7c, 0x93, 0x06, 0x15, 0x68, 0x4a, 0x12, 0x92, 0xce, 0x8b, 0x3d,
	0xba, 0x49, 0x23, 0x11, 0x79, 0x2d, 0xe9, 0x51, 0x42, 0xd2, 0x55, 0xb1, 0xc7, 0xf8, 0x3c, 0x5a,
	0xa1, 0xaa, 0x35, 0xb7, 0x9d, 0x91, 0x74, 0xea, 0x67, 0x07, 0xc1, 0xf9, 0x78, 0x55, 0x41, 0xa7,
	0x2d, 0x9d, 0x05, 0xdf, 0x88, 0xf1, 0x69, 0x34, 0x15, 0x4a, 0xd0, 0xb9, 0x57, 0xdd, 0xd3, 0x25,
	0x59, 0xd5, 0x48, 0xb4, 0xbc, 0x69, 0xe9, 0x22, 0x21, 0xe9, 0xac, 0x38, 0x08, 0x3e, 0x49, 0x08,
	0x23, 0x11, 0xe9, 0x72, 0x4c, 0x0a, 0xe8, 0x7c, 0xfe, 0xd2, 0xc7, 0x6d, 0x2b, 0xe9, 0x71, 0xf8,
	0xc1, 0xbf, 0xe0, 0x7c, 0xd5, 0x86, 0x2b, 0x7d, 0x7b, 0x43, 0x57, 0xe1, 0xa6, 0x11, 0xd7, 0xd7,
	0x5f, 0x3d, 0x23, 0xbb, 0x9e, 0x91, 0x9f, 0x9e, 0x91, 0xcf, 0x81, 0x4d, 0x76, 0x03, 0x9b, 0x7c,
	0x0f, 0x6c, 0xf2, 0x74, 0x59, 0x2b, 0xbb, 0xe9, 0xca, 0xac, 0x82, 0x26, 0x7f, 0xe0, 0x70, 0x17,
	0x7a, 0xcb, 0x5d, 0xbf, 0x1f, 0xbe, 0x61, 0xbb, 0x6d, 0x25, 0x96, 0x0b, 0x5f, 0xed, 0xd5, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x49, 0x1e, 0xa9, 0x79, 0x01, 0x00, 0x00,
}

func (m *BindingProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BindingProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BindingProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChainID != 0 {
		i = encodeVarintBindingProof(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ProofType) > 0 {
		i -= len(m.ProofType)
		copy(dAtA[i:], m.ProofType)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.ProofType)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintBindingProof(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintBindingProof(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintBindingProof(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBindingProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovBindingProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BindingProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovBindingProof(uint64(m.Version))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBindingProof(uint64(m.Timestamp))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	l = len(m.ProofType)
	if l > 0 {
		n += 1 + l + sovBindingProof(uint64(l))
	}
	if m.ChainID != 0 {
		n += 1 + sovBindingProof(uint64(m.ChainID))
	}
	return n
}

func sovBindingProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBindingProof(x uint64) (n int) {
	return sovBindingProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BindingProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBindingProof
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
			return fmt.Errorf("proto: BindingProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BindingProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
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
				return ErrInvalidLengthBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBindingProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBindingProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBindingProof
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
func skipBindingProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBindingProof
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
					return 0, ErrIntOverflowBindingProof
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
					return 0, ErrIntOverflowBindingProof
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
				return 0, ErrInvalidLengthBindingProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBindingProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBindingProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBindingProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBindingProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBindingProof = fmt.Errorf("proto: unexpected end of group")
)
