// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identityfee/v1/genesis.proto

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

// GenesisState defines the ssifee module's genesis state.
type GenesisState struct {
	CreateDidFee                *types.Coin `protobuf:"bytes,1,opt,name=create_did_fee,json=createDidFee,proto3" json:"create_did_fee,omitempty"`
	UpdateDidFee                *types.Coin `protobuf:"bytes,2,opt,name=update_did_fee,json=updateDidFee,proto3" json:"update_did_fee,omitempty"`
	DeactivateDidFee            *types.Coin `protobuf:"bytes,3,opt,name=deactivate_did_fee,json=deactivateDidFee,proto3" json:"deactivate_did_fee,omitempty"`
	CreateSchemaFee             *types.Coin `protobuf:"bytes,4,opt,name=create_schema_fee,json=createSchemaFee,proto3" json:"create_schema_fee,omitempty"`
	RegisterCredentialStatusFee *types.Coin `protobuf:"bytes,5,opt,name=register_credential_status_fee,json=registerCredentialStatusFee,proto3" json:"register_credential_status_fee,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14d65fa1d8b5cdf2, []int{0}
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

func (m *GenesisState) GetCreateDidFee() *types.Coin {
	if m != nil {
		return m.CreateDidFee
	}
	return nil
}

func (m *GenesisState) GetUpdateDidFee() *types.Coin {
	if m != nil {
		return m.UpdateDidFee
	}
	return nil
}

func (m *GenesisState) GetDeactivateDidFee() *types.Coin {
	if m != nil {
		return m.DeactivateDidFee
	}
	return nil
}

func (m *GenesisState) GetCreateSchemaFee() *types.Coin {
	if m != nil {
		return m.CreateSchemaFee
	}
	return nil
}

func (m *GenesisState) GetRegisterCredentialStatusFee() *types.Coin {
	if m != nil {
		return m.RegisterCredentialStatusFee
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "hypersignprotocol.hidnode.identityfee.GenesisState")
}

func init() { proto.RegisterFile("identityfee/v1/genesis.proto", fileDescriptor_14d65fa1d8b5cdf2) }

var fileDescriptor_14d65fa1d8b5cdf2 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd2, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc7, 0xf1, 0xc6, 0xaa, 0x43, 0x2c, 0xfe, 0x29, 0x0e, 0x5a, 0xe5, 0x10, 0x41, 0x70, 0xe9,
	0x1d, 0xd5, 0xcd, 0x45, 0xb0, 0xda, 0xee, 0x76, 0x11, 0x07, 0xc3, 0xe5, 0xee, 0x69, 0x72, 0xd0,
	0xe6, 0x42, 0xee, 0x69, 0xb1, 0x9b, 0x2f, 0xc1, 0x97, 0xe1, 0x4b, 0x71, 0xec, 0xe8, 0x28, 0xe9,
	0x1b, 0x91, 0xdc, 0xa5, 0x35, 0x5b, 0xb6, 0x84, 0xdc, 0xf7, 0x13, 0x12, 0x7e, 0xfe, 0xb9, 0x92,
	0x90, 0xa0, 0xc2, 0xc5, 0x18, 0x80, 0xcd, 0x7b, 0x2c, 0x82, 0x04, 0x8c, 0x32, 0x34, 0xcd, 0x34,
	0xea, 0xf6, 0x55, 0xbc, 0x48, 0x21, 0x33, 0x2a, 0x4a, 0xec, 0xbd, 0xd0, 0x13, 0x1a, 0x2b, 0x99,
	0x68, 0x09, 0xb4, 0xd2, 0x75, 0x8e, 0x23, 0x1d, 0x69, 0x7b, 0x82, 0x15, 0x57, 0x2e, 0xee, 0x10,
	0xa1, 0xcd, 0x54, 0x1b, 0x16, 0x72, 0x53, 0xd0, 0x21, 0x20, 0xef, 0x31, 0xa1, 0x55, 0xe2, 0x9e,
	0x5f, 0x7e, 0x34, 0xfd, 0xd6, 0xd0, 0xbd, 0x6e, 0x84, 0x1c, 0xa1, 0x7d, 0xef, 0xef, 0x8b, 0x0c,
	0x38, 0x42, 0x20, 0x95, 0x0c, 0xc6, 0x00, 0x27, 0xde, 0x85, 0x77, 0xbd, 0x77, 0x73, 0x4a, 0x9d,
	0x44, 0x0b, 0x89, 0x96, 0x12, 0xed, 0x6b, 0x95, 0x3c, 0xb7, 0x5c, 0xf0, 0xa8, 0xe4, 0x00, 0x2c,
	0x30, 0x4b, 0x65, 0x15, 0xd8, 0xaa, 0x05, 0x5c, 0x50, 0x02, 0x43, 0xbf, 0x2d, 0x81, 0x0b, 0x54,
	0xf3, 0x2a, 0xd2, 0xac, 0x43, 0x0e, 0xff, 0xa3, 0x12, 0x7a, 0xf2, 0x8f, 0xca, 0x4f, 0x31, 0x22,
	0x86, 0x29, 0xb7, 0xce, 0x76, 0x9d, 0x73, 0xe0, 0x9a, 0x91, 0x4d, 0x0a, 0xe6, 0xcd, 0x27, 0x19,
	0x44, 0xca, 0x20, 0x64, 0x81, 0xc8, 0xc0, 0xfe, 0x72, 0x3e, 0x09, 0x0c, 0x72, 0x9c, 0x19, 0x6b,
	0xee, 0xd4, 0x99, 0x67, 0x6b, 0xa0, 0xbf, 0xe9, 0x47, 0x36, 0x1f, 0x00, 0x3c, 0xbc, 0x7c, 0xe5,
	0xc4, 0xfb, 0xce, 0x89, 0xb7, 0xcc, 0x89, 0xf7, 0x9b, 0x13, 0xef, 0x73, 0x45, 0x1a, 0xcb, 0x15,
	0x69, 0xfc, 0xac, 0x48, 0xe3, 0xf5, 0x2e, 0x52, 0x18, 0xcf, 0x42, 0x2a, 0xf4, 0x94, 0x6d, 0x86,
	0xd0, 0x5d, 0x2f, 0x81, 0xc5, 0x4a, 0x76, 0x8b, 0x29, 0xb0, 0x77, 0x56, 0x1d, 0x11, 0x2e, 0x52,
	0x30, 0xe1, 0xae, 0x3d, 0x76, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xa5, 0x4e, 0xe3, 0x60,
	0x02, 0x00, 0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CreateDidFee.Equal(that1.CreateDidFee) {
		return false
	}
	if !this.UpdateDidFee.Equal(that1.UpdateDidFee) {
		return false
	}
	if !this.DeactivateDidFee.Equal(that1.DeactivateDidFee) {
		return false
	}
	if !this.CreateSchemaFee.Equal(that1.CreateSchemaFee) {
		return false
	}
	if !this.RegisterCredentialStatusFee.Equal(that1.RegisterCredentialStatusFee) {
		return false
	}
	return true
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
	if m.RegisterCredentialStatusFee != nil {
		{
			size, err := m.RegisterCredentialStatusFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.CreateSchemaFee != nil {
		{
			size, err := m.CreateSchemaFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DeactivateDidFee != nil {
		{
			size, err := m.DeactivateDidFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.UpdateDidFee != nil {
		{
			size, err := m.UpdateDidFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CreateDidFee != nil {
		{
			size, err := m.CreateDidFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
	if m.CreateDidFee != nil {
		l = m.CreateDidFee.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.UpdateDidFee != nil {
		l = m.UpdateDidFee.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.DeactivateDidFee != nil {
		l = m.DeactivateDidFee.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.CreateSchemaFee != nil {
		l = m.CreateSchemaFee.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.RegisterCredentialStatusFee != nil {
		l = m.RegisterCredentialStatusFee.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDidFee", wireType)
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
			if m.CreateDidFee == nil {
				m.CreateDidFee = &types.Coin{}
			}
			if err := m.CreateDidFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateDidFee", wireType)
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
			if m.UpdateDidFee == nil {
				m.UpdateDidFee = &types.Coin{}
			}
			if err := m.UpdateDidFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeactivateDidFee", wireType)
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
			if m.DeactivateDidFee == nil {
				m.DeactivateDidFee = &types.Coin{}
			}
			if err := m.DeactivateDidFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateSchemaFee", wireType)
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
			if m.CreateSchemaFee == nil {
				m.CreateSchemaFee = &types.Coin{}
			}
			if err := m.CreateSchemaFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterCredentialStatusFee", wireType)
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
			if m.RegisterCredentialStatusFee == nil {
				m.RegisterCredentialStatusFee = &types.Coin{}
			}
			if err := m.RegisterCredentialStatusFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
