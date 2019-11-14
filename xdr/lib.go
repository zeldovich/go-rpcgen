package gonfs

import (
	"errors"
	"io"
	"math"
)

const TRUE = true
const FALSE = false

type XdrState struct {
	// err != nil means error state
	err error

	// reader != nil means we are reading
	reader io.Reader

	// writer != nil means we are writing
	writer io.Writer
}

func MakeReader(r io.Reader) *XdrState {
	return &XdrState{
		err: nil,
		reader: r,
		writer: nil,
	}
}

func MakeWriter(w io.Writer) *XdrState {
	return &XdrState{
		err: nil,
		reader: nil,
		writer: w,
	}
}

func (xs *XdrState) EncodingSetSize(arraysz *uint32, len int) {
	if xs.err != nil {
		return
	}

	if xs.writer == nil {
		return
	}

	if len > math.MaxUint32 {
		xs.setError(errors.New("length too large"))
		return
	}

	*arraysz = uint32(len)
}

func (xs *XdrState) Decoding() bool {
	return xs.err == nil && xs.reader != nil
}

func (xs *XdrState) setError(e error) {
	xs.err = e
}

func XdrBool(xs *XdrState, v *bool) {
	if xs.err != nil {
		return
	}


}

func XdrS32(xs *XdrState, v *int32) {
	if xs.err != nil {
		return
	}


}

func XdrU32(xs *XdrState, v *uint32) {
	if xs.err != nil {
		return
	}


}

func XdrS64(xs *XdrState, v *int64) {
	if xs.err != nil {
		return
	}


}

func XdrU64(xs *XdrState, v *uint64) {
	if xs.err != nil {
		return
	}


}

func XdrVarArray(xs *XdrState, maxlen int, v *[]byte) {
	if xs.err != nil {
		return
	}


}

func XdrArray(xs *XdrState, len int, v []byte) {
	if xs.err != nil {
		return
	}


}

func XdrString(xs *XdrState, maxlen int, v *string) {
	if xs.err != nil {
		return
	}


}
