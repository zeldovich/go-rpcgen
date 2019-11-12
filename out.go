package main

import (
	"fmt"
)

func emitConst(ident string, num uint64) {
	fmt.Printf("const %s = %d\n", ident, num)
}
