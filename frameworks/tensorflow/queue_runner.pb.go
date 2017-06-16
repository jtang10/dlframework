// Code generated by protoc-gen-gogo.
// source: queue_runner.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName,proto3" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName,proto3" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []Error_Code `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,enum=tensorflow.Error_Code" json:"queue_closed_exception_types,omitempty"`
}

func (m *QueueRunnerDef) Reset()                    { *m = QueueRunnerDef{} }
func (m *QueueRunnerDef) String() string            { return proto.CompactTextString(m) }
func (*QueueRunnerDef) ProtoMessage()               {}
func (*QueueRunnerDef) Descriptor() ([]byte, []int) { return fileDescriptorQueueRunner, []int{0} }

func (m *QueueRunnerDef) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRunnerDef) GetEnqueueOpName() []string {
	if m != nil {
		return m.EnqueueOpName
	}
	return nil
}

func (m *QueueRunnerDef) GetCloseOpName() string {
	if m != nil {
		return m.CloseOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetCancelOpName() string {
	if m != nil {
		return m.CancelOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetQueueClosedExceptionTypes() []Error_Code {
	if m != nil {
		return m.QueueClosedExceptionTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueRunnerDef)(nil), "tensorflow.QueueRunnerDef")
}
func (m *QueueRunnerDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueRunnerDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.QueueName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQueueRunner(dAtA, i, uint64(len(m.QueueName)))
		i += copy(dAtA[i:], m.QueueName)
	}
	if len(m.EnqueueOpName) > 0 {
		for _, s := range m.EnqueueOpName {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.CloseOpName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintQueueRunner(dAtA, i, uint64(len(m.CloseOpName)))
		i += copy(dAtA[i:], m.CloseOpName)
	}
	if len(m.CancelOpName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintQueueRunner(dAtA, i, uint64(len(m.CancelOpName)))
		i += copy(dAtA[i:], m.CancelOpName)
	}
	if len(m.QueueClosedExceptionTypes) > 0 {
		dAtA2 := make([]byte, len(m.QueueClosedExceptionTypes)*10)
		var j1 int
		for _, num := range m.QueueClosedExceptionTypes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintQueueRunner(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeFixed64QueueRunner(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32QueueRunner(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintQueueRunner(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QueueRunnerDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.QueueName)
	if l > 0 {
		n += 1 + l + sovQueueRunner(uint64(l))
	}
	if len(m.EnqueueOpName) > 0 {
		for _, s := range m.EnqueueOpName {
			l = len(s)
			n += 1 + l + sovQueueRunner(uint64(l))
		}
	}
	l = len(m.CloseOpName)
	if l > 0 {
		n += 1 + l + sovQueueRunner(uint64(l))
	}
	l = len(m.CancelOpName)
	if l > 0 {
		n += 1 + l + sovQueueRunner(uint64(l))
	}
	if len(m.QueueClosedExceptionTypes) > 0 {
		l = 0
		for _, e := range m.QueueClosedExceptionTypes {
			l += sovQueueRunner(uint64(e))
		}
		n += 1 + sovQueueRunner(uint64(l)) + l
	}
	return n
}

func sovQueueRunner(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQueueRunner(x uint64) (n int) {
	return sovQueueRunner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueueRunnerDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueueRunner
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
			return fmt.Errorf("proto: QueueRunnerDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueRunnerDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueueRunner
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
				return ErrInvalidLengthQueueRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueueName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnqueueOpName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueueRunner
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
				return ErrInvalidLengthQueueRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnqueueOpName = append(m.EnqueueOpName, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloseOpName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueueRunner
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
				return ErrInvalidLengthQueueRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CloseOpName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelOpName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueueRunner
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
				return ErrInvalidLengthQueueRunner
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CancelOpName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v Error_Code
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQueueRunner
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Error_Code(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.QueueClosedExceptionTypes = append(m.QueueClosedExceptionTypes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQueueRunner
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
					return ErrInvalidLengthQueueRunner
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Error_Code
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueueRunner
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Error_Code(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.QueueClosedExceptionTypes = append(m.QueueClosedExceptionTypes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueClosedExceptionTypes", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueueRunner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueueRunner
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
func skipQueueRunner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueueRunner
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
					return 0, ErrIntOverflowQueueRunner
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
					return 0, ErrIntOverflowQueueRunner
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
				return 0, ErrInvalidLengthQueueRunner
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQueueRunner
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
				next, err := skipQueueRunner(dAtA[start:])
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
	ErrInvalidLengthQueueRunner = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueueRunner   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("queue_runner.proto", fileDescriptorQueueRunner) }

var fileDescriptorQueueRunner = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0xb4, 0x40,
	0x10, 0x85, 0xff, 0xfe, 0x51, 0x13, 0x5a, 0x07, 0x33, 0x2c, 0x0c, 0x1a, 0x25, 0x64, 0x62, 0x0c,
	0x2b, 0x16, 0x7a, 0x03, 0x46, 0xb7, 0x3a, 0x12, 0x13, 0x97, 0x04, 0xa1, 0x30, 0x46, 0xe8, 0xc2,
	0x6a, 0xc8, 0xe8, 0x2d, 0x3c, 0x96, 0x4b, 0x8f, 0x60, 0xf0, 0x12, 0xc6, 0x95, 0xa1, 0x1a, 0x9d,
	0xd9, 0xbe, 0xfe, 0xbe, 0xd7, 0xa9, 0x27, 0xdd, 0xa7, 0x0e, 0x3a, 0x48, 0xa9, 0x53, 0x0a, 0x28,
	0x6a, 0x08, 0x5b, 0x74, 0x65, 0x0b, 0x4a, 0x23, 0x95, 0x15, 0x2e, 0x0f, 0xa6, 0x40, 0x84, 0x94,
	0xe6, 0x58, 0x80, 0x36, 0xcf, 0xb3, 0x6f, 0x21, 0x9d, 0xeb, 0xc1, 0x4a, 0x58, 0x3a, 0x87, 0xd2,
	0x3d, 0x92, 0xd2, 0xf4, 0xa8, 0xac, 0x06, 0x4f, 0x04, 0x22, 0xb4, 0x13, 0x9b, 0x93, 0xcb, 0xac,
	0x06, 0xf7, 0x44, 0xee, 0x82, 0x32, 0x00, 0x36, 0x86, 0xf9, 0x1f, 0x58, 0xa1, 0x9d, 0x4c, 0xc6,
	0xf8, 0xaa, 0x61, 0x6e, 0x26, 0x27, 0x79, 0x85, 0x7a, 0x45, 0x59, 0xdc, 0xb4, 0xcd, 0xe1, 0xc8,
	0x1c, 0x4b, 0x27, 0xcf, 0x54, 0x0e, 0xd5, 0x1f, 0xb4, 0xc1, 0xd0, 0x8e, 0x49, 0x47, 0xea, 0x56,
	0x1e, 0x9a, 0xff, 0x58, 0x2d, 0x52, 0x78, 0xce, 0xa1, 0x69, 0x1f, 0x50, 0xa5, 0xed, 0x4b, 0x03,
	0xda, 0xdb, 0x0c, 0xac, 0xd0, 0x39, 0xdd, 0x8b, 0x56, 0x97, 0x46, 0x7c, 0x68, 0x34, 0xc7, 0x02,
	0x92, 0x7d, 0x76, 0xe7, 0xac, 0x5e, 0xfc, 0x9a, 0x37, 0x83, 0x18, 0xc7, 0x6f, 0xbd, 0x2f, 0xde,
	0x7b, 0x5f, 0x7c, 0xf4, 0xbe, 0x78, 0xfd, 0xf4, 0xff, 0x49, 0x0f, 0xe9, 0x7e, 0xbd, 0xa7, 0xa4,
	0xac, 0x86, 0x25, 0xd2, 0x63, 0x3c, 0x5d, 0x5b, 0x69, 0x31, 0x4c, 0xa7, 0x17, 0xe2, 0x4b, 0x88,
	0xbb, 0x2d, 0xde, 0xf1, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x62, 0x8e, 0x11, 0x48, 0x7c, 0x01,
	0x00, 0x00,
}