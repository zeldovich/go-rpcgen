package rfc1057

import (
	"io"
)

type rwBuffer struct {
	buf []byte
}

func (rw *rwBuffer) Write(data []byte) (n int, err error) {
	rw.buf = append(rw.buf, data...)
	return len(data), nil
}

func (rw *rwBuffer) Read(buf []byte) (n int, err error) {
	copy(buf, rw.buf)
	rw.buf = rw.buf[len(buf):]
	n = len(buf)
	if n == 0 {
		err = io.EOF
	}
	return
}
