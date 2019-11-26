package rfc1057

import "github.com/zeldovich/go-rpcgen/xdr"

func (v *Auth_flavor) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Opaque_auth) Xdr(xs *xdr.XdrState) {
	(*Auth_flavor)(&((v).Flavor)).Xdr(xs)
	xdr.XdrVarArray(xs, int(400), (*[]byte)(&((v).Body)))
}
func (v *Msg_type) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Reply_stat) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Accept_stat) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Reject_stat) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Auth_stat) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Rpc_msg) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Xid)))
	(*Msg_type)(&((&((v).Body)).Mtype)).Xdr(xs)
	switch (&((v).Body)).Mtype {
	case CALL:
		(*Call_body)(&((&((v).Body)).Cbody)).Xdr(xs)
	case REPLY:
		(*Reply_body)(&((&((v).Body)).Rbody)).Xdr(xs)
	}
}
func (v *Call_body) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Rpcvers)))
	xdr.XdrU32(xs, (*uint32)(&((v).Prog)))
	xdr.XdrU32(xs, (*uint32)(&((v).Vers)))
	xdr.XdrU32(xs, (*uint32)(&((v).Proc)))
	(*Opaque_auth)(&((v).Cred)).Xdr(xs)
	(*Opaque_auth)(&((v).Verf)).Xdr(xs)
}
func (v *Reply_body) Xdr(xs *xdr.XdrState) {
	(*Reply_stat)(&((v).Stat)).Xdr(xs)
	switch (v).Stat {
	case MSG_ACCEPTED:
		(*Accepted_reply)(&((v).Areply)).Xdr(xs)
	case MSG_DENIED:
		(*Rejected_reply)(&((v).Rreply)).Xdr(xs)
	}
}
func (v *Accepted_reply) Xdr(xs *xdr.XdrState) {
	(*Opaque_auth)(&((v).Verf)).Xdr(xs)
	(*Accept_stat)(&((&((v).Reply_data)).Stat)).Xdr(xs)
	switch (&((v).Reply_data)).Stat {
	case SUCCESS:
		xdr.XdrArray(xs, (*&((&((v).Reply_data)).Results))[:])
	case PROG_MISMATCH:
		xdr.XdrU32(xs, (*uint32)(&((&((&((v).Reply_data)).Mismatch_info)).Low)))
		xdr.XdrU32(xs, (*uint32)(&((&((&((v).Reply_data)).Mismatch_info)).High)))
	default:
	}
}
func (v *Rejected_reply) Xdr(xs *xdr.XdrState) {
	(*Reject_stat)(&((v).Stat)).Xdr(xs)
	switch (v).Stat {
	case RPC_MISMATCH:
		xdr.XdrU32(xs, (*uint32)(&((&((v).Mismatch_info)).Low)))
		xdr.XdrU32(xs, (*uint32)(&((&((v).Mismatch_info)).High)))
	case AUTH_ERROR:
		(*Auth_stat)(&((v).Astat)).Xdr(xs)
	}
}
func (v *Auth_unix) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Stamp)))
	xdr.XdrString(xs, int(255), (*string)(&((v).Machinename)))
	xdr.XdrU32(xs, (*uint32)(&((v).Uid)))
	xdr.XdrU32(xs, (*uint32)(&((v).Gid)))
	{
		var __arraysz uint32
		xs.EncodingSetSize(&__arraysz, len(*&((v).Gids)))
		xdr.XdrU32(xs, (*uint32)(&__arraysz))

		if __arraysz > 16 {
			xs.SetError("array too large")
		} else {
			if xs.Decoding() {
				*&((v).Gids) = make([]uint32, __arraysz)
			}
			for i := uint64(0); i < uint64(__arraysz); i++ {
				xdr.XdrU32(xs, (*uint32)(&((*(&((v).Gids)))[i])))

			}
		}
	}
}
func (v *Mapping) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Prog)))
	xdr.XdrU32(xs, (*uint32)(&((v).Vers)))
	xdr.XdrU32(xs, (*uint32)(&((v).Prot)))
	xdr.XdrU32(xs, (*uint32)(&((v).Port)))
}
func (v *Pmaplist) Xdr(xs *xdr.XdrState) {
	if xs.Encoding() {
		opted := *(&v.P) != nil
		xdr.XdrBool(xs, (*bool)(&opted))
		if opted {
			(*Pmaplistelem)(*(&v.P)).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, (*bool)(&opted))
		if opted {
			*(&v.P) = new(Pmaplistelem)
			(*Pmaplistelem)(*(&v.P)).Xdr(xs)
		}
	}
}
func (v *Pmaplistelem) Xdr(xs *xdr.XdrState) {
	(*Mapping)(&((v).Map)).Xdr(xs)
	(*Pmaplist)(&((v).Next)).Xdr(xs)
}
func (v *Call_args) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Prog)))
	xdr.XdrU32(xs, (*uint32)(&((v).Vers)))
	xdr.XdrU32(xs, (*uint32)(&((v).Proc)))
	xdr.XdrVarArray(xs, int(-1), (*[]byte)(&((v).Args)))
}
func (v *Call_result) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(&((v).Port)))
	xdr.XdrVarArray(xs, int(-1), (*[]byte)(&((v).Res)))
}
func (v *Uint32) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}
func (v *Xbool) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, (*bool)(v))
}

type PMAP_PROG_PMAP_VERS_handler interface {
	PMAPPROC_NULL()
	PMAPPROC_SET(Mapping) Xbool
	PMAPPROC_UNSET(Mapping) Xbool
	PMAPPROC_GETPORT(Mapping) Uint32
	PMAPPROC_DUMP() Pmaplist
	PMAPPROC_CALLIT(Call_args) Call_result
}
type PMAP_PROG_PMAP_VERS_handler_wrapper struct {
	h PMAP_PROG_PMAP_VERS_handler
}

func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_NULL(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var out xdr.Void
	w.h.PMAPPROC_NULL()
	return &out, nil
}
func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_SET(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var in Mapping
	in.Xdr(args)
	err = args.Error()
	if err != nil {
		return
	}
	var out Xbool
	out = w.h.PMAPPROC_SET(in)
	return &out, nil
}
func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_UNSET(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var in Mapping
	in.Xdr(args)
	err = args.Error()
	if err != nil {
		return
	}
	var out Xbool
	out = w.h.PMAPPROC_UNSET(in)
	return &out, nil
}
func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_GETPORT(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var in Mapping
	in.Xdr(args)
	err = args.Error()
	if err != nil {
		return
	}
	var out Uint32
	out = w.h.PMAPPROC_GETPORT(in)
	return &out, nil
}
func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_DUMP(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var out Pmaplist
	out = w.h.PMAPPROC_DUMP()
	return &out, nil
}
func (w *PMAP_PROG_PMAP_VERS_handler_wrapper) PMAPPROC_CALLIT(args *xdr.XdrState) (res xdr.Xdrable, err error) {
	var in Call_args
	in.Xdr(args)
	err = args.Error()
	if err != nil {
		return
	}
	var out Call_result
	out = w.h.PMAPPROC_CALLIT(in)
	return &out, nil
}
func PMAP_PROG_PMAP_VERS_regs(h PMAP_PROG_PMAP_VERS_handler) []xdr.ProcRegistration {
	w := &PMAP_PROG_PMAP_VERS_handler_wrapper{h}
	return []xdr.ProcRegistration{
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_NULL,
			Handler: w.PMAPPROC_NULL,
		},
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_SET,
			Handler: w.PMAPPROC_SET,
		},
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_UNSET,
			Handler: w.PMAPPROC_UNSET,
		},
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_GETPORT,
			Handler: w.PMAPPROC_GETPORT,
		},
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_DUMP,
			Handler: w.PMAPPROC_DUMP,
		},
		xdr.ProcRegistration{
			Prog:    PMAP_PROG,
			Vers:    PMAP_VERS,
			Proc:    PMAPPROC_CALLIT,
			Handler: w.PMAPPROC_CALLIT,
		},
	}
}
