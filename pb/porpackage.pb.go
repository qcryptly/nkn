// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/porpackage.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PorPackage struct {
	VoteForHeight uint32    `protobuf:"varint,1,opt,name=VoteForHeight,proto3" json:"VoteForHeight,omitempty"`
	Owner         []byte    `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	BlockHash     []byte    `protobuf:"bytes,3,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	TxHash        []byte    `protobuf:"bytes,4,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
	SigHash       []byte    `protobuf:"bytes,5,opt,name=SigHash,proto3" json:"SigHash,omitempty"`
	SigChain      *SigChain `protobuf:"bytes,6,opt,name=SigChain" json:"SigChain,omitempty"`
}

func (m *PorPackage) Reset()      { *m = PorPackage{} }
func (*PorPackage) ProtoMessage() {}
func (*PorPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_porpackage_539ed5359ba64c95, []int{0}
}
func (m *PorPackage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PorPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PorPackage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PorPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PorPackage.Merge(dst, src)
}
func (m *PorPackage) XXX_Size() int {
	return m.Size()
}
func (m *PorPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_PorPackage.DiscardUnknown(m)
}

var xxx_messageInfo_PorPackage proto.InternalMessageInfo

func (m *PorPackage) GetVoteForHeight() uint32 {
	if m != nil {
		return m.VoteForHeight
	}
	return 0
}

func (m *PorPackage) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PorPackage) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *PorPackage) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *PorPackage) GetSigHash() []byte {
	if m != nil {
		return m.SigHash
	}
	return nil
}

func (m *PorPackage) GetSigChain() *SigChain {
	if m != nil {
		return m.SigChain
	}
	return nil
}

func init() {
	proto.RegisterType((*PorPackage)(nil), "pb.PorPackage")
}
func (this *PorPackage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PorPackage)
	if !ok {
		that2, ok := that.(PorPackage)
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
	if this.VoteForHeight != that1.VoteForHeight {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if !bytes.Equal(this.BlockHash, that1.BlockHash) {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	if !bytes.Equal(this.SigHash, that1.SigHash) {
		return false
	}
	if !this.SigChain.Equal(that1.SigChain) {
		return false
	}
	return true
}
func (this *PorPackage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&pb.PorPackage{")
	s = append(s, "VoteForHeight: "+fmt.Sprintf("%#v", this.VoteForHeight)+",\n")
	s = append(s, "Owner: "+fmt.Sprintf("%#v", this.Owner)+",\n")
	s = append(s, "BlockHash: "+fmt.Sprintf("%#v", this.BlockHash)+",\n")
	s = append(s, "TxHash: "+fmt.Sprintf("%#v", this.TxHash)+",\n")
	s = append(s, "SigHash: "+fmt.Sprintf("%#v", this.SigHash)+",\n")
	if this.SigChain != nil {
		s = append(s, "SigChain: "+fmt.Sprintf("%#v", this.SigChain)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPorpackage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PorPackage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PorPackage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.VoteForHeight != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(m.VoteForHeight))
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.BlockHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(len(m.BlockHash)))
		i += copy(dAtA[i:], m.BlockHash)
	}
	if len(m.TxHash) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if len(m.SigHash) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(len(m.SigHash)))
		i += copy(dAtA[i:], m.SigHash)
	}
	if m.SigChain != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPorpackage(dAtA, i, uint64(m.SigChain.Size()))
		n1, err := m.SigChain.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintPorpackage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPorPackage(r randyPorpackage, easy bool) *PorPackage {
	this := &PorPackage{}
	this.VoteForHeight = uint32(r.Uint32())
	v1 := r.Intn(100)
	this.Owner = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Owner[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(100)
	this.BlockHash = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.BlockHash[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.TxHash = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.TxHash[i] = byte(r.Intn(256))
	}
	v4 := r.Intn(100)
	this.SigHash = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.SigHash[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		this.SigChain = NewPopulatedSigChain(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPorpackage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePorpackage(r randyPorpackage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPorpackage(r randyPorpackage) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RunePorpackage(r)
	}
	return string(tmps)
}
func randUnrecognizedPorpackage(r randyPorpackage, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPorpackage(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPorpackage(dAtA []byte, r randyPorpackage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePorpackage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePorpackage(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PorPackage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteForHeight != 0 {
		n += 1 + sovPorpackage(uint64(m.VoteForHeight))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovPorpackage(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovPorpackage(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovPorpackage(uint64(l))
	}
	l = len(m.SigHash)
	if l > 0 {
		n += 1 + l + sovPorpackage(uint64(l))
	}
	if m.SigChain != nil {
		l = m.SigChain.Size()
		n += 1 + l + sovPorpackage(uint64(l))
	}
	return n
}

func sovPorpackage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPorpackage(x uint64) (n int) {
	return sovPorpackage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PorPackage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PorPackage{`,
		`VoteForHeight:` + fmt.Sprintf("%v", this.VoteForHeight) + `,`,
		`Owner:` + fmt.Sprintf("%v", this.Owner) + `,`,
		`BlockHash:` + fmt.Sprintf("%v", this.BlockHash) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`SigHash:` + fmt.Sprintf("%v", this.SigHash) + `,`,
		`SigChain:` + strings.Replace(fmt.Sprintf("%v", this.SigChain), "SigChain", "SigChain", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPorpackage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PorPackage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPorpackage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PorPackage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PorPackage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteForHeight", wireType)
			}
			m.VoteForHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteForHeight |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPorpackage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPorpackage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPorpackage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPorpackage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigHash = append(m.SigHash[:0], dAtA[iNdEx:postIndex]...)
			if m.SigHash == nil {
				m.SigHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPorpackage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SigChain == nil {
				m.SigChain = &SigChain{}
			}
			if err := m.SigChain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPorpackage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPorpackage
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
func skipPorpackage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPorpackage
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
					return 0, ErrIntOverflowPorpackage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPorpackage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPorpackage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPorpackage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPorpackage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPorpackage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPorpackage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/porpackage.proto", fileDescriptor_porpackage_539ed5359ba64c95) }

var fileDescriptor_porpackage_539ed5359ba64c95 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0x80, 0xef, 0x55, 0x5b, 0xf5, 0x6c, 0x07, 0x4f, 0x91, 0x50, 0xe4, 0x11, 0xc4, 0x21, 0x8b,
	0x09, 0xe8, 0xea, 0x54, 0x41, 0xba, 0x59, 0x52, 0x71, 0xcf, 0x85, 0x78, 0x39, 0xaa, 0xbd, 0xe3,
	0x9a, 0xa2, 0xa3, 0x3f, 0xc1, 0x9f, 0xe1, 0x4f, 0x70, 0x76, 0x72, 0xcc, 0xd8, 0xd1, 0x5c, 0x16,
	0xc7, 0x8e, 0x8e, 0xe2, 0xc5, 0x2a, 0x6e, 0xf7, 0x7d, 0x1f, 0xf7, 0x78, 0x3c, 0xba, 0xab, 0x79,
	0xa4, 0x95, 0xd1, 0x49, 0x3a, 0x49, 0x44, 0x16, 0x6a, 0xa3, 0x0a, 0xc5, 0x5a, 0x9a, 0xf7, 0x8f,
	0x85, 0x2c, 0xf2, 0x39, 0x0f, 0x53, 0x75, 0x17, 0x09, 0x25, 0x54, 0xe4, 0x12, 0x9f, 0xdf, 0x38,
	0x72, 0xe0, 0x5e, 0xcd, 0x97, 0xfe, 0x8e, 0xe6, 0xd1, 0x4c, 0x8a, 0x34, 0x4f, 0xe4, 0xb4, 0x51,
	0x87, 0xaf, 0x40, 0xe9, 0x48, 0x99, 0x51, 0x33, 0x9a, 0x1d, 0xd1, 0xde, 0xb5, 0x2a, 0xb2, 0x0b,
	0x65, 0x86, 0x99, 0x14, 0x79, 0xe1, 0x81, 0x0f, 0x41, 0x2f, 0xfe, 0x2f, 0xd9, 0x1e, 0x6d, 0x5f,
	0xde, 0x4f, 0x33, 0xe3, 0xb5, 0x7c, 0x08, 0xba, 0x71, 0x03, 0xec, 0x80, 0x6e, 0x0d, 0x6e, 0x55,
	0x3a, 0x19, 0x26, 0xb3, 0xdc, 0x5b, 0x73, 0xe5, 0x4f, 0xb0, 0x7d, 0xda, 0xb9, 0x7a, 0x70, 0x69,
	0xdd, 0xa5, 0x1f, 0x62, 0x1e, 0xdd, 0x18, 0x4b, 0xe1, 0x42, 0xdb, 0x85, 0x15, 0xb2, 0x80, 0x6e,
	0x8e, 0xa5, 0x38, 0xff, 0x5e, 0xd6, 0xeb, 0xf8, 0x10, 0x6c, 0x9f, 0x74, 0x43, 0xcd, 0xc3, 0x95,
	0x8b, 0x7f, 0xeb, 0xe0, 0xac, 0xac, 0x90, 0x2c, 0x2a, 0x24, 0xcb, 0x0a, 0xe1, 0xb3, 0x42, 0x78,
	0xb4, 0x08, 0xcf, 0x16, 0xe1, 0xc5, 0x22, 0xbc, 0x59, 0x84, 0xd2, 0x22, 0xbc, 0x5b, 0x84, 0x0f,
	0x8b, 0x64, 0x69, 0x11, 0x9e, 0x6a, 0x24, 0x65, 0x8d, 0x64, 0x51, 0x23, 0xe1, 0x1d, 0x77, 0x89,
	0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x26, 0x19, 0xd9, 0x66, 0x01, 0x00, 0x00,
}