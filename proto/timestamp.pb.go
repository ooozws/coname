// Code generated by protoc-gen-gogo.
// source: timestamp.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Timestamp represents a point in time independent of any time zone
// or calendar, represented as seconds and fractions of seconds at
// nanosecond resolution in UTC Epoch time. It is encoded using the
// Proleptic Gregorian Calendar which extends the Gregorian calendar
// backwards to year one. It is encoded assuming all minutes are 60
// seconds long, i.e. leap seconds are "smeared" so that no leap second
// table is needed for interpretation. Range is from
// 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z.
// By restricting to that range, we ensure that we can convert to
// and from  RFC 3339 date strings.
// See [https://www.ietf.org/rfc/rfc3339.txt](https://www.ietf.org/rfc/rfc3339.txt).
//
// Example 1: Compute Timestamp from POSIX `time()`.
//
//     Timestamp timestamp;
//     timestamp.set_seconds(time(NULL));
//     timestamp.set_nanos(0);
//
// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
//
//     struct timeval tv;
//     gettimeofday(&tv, NULL);
//
//     Timestamp timestamp;
//     timestamp.set_seconds(tv.tv_sec);
//     timestamp.set_nanos(tv.tv_usec * 1000);
//
// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
//
//     FILETIME ft;
//     GetSystemTimeAsFileTime(&ft);
//     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
//
//     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
//     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
//     Timestamp timestamp;
//     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
//     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
//
// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
//
//     long millis = System.currentTimeMillis();
//
//     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
//         .setNanos((int) ((millis % 1000) * 1000000)).build();
//
// Example 5: Compute Timestamp from Python `datetime.datetime`.
//
//     now = datetime.datetime.utcnow()
//     seconds = int(time.mktime(now.timetuple()))
//     nanos = now.microsecond * 1000
//     timestamp = Timestamp(seconds=seconds, nanos=nanos)
//
type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptorTimestamp, []int{0} }

func init() {
	proto1.RegisterType((*Timestamp)(nil), "proto.Timestamp")
}
func (this *Timestamp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Timestamp)
	if !ok {
		that2, ok := that.(Timestamp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Timestamp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Timestamp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Timestamp but is not nil && this == nil")
	}
	if this.Seconds != that1.Seconds {
		return fmt.Errorf("Seconds this(%v) Not Equal that(%v)", this.Seconds, that1.Seconds)
	}
	if this.Nanos != that1.Nanos {
		return fmt.Errorf("Nanos this(%v) Not Equal that(%v)", this.Nanos, that1.Nanos)
	}
	return nil
}
func (this *Timestamp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Timestamp)
	if !ok {
		that2, ok := that.(Timestamp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Seconds != that1.Seconds {
		return false
	}
	if this.Nanos != that1.Nanos {
		return false
	}
	return true
}
func (this *Timestamp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&proto.Timestamp{")
	s = append(s, "Seconds: "+fmt.Sprintf("%#v", this.Seconds)+",\n")
	s = append(s, "Nanos: "+fmt.Sprintf("%#v", this.Nanos)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTimestamp(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringTimestamp(e map[int32]github_com_maditya_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *Timestamp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Timestamp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seconds != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintTimestamp(data, i, uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintTimestamp(data, i, uint64(m.Nanos))
	}
	return i, nil
}

func encodeFixed64Timestamp(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timestamp(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimestamp(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTimestamp(r randyTimestamp, easy bool) *Timestamp {
	this := &Timestamp{}
	this.Seconds = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Seconds *= -1
	}
	this.Nanos = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Nanos *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTimestamp interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTimestamp(r randyTimestamp) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTimestamp(r randyTimestamp) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTimestamp(r)
	}
	return string(tmps)
}
func randUnrecognizedTimestamp(r randyTimestamp, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldTimestamp(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldTimestamp(data []byte, r randyTimestamp, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateTimestamp(data, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateTimestamp(data, uint64(v2))
	case 1:
		data = encodeVarintPopulateTimestamp(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateTimestamp(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateTimestamp(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateTimestamp(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateTimestamp(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Timestamp) Size() (n int) {
	var l int
	_ = l
	if m.Seconds != 0 {
		n += 1 + sovTimestamp(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovTimestamp(uint64(m.Nanos))
	}
	return n
}

func sovTimestamp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimestamp(x uint64) (n int) {
	return sovTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Timestamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Timestamp{`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTimestamp(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Timestamp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimestamp
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
func skipTimestamp(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimestamp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimestamp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipTimestamp(data[start:])
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
	ErrInvalidLengthTimestamp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimestamp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("timestamp.proto", fileDescriptorTimestamp) }

var fileDescriptorTimestamp = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xd6, 0x5c, 0x9c, 0x21, 0x30, 0x19, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd4, 0xe4, 0xfc, 0xbc, 0x94,
	0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0x35, 0x2f, 0x31,
	0x2f, 0xbf, 0x58, 0x82, 0x09, 0x28, 0xce, 0x1a, 0x04, 0xe1, 0x38, 0x59, 0x5c, 0x78, 0x28, 0xc7,
	0x70, 0x03, 0x88, 0x1f, 0x3c, 0x94, 0x63, 0xfc, 0x00, 0xc4, 0x3f, 0x80, 0xb8, 0xe1, 0x91, 0x1c,
	0xe3, 0x0a, 0x20, 0xde, 0x01, 0xc4, 0x07, 0x80, 0xf8, 0x04, 0x10, 0x5f, 0x00, 0xe2, 0x07, 0x40,
	0xfc, 0xe2, 0x91, 0x1c, 0xc3, 0x07, 0x20, 0x9d, 0xc4, 0x06, 0xb6, 0xdd, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xce, 0xd2, 0x68, 0xf3, 0x97, 0x00, 0x00, 0x00,
}
