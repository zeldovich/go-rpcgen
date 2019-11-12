package main

import (
	"fmt"
)

type decl interface{}

type declVoid struct {
}

type declName struct {
	t declType
	n string
}

type declType interface{}

type declTypeTypespec struct {
	t typespec
}

type declTypeArray struct {
	t typespec
	sz string
}

type declTypeVarArray struct {
	t typespec
	sz string
}

type declTypeOpaqueArray struct {
	sz string
}

type declTypeOpaqueVarArray struct {
	sz string
}

type declTypeString struct {
	sz string
}

type declTypePtr struct {
	t typespec
}

type typespec interface{}

type typeInt struct {
	unsig bool
}

type typeHyper struct {
	unsig bool
}

type typeFloat struct {}
type typeDouble struct {}
type typeQuadruple struct {}
type typeBool struct {}

type typeEnum struct {
	items []enumItem
}

type typeStruct struct {
	items []decl
}

type typeUnion struct {
	switchDecl decl
	cases unionCasesDef
}

type typeIdent struct {
	n string
}

type enumItem struct {
	name string
	val string
}

type unionCasesDef struct {
	cases []unionCaseDecl
	def decl
}

type unionCaseDecl struct {
	cases []string
	body decl
}

func emitConst(ident string, val string) {
	fmt.Printf("const %s = %s\n", ident, val)
}

func emitTypedef(val decl) {
}

func emitEnum(ident string, val []enumItem) {
}

func emitStruct(ident string, val []decl) {
}

func emitUnion(ident string, val typeUnion) {
}
