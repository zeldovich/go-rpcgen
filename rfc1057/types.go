package rfc1057

type Auth_flavor int32

const AUTH_NONE Auth_flavor = 0
const AUTH_UNIX Auth_flavor = 1
const AUTH_SHORT Auth_flavor = 2
const AUTH_DES Auth_flavor = 3

type Opaque_auth struct {
	Flavor Auth_flavor
	Body   []byte
}
type Msg_type int32

const CALL Msg_type = 0
const REPLY Msg_type = 1

type Reply_stat int32

const MSG_ACCEPTED Reply_stat = 0
const MSG_DENIED Reply_stat = 1

type Accept_stat int32

const SUCCESS Accept_stat = 0
const PROG_UNAVAIL Accept_stat = 1
const PROG_MISMATCH Accept_stat = 2
const PROC_UNAVAIL Accept_stat = 3
const GARBAGE_ARGS Accept_stat = 4

type Reject_stat int32

const RPC_MISMATCH Reject_stat = 0
const AUTH_ERROR Reject_stat = 1

type Auth_stat int32

const AUTH_BADCRED Auth_stat = 1
const AUTH_REJECTEDCRED Auth_stat = 2
const AUTH_BADVERF Auth_stat = 3
const AUTH_REJECTEDVERF Auth_stat = 4
const AUTH_TOOWEAK Auth_stat = 5

type Rpc_msg struct {
	Xid  uint32
	Body struct {
		Mtype Msg_type
		Cbody Call_body
		Rbody Reply_body
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
type Reply_body struct {
	Stat   Reply_stat
	Areply Accepted_reply
	Rreply Rejected_reply
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
type Rejected_reply struct {
	Stat          Reject_stat
	Mismatch_info struct {
		Low  uint32
		High uint32
	}
	Astat Auth_stat
}
type Auth_unix struct {
	Stamp       uint32
	Machinename string
	Uid         uint32
	Gid         uint32
	Gids        []uint32
}

const PMAP_PORT = 111

type Mapping struct {
	Prog uint32
	Vers uint32
	Prot uint32
	Port uint32
}

const IPPROTO_TCP = 6
const IPPROTO_UDP = 17

type Pmaplist struct{ P *Pmaplistelem }
type Pmaplistelem struct {
	Map  Mapping
	Next Pmaplist
}
type Call_args struct {
	Prog uint32
	Vers uint32
	Proc uint32
	Args []byte
}
type Call_result struct {
	Port uint32
	Res  []byte
}

const PMAP_PROG uint32 = 100000
const PMAP_VERS uint32 = 2
const PMAPPROC_NULL uint32 = 0
const PMAPPROC_SET uint32 = 1
const PMAPPROC_UNSET uint32 = 2
const PMAPPROC_GETPORT uint32 = 3
const PMAPPROC_DUMP uint32 = 4
const PMAPPROC_CALLIT uint32 = 5
