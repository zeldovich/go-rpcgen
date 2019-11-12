package gonfs

import (
	"io"
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
