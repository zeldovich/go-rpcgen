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

func (xs *XdrState) encodingSetSize(arraysz *uint32, len int) {
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

func (xs *XdrState) decoding() bool {
	return xs.err == nil && xs.reader != nil
}

func (xs *XdrState) setError(e error) {
	xs.err = e
}

func xdrBool(xs *XdrState, v *bool) {
}

func xdrS32(xs *XdrState, v *int32) {
}

func xdrU32(xs *XdrState, v *uint32) {
}

func xdrS64(xs *XdrState, v *int64) {
}

func xdrU64(xs *XdrState, v *uint64) {
}

func xdrVarArray(xs *XdrState, maxlen int, v *[]byte) {
}

func xdrArray(xs *XdrState, len int, v []byte) {
}

func xdrString(xs *XdrState, maxlen int, v *string) {
}
