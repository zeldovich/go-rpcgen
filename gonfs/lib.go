package gonfs

import (
	"errors"
	"io"
	"math"
)

const TRUE = true
const FALSE = false

type xdrState struct {
	// err != nil means error state
	err error

	// reader != nil means we are reading
	reader io.Reader

	// writer != nil means we are writing
	writer io.Writer
}

func (xs *xdrState) encodingSetSize(arraysz *uint32, len int) {
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

func (xs *xdrState) decoding() bool {
	return xs.err == nil && xs.reader != nil
}

func (xs *xdrState) setError(e error) {
	xs.err = e
}

func xdrBool(xs *xdrState, v *bool) {
}

func xdrS32(xs *xdrState, v *int32) {
}

func xdrU32(xs *xdrState, v *uint32) {
}

func xdrS64(xs *xdrState, v *int64) {
}

func xdrU64(xs *xdrState, v *uint64) {
}

func xdrVarArray(xs *xdrState, maxlen int, v *[]byte) {
}

func xdrArray(xs *xdrState, len int, v []byte) {
}

func xdrString(xs *xdrState, maxlen int, v *string) {
}
