package rfc1057

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/zeldovich/go-rpcgen/xdr"
)

type ProcHandler func (args *xdr.XdrState) (res xdr.Xdrable, err error)

type Server struct {
	rw       io.ReadWriter
	handlers map[uint32]map[uint32]map[uint32]ProcHandler
	writeMu  sync.Mutex
}

func MakeServer(rw io.ReadWriter) *Server {
	return &Server{
		rw:   rw,
		handlers: make(map[uint32]map[uint32]map[uint32]ProcHandler),
	}
}

func (s *Server) Register(prog, vers, proc uint32, handler ProcHandler) {
	_, progok := s.handlers[prog]
	if !progok {
		s.handlers[prog] = make(map[uint32]map[uint32]ProcHandler)
	}

	_, versok := s.handlers[prog][vers]
	if !versok {
		s.handlers[prog][vers] = make(map[uint32]ProcHandler)
	}

	s.handlers[prog][vers][proc] = handler
}

func (s *Server) Run() error {
	for {
		var hdr [4]byte
		_, err := io.ReadFull(s.rw, hdr[:])
		if err != nil {
			return err
		}

		hlen := binary.BigEndian.Uint32(hdr[:])
		if hlen&(1<<31) == 0 {
			return fmt.Errorf("fragments not supported")
		}

		buf := make([]byte, hlen&0x7fffffff)
		_, err = io.ReadFull(s.rw, buf)
		if err != nil {
			return err
		}

		go s.handleReq(buf)
	}
}

func (s *Server) handleReq(buf []byte) {
	err := s.handleReqErr(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func (s *Server) handleReqErr(buf []byte) error {
	rb := &rwBuffer{buf}
	rd := xdr.MakeReader(rb)

	var req Rpc_msg
	req.Xdr(rd)
	err := rd.Error()
	if err != nil {
		return err
	}

	if req.Body.Mtype != CALL {
		return fmt.Errorf("request mtype %d != CALL", req.Body.Mtype)
	}

	var res Rpc_msg
	var resdata xdr.Xdrable
	res.Xid = req.Xid
	res.Body.Mtype = REPLY

	if req.Body.Cbody.Rpcvers != 2 {
		res.Body.Rbody.Stat = MSG_DENIED
		res.Body.Rbody.Rreply.Stat = RPC_MISMATCH
	} else {
		res.Body.Rbody.Stat = MSG_ACCEPTED
		vermap, progok := s.handlers[req.Body.Cbody.Prog]
		if !progok {
			res.Body.Rbody.Areply.Reply_data.Stat = PROG_UNAVAIL
			goto reply
		}

		procmap, verok := vermap[req.Body.Cbody.Vers]
		if !verok {
			res.Body.Rbody.Areply.Reply_data.Stat = PROG_MISMATCH
			goto reply
		}

		h, procok := procmap[req.Body.Cbody.Proc]
		if !procok {
			res.Body.Rbody.Areply.Reply_data.Stat = PROC_UNAVAIL
			goto reply
		}

		resdata, err = h(rd)
		if err != nil {
			res.Body.Rbody.Areply.Reply_data.Stat = GARBAGE_ARGS
			goto reply
		}

		res.Body.Rbody.Areply.Reply_data.Stat = SUCCESS
	}

reply:
	wb := &rwBuffer{}
	wr := xdr.MakeWriter(wb)
	res.Xdr(wr)
	err = wr.Error()
	if err != nil {
		return err
	}

	if resdata != nil {
		resdata.Xdr(wr)
		err = wr.Error()
		if err != nil {
			return err
		}
	}

	var hdr [4]byte
	binary.BigEndian.PutUint32(hdr[:], (1<<31)|uint32(len(wb.buf)))
	resbuf := append(hdr[:], wb.buf...)

	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	_, err = s.rw.Write(resbuf)
	return err
}