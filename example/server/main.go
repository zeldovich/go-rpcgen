package main

import (
	"fmt"
	"net"

	"github.com/zeldovich/go-rpcgen/rfc1057"
	"github.com/zeldovich/go-rpcgen/rfc1813"
	"github.com/zeldovich/go-rpcgen/xdr"
)

func pmap_set_unset(prog, vers, port uint32, setit bool) bool {
	var cred rfc1057.Opaque_auth
	cred.Flavor = rfc1057.AUTH_NONE

	pmapc, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", rfc1057.PMAP_PORT))
	if err != nil {
		panic(err)
	}
	defer pmapc.Close()
	pmap := rfc1057.MakeClient(pmapc, rfc1057.PMAP_PROG, rfc1057.PMAP_VERS)

	arg := rfc1057.Mapping{
		Prog: prog,
		Vers: vers,
		Prot: rfc1057.IPPROTO_TCP,
		Port: port,
	}

	var res xdr.Bool
	var proc uint32
	if setit {
		proc = rfc1057.PMAPPROC_SET
	} else {
		proc = rfc1057.PMAPPROC_UNSET
	}

	err = pmap.Call(proc, cred, cred, &arg, &res)
	if err != nil {
		panic(err)
	}

	return bool(res)
}

func main() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	port := uint32(listener.Addr().(*net.TCPAddr).Port)
	ok := pmap_set_unset(rfc1813.NFS_PROGRAM, rfc1813.NFS_V3, port, true)
	if !ok {
		panic("Could not set pmap mapping")
	}

	defer pmap_set_unset(rfc1813.NFS_PROGRAM, rfc1813.NFS_V3, port, false)
}
