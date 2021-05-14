package xdr

func EncodeBuf(v Xdrable) (res []byte, err error) {
	x := MakeWriter(nil)
	v.Xdr(x)
	return x.WriteBuf(), x.Error()
}

func DecodeBuf(buf []byte, v Xdrable) error {
	x := MakeReader(buf)
	v.Xdr(x)
	return x.Error()
}
