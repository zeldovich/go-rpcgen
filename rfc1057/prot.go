package rfc1057

import . "github.com/zeldovich/go-rpcgen/xdr"

type Auth_flavor int32

func (v *Auth_flavor) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const AUTH_NONE = 0
const AUTH_UNIX = 1
const AUTH_SHORT = 2
const AUTH_DES = 3

type Opaque_auth struct {
	Flavor Auth_flavor
	Body   []byte
}

func (v *Opaque_auth) Xdr(xs *XdrState) {
	(*Auth_flavor)(&((v).Flavor)).Xdr(xs)
	XdrVarArray(xs, 400, (*[]byte)(&((v).Body)))
}

type Msg_type int32

func (v *Msg_type) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const CALL = 0
const REPLY = 1

type Reply_stat int32

func (v *Reply_stat) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const MSG_ACCEPTED = 0
const MSG_DENIED = 1

type Accept_stat int32

func (v *Accept_stat) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const SUCCESS = 0
const PROG_UNAVAIL = 1
const PROG_MISMATCH = 2
const PROC_UNAVAIL = 3
const GARBAGE_ARGS = 4

type Reject_stat int32

func (v *Reject_stat) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const RPC_MISMATCH = 0
const AUTH_ERROR = 1

type Auth_stat int32

func (v *Auth_stat) Xdr(xs *XdrState) {
	XdrS32(xs, (*int32)(v))
}

const AUTH_BADCRED = 1
const AUTH_REJECTEDCRED = 2
const AUTH_BADVERF = 3
const AUTH_REJECTEDVERF = 4
const AUTH_TOOWEAK = 5

type Rpc_msg struct {
	Xid  uint32
	Body struct {
		Mtype Msg_type
		Cbody Call_body
		Rbody Reply_body
	}
}

func (v *Rpc_msg) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Xid)))
	(*Msg_type)(&((&((v).Body)).Mtype)).Xdr(xs)
	switch (&((v).Body)).Mtype {
	case CALL:
		(*Call_body)(&((&((v).Body)).Cbody)).Xdr(xs)
	case REPLY:
		(*Reply_body)(&((&((v).Body)).Rbody)).Xdr(xs)
	}
}

type Call_body struct {
	Rpcvers uint32
	Prog    uint32
	Vers    uint32
	Proc    uint32
	Cred    Opaque_auth
	Verf    Opaque_auth
}

func (v *Call_body) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Rpcvers)))
	XdrU32(xs, (*uint32)(&((v).Prog)))
	XdrU32(xs, (*uint32)(&((v).Vers)))
	XdrU32(xs, (*uint32)(&((v).Proc)))
	(*Opaque_auth)(&((v).Cred)).Xdr(xs)
	(*Opaque_auth)(&((v).Verf)).Xdr(xs)
}

type Reply_body struct {
	Stat   Reply_stat
	Areply Accepted_reply
	Rreply Rejected_reply
}

func (v *Reply_body) Xdr(xs *XdrState) {
	(*Reply_stat)(&((v).Stat)).Xdr(xs)
	switch (v).Stat {
	case MSG_ACCEPTED:
		(*Accepted_reply)(&((v).Areply)).Xdr(xs)
	case MSG_DENIED:
		(*Rejected_reply)(&((v).Rreply)).Xdr(xs)
	}
}

type Accepted_reply struct {
	Verf       Opaque_auth
	Reply_data struct {
		Stat          Accept_stat
		Results       [0]byte
		Mismatch_info struct {
			Low  uint32
			High uint32
		}
	}
}

func (v *Accepted_reply) Xdr(xs *XdrState) {
	(*Opaque_auth)(&((v).Verf)).Xdr(xs)
	(*Accept_stat)(&((&((v).Reply_data)).Stat)).Xdr(xs)
	switch (&((v).Reply_data)).Stat {
	case SUCCESS:
		XdrArray(xs, (*&((&((v).Reply_data)).Results))[:])
	case PROG_MISMATCH:
		XdrU32(xs, (*uint32)(&((&((&((v).Reply_data)).Mismatch_info)).Low)))
		XdrU32(xs, (*uint32)(&((&((&((v).Reply_data)).Mismatch_info)).High)))
	default:
	}
}

type Rejected_reply struct {
	Stat          Reject_stat
	Mismatch_info struct {
		Low  uint32
		High uint32
	}
	Astat Auth_stat
}

func (v *Rejected_reply) Xdr(xs *XdrState) {
	(*Reject_stat)(&((v).Stat)).Xdr(xs)
	switch (v).Stat {
	case RPC_MISMATCH:
		XdrU32(xs, (*uint32)(&((&((v).Mismatch_info)).Low)))
		XdrU32(xs, (*uint32)(&((&((v).Mismatch_info)).High)))
	case AUTH_ERROR:
		(*Auth_stat)(&((v).Astat)).Xdr(xs)
	}
}

type Auth_unix struct {
	Stamp       uint32
	Machinename string
	Uid         uint32
	Gid         uint32
	Gids        []uint32
}

func (v *Auth_unix) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Stamp)))
	XdrString(xs, 255, (*string)(&((v).Machinename)))
	XdrU32(xs, (*uint32)(&((v).Uid)))
	XdrU32(xs, (*uint32)(&((v).Gid)))
	{
		var __arraysz uint32
		xs.EncodingSetSize(&__arraysz, len(*&((v).Gids)))
		XdrU32(xs, (*uint32)(&__arraysz))

		if __arraysz > 16 {
			xs.SetError("array too large")
		} else {
			if xs.Decoding() {
				*&((v).Gids) = make([]uint32, __arraysz)
			}
			for i := uint64(0); i < uint64(__arraysz); i++ {
				XdrU32(xs, (*uint32)(&((*(&((v).Gids)))[i])))

			}
		}
	}
}

const PMAP_PORT = 111

type Mapping struct {
	Prog uint32
	Vers uint32
	Prot uint32
	Port uint32
}

func (v *Mapping) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Prog)))
	XdrU32(xs, (*uint32)(&((v).Vers)))
	XdrU32(xs, (*uint32)(&((v).Prot)))
	XdrU32(xs, (*uint32)(&((v).Port)))
}

const IPPROTO_TCP = 6
const IPPROTO_UDP = 17

type Pmaplist struct{ P *Pmaplistelem }

func (v *Pmaplist) Xdr(xs *XdrState) {
	if xs.Encoding() {
		opted := (v).P != nil
		XdrBool(xs, &opted)
		if opted {
			(*Pmaplistelem)((v).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		XdrBool(xs, &opted)
		if opted {
			(v).P = new(Pmaplistelem)
			(*Pmaplistelem)((v).P).Xdr(xs)
		}
	}
}

type Pmaplistelem struct {
	Map  Mapping
	Next Pmaplist
}

func (v *Pmaplistelem) Xdr(xs *XdrState) {
	(*Mapping)(&((v).Map)).Xdr(xs)
	(*Pmaplist)(&((v).Next)).Xdr(xs)
}

type Call_args struct {
	Prog uint32
	Vers uint32
	Proc uint32
	Args []byte
}

func (v *Call_args) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Prog)))
	XdrU32(xs, (*uint32)(&((v).Vers)))
	XdrU32(xs, (*uint32)(&((v).Proc)))
	XdrVarArray(xs, -1, (*[]byte)(&((v).Args)))
}

type Call_result struct {
	Port uint32
	Res  []byte
}

func (v *Call_result) Xdr(xs *XdrState) {
	XdrU32(xs, (*uint32)(&((v).Port)))
	XdrVarArray(xs, -1, (*[]byte)(&((v).Res)))
}

const PMAP_PROG = 100000
const PMAP_VERS = 2
const PMAPPROC_NULL = 0
const PMAPPROC_SET = 1
const PMAPPROC_UNSET = 2
const PMAPPROC_GETPORT = 3
const PMAPPROC_DUMP = 4
const PMAPPROC_CALLIT = 5
