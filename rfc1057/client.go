package rfc1057

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeldovich/go-rpcgen/xdr"
)

type Client struct {
	rw  io.ReadWriter
	xid uint32
}

func MakeClient(rw io.ReadWriter) *Client {
	return &Client{
		rw:  rw,
		xid: 0,
	}
}

type rwBuffer struct {
	buf []byte
}

func (rw *rwBuffer) Write(data []byte) (n int, err error) {
	rw.buf = append(rw.buf, data...)
	return len(data), nil
}

func (rw *rwBuffer) Read(buf []byte) (n int, err error) {
	copy(buf, rw.buf)
	rw.buf = rw.buf[len(buf):]
	n = len(buf)
	if n == 0 {
		err = io.EOF
	}
	return
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

	wb := &rwBuffer{}
	wr := xdr.MakeWriter(wb)
	req.Xdr(wr)
	err := wr.Error()
	if err != nil {
		return err
	}

	args.Xdr(wr)
	err = wr.Error()
	if err != nil {
		return err
	}

	var hdr [4]byte
	binary.BigEndian.PutUint32(hdr[:], uint32(len(wb.buf)))
	_, err = c.rw.Write(append(hdr[:], wb.buf...))
	if err != nil {
		return err
	}

	_, err = io.ReadFull(c.rw, hdr[:])
	if err != nil {
		return err
	}

	hlen := binary.BigEndian.Uint32(hdr[:])
	if hlen & (1 << 31) == 0 {
		return fmt.Errorf("fragments not supported")
	}

	buf := make([]byte, hlen & 0x7fffffff)
	_, err = io.ReadFull(c.rw, buf)
	if err != nil {
		return err
	}

	rb := &rwBuffer{buf}
	rd := xdr.MakeReader(rb)
	var res Rpc_msg
	res.Xdr(rd)
	err = rd.Error()
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

	resp.Xdr(rd)
	err = rd.Error()
	if err != nil {
		return err
	}

	return nil
}
