GOPATH	:= $(shell go env GOPATH)

all: rfc1813/xdr.go rfc1057/xdr.go

$(GOPATH)/bin/go-rpcgen: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go install .

%/xdr.go %/types.go: %/prot.x $(GOPATH)/bin/go-rpcgen
	$(GOPATH)/bin/go-rpcgen -i $< -o $@ -t $(@D)/types.go -p $(@D) -unsigned-enum -const-type uint32
	go vet ./$(@D)
