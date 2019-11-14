package main

import (
	"flag"
	"fmt"
	"go/format"
	"go/token"
	"io"
	"io/ioutil"
	"os"
)

var inputFile = flag.String("i", "", "Input file (.x)")
var outputFile = flag.String("o", "", "Output file (.go)")
var out io.Writer

func main() {
	flag.Parse()

	if *inputFile == "" {
		fmt.Fprintf(os.Stderr, "Must specify input file (-i)\n")
		os.Exit(1)
	}

	if *outputFile == "" {
		fmt.Fprintf(os.Stderr, "Must specify output file (-o)\n")
		os.Exit(1)
	}

	src, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f := fset.AddFile(*inputFile, -1, len(src))

	var l lexer
	l.s.Init(f, src, nil, 0)

	outf, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	out = outf
	fmt.Fprintf(out, "package gonfs\n")
	fmt.Fprintf(out, "import . \"github.com/zeldovich/go-rpcgen/xdr\"\n")
	xdrParse(&l)
	outf.Close()

	buf, err := ioutil.ReadFile(*outputFile)
	if err != nil {
		panic(err)
	}

	buf, err = format.Source(buf)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outputFile, buf, 0666)
	if err != nil {
		panic(err)
	}
}
