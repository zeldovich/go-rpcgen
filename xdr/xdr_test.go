package xdr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXdrString(t *testing.T) {
	xw := MakeWriter(nil)
	s := "hello"
	XdrString(xw, 100, &s)
	assert.NoError(t, xw.Error())
	buf := xw.WriteBuf()

	xr := MakeReader(buf)
	var newString string
	XdrString(xr, 100, &newString)
	assert.NoError(t, xr.Error())
	assert.Equal(t, s, newString)
}

func TestLargeStringLength(t *testing.T) {
	xw := MakeWriter(nil)
	s := "hello world"
	XdrString(xw, 100, &s)
	buf := xw.WriteBuf()
	// cut off the actual string (but leave the length as len(s))
	//
	// 6 is enough for the length prefix (4 bytes) and a little bit for the
	// string
	buf = buf[:6]

	xr := MakeReader(buf)
	var newString string
	XdrString(xr, 100, &newString)
	assert.Error(t, xr.Error())
}

func TestXdrVarArray(t *testing.T) {
	xw := MakeWriter(nil)
	v := []byte{1, 2, 3}
	XdrVarArray(xw, 100, &v)
	assert.NoError(t, xw.Error())
	buf := xw.WriteBuf()

	xr := MakeReader(buf)
	var newArray []byte
	XdrVarArray(xr, 100, &newArray)
	assert.NoError(t, xr.Error())
	assert.Equal(t, v, newArray)
}

func TestLargeVarArrayLength(t *testing.T) {
	xw := MakeWriter(nil)
	v := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	XdrVarArray(xw, 100, &v)
	buf := xw.WriteBuf()
	buf = buf[:6]

	xr := MakeReader(buf)
	var newArray []byte
	XdrVarArray(xr, 100, &newArray)
	assert.Error(t, xr.Error())
}
