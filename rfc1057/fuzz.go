// +build gofuzz

package rfc1057

import (
	"github.com/zeldovich/go-rpcgen/xdr"
)

func Fuzz(data []byte) int {
	var x Rpc_msg
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
