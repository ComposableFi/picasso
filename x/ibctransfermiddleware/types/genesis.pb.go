// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: composable/ibctransfermiddleware/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the ibctransfermiddleware module's genesis state.
type GenesisState struct {
	Params                Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TakenFeeByIbcSequence []types.Coin `protobuf:"bytes,2,rep,name=taken_fee_by_ibc_sequence,json=takenFeeByIbcSequence,proto3" json:"taken_fee_by_ibc_sequence"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab9a6edd8a683ba6, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTakenFeeByIbcSequence() []types.Coin {
	if m != nil {
		return m.TakenFeeByIbcSequence
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "composable.ibctransfermiddleware.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("composable/ibctransfermiddleware/v1beta1/genesis.proto", fileDescriptor_ab9a6edd8a683ba6)
}

var fileDescriptor_ab9a6edd8a683ba6 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3d, 0x4b, 0x03, 0x31,
	0x18, 0xc7, 0x2f, 0x2a, 0x05, 0xaf, 0x2e, 0x16, 0x85, 0xb6, 0x60, 0x2c, 0x4e, 0x87, 0x43, 0x62,
	0x2b, 0xe8, 0x7e, 0x8a, 0xe2, 0x22, 0xd2, 0x6e, 0x0e, 0x1e, 0x49, 0xfa, 0xf4, 0x08, 0xf6, 0x92,
	0xf3, 0x12, 0x5f, 0xee, 0x5b, 0xf8, 0x31, 0x1c, 0xfd, 0x04, 0xce, 0x1d, 0x3b, 0x3a, 0x89, 0xdc,
	0x0d, 0x7e, 0x0d, 0xb9, 0x17, 0xec, 0x52, 0xa1, 0x4b, 0x08, 0x09, 0xbf, 0xdf, 0xf3, 0xfc, 0xff,
	0xee, 0x89, 0xd0, 0x51, 0xac, 0x0d, 0xe3, 0x53, 0xa0, 0x92, 0x0b, 0x9b, 0x30, 0x65, 0x26, 0x90,
	0x44, 0x72, 0x3c, 0x9e, 0xc2, 0x33, 0x4b, 0x80, 0x3e, 0xf5, 0x39, 0x58, 0xd6, 0xa7, 0x21, 0x28,
	0x30, 0xd2, 0x90, 0x38, 0xd1, 0x56, 0xb7, 0xbc, 0x05, 0x47, 0x96, 0x72, 0xa4, 0xe6, 0xba, 0x3b,
	0xa1, 0x0e, 0x75, 0x09, 0xd1, 0xe2, 0x56, 0xf1, 0xdd, 0xf3, 0x95, 0xe7, 0x2e, 0xb7, 0x57, 0x96,
	0x6d, 0x16, 0x49, 0xa5, 0x69, 0x79, 0xd6, 0x4f, 0x58, 0x68, 0x13, 0x69, 0x43, 0x39, 0x33, 0x0b,
	0x87, 0xd0, 0x52, 0x55, 0xff, 0x07, 0x1f, 0xc8, 0xdd, 0xba, 0xac, 0xa2, 0x8c, 0x2c, 0xb3, 0xd0,
	0xba, 0x76, 0x1b, 0x31, 0x4b, 0x58, 0x64, 0xda, 0xa8, 0x87, 0xbc, 0xe6, 0xe0, 0x88, 0xac, 0x1a,
	0x8d, 0xdc, 0x94, 0x9c, 0xbf, 0x31, 0xfb, 0xda, 0x77, 0x86, 0xb5, 0xa5, 0x75, 0xe7, 0x76, 0x2c,
	0xbb, 0x07, 0x15, 0x4c, 0x00, 0x02, 0x9e, 0x06, 0x92, 0x8b, 0xc0, 0xc0, 0xc3, 0x23, 0x28, 0x01,
	0xed, 0xb5, 0xde, 0xba, 0xd7, 0x1c, 0x74, 0x48, 0xb5, 0x24, 0x29, 0x96, 0xfc, 0xb3, 0x9d, 0x69,
	0xa9, 0xfc, 0xcd, 0xc2, 0xf5, 0xf6, 0xf3, 0x7e, 0x88, 0x86, 0xbb, 0xa5, 0xe6, 0x02, 0xc0, 0x4f,
	0xaf, 0xb8, 0x18, 0xd5, 0x0a, 0xff, 0x74, 0x96, 0x61, 0x34, 0xcf, 0x30, 0xfa, 0xce, 0x30, 0x7a,
	0xcd, 0xb1, 0x33, 0xcf, 0xb1, 0xf3, 0x99, 0x63, 0xe7, 0x76, 0xef, 0xe5, 0x9f, 0x2a, 0x6d, 0x1a,
	0x83, 0xe1, 0x8d, 0xb2, 0x80, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x69, 0xde, 0x66,
	0xf3, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TakenFeeByIbcSequence) > 0 {
		for iNdEx := len(m.TakenFeeByIbcSequence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TakenFeeByIbcSequence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TakenFeeByIbcSequence) > 0 {
		for _, e := range m.TakenFeeByIbcSequence {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakenFeeByIbcSequence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TakenFeeByIbcSequence = append(m.TakenFeeByIbcSequence, types.Coin{})
			if err := m.TakenFeeByIbcSequence[len(m.TakenFeeByIbcSequence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
