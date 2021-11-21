package xdr

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type XdrState struct {
	// err != nil means error state
	err error

	// reader means we are reading; otherwise writing
	reader bool

	// buf has the data to be read (or already written)
	buf []byte
}

func MakeReader(buf []byte) *XdrState {
	return &XdrState{
		err:    nil,
		reader: true,
		buf:    buf,
	}
}

func MakeWriter(buf []byte) *XdrState {
	return &XdrState{
		err:    nil,
		reader: false,
		buf:    buf,
	}
}

func (xs *XdrState) WriteBuf() []byte {
	return xs.buf
}

func (xs *XdrState) EncodingSetSize(arraysz *uint32, len int) {
	if xs.err != nil {
		return
	}

	if xs.reader {
		return
	}

	if len > math.MaxUint32 {
		xs.SetError("length too large")
		return
	}

	*arraysz = uint32(len)
}

func (xs *XdrState) Encoding() bool {
	return xs.err == nil && xs.reader == false
}

func (xs *XdrState) Decoding() bool {
	return xs.err == nil && xs.reader == true
}

func (xs *XdrState) SetError(s string) {
	xs.err = errors.New(s)
}

func (xs *XdrState) Error() error {
	return xs.err
}

func xdrRW(xs *XdrState, v []byte) {
	if xs.err != nil {
		return
	}

	if xs.reader {
		n := copy(v, xs.buf)
		xs.buf = xs.buf[n:]

		if n < len(v) {
			xs.err = fmt.Errorf("Not enough bytes: wanted %d, got %d", len(v), n)
		}
	} else {
		xs.buf = append(xs.buf, v...)
	}
}

func xdrR(xs *XdrState, n int) []byte {
	if len(xs.buf) < n {
		xs.err = fmt.Errorf("Not enough bytes: wanted %d, have %d", n, len(xs.buf))
	}

	if xs.err != nil {
		return nil
	}

	res := xs.buf[:n:n]
	xs.buf = xs.buf[n:]
	return res
}

func XdrBool(xs *XdrState, v *bool) {
	if xs.err != nil {
		return
	}

	var buf [4]byte
	if xs.Encoding() {
		if *v {
			binary.BigEndian.PutUint32(buf[:], 1)
		} else {
			binary.BigEndian.PutUint32(buf[:], 0)
		}
	}

	xdrRW(xs, buf[:])

	if xs.Decoding() {
		r := binary.BigEndian.Uint32(buf[:])
		if r == 0 {
			*v = false
		} else {
			*v = true
		}
	}
}

func XdrS32(xs *XdrState, v *int32) {
	if xs.err != nil {
		return
	}

	var buf [4]byte
	if xs.Encoding() {
		binary.BigEndian.PutUint32(buf[:], uint32(*v))
	}
	xdrRW(xs, buf[:])
	if xs.Decoding() {
		*v = int32(binary.BigEndian.Uint32(buf[:]))
	}
}

func XdrU32(xs *XdrState, v *uint32) {
	if xs.err != nil {
		return
	}

	var buf [4]byte
	if xs.Encoding() {
		binary.BigEndian.PutUint32(buf[:], *v)
	}
	xdrRW(xs, buf[:])
	if xs.Decoding() {
		*v = binary.BigEndian.Uint32(buf[:])
	}
}

func XdrS64(xs *XdrState, v *int64) {
	if xs.err != nil {
		return
	}

	var buf [8]byte
	if xs.Encoding() {
		binary.BigEndian.PutUint64(buf[:], uint64(*v))
	}
	xdrRW(xs, buf[:])
	if xs.Decoding() {
		*v = int64(binary.BigEndian.Uint64(buf[:]))
	}
}

func XdrU64(xs *XdrState, v *uint64) {
	if xs.err != nil {
		return
	}

	var buf [8]byte
	if xs.Encoding() {
		binary.BigEndian.PutUint64(buf[:], *v)
	}
	xdrRW(xs, buf[:])
	if xs.Decoding() {
		*v = binary.BigEndian.Uint64(buf[:])
	}
}

func XdrVarArray(xs *XdrState, maxlen int, v *[]byte) {
	if xs.err != nil {
		return
	}

	if xs.Encoding() {
		if len(*v) > math.MaxUint32 || (maxlen >= 0 && len(*v) > maxlen) {
			xs.SetError("var array too large")
			return
		}

		var szbuf [4]byte
		binary.BigEndian.PutUint32(szbuf[:], uint32(len(*v)))
		xdrRW(xs, szbuf[:])
		xdrRW(xs, *v)
	} else {
		var szbuf [4]byte
		xdrRW(xs, szbuf[:])
		sz32 := binary.BigEndian.Uint32(szbuf[:])
		sz := int(sz32)

		if (maxlen >= 0 && sz > maxlen) || sz < 0 {
			xs.SetError("var array too large")
			return
		}

		*v = xdrR(xs, sz)
	}

	if len(*v)%4 != 0 {
		xdrRW(xs, make([]byte, (4-len(*v)%4)%4))
	}
}

func XdrArray(xs *XdrState, v []byte) {
	if xs.err != nil {
		return
	}

	xdrRW(xs, v)
	if len(v)%4 != 0 {
		xdrRW(xs, make([]byte, (4-len(v)%4)%4))
	}

	// Check that the padding values are zero?
}

func XdrString(xs *XdrState, maxlen int, v *string) {
	if xs.err != nil {
		return
	}

	if xs.Encoding() {
		if len(*v) > math.MaxUint32 || (maxlen >= 0 && len(*v) > maxlen) {
			xs.SetError("string too large")
			return
		}

		var szbuf [4]byte
		binary.BigEndian.PutUint32(szbuf[:], uint32(len(*v)))
		xdrRW(xs, szbuf[:])

		xdrRW(xs, []byte(*v))
		xdrRW(xs, make([]byte, (4-len(*v)%4)%4))
	} else {
		var szbuf [4]byte
		xdrRW(xs, szbuf[:])
		sz32 := binary.BigEndian.Uint32(szbuf[:])
		sz := int(sz32)

		if (maxlen >= 0 && sz > maxlen) || sz < 0 {
			xs.SetError("string too large")
			return
		}

		buf := xdrR(xs, sz)
		*v = string(buf)

		xdrRW(xs, make([]byte, (4-len(*v)%4)%4))
	}
}
