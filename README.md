# RFC4506 XDR generator for Go

[![CI](https://github.com/zeldovich/go-rpcgen/actions/workflows/build.yml/badge.svg)](https://github.com/zeldovich/go-rpcgen/actions/workflows/build.yml)

Run `make` to build the `go-rpcgen` tool and compile several specs,
including NFS (rfc1813) and SUNRPC (rfc1057).

You will likely need to `go install golang.org/x/tools/cmd/goyacc@latest`.

There is an example client in `example/client/main.go` that connects to
an NFS server and issues some NFS RPCs, and an example server in
`example/server/main.go`.

## Fuzzing the decoders

```
go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
( cd rfc1057 && go-fuzz-build && ulimit -d 1048576 && go-fuzz )
```

## TODO

- Validate that encoded/decoded enums match one of the allowed values.
