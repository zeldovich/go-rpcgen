package xdr

import (
	"io"
)

type helperBuf struct {
	buf []byte
}

func (h *helperBuf) Write(data []byte) (n int, err error) {
	h.buf = append(h.buf, data...)
	return len(data), nil
}

func (h *helperBuf) Read(buf []byte) (n int, err error) {
	copy(buf, h.buf)
	h.buf = h.buf[len(buf):]
	n = len(buf)
	if n == 0 {
		err = io.EOF
	}
	return
}

func EncodeBuf(v Xdrable) (res []byte, err error) {
	h := &helperBuf{}
	x := MakeWriter(h)
	v.Xdr(x)
	return h.buf, x.Error()
}

func DecodeBuf(buf []byte, v Xdrable) error {
	h := &helperBuf{buf}
	x := MakeReader(h)
	v.Xdr(x)
	return x.Error()
}
