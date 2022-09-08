package rfc1057

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zeldovich/go-rpcgen/xdr"
)

func FuzzDecodeEncode(f *testing.F) {
	var m Rpc_msg
	res, err := xdr.EncodeBuf(&m)
	if err != nil {
		panic(err)
	}
	f.Add(res)

	m.Xid = 1
	m.Body.Mtype = 2
	m.Body.Cbody.Rpcvers = 3
	m.Body.Cbody.Prog = 4
	m.Body.Cbody.Cred.Flavor = AUTH_UNIX
	m.Body.Cbody.Cred.Body = []byte{1, 2, 3, 4}
	res, err = xdr.EncodeBuf(&m)
	if err != nil {
		panic(err)
	}
	f.Add(res)

	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		var x Rpc_msg

		err := xdr.DecodeBuf(data, &x)
		if err != nil {
			return
		}

		var res []byte
		res, err = xdr.EncodeBuf(&x)
		if err != nil {
			panic(err)
		}

		var y Rpc_msg
		xdr.DecodeBuf(res, &y)
		if !reflect.DeepEqual(x, y) {
			panic(fmt.Sprintf("%+v does not roundtrip; got %+v", x, y))
		}
	})
}
