// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/maichain/eth-indexer/indexer/pb/block_header.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		github.com/maichain/eth-indexer/indexer/pb/block_header.proto
		github.com/maichain/eth-indexer/indexer/pb/transaction.proto

	It has these top-level messages:
		BlockHeader
		Transaction
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id,omitempty" gorm:"primary_key" deepequal:"-"`
	ParentHash  string `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty" gorm:"column:parent_hash;size:255"`
	UncleHash   string `protobuf:"bytes,3,opt,name=uncle_hash,json=uncleHash,proto3" json:"uncle_hash,omitempty" gorm:"column:uncle_hash;size:255"`
	Coinbase    string `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty" gorm:"column:coinbase;size:255"`
	Root        string `protobuf:"bytes,5,opt,name=root,proto3" json:"root,omitempty" gorm:"column:root;size:255"`
	TxHash      string `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" gorm:"column:tx_hash;size:255"`
	ReceiptHash string `protobuf:"bytes,7,opt,name=receipt_hash,json=receiptHash,proto3" json:"receipt_hash,omitempty" gorm:"column:receipt_hash;size:255"`
	Bloom       string `protobuf:"bytes,8,opt,name=bloom,proto3" json:"bloom,omitempty" gorm:"column:bloom"`
	Difficulty  int64  `protobuf:"varint,9,opt,name=difficulty,proto3" json:"difficulty,omitempty" gorm:"column:difficulty"`
	Number      int64  `protobuf:"varint,10,opt,name=number,proto3" json:"number,omitempty" gorm:"column:number"`
	GasLimit    int64  `protobuf:"varint,11,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty" gorm:"column:gas_limit"`
	GasUsed     int64  `protobuf:"varint,12,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty" gorm:"column:gas_used"`
	ExtraData   string `protobuf:"bytes,13,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty" gorm:"column:extra_data"`
	Nonce       string `protobuf:"bytes,14,opt,name=nonce,proto3" json:"nonce,omitempty" gorm:"column:nonce"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlockHeader, []int{0} }

func init() {
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
}
func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.ID))
	}
	if len(m.ParentHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ParentHash)))
		i += copy(dAtA[i:], m.ParentHash)
	}
	if len(m.UncleHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.UncleHash)))
		i += copy(dAtA[i:], m.UncleHash)
	}
	if len(m.Coinbase) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Coinbase)))
		i += copy(dAtA[i:], m.Coinbase)
	}
	if len(m.Root) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Root)))
		i += copy(dAtA[i:], m.Root)
	}
	if len(m.TxHash) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if len(m.ReceiptHash) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ReceiptHash)))
		i += copy(dAtA[i:], m.ReceiptHash)
	}
	if len(m.Bloom) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Bloom)))
		i += copy(dAtA[i:], m.Bloom)
	}
	if m.Difficulty != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.Difficulty))
	}
	if m.Number != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.Number))
	}
	if m.GasLimit != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(m.GasUsed))
	}
	if len(m.ExtraData) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.ExtraData)))
		i += copy(dAtA[i:], m.ExtraData)
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintBlockHeader(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	return i, nil
}

func encodeVarintBlockHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockHeader) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovBlockHeader(uint64(m.ID))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.UncleHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Coinbase)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.ReceiptHash)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	if m.Difficulty != 0 {
		n += 1 + sovBlockHeader(uint64(m.Difficulty))
	}
	if m.Number != 0 {
		n += 1 + sovBlockHeader(uint64(m.Number))
	}
	if m.GasLimit != 0 {
		n += 1 + sovBlockHeader(uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		n += 1 + sovBlockHeader(uint64(m.GasUsed))
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovBlockHeader(uint64(l))
	}
	return n
}

func sovBlockHeader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlockHeader(x uint64) (n int) {
	return sovBlockHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BlockHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlockHeader{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`ParentHash:` + fmt.Sprintf("%v", this.ParentHash) + `,`,
		`UncleHash:` + fmt.Sprintf("%v", this.UncleHash) + `,`,
		`Coinbase:` + fmt.Sprintf("%v", this.Coinbase) + `,`,
		`Root:` + fmt.Sprintf("%v", this.Root) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`ReceiptHash:` + fmt.Sprintf("%v", this.ReceiptHash) + `,`,
		`Bloom:` + fmt.Sprintf("%v", this.Bloom) + `,`,
		`Difficulty:` + fmt.Sprintf("%v", this.Difficulty) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasUsed:` + fmt.Sprintf("%v", this.GasUsed) + `,`,
		`ExtraData:` + fmt.Sprintf("%v", this.ExtraData) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlockHeader(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockHeader
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncleHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UncleHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coinbase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coinbase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Difficulty", wireType)
			}
			m.Difficulty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Difficulty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockHeader
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
func skipBlockHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockHeader
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
					return 0, ErrIntOverflowBlockHeader
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
					return 0, ErrIntOverflowBlockHeader
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
				return 0, ErrInvalidLengthBlockHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlockHeader
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
				next, err := skipBlockHeader(dAtA[start:])
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
	ErrInvalidLengthBlockHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockHeader   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/maichain/eth-indexer/indexer/pb/block_header.proto", fileDescriptorBlockHeader)
}

var fileDescriptorBlockHeader = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd4, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0x07, 0xf0, 0xa5, 0xdb, 0xba, 0xd6, 0xdd, 0xbb, 0x17, 0xc2, 0x18, 0x11, 0x48, 0x49, 0x31,
	0x97, 0x09, 0x6d, 0x2b, 0x0c, 0x26, 0x44, 0xd1, 0x04, 0x54, 0x45, 0xda, 0x04, 0x12, 0x28, 0xb0,
	0x1b, 0x52, 0x71, 0x12, 0x37, 0xb1, 0x96, 0xc4, 0x21, 0x71, 0xa4, 0x96, 0x13, 0x1f, 0x83, 0x8f,
	0xc3, 0x71, 0x47, 0x8e, 0x9c, 0x22, 0xd6, 0xdd, 0x72, 0xe4, 0xcc, 0x01, 0xe5, 0x49, 0xbb, 0xc4,
	0x42, 0xda, 0xa9, 0xf6, 0xdf, 0x8f, 0x7f, 0xae, 0xf3, 0x44, 0x41, 0x87, 0x2e, 0x13, 0x5e, 0x6a,
	0xed, 0xd9, 0x3c, 0xe8, 0x05, 0x84, 0xd9, 0x1e, 0x61, 0x61, 0x8f, 0x0a, 0x6f, 0x97, 0x85, 0x0e,
	0x9d, 0xd0, 0xb8, 0xb7, 0xf8, 0x8d, 0xac, 0x9e, 0xe5, 0x73, 0xfb, 0x74, 0xe4, 0x51, 0xe2, 0xd0,
	0x78, 0x2f, 0x8a, 0xb9, 0xe0, 0x6a, 0x23, 0xb2, 0x6e, 0xef, 0xd6, 0x08, 0x97, 0xbb, 0xbc, 0x07,
	0x4b, 0x56, 0x3a, 0x86, 0x19, 0x4c, 0x60, 0x54, 0x6e, 0xc1, 0x7f, 0x5a, 0xa8, 0x33, 0x28, 0xa4,
	0x23, 0x80, 0xd4, 0x21, 0x6a, 0x1c, 0x0f, 0x35, 0xa5, 0xab, 0x6c, 0xaf, 0x0c, 0x1e, 0xe7, 0x99,
	0xb1, 0xce, 0x9c, 0x1d, 0x1e, 0x30, 0x41, 0x83, 0x48, 0x4c, 0x7f, 0x67, 0x46, 0xd7, 0xe5, 0x71,
	0xd0, 0xc7, 0x51, 0xcc, 0x02, 0x12, 0x4f, 0x47, 0xa7, 0x74, 0x8a, 0xbb, 0x0e, 0xa5, 0x11, 0xfd,
	0x9c, 0x12, 0xbf, 0x8f, 0x77, 0xb1, 0xd9, 0x38, 0x1e, 0xaa, 0x9f, 0x50, 0x27, 0x22, 0x31, 0x0d,
	0xc5, 0xc8, 0x23, 0x89, 0xa7, 0x35, 0xba, 0xca, 0x76, 0x7b, 0xf0, 0x3c, 0xcf, 0x8c, 0x9b, 0xb5,
	0x58, 0x72, 0x71, 0xe9, 0xda, 0xdc, 0x4f, 0x83, 0xb0, 0x5f, 0xab, 0x7a, 0x96, 0xb0, 0x2f, 0xb4,
	0xbf, 0x7f, 0x70, 0x80, 0x4d, 0x54, 0xc6, 0x47, 0x24, 0xf1, 0xd4, 0x8f, 0x08, 0xa5, 0xa1, 0xed,
	0xd3, 0xf2, 0x80, 0x65, 0x38, 0xe0, 0x30, 0xcf, 0x8c, 0xcd, 0x2a, 0x95, 0xfc, 0xbb, 0x92, 0x5f,
	0x15, 0xd5, 0xf8, 0x36, 0xa4, 0xa0, 0x9f, 0xa0, 0x96, 0xcd, 0x59, 0x68, 0x91, 0x84, 0x6a, 0x2b,
	0x60, 0x3f, 0xcd, 0x33, 0x43, 0x5d, 0x64, 0x92, 0x6c, 0x48, 0xf2, 0xa2, 0xa4, 0xe6, 0x5e, 0x52,
	0xea, 0x2b, 0xb4, 0x12, 0x73, 0x2e, 0xb4, 0x55, 0x20, 0x1f, 0xe6, 0x99, 0xb1, 0x51, 0xcc, 0x25,
	0xee, 0x8e, 0xc4, 0x15, 0xcb, 0x35, 0x0a, 0xb6, 0xab, 0xef, 0xd0, 0x9a, 0x98, 0x94, 0x17, 0x6f,
	0x82, 0xf4, 0x24, 0xcf, 0x8c, 0xeb, 0xf3, 0x48, 0xc2, 0x74, 0x09, 0x9b, 0x57, 0xd4, 0xbc, 0xa6,
	0x98, 0xc0, 0x7d, 0x1d, 0xb4, 0x1e, 0x53, 0x9b, 0xb2, 0x68, 0xde, 0xb0, 0x35, 0x60, 0x5f, 0xe6,
	0x99, 0xb1, 0x55, 0xcf, 0x25, 0xfb, 0x9e, 0xfc, 0x47, 0x6b, 0x65, 0xb5, 0x03, 0x3a, 0xf3, 0x1c,
	0x4e, 0x79, 0x81, 0x56, 0x2d, 0x9f, 0xf3, 0x40, 0x6b, 0x01, 0x7f, 0x3f, 0xcf, 0x8c, 0xff, 0x21,
	0x90, 0xdc, 0x1b, 0x92, 0x0b, 0xeb, 0xd8, 0x2c, 0x37, 0xaa, 0x1f, 0x10, 0x72, 0xd8, 0x78, 0xcc,
	0xec, 0xd4, 0x17, 0x53, 0xad, 0xdd, 0x55, 0xb6, 0x97, 0xe1, 0x2d, 0xdd, 0xac, 0x52, 0xc9, 0xd2,
	0x24, 0xab, 0x2a, 0xc2, 0x66, 0xcd, 0x51, 0x87, 0xa8, 0x19, 0xa6, 0x81, 0x45, 0x63, 0x0d, 0x81,
	0xb8, 0x93, 0x67, 0xc6, 0xb5, 0x32, 0x91, 0xb4, 0x4d, 0x49, 0x2b, 0x0b, 0xb0, 0x39, 0xdf, 0xab,
	0xbe, 0x45, 0x6d, 0x97, 0x24, 0x23, 0x9f, 0x05, 0x4c, 0x68, 0x1d, 0x80, 0xf6, 0xf3, 0xe2, 0x3a,
	0x8b, 0x50, 0xb2, 0x6e, 0x49, 0xd6, 0x65, 0x0d, 0x36, 0x5b, 0x2e, 0x49, 0xde, 0x14, 0x43, 0xf5,
	0x35, 0x2a, 0xc6, 0xa3, 0x34, 0xa1, 0x8e, 0xb6, 0x0e, 0xde, 0x83, 0xe2, 0x25, 0x5c, 0x64, 0x12,
	0xb7, 0xf5, 0x0f, 0x57, 0x94, 0x60, 0x73, 0xcd, 0x25, 0xc9, 0x49, 0x42, 0x1d, 0xf5, 0x3d, 0x42,
	0x74, 0x22, 0x62, 0x32, 0x72, 0x88, 0x20, 0xda, 0x7f, 0xd0, 0x00, 0x78, 0x72, 0x55, 0x7a, 0xc5,
	0x93, 0xab, 0x8a, 0xb0, 0xd9, 0x86, 0xc9, 0x90, 0x08, 0x52, 0x34, 0x34, 0xe4, 0xa1, 0x4d, 0xb5,
	0x8d, 0xaa, 0xa1, 0x10, 0x5c, 0xd1, 0x50, 0x58, 0xc7, 0x66, 0xb9, 0x71, 0xd0, 0x3d, 0x3b, 0xd7,
	0x97, 0x7e, 0x9e, 0xeb, 0x4b, 0x5f, 0x67, 0xba, 0x72, 0x36, 0xd3, 0x95, 0x1f, 0x33, 0x5d, 0xf9,
	0x35, 0xd3, 0x95, 0x6f, 0x17, 0xfa, 0xd2, 0xf7, 0x0b, 0x5d, 0xb1, 0x9a, 0xf0, 0x9d, 0x7a, 0xf4,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xb7, 0x81, 0x3f, 0x1b, 0x05, 0x00, 0x00,
}
