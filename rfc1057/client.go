package rfc1057

import (
	"fmt"
	"io"

	"github.com/zeldovich/go-rpcgen/xdr"
)

type Client struct {
	rcv *xdr.XdrState
	snd *xdr.XdrState
	xid uint32
}

func MakeClient(rw io.ReadWriter) *Client {
	return &Client{
		rcv: xdr.MakeReader(rw),
		snd: xdr.MakeWriter(rw),
		xid: 0,
	}
}

func (c *Client) Call(prog, vers, proc uint32, cred, verf Opaque_auth, args xdr.Xdrable, resp xdr.Xdrable) error {
	c.xid++

	var req Rpc_msg
	req.Xid = c.xid
	req.Body.Mtype = CALL
	req.Body.Cbody.Rpcvers = 2
	req.Body.Cbody.Prog = prog
	req.Body.Cbody.Vers = vers
	req.Body.Cbody.Proc = proc
	req.Body.Cbody.Cred = cred
	req.Body.Cbody.Verf = verf

	req.Xdr(c.snd)
	err := c.snd.Error()
	if err != nil {
		return err
	}

	args.Xdr(c.snd)
	err = c.snd.Error()
	if err != nil {
		return err
	}

	var res Rpc_msg
	res.Xdr(c.rcv)
	err = c.rcv.Error()
	if err != nil {
		return err
	}

	if res.Xid != req.Xid {
		return fmt.Errorf("xid mismatch: %d != %d", res.Xid, req.Xid)
	}

	if res.Body.Mtype != REPLY {
		return fmt.Errorf("expected REPLY, got %d", res.Body.Mtype)
	}

	if res.Body.Rbody.Stat != MSG_ACCEPTED {
		return fmt.Errorf("MSG_DENIED stat %d", res.Body.Rbody.Rreply.Stat)
	}

	if res.Body.Rbody.Areply.Reply_data.Stat != SUCCESS {
		return fmt.Errorf("accept_stat %d", res.Body.Rbody.Areply.Reply_data.Stat)
	}

	resp.Xdr(c.rcv)
	err = c.rcv.Error()
	if err != nil {
		return err
	}

	return nil
}
