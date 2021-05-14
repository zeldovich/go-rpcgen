package rfc1057

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/zeldovich/go-rpcgen/xdr"
)

// reqBufPool holds temporary byte slice buffers used for incoming requests.
var reqBufPool sync.Pool

func getReqBuf(buflen int) []byte {
	bufi := reqBufPool.Get()
	if bufi != nil {
		buf := bufi.([]byte)
		if buflen <= cap(buf) {
			return buf[:buflen]
		}
	}

	return make([]byte, buflen)
}

func putReqBuf(s []byte) {
	reqBufPool.Put(s)
}

type ProcHandler func(args *xdr.XdrState) (res xdr.Xdrable, err error)

type Server struct {
	handlers map[uint32]map[uint32]map[uint32]ProcHandler
}

type serverConn struct {
	s       *Server
	rw      io.ReadWriter
	writeMu sync.Mutex
}

func MakeServer() *Server {
	return &Server{
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

func (s *Server) RegisterMany(regs []xdr.ProcRegistration) {
	for _, r := range regs {
		s.Register(r.Prog, r.Vers, r.Proc, r.Handler)
	}
}

func (s *Server) Run(rw io.ReadWriter) error {
	sc := &serverConn{
		s:  s,
		rw: rw,
	}

	for {
		var hdr [4]byte
		_, err := io.ReadFull(sc.rw, hdr[:])
		if err != nil {
			return err
		}

		hlen := binary.BigEndian.Uint32(hdr[:])
		if hlen&(1<<31) == 0 {
			return fmt.Errorf("fragments not supported")
		}

		buflen := int(hlen & 0x7fffffff)
		buf := getReqBuf(buflen)
		_, err = io.ReadFull(sc.rw, buf)
		if err != nil {
			return err
		}

		go sc.handleReq(buf)
	}
}

func (sc *serverConn) handleReq(buf []byte) {
	defer putReqBuf(buf)

	err := sc.handleReqErr(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func (sc *serverConn) handleReqErr(buf []byte) error {
	rb := &readBuffer{buf}
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
		vermap, progok := sc.s.handlers[req.Body.Cbody.Prog]
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
	wb := &writeBuffer{}

	// Reserve 4 bytes at the front for the length
	var reserveLen [4]byte
	wb.Write(reserveLen[:])

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

	binary.BigEndian.PutUint32(wb.buf[0:4], (1<<31)|uint32(len(wb.buf)-4))

	sc.writeMu.Lock()
	defer sc.writeMu.Unlock()
	_, err = sc.rw.Write(wb.buf)
	return err
}
