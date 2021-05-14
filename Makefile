all: rfc1813/xdr.go rfc1057/xdr.go

go-rpcgen: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go build .

%/xdr.go %/types.go: %/prot.x ./go-rpcgen
	./go-rpcgen -i $< -o $@ -t $(@D)/types.go -p $(@D) -unsigned-enum -const-type uint32
	go vet ./$(@D)

clean:
	@echo CLEAN
	@rm -f go-rpcgen xdr.go y.output
