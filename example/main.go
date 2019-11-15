package main

import (
	"fmt"
	"net"

	"github.com/zeldovich/go-rpcgen/rfc1057"
	"github.com/zeldovich/go-rpcgen/rfc4506"
	"github.com/zeldovich/go-rpcgen/xdr"
)

func pmap_client(host string, prog, vers uint32) *rfc1057.Client {
	var cred rfc1057.Opaque_auth
	cred.Flavor = rfc1057.AUTH_NONE

	pmapc, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, rfc1057.PMAP_PORT))
	if err != nil {
		panic(err)
	}
	defer pmapc.Close()
	pmap := rfc1057.MakeClient(pmapc, rfc1057.PMAP_PROG, rfc1057.PMAP_VERS)

	arg := rfc1057.Mapping{
		Prog: prog,
		Vers: vers,
		Prot: rfc1057.IPPROTO_TCP,
	}
	var res xdr.Uint32
	err = pmap.Call(rfc1057.PMAPPROC_GETPORT, cred, cred, &arg, &res)
	if err != nil {
		panic(err)
	}

	svcc, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, res))
	if err != nil {
		panic(err)
	}
	return rfc1057.MakeClient(svcc, prog, vers)
}

func lookup(c *rfc1057.Client, cred, verf rfc1057.Opaque_auth, fh rfc4506.Nfs_fh3, name string) rfc4506.Nfs_fh3 {
	var arg rfc4506.LOOKUP3args
	var res rfc4506.LOOKUP3res
	arg.What.Dir = fh
	arg.What.Name = rfc4506.Filename3(name)
	err := c.Call(rfc4506.NFSPROC3_LOOKUP, cred, verf, &arg, &res)
	if err != nil {
		panic(err)
	}

	if res.Status != rfc4506.NFS3_OK {
		panic(fmt.Sprintf("lookup status %d", res.Status))
	}

	return res.Resok.Object
}

func main() {
	var err error

	var unix rfc1057.Auth_unix
	var cred_unix rfc1057.Opaque_auth
	cred_unix.Flavor = rfc1057.AUTH_UNIX
	cred_unix.Body, err = xdr.EncodeBuf(&unix)
	if err != nil {
		panic(err)
	}

	var cred_none rfc1057.Opaque_auth
	cred_none.Flavor = rfc1057.AUTH_NONE

	mnt := pmap_client("localhost", rfc4506.MOUNT_PROGRAM, rfc4506.MOUNT_V3)

	arg := rfc4506.Dirpath3("/")
	var res rfc4506.Mountres3
	err = mnt.Call(rfc4506.MOUNTPROC3_MNT, cred_none, cred_none, &arg, &res)
	if err != nil {
		panic(err)
	}

	if res.Fhs_status != rfc4506.MNT3_OK {
		panic(fmt.Sprintf("mount status %d", res.Fhs_status))
	}

	var root_fh rfc4506.Nfs_fh3
	root_fh.Data = res.Mountinfo.Fhandle

	for _, flavor := range res.Mountinfo.Auth_flavors {
		fmt.Printf("flavor %d\n", flavor)
	}

	fmt.Printf("root fh %v\n", root_fh)

	nfs := pmap_client("localhost", rfc4506.NFS_PROGRAM, rfc4506.NFS_V3)

	foo := lookup(nfs, cred_unix, cred_none, root_fh, "foo")
	bar := lookup(nfs, cred_unix, cred_none, foo, "bar")

	fmt.Printf("bar = %v\n", bar)

	var arg1 rfc4506.RENAME3args
	var res1 rfc4506.RENAME3res
	arg1.From.Dir = root_fh
	arg1.From.Name = "foo"
	arg1.To.Dir = bar
	arg1.To.Name = "foonew"
	err = nfs.Call(rfc4506.NFSPROC3_RENAME, cred_unix, cred_none, &arg1, &res1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("res = %v\n", res1)
}
