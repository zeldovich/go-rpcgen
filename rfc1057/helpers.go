package rfc1057

import (
	"io"
)

type writeBuffer struct {
	buf []byte
}

func (b *writeBuffer) Write(data []byte) (n int, err error) {
	b.buf = append(b.buf, data...)
	return len(data), nil
}

type readBuffer struct {
	buf []byte
}

func (b *readBuffer) Read(buf []byte) (n int, err error) {
	copy(buf, b.buf)
	b.buf = b.buf[len(buf):]
	n = len(buf)
	if n == 0 {
		err = io.EOF
	}
	return
}
