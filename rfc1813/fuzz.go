//go:build gofuzz
// +build gofuzz

package rfc1813

import (
	"fmt"
	"reflect"

	"github.com/zeldovich/go-rpcgen/xdr"
)

// type FUZZ_T = CREATE3args
type FUZZ_T = READDIRPLUS3res

func Fuzz(data []byte) int {
	var x FUZZ_T
	err := xdr.DecodeBuf(data, &x)
	if err != nil {
		return 0
	}

	_, err = xdr.EncodeBuf(&x)
	if err != nil {
		panic(err)
	}

	var y FUZZ_T
	xdr.DecodeBuf(data, &y)
	if !reflect.DeepEqual(x, y) {
		panic(fmt.Sprintf("%+v does not roundtrip; got %+v", x, y))
	}

	return 1
}
