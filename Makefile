GOPATH	:= $(shell go env GOPATH)

all: gonfs/gen.go

gonfs/gen.go: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go install .
	$(GOPATH)/bin/go-rpcgen -i rpc_nfs3_prot.x -o gonfs/gen.go
