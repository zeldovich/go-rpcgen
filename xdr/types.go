package xdr

type Xdrable interface {
	Xdr(xs *XdrState)
}

type Bool bool
type Uint32 uint32
type Int32 int32
type Uint64 uint64
type Int64 int64
type Void struct{}

const TRUE Bool = true
const FALSE Bool = false

func (v *Bool) Xdr(xs *XdrState)   { XdrBool(xs, (*bool)(v)) }
func (v *Uint32) Xdr(xs *XdrState) { XdrU32(xs, (*uint32)(v)) }
func (v *Int32) Xdr(xs *XdrState)  { XdrS32(xs, (*int32)(v)) }
func (v *Uint64) Xdr(xs *XdrState) { XdrU64(xs, (*uint64)(v)) }
func (v *Int64) Xdr(xs *XdrState)  { XdrS64(xs, (*int64)(v)) }
func (v *Void) Xdr(xs *XdrState)   {}
