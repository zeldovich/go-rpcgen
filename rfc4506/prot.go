package rfc4506

import "github.com/zeldovich/go-rpcgen/xdr"

const PROGRAM = 100003
const VERSION = 3
const NFS3_FHSIZE = 64
const NFS3_COOKIEVERFSIZE = 8
const NFS3_CREATEVERFSIZE = 8
const NFS3_WRITEVERFSIZE = 8

type Uint64 uint64

func (v *Uint64) Xdr(xs *xdr.XdrState) {
	xdr.XdrU64(xs, (*uint64)(v))
}

type Int64 int64

func (v *Int64) Xdr(xs *xdr.XdrState) {
	xdr.XdrS64(xs, (*int64)(v))
}

type Uint32 uint32

func (v *Uint32) Xdr(xs *xdr.XdrState) {
	xdr.XdrU32(xs, (*uint32)(v))
}

type Int32 int32

func (v *Int32) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

type Filename3 string

func (v *Filename3) Xdr(xs *xdr.XdrState) {
	xdr.XdrString(xs, -1, (*string)(v))
}

type Nfspath3 string

func (v *Nfspath3) Xdr(xs *xdr.XdrState) {
	xdr.XdrString(xs, -1, (*string)(v))
}

type Fileid3 Uint64

func (v *Fileid3) Xdr(xs *xdr.XdrState) {
	(*Uint64)(v).Xdr(xs)
}

type Cookie3 Uint64

func (v *Cookie3) Xdr(xs *xdr.XdrState) {
	(*Uint64)(v).Xdr(xs)
}

type Cookieverf3 [NFS3_COOKIEVERFSIZE]byte

func (v *Cookieverf3) Xdr(xs *xdr.XdrState) {
	xdr.XdrArray(xs, (*v)[:])
}

type Createverf3 [NFS3_CREATEVERFSIZE]byte

func (v *Createverf3) Xdr(xs *xdr.XdrState) {
	xdr.XdrArray(xs, (*v)[:])
}

type Writeverf3 [NFS3_WRITEVERFSIZE]byte

func (v *Writeverf3) Xdr(xs *xdr.XdrState) {
	xdr.XdrArray(xs, (*v)[:])
}

type Uid3 Uint32

func (v *Uid3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(v).Xdr(xs)
}

type Gid3 Uint32

func (v *Gid3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(v).Xdr(xs)
}

type Size3 Uint64

func (v *Size3) Xdr(xs *xdr.XdrState) {
	(*Uint64)(v).Xdr(xs)
}

type Offset3 Uint64

func (v *Offset3) Xdr(xs *xdr.XdrState) {
	(*Uint64)(v).Xdr(xs)
}

type Mode3 Uint32

func (v *Mode3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(v).Xdr(xs)
}

type Count3 Uint32

func (v *Count3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(v).Xdr(xs)
}

type Nfsstat3 int32

func (v *Nfsstat3) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const NFS3_OK Nfsstat3 = 0
const NFS3ERR_PERM Nfsstat3 = 1
const NFS3ERR_NOENT Nfsstat3 = 2
const NFS3ERR_IO Nfsstat3 = 5
const NFS3ERR_NXIO Nfsstat3 = 6
const NFS3ERR_ACCES Nfsstat3 = 13
const NFS3ERR_EXIST Nfsstat3 = 17
const NFS3ERR_XDEV Nfsstat3 = 18
const NFS3ERR_NODEV Nfsstat3 = 19
const NFS3ERR_NOTDIR Nfsstat3 = 20
const NFS3ERR_ISDIR Nfsstat3 = 21
const NFS3ERR_INVAL Nfsstat3 = 22
const NFS3ERR_FBIG Nfsstat3 = 27
const NFS3ERR_NOSPC Nfsstat3 = 28
const NFS3ERR_ROFS Nfsstat3 = 30
const NFS3ERR_MLINK Nfsstat3 = 31
const NFS3ERR_NAMETOOLONG Nfsstat3 = 63
const NFS3ERR_NOTEMPTY Nfsstat3 = 66
const NFS3ERR_DQUOT Nfsstat3 = 69
const NFS3ERR_STALE Nfsstat3 = 70
const NFS3ERR_REMOTE Nfsstat3 = 71
const NFS3ERR_BADHANDLE Nfsstat3 = 10001
const NFS3ERR_NOT_SYNC Nfsstat3 = 10002
const NFS3ERR_BAD_COOKIE Nfsstat3 = 10003
const NFS3ERR_NOTSUPP Nfsstat3 = 10004
const NFS3ERR_TOOSMALL Nfsstat3 = 10005
const NFS3ERR_SERVERFAULT Nfsstat3 = 10006
const NFS3ERR_BADTYPE Nfsstat3 = 10007
const NFS3ERR_JUKEBOX Nfsstat3 = 10008

type Ftype3 int32

func (v *Ftype3) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const NF3REG Ftype3 = 1
const NF3DIR Ftype3 = 2
const NF3BLK Ftype3 = 3
const NF3CHR Ftype3 = 4
const NF3LNK Ftype3 = 5
const NF3SOCK Ftype3 = 6
const NF3FIFO Ftype3 = 7

type Specdata3 struct {
	Specdata1 Uint32
	Specdata2 Uint32
}

func (v *Specdata3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(&((v).Specdata1)).Xdr(xs)
	(*Uint32)(&((v).Specdata2)).Xdr(xs)
}

type Nfs_fh3 struct {
	Data []byte
}

func (v *Nfs_fh3) Xdr(xs *xdr.XdrState) {
	xdr.XdrVarArray(xs, NFS3_FHSIZE, (*[]byte)(&((v).Data)))
}

type Nfstime3 struct {
	Seconds  Uint32
	Nseconds Uint32
}

func (v *Nfstime3) Xdr(xs *xdr.XdrState) {
	(*Uint32)(&((v).Seconds)).Xdr(xs)
	(*Uint32)(&((v).Nseconds)).Xdr(xs)
}

type Fattr3 struct {
	Ftype  Ftype3
	Mode   Mode3
	Nlink  Uint32
	Uid    Uid3
	Gid    Gid3
	Size   Size3
	Used   Size3
	Rdev   Specdata3
	Fsid   Uint64
	Fileid Fileid3
	Atime  Nfstime3
	Mtime  Nfstime3
	Ctime  Nfstime3
}

func (v *Fattr3) Xdr(xs *xdr.XdrState) {
	(*Ftype3)(&((v).Ftype)).Xdr(xs)
	(*Mode3)(&((v).Mode)).Xdr(xs)
	(*Uint32)(&((v).Nlink)).Xdr(xs)
	(*Uid3)(&((v).Uid)).Xdr(xs)
	(*Gid3)(&((v).Gid)).Xdr(xs)
	(*Size3)(&((v).Size)).Xdr(xs)
	(*Size3)(&((v).Used)).Xdr(xs)
	(*Specdata3)(&((v).Rdev)).Xdr(xs)
	(*Uint64)(&((v).Fsid)).Xdr(xs)
	(*Fileid3)(&((v).Fileid)).Xdr(xs)
	(*Nfstime3)(&((v).Atime)).Xdr(xs)
	(*Nfstime3)(&((v).Mtime)).Xdr(xs)
	(*Nfstime3)(&((v).Ctime)).Xdr(xs)
}

type Post_op_attr struct {
	Attributes_follow bool
	Attributes        Fattr3
}

func (v *Post_op_attr) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Attributes_follow))
	switch (v).Attributes_follow {
	case true:
		(*Fattr3)(&((v).Attributes)).Xdr(xs)
	case false:
	}
}

type Wcc_attr struct {
	Size  Size3
	Mtime Nfstime3
	Ctime Nfstime3
}

func (v *Wcc_attr) Xdr(xs *xdr.XdrState) {
	(*Size3)(&((v).Size)).Xdr(xs)
	(*Nfstime3)(&((v).Mtime)).Xdr(xs)
	(*Nfstime3)(&((v).Ctime)).Xdr(xs)
}

type Pre_op_attr struct {
	Attributes_follow bool
	Attributes        Wcc_attr
}

func (v *Pre_op_attr) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Attributes_follow))
	switch (v).Attributes_follow {
	case true:
		(*Wcc_attr)(&((v).Attributes)).Xdr(xs)
	case false:
	}
}

type Wcc_data struct {
	Before Pre_op_attr
	After  Post_op_attr
}

func (v *Wcc_data) Xdr(xs *xdr.XdrState) {
	(*Pre_op_attr)(&((v).Before)).Xdr(xs)
	(*Post_op_attr)(&((v).After)).Xdr(xs)
}

type Post_op_fh3 struct {
	Handle_follows bool
	Handle         Nfs_fh3
}

func (v *Post_op_fh3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Handle_follows))
	switch (v).Handle_follows {
	case true:
		(*Nfs_fh3)(&((v).Handle)).Xdr(xs)
	case false:
	}
}

type Time_how int32

func (v *Time_how) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const DONT_CHANGE Time_how = 0
const SET_TO_SERVER_TIME Time_how = 1
const SET_TO_CLIENT_TIME Time_how = 2

type Set_mode3 struct {
	Set_it bool
	Mode   Mode3
}

func (v *Set_mode3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Set_it))
	switch (v).Set_it {
	case true:
		(*Mode3)(&((v).Mode)).Xdr(xs)
	default:
	}
}

type Set_uid3 struct {
	Set_it bool
	Uid    Uid3
}

func (v *Set_uid3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Set_it))
	switch (v).Set_it {
	case true:
		(*Uid3)(&((v).Uid)).Xdr(xs)
	default:
	}
}

type Set_gid3 struct {
	Set_it bool
	Gid    Gid3
}

func (v *Set_gid3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Set_it))
	switch (v).Set_it {
	case true:
		(*Gid3)(&((v).Gid)).Xdr(xs)
	default:
	}
}

type Set_size3 struct {
	Set_it bool
	Size   Size3
}

func (v *Set_size3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Set_it))
	switch (v).Set_it {
	case true:
		(*Size3)(&((v).Size)).Xdr(xs)
	default:
	}
}

type Set_atime struct {
	Set_it Time_how
	Atime  Nfstime3
}

func (v *Set_atime) Xdr(xs *xdr.XdrState) {
	(*Time_how)(&((v).Set_it)).Xdr(xs)
	switch (v).Set_it {
	case SET_TO_CLIENT_TIME:
		(*Nfstime3)(&((v).Atime)).Xdr(xs)
	default:
	}
}

type Set_mtime struct {
	Set_it Time_how
	Mtime  Nfstime3
}

func (v *Set_mtime) Xdr(xs *xdr.XdrState) {
	(*Time_how)(&((v).Set_it)).Xdr(xs)
	switch (v).Set_it {
	case SET_TO_CLIENT_TIME:
		(*Nfstime3)(&((v).Mtime)).Xdr(xs)
	default:
	}
}

type Sattr3 struct {
	Mode  Set_mode3
	Uid   Set_uid3
	Gid   Set_gid3
	Size  Set_size3
	Atime Set_atime
	Mtime Set_mtime
}

func (v *Sattr3) Xdr(xs *xdr.XdrState) {
	(*Set_mode3)(&((v).Mode)).Xdr(xs)
	(*Set_uid3)(&((v).Uid)).Xdr(xs)
	(*Set_gid3)(&((v).Gid)).Xdr(xs)
	(*Set_size3)(&((v).Size)).Xdr(xs)
	(*Set_atime)(&((v).Atime)).Xdr(xs)
	(*Set_mtime)(&((v).Mtime)).Xdr(xs)
}

type Diropargs3 struct {
	Dir  Nfs_fh3
	Name Filename3
}

func (v *Diropargs3) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Dir)).Xdr(xs)
	(*Filename3)(&((v).Name)).Xdr(xs)
}

const NFS_PROGRAM uint32 = 100003
const NFS_V3 uint32 = 3
const NFSPROC3_NULL uint32 = 0
const NFSPROC3_GETATTR uint32 = 1
const NFSPROC3_SETATTR uint32 = 2
const NFSPROC3_LOOKUP uint32 = 3
const NFSPROC3_ACCESS uint32 = 4
const NFSPROC3_READLINK uint32 = 5
const NFSPROC3_READ uint32 = 6
const NFSPROC3_WRITE uint32 = 7
const NFSPROC3_CREATE uint32 = 8
const NFSPROC3_MKDIR uint32 = 9
const NFSPROC3_SYMLINK uint32 = 10
const NFSPROC3_MKNOD uint32 = 11
const NFSPROC3_REMOVE uint32 = 12
const NFSPROC3_RMDIR uint32 = 13
const NFSPROC3_RENAME uint32 = 14
const NFSPROC3_LINK uint32 = 15
const NFSPROC3_READDIR uint32 = 16
const NFSPROC3_READDIRPLUS uint32 = 17
const NFSPROC3_FSSTAT uint32 = 18
const NFSPROC3_FSINFO uint32 = 19
const NFSPROC3_PATHCONF uint32 = 20
const NFSPROC3_COMMIT uint32 = 21

type GETATTR3args struct {
	Object Nfs_fh3
}

func (v *GETATTR3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Object)).Xdr(xs)
}

type GETATTR3resok struct {
	Obj_attributes Fattr3
}

func (v *GETATTR3resok) Xdr(xs *xdr.XdrState) {
	(*Fattr3)(&((v).Obj_attributes)).Xdr(xs)
}

type GETATTR3res struct {
	Status Nfsstat3
	Resok  GETATTR3resok
}

func (v *GETATTR3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*GETATTR3resok)(&((v).Resok)).Xdr(xs)
	default:
	}
}

type Sattrguard3 struct {
	Check     bool
	Obj_ctime Nfstime3
}

func (v *Sattrguard3) Xdr(xs *xdr.XdrState) {
	xdr.XdrBool(xs, &((v).Check))
	switch (v).Check {
	case true:
		(*Nfstime3)(&((v).Obj_ctime)).Xdr(xs)
	case false:
	}
}

type SETATTR3args struct {
	Object         Nfs_fh3
	New_attributes Sattr3
	Guard          Sattrguard3
}

func (v *SETATTR3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Object)).Xdr(xs)
	(*Sattr3)(&((v).New_attributes)).Xdr(xs)
	(*Sattrguard3)(&((v).Guard)).Xdr(xs)
}

type SETATTR3resok struct {
	Obj_wcc Wcc_data
}

func (v *SETATTR3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Obj_wcc)).Xdr(xs)
}

type SETATTR3resfail struct {
	Obj_wcc Wcc_data
}

func (v *SETATTR3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Obj_wcc)).Xdr(xs)
}

type SETATTR3res struct {
	Status  Nfsstat3
	Resok   SETATTR3resok
	Resfail SETATTR3resfail
}

func (v *SETATTR3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*SETATTR3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*SETATTR3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type LOOKUP3args struct {
	What Diropargs3
}

func (v *LOOKUP3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).What)).Xdr(xs)
}

type LOOKUP3resok struct {
	Object         Nfs_fh3
	Obj_attributes Post_op_attr
	Dir_attributes Post_op_attr
}

func (v *LOOKUP3resok) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Object)).Xdr(xs)
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
}

type LOOKUP3resfail struct {
	Dir_attributes Post_op_attr
}

func (v *LOOKUP3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
}

type LOOKUP3res struct {
	Status  Nfsstat3
	Resok   LOOKUP3resok
	Resfail LOOKUP3resfail
}

func (v *LOOKUP3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*LOOKUP3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*LOOKUP3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

const ACCESS3_READ = 0x0001
const ACCESS3_LOOKUP = 0x0002
const ACCESS3_MODIFY = 0x0004
const ACCESS3_EXTEND = 0x0008
const ACCESS3_DELETE = 0x0010
const ACCESS3_EXECUTE = 0x0020

type ACCESS3args struct {
	Object Nfs_fh3
	Access Uint32
}

func (v *ACCESS3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Object)).Xdr(xs)
	(*Uint32)(&((v).Access)).Xdr(xs)
}

type ACCESS3resok struct {
	Obj_attributes Post_op_attr
	Access         Uint32
}

func (v *ACCESS3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Uint32)(&((v).Access)).Xdr(xs)
}

type ACCESS3resfail struct {
	Obj_attributes Post_op_attr
}

func (v *ACCESS3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
}

type ACCESS3res struct {
	Status  Nfsstat3
	Resok   ACCESS3resok
	Resfail ACCESS3resfail
}

func (v *ACCESS3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*ACCESS3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*ACCESS3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type READLINK3args struct {
	Symlink Nfs_fh3
}

func (v *READLINK3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Symlink)).Xdr(xs)
}

type READLINK3resok struct {
	Symlink_attributes Post_op_attr
	Data               Nfspath3
}

func (v *READLINK3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Symlink_attributes)).Xdr(xs)
	(*Nfspath3)(&((v).Data)).Xdr(xs)
}

type READLINK3resfail struct {
	Symlink_attributes Post_op_attr
}

func (v *READLINK3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Symlink_attributes)).Xdr(xs)
}

type READLINK3res struct {
	Status  Nfsstat3
	Resok   READLINK3resok
	Resfail READLINK3resfail
}

func (v *READLINK3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*READLINK3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*READLINK3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type READ3args struct {
	File   Nfs_fh3
	Offset Offset3
	Count  Count3
}

func (v *READ3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).File)).Xdr(xs)
	(*Offset3)(&((v).Offset)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
}

type READ3resok struct {
	File_attributes Post_op_attr
	Count           Count3
	Eof             bool
	Data            []byte
}

func (v *READ3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).File_attributes)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
	xdr.XdrBool(xs, &((v).Eof))
	xdr.XdrVarArray(xs, -1, (*[]byte)(&((v).Data)))
}

type READ3resfail struct {
	File_attributes Post_op_attr
}

func (v *READ3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).File_attributes)).Xdr(xs)
}

type READ3res struct {
	Status  Nfsstat3
	Resok   READ3resok
	Resfail READ3resfail
}

func (v *READ3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*READ3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*READ3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type Stable_how int32

func (v *Stable_how) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const UNSTABLE Stable_how = 0
const DATA_SYNC Stable_how = 1
const FILE_SYNC Stable_how = 2

type WRITE3args struct {
	File   Nfs_fh3
	Offset Offset3
	Count  Count3
	Stable Stable_how
	Data   []byte
}

func (v *WRITE3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).File)).Xdr(xs)
	(*Offset3)(&((v).Offset)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
	(*Stable_how)(&((v).Stable)).Xdr(xs)
	xdr.XdrVarArray(xs, -1, (*[]byte)(&((v).Data)))
}

type WRITE3resok struct {
	File_wcc  Wcc_data
	Count     Count3
	Committed Stable_how
	Verf      Writeverf3
}

func (v *WRITE3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).File_wcc)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
	(*Stable_how)(&((v).Committed)).Xdr(xs)
	(*Writeverf3)(&((v).Verf)).Xdr(xs)
}

type WRITE3resfail struct {
	File_wcc Wcc_data
}

func (v *WRITE3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).File_wcc)).Xdr(xs)
}

type WRITE3res struct {
	Status  Nfsstat3
	Resok   WRITE3resok
	Resfail WRITE3resfail
}

func (v *WRITE3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*WRITE3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*WRITE3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type Createmode3 int32

func (v *Createmode3) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const UNCHECKED Createmode3 = 0
const GUARDED Createmode3 = 1
const EXCLUSIVE Createmode3 = 2

type Createhow3 struct {
	Mode           Createmode3
	Obj_attributes Sattr3
	Verf           Createverf3
}

func (v *Createhow3) Xdr(xs *xdr.XdrState) {
	(*Createmode3)(&((v).Mode)).Xdr(xs)
	switch (v).Mode {
	case UNCHECKED:
		fallthrough
	case GUARDED:
		(*Sattr3)(&((v).Obj_attributes)).Xdr(xs)
	case EXCLUSIVE:
		(*Createverf3)(&((v).Verf)).Xdr(xs)
	}
}

type CREATE3args struct {
	Where Diropargs3
	How   Createhow3
}

func (v *CREATE3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Where)).Xdr(xs)
	(*Createhow3)(&((v).How)).Xdr(xs)
}

type CREATE3resok struct {
	Obj            Post_op_fh3
	Obj_attributes Post_op_attr
	Dir_wcc        Wcc_data
}

func (v *CREATE3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_fh3)(&((v).Obj)).Xdr(xs)
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type CREATE3resfail struct {
	Dir_wcc Wcc_data
}

func (v *CREATE3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type CREATE3res struct {
	Status  Nfsstat3
	Resok   CREATE3resok
	Resfail CREATE3resfail
}

func (v *CREATE3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*CREATE3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*CREATE3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type MKDIR3args struct {
	Where      Diropargs3
	Attributes Sattr3
}

func (v *MKDIR3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Where)).Xdr(xs)
	(*Sattr3)(&((v).Attributes)).Xdr(xs)
}

type MKDIR3resok struct {
	Obj            Post_op_fh3
	Obj_attributes Post_op_attr
	Dir_wcc        Wcc_data
}

func (v *MKDIR3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_fh3)(&((v).Obj)).Xdr(xs)
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type MKDIR3resfail struct {
	Dir_wcc Wcc_data
}

func (v *MKDIR3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type MKDIR3res struct {
	Status  Nfsstat3
	Resok   MKDIR3resok
	Resfail MKDIR3resfail
}

func (v *MKDIR3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*MKDIR3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*MKDIR3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type Symlinkdata3 struct {
	Symlink_attributes Sattr3
	Symlink_data       Nfspath3
}

func (v *Symlinkdata3) Xdr(xs *xdr.XdrState) {
	(*Sattr3)(&((v).Symlink_attributes)).Xdr(xs)
	(*Nfspath3)(&((v).Symlink_data)).Xdr(xs)
}

type SYMLINK3args struct {
	Where   Diropargs3
	Symlink Symlinkdata3
}

func (v *SYMLINK3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Where)).Xdr(xs)
	(*Symlinkdata3)(&((v).Symlink)).Xdr(xs)
}

type SYMLINK3resok struct {
	Obj            Post_op_fh3
	Obj_attributes Post_op_attr
	Dir_wcc        Wcc_data
}

func (v *SYMLINK3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_fh3)(&((v).Obj)).Xdr(xs)
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type SYMLINK3resfail struct {
	Dir_wcc Wcc_data
}

func (v *SYMLINK3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type SYMLINK3res struct {
	Status  Nfsstat3
	Resok   SYMLINK3resok
	Resfail SYMLINK3resfail
}

func (v *SYMLINK3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*SYMLINK3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*SYMLINK3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type Devicedata3 struct {
	Dev_attributes Sattr3
	Spec           Specdata3
}

func (v *Devicedata3) Xdr(xs *xdr.XdrState) {
	(*Sattr3)(&((v).Dev_attributes)).Xdr(xs)
	(*Specdata3)(&((v).Spec)).Xdr(xs)
}

type Mknoddata3 struct {
	Ftype           Ftype3
	Device          Devicedata3
	Pipe_attributes Sattr3
}

func (v *Mknoddata3) Xdr(xs *xdr.XdrState) {
	(*Ftype3)(&((v).Ftype)).Xdr(xs)
	switch (v).Ftype {
	case NF3CHR:
		fallthrough
	case NF3BLK:
		(*Devicedata3)(&((v).Device)).Xdr(xs)
	case NF3SOCK:
		fallthrough
	case NF3FIFO:
		(*Sattr3)(&((v).Pipe_attributes)).Xdr(xs)
	default:
	}
}

type MKNOD3args struct {
	Where Diropargs3
	What  Mknoddata3
}

func (v *MKNOD3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Where)).Xdr(xs)
	(*Mknoddata3)(&((v).What)).Xdr(xs)
}

type MKNOD3resok struct {
	Obj            Post_op_fh3
	Obj_attributes Post_op_attr
	Dir_wcc        Wcc_data
}

func (v *MKNOD3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_fh3)(&((v).Obj)).Xdr(xs)
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type MKNOD3resfail struct {
	Dir_wcc Wcc_data
}

func (v *MKNOD3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type MKNOD3res struct {
	Status  Nfsstat3
	Resok   MKNOD3resok
	Resfail MKNOD3resfail
}

func (v *MKNOD3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*MKNOD3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*MKNOD3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type REMOVE3args struct {
	Object Diropargs3
}

func (v *REMOVE3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Object)).Xdr(xs)
}

type REMOVE3resok struct {
	Dir_wcc Wcc_data
}

func (v *REMOVE3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type REMOVE3resfail struct {
	Dir_wcc Wcc_data
}

func (v *REMOVE3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type REMOVE3res struct {
	Status  Nfsstat3
	Resok   REMOVE3resok
	Resfail REMOVE3resfail
}

func (v *REMOVE3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*REMOVE3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*REMOVE3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type RMDIR3args struct {
	Object Diropargs3
}

func (v *RMDIR3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).Object)).Xdr(xs)
}

type RMDIR3resok struct {
	Dir_wcc Wcc_data
}

func (v *RMDIR3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type RMDIR3resfail struct {
	Dir_wcc Wcc_data
}

func (v *RMDIR3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Dir_wcc)).Xdr(xs)
}

type RMDIR3res struct {
	Status  Nfsstat3
	Resok   RMDIR3resok
	Resfail RMDIR3resfail
}

func (v *RMDIR3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*RMDIR3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*RMDIR3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type RENAME3args struct {
	From Diropargs3
	To   Diropargs3
}

func (v *RENAME3args) Xdr(xs *xdr.XdrState) {
	(*Diropargs3)(&((v).From)).Xdr(xs)
	(*Diropargs3)(&((v).To)).Xdr(xs)
}

type RENAME3resok struct {
	Fromdir_wcc Wcc_data
	Todir_wcc   Wcc_data
}

func (v *RENAME3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Fromdir_wcc)).Xdr(xs)
	(*Wcc_data)(&((v).Todir_wcc)).Xdr(xs)
}

type RENAME3resfail struct {
	Fromdir_wcc Wcc_data
	Todir_wcc   Wcc_data
}

func (v *RENAME3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).Fromdir_wcc)).Xdr(xs)
	(*Wcc_data)(&((v).Todir_wcc)).Xdr(xs)
}

type RENAME3res struct {
	Status  Nfsstat3
	Resok   RENAME3resok
	Resfail RENAME3resfail
}

func (v *RENAME3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*RENAME3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*RENAME3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type LINK3args struct {
	File Nfs_fh3
	Link Diropargs3
}

func (v *LINK3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).File)).Xdr(xs)
	(*Diropargs3)(&((v).Link)).Xdr(xs)
}

type LINK3resok struct {
	File_attributes Post_op_attr
	Linkdir_wcc     Wcc_data
}

func (v *LINK3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).File_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Linkdir_wcc)).Xdr(xs)
}

type LINK3resfail struct {
	File_attributes Post_op_attr
	Linkdir_wcc     Wcc_data
}

func (v *LINK3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).File_attributes)).Xdr(xs)
	(*Wcc_data)(&((v).Linkdir_wcc)).Xdr(xs)
}

type LINK3res struct {
	Status  Nfsstat3
	Resok   LINK3resok
	Resfail LINK3resfail
}

func (v *LINK3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*LINK3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*LINK3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type READDIR3args struct {
	Dir        Nfs_fh3
	Cookie     Cookie3
	Cookieverf Cookieverf3
	Count      Count3
}

func (v *READDIR3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Dir)).Xdr(xs)
	(*Cookie3)(&((v).Cookie)).Xdr(xs)
	(*Cookieverf3)(&((v).Cookieverf)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
}

type Entry3 struct {
	Fileid    Fileid3
	Name      Filename3
	Cookie    Cookie3
	Nextentry struct{ P *Entry3 }
}

func (v *Entry3) Xdr(xs *xdr.XdrState) {
	(*Fileid3)(&((v).Fileid)).Xdr(xs)
	(*Filename3)(&((v).Name)).Xdr(xs)
	(*Cookie3)(&((v).Cookie)).Xdr(xs)
	if xs.Encoding() {
		opted := (&((v).Nextentry)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Entry3)((&((v).Nextentry)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Nextentry)).P = new(Entry3)
			(*Entry3)((&((v).Nextentry)).P).Xdr(xs)
		}
	}
}

type Dirlist3 struct {
	Entries struct{ P *Entry3 }
	Eof     bool
}

func (v *Dirlist3) Xdr(xs *xdr.XdrState) {
	if xs.Encoding() {
		opted := (&((v).Entries)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Entry3)((&((v).Entries)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Entries)).P = new(Entry3)
			(*Entry3)((&((v).Entries)).P).Xdr(xs)
		}
	}
	xdr.XdrBool(xs, &((v).Eof))
}

type READDIR3resok struct {
	Dir_attributes Post_op_attr
	Cookieverf     Cookieverf3
	Reply          Dirlist3
}

func (v *READDIR3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
	(*Cookieverf3)(&((v).Cookieverf)).Xdr(xs)
	(*Dirlist3)(&((v).Reply)).Xdr(xs)
}

type READDIR3resfail struct {
	Dir_attributes Post_op_attr
}

func (v *READDIR3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
}

type READDIR3res struct {
	Status  Nfsstat3
	Resok   READDIR3resok
	Resfail READDIR3resfail
}

func (v *READDIR3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*READDIR3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*READDIR3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type READDIRPLUS3args struct {
	Dir        Nfs_fh3
	Cookie     Cookie3
	Cookieverf Cookieverf3
	Dircount   Count3
	Maxcount   Count3
}

func (v *READDIRPLUS3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Dir)).Xdr(xs)
	(*Cookie3)(&((v).Cookie)).Xdr(xs)
	(*Cookieverf3)(&((v).Cookieverf)).Xdr(xs)
	(*Count3)(&((v).Dircount)).Xdr(xs)
	(*Count3)(&((v).Maxcount)).Xdr(xs)
}

type Entryplus3 struct {
	Fileid          Fileid3
	Name            Filename3
	Cookie          Cookie3
	Name_attributes Post_op_attr
	Name_handle     Post_op_fh3
	Nextentry       struct{ P *Entryplus3 }
}

func (v *Entryplus3) Xdr(xs *xdr.XdrState) {
	(*Fileid3)(&((v).Fileid)).Xdr(xs)
	(*Filename3)(&((v).Name)).Xdr(xs)
	(*Cookie3)(&((v).Cookie)).Xdr(xs)
	(*Post_op_attr)(&((v).Name_attributes)).Xdr(xs)
	(*Post_op_fh3)(&((v).Name_handle)).Xdr(xs)
	if xs.Encoding() {
		opted := (&((v).Nextentry)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Entryplus3)((&((v).Nextentry)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Nextentry)).P = new(Entryplus3)
			(*Entryplus3)((&((v).Nextentry)).P).Xdr(xs)
		}
	}
}

type Dirlistplus3 struct {
	Entries struct{ P *Entryplus3 }
	Eof     bool
}

func (v *Dirlistplus3) Xdr(xs *xdr.XdrState) {
	if xs.Encoding() {
		opted := (&((v).Entries)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Entryplus3)((&((v).Entries)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Entries)).P = new(Entryplus3)
			(*Entryplus3)((&((v).Entries)).P).Xdr(xs)
		}
	}
	xdr.XdrBool(xs, &((v).Eof))
}

type READDIRPLUS3resok struct {
	Dir_attributes Post_op_attr
	Cookieverf     Cookieverf3
	Reply          Dirlistplus3
}

func (v *READDIRPLUS3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
	(*Cookieverf3)(&((v).Cookieverf)).Xdr(xs)
	(*Dirlistplus3)(&((v).Reply)).Xdr(xs)
}

type READDIRPLUS3resfail struct {
	Dir_attributes Post_op_attr
}

func (v *READDIRPLUS3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Dir_attributes)).Xdr(xs)
}

type READDIRPLUS3res struct {
	Status  Nfsstat3
	Resok   READDIRPLUS3resok
	Resfail READDIRPLUS3resfail
}

func (v *READDIRPLUS3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*READDIRPLUS3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*READDIRPLUS3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type FSSTAT3args struct {
	Fsroot Nfs_fh3
}

func (v *FSSTAT3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Fsroot)).Xdr(xs)
}

type FSSTAT3resok struct {
	Obj_attributes Post_op_attr
	Tbytes         Size3
	Fbytes         Size3
	Abytes         Size3
	Tfiles         Size3
	Ffiles         Size3
	Afiles         Size3
	Invarsec       Uint32
}

func (v *FSSTAT3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Size3)(&((v).Tbytes)).Xdr(xs)
	(*Size3)(&((v).Fbytes)).Xdr(xs)
	(*Size3)(&((v).Abytes)).Xdr(xs)
	(*Size3)(&((v).Tfiles)).Xdr(xs)
	(*Size3)(&((v).Ffiles)).Xdr(xs)
	(*Size3)(&((v).Afiles)).Xdr(xs)
	(*Uint32)(&((v).Invarsec)).Xdr(xs)
}

type FSSTAT3resfail struct {
	Obj_attributes Post_op_attr
}

func (v *FSSTAT3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
}

type FSSTAT3res struct {
	Status  Nfsstat3
	Resok   FSSTAT3resok
	Resfail FSSTAT3resfail
}

func (v *FSSTAT3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*FSSTAT3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*FSSTAT3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

const FSF3_LINK = 0x0001
const FSF3_SYMLINK = 0x0002
const FSF3_HOMOGENEOUS = 0x0008
const FSF3_CANSETTIME = 0x0010

type FSINFO3args struct {
	Fsroot Nfs_fh3
}

func (v *FSINFO3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Fsroot)).Xdr(xs)
}

type FSINFO3resok struct {
	Obj_attributes Post_op_attr
	Rtmax          Uint32
	Rtpref         Uint32
	Rtmult         Uint32
	Wtmax          Uint32
	Wtpref         Uint32
	Wtmult         Uint32
	Dtpref         Uint32
	Maxfilesize    Size3
	Time_delta     Nfstime3
	Properties     Uint32
}

func (v *FSINFO3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Uint32)(&((v).Rtmax)).Xdr(xs)
	(*Uint32)(&((v).Rtpref)).Xdr(xs)
	(*Uint32)(&((v).Rtmult)).Xdr(xs)
	(*Uint32)(&((v).Wtmax)).Xdr(xs)
	(*Uint32)(&((v).Wtpref)).Xdr(xs)
	(*Uint32)(&((v).Wtmult)).Xdr(xs)
	(*Uint32)(&((v).Dtpref)).Xdr(xs)
	(*Size3)(&((v).Maxfilesize)).Xdr(xs)
	(*Nfstime3)(&((v).Time_delta)).Xdr(xs)
	(*Uint32)(&((v).Properties)).Xdr(xs)
}

type FSINFO3resfail struct {
	Obj_attributes Post_op_attr
}

func (v *FSINFO3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
}

type FSINFO3res struct {
	Status  Nfsstat3
	Resok   FSINFO3resok
	Resfail FSINFO3resfail
}

func (v *FSINFO3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*FSINFO3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*FSINFO3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type PATHCONF3args struct {
	Object Nfs_fh3
}

func (v *PATHCONF3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).Object)).Xdr(xs)
}

type PATHCONF3resok struct {
	Obj_attributes   Post_op_attr
	Linkmax          Uint32
	Name_max         Uint32
	No_trunc         bool
	Chown_restricted bool
	Case_insensitive bool
	Case_preserving  bool
}

func (v *PATHCONF3resok) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
	(*Uint32)(&((v).Linkmax)).Xdr(xs)
	(*Uint32)(&((v).Name_max)).Xdr(xs)
	xdr.XdrBool(xs, &((v).No_trunc))
	xdr.XdrBool(xs, &((v).Chown_restricted))
	xdr.XdrBool(xs, &((v).Case_insensitive))
	xdr.XdrBool(xs, &((v).Case_preserving))
}

type PATHCONF3resfail struct {
	Obj_attributes Post_op_attr
}

func (v *PATHCONF3resfail) Xdr(xs *xdr.XdrState) {
	(*Post_op_attr)(&((v).Obj_attributes)).Xdr(xs)
}

type PATHCONF3res struct {
	Status  Nfsstat3
	Resok   PATHCONF3resok
	Resfail PATHCONF3resfail
}

func (v *PATHCONF3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*PATHCONF3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*PATHCONF3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

type COMMIT3args struct {
	File   Nfs_fh3
	Offset Offset3
	Count  Count3
}

func (v *COMMIT3args) Xdr(xs *xdr.XdrState) {
	(*Nfs_fh3)(&((v).File)).Xdr(xs)
	(*Offset3)(&((v).Offset)).Xdr(xs)
	(*Count3)(&((v).Count)).Xdr(xs)
}

type COMMIT3resok struct {
	File_wcc Wcc_data
	Verf     Writeverf3
}

func (v *COMMIT3resok) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).File_wcc)).Xdr(xs)
	(*Writeverf3)(&((v).Verf)).Xdr(xs)
}

type COMMIT3resfail struct {
	File_wcc Wcc_data
}

func (v *COMMIT3resfail) Xdr(xs *xdr.XdrState) {
	(*Wcc_data)(&((v).File_wcc)).Xdr(xs)
}

type COMMIT3res struct {
	Status  Nfsstat3
	Resok   COMMIT3resok
	Resfail COMMIT3resfail
}

func (v *COMMIT3res) Xdr(xs *xdr.XdrState) {
	(*Nfsstat3)(&((v).Status)).Xdr(xs)
	switch (v).Status {
	case NFS3_OK:
		(*COMMIT3resok)(&((v).Resok)).Xdr(xs)
	default:
		(*COMMIT3resfail)(&((v).Resfail)).Xdr(xs)
	}
}

const MNTPATHLEN3 = 1024
const MNTNAMLEN3 = 255
const FHSIZE3 = 64

type Fhandle3 []byte

func (v *Fhandle3) Xdr(xs *xdr.XdrState) {
	xdr.XdrVarArray(xs, FHSIZE3, (*[]byte)(v))
}

type Dirpath3 string

func (v *Dirpath3) Xdr(xs *xdr.XdrState) {
	xdr.XdrString(xs, MNTPATHLEN3, (*string)(v))
}

type Name3 string

func (v *Name3) Xdr(xs *xdr.XdrState) {
	xdr.XdrString(xs, MNTNAMLEN3, (*string)(v))
}

type Mountstat3 int32

func (v *Mountstat3) Xdr(xs *xdr.XdrState) {
	xdr.XdrS32(xs, (*int32)(v))
}

const MNT3_OK Mountstat3 = 0
const MNT3ERR_PERM Mountstat3 = 1
const MNT3ERR_NOENT Mountstat3 = 2
const MNT3ERR_IO Mountstat3 = 5
const MNT3ERR_ACCES Mountstat3 = 13
const MNT3ERR_NOTDIR Mountstat3 = 20
const MNT3ERR_INVAL Mountstat3 = 22
const MNT3ERR_NAMETOOLONG Mountstat3 = 63
const MNT3ERR_NOTSUPP Mountstat3 = 10004
const MNT3ERR_SERVERFAULT Mountstat3 = 10006
const MOUNT_PROGRAM uint32 = 100005
const MOUNT_V3 uint32 = 3
const MOUNTPROC3_NULL uint32 = 0
const MOUNTPROC3_MNT uint32 = 1
const MOUNTPROC3_DUMP uint32 = 2
const MOUNTPROC3_UMNT uint32 = 3
const MOUNTPROC3_UMNTALL uint32 = 4
const MOUNTPROC3_EXPORT uint32 = 5

type Mountres3_ok struct {
	Fhandle      Fhandle3
	Auth_flavors []int32
}

func (v *Mountres3_ok) Xdr(xs *xdr.XdrState) {
	(*Fhandle3)(&((v).Fhandle)).Xdr(xs)
	{
		var __arraysz uint32
		xs.EncodingSetSize(&__arraysz, len(*&((v).Auth_flavors)))
		xdr.XdrU32(xs, (*uint32)(&__arraysz))

		if xs.Decoding() {
			*&((v).Auth_flavors) = make([]int32, __arraysz)
		}
		for i := uint64(0); i < uint64(__arraysz); i++ {
			xdr.XdrS32(xs, (*int32)(&((*(&((v).Auth_flavors)))[i])))

		}
	}
}

type Mountres3 struct {
	Fhs_status Mountstat3
	Mountinfo  Mountres3_ok
}

func (v *Mountres3) Xdr(xs *xdr.XdrState) {
	(*Mountstat3)(&((v).Fhs_status)).Xdr(xs)
	switch (v).Fhs_status {
	case MNT3_OK:
		(*Mountres3_ok)(&((v).Mountinfo)).Xdr(xs)
	default:
	}
}

type Mount3 struct {
	Ml_hostname  Name3
	Ml_directory Dirpath3
	Ml_next      struct{ P *Mount3 }
}

func (v *Mount3) Xdr(xs *xdr.XdrState) {
	(*Name3)(&((v).Ml_hostname)).Xdr(xs)
	(*Dirpath3)(&((v).Ml_directory)).Xdr(xs)
	if xs.Encoding() {
		opted := (&((v).Ml_next)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Mount3)((&((v).Ml_next)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Ml_next)).P = new(Mount3)
			(*Mount3)((&((v).Ml_next)).P).Xdr(xs)
		}
	}
}

type Mountopt3 struct{ P *Mount3 }

func (v *Mountopt3) Xdr(xs *xdr.XdrState) {
	if xs.Encoding() {
		opted := (v).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Mount3)((v).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(v).P = new(Mount3)
			(*Mount3)((v).P).Xdr(xs)
		}
	}
}

type Groups3 struct {
	Gr_name Name3
	Gr_next struct{ P *Groups3 }
}

func (v *Groups3) Xdr(xs *xdr.XdrState) {
	(*Name3)(&((v).Gr_name)).Xdr(xs)
	if xs.Encoding() {
		opted := (&((v).Gr_next)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Groups3)((&((v).Gr_next)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Gr_next)).P = new(Groups3)
			(*Groups3)((&((v).Gr_next)).P).Xdr(xs)
		}
	}
}

type Exports3 struct {
	Ex_dir    Dirpath3
	Ex_groups struct{ P *Groups3 }
	Ex_next   struct{ P *Exports3 }
}

func (v *Exports3) Xdr(xs *xdr.XdrState) {
	(*Dirpath3)(&((v).Ex_dir)).Xdr(xs)
	if xs.Encoding() {
		opted := (&((v).Ex_groups)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Groups3)((&((v).Ex_groups)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Ex_groups)).P = new(Groups3)
			(*Groups3)((&((v).Ex_groups)).P).Xdr(xs)
		}
	}
	if xs.Encoding() {
		opted := (&((v).Ex_next)).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Exports3)((&((v).Ex_next)).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(&((v).Ex_next)).P = new(Exports3)
			(*Exports3)((&((v).Ex_next)).P).Xdr(xs)
		}
	}
}

type Exportsopt3 struct{ P *Exports3 }

func (v *Exportsopt3) Xdr(xs *xdr.XdrState) {
	if xs.Encoding() {
		opted := (v).P != nil
		xdr.XdrBool(xs, &opted)
		if opted {
			(*Exports3)((v).P).Xdr(xs)
		}
	}
	if xs.Decoding() {
		var opted bool
		xdr.XdrBool(xs, &opted)
		if opted {
			(v).P = new(Exports3)
			(*Exports3)((v).P).Xdr(xs)
		}
	}
}
