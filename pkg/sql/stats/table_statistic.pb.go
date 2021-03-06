// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/stats/table_statistic.proto

package stats

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_cockroachdb_cockroach_pkg_sql_sqlbase "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A TableStatistic object holds a statistic for a particular column or group
// of columns. It mirrors the structure of the system.table_statistics table.
// It is also used as the format in which table statistics are serialized
// in a backup.
type TableStatistic struct {
	// The ID of the table.
	TableID github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID `protobuf:"varint,1,opt,name=table_id,json=tableId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/sql/sqlbase.ID" json:"table_id,omitempty"`
	// The ID for this statistic.  It need not be globally unique,
	// but must be unique for this table.
	StatisticID uint64 `protobuf:"varint,2,opt,name=statistic_id,json=statisticId,proto3" json:"statistic_id,omitempty"`
	// Optional user-defined name for the statistic.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The column ID(s) for which this statistic is generated.
	ColumnIDs []github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID `protobuf:"varint,4,rep,packed,name=column_ids,json=columnIds,proto3,casttype=github.com/cockroachdb/cockroach/pkg/sql/sqlbase.ColumnID" json:"column_ids,omitempty"`
	// The time at which the statistic was created.
	CreatedAt time.Time `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	// The total number of rows in the table.
	RowCount uint64 `protobuf:"varint,6,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	// The estimated number of distinct values of the columns in ColumnIDs.
	DistinctCount uint64 `protobuf:"varint,7,opt,name=distinct_count,json=distinctCount,proto3" json:"distinct_count,omitempty"`
	// The number of rows that have a NULL in any of the columns in ColumnIDs.
	NullCount uint64 `protobuf:"varint,8,opt,name=null_count,json=nullCount,proto3" json:"null_count,omitempty"`
	// Histogram (if available)
	Histogram *HistogramData `protobuf:"bytes,9,opt,name=histogram,proto3" json:"histogram,omitempty"`
}

func (m *TableStatistic) Reset()         { *m = TableStatistic{} }
func (m *TableStatistic) String() string { return proto.CompactTextString(m) }
func (*TableStatistic) ProtoMessage()    {}
func (*TableStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_statistic_fb3b5153f9b658f1, []int{0}
}
func (m *TableStatistic) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *TableStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableStatistic.Merge(dst, src)
}
func (m *TableStatistic) XXX_Size() int {
	return m.Size()
}
func (m *TableStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_TableStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_TableStatistic proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TableStatistic)(nil), "cockroach.sql.stats.TableStatistic")
}
func (m *TableStatistic) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableStatistic) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TableID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.TableID))
	}
	if m.StatisticID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.StatisticID))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.ColumnIDs) > 0 {
		dAtA2 := make([]byte, len(m.ColumnIDs)*10)
		var j1 int
		for _, num := range m.ColumnIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintTableStatistic(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.RowCount != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.RowCount))
	}
	if m.DistinctCount != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.DistinctCount))
	}
	if m.NullCount != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.NullCount))
	}
	if m.Histogram != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintTableStatistic(dAtA, i, uint64(m.Histogram.Size()))
		n4, err := m.Histogram.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintTableStatistic(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TableStatistic) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TableID != 0 {
		n += 1 + sovTableStatistic(uint64(m.TableID))
	}
	if m.StatisticID != 0 {
		n += 1 + sovTableStatistic(uint64(m.StatisticID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTableStatistic(uint64(l))
	}
	if len(m.ColumnIDs) > 0 {
		l = 0
		for _, e := range m.ColumnIDs {
			l += sovTableStatistic(uint64(e))
		}
		n += 1 + sovTableStatistic(uint64(l)) + l
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovTableStatistic(uint64(l))
	if m.RowCount != 0 {
		n += 1 + sovTableStatistic(uint64(m.RowCount))
	}
	if m.DistinctCount != 0 {
		n += 1 + sovTableStatistic(uint64(m.DistinctCount))
	}
	if m.NullCount != 0 {
		n += 1 + sovTableStatistic(uint64(m.NullCount))
	}
	if m.Histogram != nil {
		l = m.Histogram.Size()
		n += 1 + l + sovTableStatistic(uint64(l))
	}
	return n
}

func sovTableStatistic(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTableStatistic(x uint64) (n int) {
	return sovTableStatistic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableStatistic) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTableStatistic
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
			return fmt.Errorf("proto: TableStatistic: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableStatistic: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableID", wireType)
			}
			m.TableID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableID |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatisticID", wireType)
			}
			m.StatisticID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatisticID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
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
				return ErrInvalidLengthTableStatistic
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTableStatistic
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ColumnIDs = append(m.ColumnIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTableStatistic
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTableStatistic
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ColumnIDs) == 0 {
					m.ColumnIDs = make([]github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTableStatistic
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ColumnID(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ColumnIDs = append(m.ColumnIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnIDs", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
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
				return ErrInvalidLengthTableStatistic
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowCount", wireType)
			}
			m.RowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RowCount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistinctCount", wireType)
			}
			m.DistinctCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistinctCount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullCount", wireType)
			}
			m.NullCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NullCount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Histogram", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTableStatistic
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
				return ErrInvalidLengthTableStatistic
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Histogram == nil {
				m.Histogram = &HistogramData{}
			}
			if err := m.Histogram.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTableStatistic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTableStatistic
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
func skipTableStatistic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTableStatistic
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
					return 0, ErrIntOverflowTableStatistic
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
					return 0, ErrIntOverflowTableStatistic
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
				return 0, ErrInvalidLengthTableStatistic
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTableStatistic
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
				next, err := skipTableStatistic(dAtA[start:])
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
	ErrInvalidLengthTableStatistic = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTableStatistic   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/stats/table_statistic.proto", fileDescriptor_table_statistic_fb3b5153f9b658f1)
}

var fileDescriptor_table_statistic_fb3b5153f9b658f1 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x8e, 0x9b, 0x40,
	0x10, 0xc6, 0xd9, 0x9c, 0xef, 0x6c, 0xd6, 0xf1, 0x45, 0xda, 0xa4, 0x20, 0x8e, 0xc2, 0xa2, 0x93,
	0xa2, 0x50, 0x2d, 0xd2, 0x5d, 0x95, 0x2e, 0xc1, 0x14, 0x21, 0x25, 0xb9, 0x2a, 0x52, 0x64, 0x2d,
	0x0b, 0xc1, 0xe8, 0x80, 0xb5, 0xd9, 0x45, 0xf7, 0x12, 0x29, 0xee, 0xb1, 0x5c, 0x5e, 0x79, 0x15,
	0x49, 0xf0, 0x5b, 0xb8, 0x8a, 0x58, 0xfe, 0xb8, 0x49, 0x93, 0x6e, 0x76, 0xe6, 0x37, 0x33, 0xdf,
	0x37, 0x00, 0xb1, 0xd8, 0x65, 0x8e, 0x90, 0x54, 0x0a, 0x47, 0xd2, 0x30, 0x8b, 0xd7, 0x6d, 0x9c,
	0x0a, 0x99, 0x32, 0xb2, 0x2d, 0xb9, 0xe4, 0xe8, 0x25, 0xe3, 0xec, 0xae, 0xe4, 0x94, 0x6d, 0x88,
	0xd8, 0x65, 0x44, 0xa1, 0xcb, 0x57, 0x09, 0x4f, 0xb8, 0xaa, 0x3b, 0x6d, 0xd4, 0xa1, 0x4b, 0x9c,
	0x70, 0x9e, 0x64, 0xb1, 0xa3, 0x5e, 0x61, 0xf5, 0xc3, 0x91, 0x69, 0x1e, 0x0b, 0x49, 0xf3, 0x6d,
	0x0f, 0xbc, 0x3e, 0x2d, 0xdb, 0xa4, 0x42, 0xf2, 0xa4, 0xa4, 0x79, 0x57, 0xba, 0xfa, 0x39, 0x81,
	0x97, 0xb7, 0xad, 0x80, 0xaf, 0xc3, 0x7e, 0xf4, 0x1d, 0xce, 0x3a, 0x49, 0x69, 0x64, 0x00, 0x0b,
	0xd8, 0x0b, 0xd7, 0x6d, 0x6a, 0x3c, 0x55, 0x94, 0xef, 0x1d, 0x6b, 0x7c, 0x93, 0xa4, 0x72, 0x53,
	0x85, 0x84, 0xf1, 0xdc, 0x19, 0x55, 0x46, 0xe1, 0x29, 0x76, 0xb6, 0x77, 0x89, 0xa3, 0x76, 0xee,
	0xb2, 0x90, 0x8a, 0x98, 0xf8, 0x5e, 0x30, 0x55, 0x33, 0xfd, 0x08, 0x5d, 0xc3, 0xe7, 0xa3, 0xd7,
	0x76, 0xc5, 0x33, 0x0b, 0xd8, 0x13, 0xf7, 0x45, 0x53, 0xe3, 0xf9, 0xa8, 0xc1, 0xf7, 0x82, 0xf9,
	0x08, 0xf9, 0x11, 0x42, 0x70, 0x52, 0xd0, 0x3c, 0x36, 0xce, 0x2c, 0x60, 0xeb, 0x81, 0x8a, 0x51,
	0x0a, 0x21, 0xe3, 0x59, 0x95, 0x17, 0xeb, 0x34, 0x12, 0xc6, 0xc4, 0x3a, 0xb3, 0x17, 0xee, 0x97,
	0xa6, 0xc6, 0xfa, 0x4a, 0x65, 0x7d, 0x4f, 0x1c, 0x6b, 0xfc, 0xe1, 0xbf, 0xa5, 0x0e, 0xdd, 0x81,
	0xde, 0x4d, 0xf7, 0x23, 0x81, 0x56, 0x10, 0xb2, 0x32, 0xa6, 0x32, 0x8e, 0xd6, 0x54, 0x1a, 0xe7,
	0x16, 0xb0, 0xe7, 0xd7, 0x4b, 0xd2, 0x5d, 0x9d, 0x0c, 0x57, 0x27, 0xb7, 0xc3, 0xd5, 0xdd, 0xd9,
	0xbe, 0xc6, 0xda, 0xc3, 0x2f, 0x0c, 0x02, 0xbd, 0xef, 0xfb, 0x24, 0xd1, 0x1b, 0xa8, 0x97, 0xfc,
	0x7e, 0xcd, 0x78, 0x55, 0x48, 0xe3, 0xa2, 0x35, 0x1d, 0xcc, 0x4a, 0x7e, 0xbf, 0x6a, 0xdf, 0xe8,
	0x1d, 0xbc, 0x8c, 0x5a, 0xb3, 0x05, 0x93, 0x3d, 0x31, 0x55, 0xc4, 0x62, 0xc8, 0x76, 0xd8, 0x5b,
	0x08, 0x8b, 0x2a, 0xcb, 0x7a, 0x64, 0xa6, 0x10, 0xbd, 0xcd, 0x74, 0xe5, 0x8f, 0x50, 0x1f, 0xbf,
	0xaf, 0xa1, 0x2b, 0x99, 0x57, 0xe4, 0x1f, 0xff, 0x11, 0xf9, 0x3c, 0x50, 0x1e, 0x95, 0x34, 0x38,
	0x35, 0xb9, 0xef, 0xf7, 0x7f, 0x4c, 0x6d, 0xdf, 0x98, 0xe0, 0xb1, 0x31, 0xc1, 0x53, 0x63, 0x82,
	0xdf, 0x8d, 0x09, 0x1e, 0x0e, 0xa6, 0xf6, 0x78, 0x30, 0xb5, 0xa7, 0x83, 0xa9, 0x7d, 0x3b, 0x57,
	0x13, 0xc2, 0x0b, 0x65, 0xfb, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xbe, 0x83, 0x6a,
	0xc8, 0x02, 0x00, 0x00,
}
