GOPATH	:= $(shell go env GOPATH)
GOPATH1	:= $(firstword $(subst :, ,$(GOPATH)))

all: rfc1813/xdr.go rfc1057/xdr.go

go-rpcgen: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	PATH=$(PATH):$(GOPATH1)/bin go generate
	go build .

%/xdr.go %/types.go: %/prot.x ./go-rpcgen
	./go-rpcgen -i $< -o $@ -t $(@D)/types.go -p $(@D) -unsigned-enum -const-type uint32
	go vet ./$(@D)

clean:
	@echo CLEAN
	@rm -f go-rpcgen xdr.go y.output
