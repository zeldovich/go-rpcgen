//go:build gofuzz
// +build gofuzz

package rfc1813

import (
	"github.com/zeldovich/go-rpcgen/xdr"
)

func Fuzz(data []byte) int {
	// var x CREATE3args
	var x READDIRPLUS3res
	err := xdr.DecodeBuf(data, &x)
	if err != nil {
		return 0
	}

	_, err = xdr.EncodeBuf(&x)
	if err != nil {
		panic(err)
	}

	return 1
}
