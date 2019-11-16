GOPATH	:= $(shell go env GOPATH)

all: rfc1813/prot.go rfc1057/prot.go

$(GOPATH)/bin/go-rpcgen: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go install .

%/prot.go: %/prot.x $(GOPATH)/bin/go-rpcgen
	$(GOPATH)/bin/go-rpcgen -i $< -o $@ -p $(patsubst %/prot.go,%,$@)
	go vet $@
