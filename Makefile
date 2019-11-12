all: gonfs/gen.go

gonfs/gen.go: $(wildcard *.go) $(wildcard *.y) $(wildcard *.x)
	go generate
	go run . | gofmt > gonfs/gen.go
