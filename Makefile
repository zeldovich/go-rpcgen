GOPATH	:= $(shell go env GOPATH)

all: rfc4506/prot.go

$(GOPATH)/bin/go-rpcgen: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go install .

%.go: %.x $(GOPATH)/bin/go-rpcgen
	$(GOPATH)/bin/go-rpcgen -i $< -o $@
