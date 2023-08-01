// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/auction/auction.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Auction struct {
	Id           uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seller       string     `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	TokenId      uint64     `protobuf:"varint,3,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	StartPrice   types.Coin `protobuf:"bytes,4,opt,name=startPrice,proto3" json:"startPrice"`
	CurrentPrice types.Coin `protobuf:"bytes,5,opt,name=currentPrice,proto3" json:"currentPrice"`
	AuctionDate  string     `protobuf:"bytes,6,opt,name=auctionDate,proto3" json:"auctionDate,omitempty"`
	Deadline     string     `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	BidIds       []uint64   `protobuf:"varint,8,rep,packed,name=bidIds,proto3" json:"bidIds,omitempty"`
	IsCancel     bool       `protobuf:"varint,9,opt,name=isCancel,proto3" json:"isCancel,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_028c42bd7eb07429, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Auction) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *Auction) GetTokenId() uint64 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *Auction) GetStartPrice() types.Coin {
	if m != nil {
		return m.StartPrice
	}
	return types.Coin{}
}

func (m *Auction) GetCurrentPrice() types.Coin {
	if m != nil {
		return m.CurrentPrice
	}
	return types.Coin{}
}

func (m *Auction) GetAuctionDate() string {
	if m != nil {
		return m.AuctionDate
	}
	return ""
}

func (m *Auction) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Auction) GetBidIds() []uint64 {
	if m != nil {
		return m.BidIds
	}
	return nil
}

func (m *Auction) GetIsCancel() bool {
	if m != nil {
		return m.IsCancel
	}
	return false
}

func init() {
	proto.RegisterType((*Auction)(nil), "auction.auction.Auction")
}

func init() { proto.RegisterFile("auction/auction/auction.proto", fileDescriptor_028c42bd7eb07429) }

var fileDescriptor_028c42bd7eb07429 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0x34, 0x5f, 0xff, 0xb8, 0x9f, 0x40, 0xb2, 0x10, 0x32, 0x95, 0x30, 0x11, 0x53,
	0x26, 0x47, 0x85, 0x07, 0x40, 0xb4, 0x0c, 0x74, 0xab, 0x22, 0x26, 0x36, 0xc7, 0xb6, 0x8a, 0x45,
	0x1a, 0x57, 0xb1, 0x53, 0xc1, 0x5b, 0xf0, 0x4e, 0x2c, 0x1d, 0x3b, 0x32, 0x21, 0xd4, 0xbe, 0x08,
	0x4a, 0xea, 0x46, 0x85, 0x89, 0xe9, 0xde, 0x5f, 0xee, 0x3d, 0x47, 0x27, 0xbe, 0xf0, 0x9c, 0x95,
	0xdc, 0x2a, 0x9d, 0xc7, 0xbf, 0x2a, 0x5d, 0x14, 0xda, 0x6a, 0x74, 0xbc, 0x47, 0x57, 0x07, 0x27,
	0x33, 0x3d, 0xd3, 0xf5, 0x2c, 0xae, 0xba, 0xdd, 0xda, 0x80, 0x70, 0x6d, 0xe6, 0xda, 0xc4, 0x29,
	0x33, 0x32, 0x5e, 0x0e, 0x53, 0x69, 0xd9, 0x30, 0xe6, 0x5a, 0x39, 0x9b, 0xcb, 0x77, 0x1f, 0x76,
	0x6e, 0x77, 0x0e, 0xe8, 0x08, 0xfa, 0x4a, 0x60, 0x10, 0x82, 0x28, 0x48, 0x7c, 0x25, 0xd0, 0x29,
	0x6c, 0x1b, 0x99, 0x65, 0xb2, 0xc0, 0x7e, 0x08, 0xa2, 0x5e, 0xe2, 0x08, 0x61, 0xd8, 0xb1, 0xfa,
	0x59, 0xe6, 0x13, 0x81, 0x5b, 0xf5, 0xf2, 0x1e, 0xd1, 0x0d, 0x84, 0xc6, 0xb2, 0xc2, 0x4e, 0x0b,
	0xc5, 0x25, 0x0e, 0x42, 0x10, 0xf5, 0xaf, 0xce, 0xe8, 0x2e, 0x02, 0xad, 0x22, 0x50, 0x17, 0x81,
	0x8e, 0xb5, 0xca, 0x47, 0xc1, 0xea, 0xf3, 0xc2, 0x4b, 0x0e, 0x24, 0x68, 0x0c, 0xff, 0xf3, 0xb2,
	0x28, 0x64, 0xee, 0x2c, 0xfe, 0xfd, 0xcd, 0xe2, 0x87, 0x08, 0x85, 0xb0, 0xef, 0x1e, 0xe5, 0x8e,
	0x59, 0x89, 0xdb, 0x75, 0xf8, 0xc3, 0x4f, 0x68, 0x00, 0xbb, 0x42, 0x32, 0x91, 0xa9, 0x5c, 0xe2,
	0x4e, 0x3d, 0x6e, 0xb8, 0xfa, 0xeb, 0x54, 0x89, 0x89, 0x30, 0xb8, 0x1b, 0xb6, 0xa2, 0x20, 0x71,
	0x54, 0x69, 0x94, 0x19, 0xb3, 0x9c, 0xcb, 0x0c, 0xf7, 0x42, 0x10, 0x75, 0x93, 0x86, 0x47, 0xf7,
	0xab, 0x0d, 0x01, 0xeb, 0x0d, 0x01, 0x5f, 0x1b, 0x02, 0xde, 0xb6, 0xc4, 0x5b, 0x6f, 0x89, 0xf7,
	0xb1, 0x25, 0xde, 0x23, 0x9d, 0x29, 0xfb, 0x54, 0xa6, 0x94, 0xeb, 0x79, 0x3c, 0x65, 0x4b, 0x99,
	0x3d, 0xb0, 0x94, 0xf1, 0xb2, 0x39, 0xea, 0x4b, 0xd3, 0xd9, 0xd7, 0x85, 0x34, 0x69, 0xbb, 0x3e,
	0xcb, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xb9, 0x9f, 0x04, 0xfe, 0x01, 0x00, 0x00,
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsCancel {
		i--
		if m.IsCancel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.BidIds) > 0 {
		dAtA2 := make([]byte, len(m.BidIds)*10)
		var j1 int
		for _, num := range m.BidIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAuction(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.AuctionDate) > 0 {
		i -= len(m.AuctionDate)
		copy(dAtA[i:], m.AuctionDate)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.AuctionDate)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.CurrentPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.StartPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.TokenId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.TokenId != 0 {
		n += 1 + sovAuction(uint64(m.TokenId))
	}
	l = m.StartPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.CurrentPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = len(m.AuctionDate)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if len(m.BidIds) > 0 {
		l = 0
		for _, e := range m.BidIds {
			l += sovAuction(uint64(e))
		}
		n += 1 + sovAuction(uint64(l)) + l
	}
	if m.IsCancel {
		n += 2
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuction
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BidIds = append(m.BidIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuction
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAuction
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAuction
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BidIds) == 0 {
					m.BidIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAuction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BidIds = append(m.BidIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BidIds", wireType)
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCancel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsCancel = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
