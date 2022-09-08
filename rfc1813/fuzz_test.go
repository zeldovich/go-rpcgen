package rfc1813

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zeldovich/go-rpcgen/xdr"
)

func FuzzDecodeEncode(f *testing.F) {
	var r READDIRPLUS3res
	res, err := xdr.EncodeBuf(&r)
	if err != nil {
		panic(err)
	}
	f.Add(res)

	r.Resok.Reply.Entries = &Entryplus3{}
	r.Resok.Reply.Entries.Nextentry = &Entryplus3{}
	res, err = xdr.EncodeBuf(&r)
	if err != nil {
		panic(err)
	}
	f.Add(res)

	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		var x READDIRPLUS3res
		err := xdr.DecodeBuf(data, &x)
		if err != nil {
			return
		}

		var res []byte
		res, err = xdr.EncodeBuf(&x)
		if err != nil {
			panic(err)
		}

		var y READDIRPLUS3res
		xdr.DecodeBuf(res, &y)
		if !reflect.DeepEqual(x, y) {
			panic(fmt.Sprintf("%+v does not roundtrip; got %+v", x, y))
		}
	})
}
