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
