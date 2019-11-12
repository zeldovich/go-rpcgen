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

type declType interface{
	goType() string
}

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

type typespec interface{
	goType() string
}

type typeInt struct {
	unsig bool
}

func (t typeInt) goType() string {
	if t.unsig {
		return "uint32"
	} else {
		return "int32"
	}
}

type typeHyper struct {
	unsig bool
}

func (t typeHyper) goType() string {
	if t.unsig {
		return "uint64"
	} else {
		return "int64"
	}
}

type typeFloat struct {}
func (t typeFloat) goType() string { return "float32" }

type typeDouble struct {}
func (t typeDouble) goType() string { return "float64" }

type typeQuadruple struct {}
func (t typeQuadruple) goType() string { panic("quadruple") }

type typeBool struct {}
func (t typeBool) goType() string { return "bool" }

type typeEnum struct {
	items []enumItem
}

func (t typeEnum) goType() string { return "int32" }

func declToNameGotype(d decl) string {
	switch v := d.(type) {
	case declName:
		return fmt.Sprintf("%s %s;", v.n, v.t.goType())
	}

	return ""
}

type typeStruct struct {
	items []decl
}

func (t typeStruct) goType() string {
	res := "struct { "
	for _, i := range t.items {
		res += declToNameGotype(i)
	}
	res += "}"
	return res
}

type typeUnion struct {
	switchDecl decl
	cases unionCasesDef
}

func (t typeUnion) goType() string {
	res := "struct { "
	res += declToNameGotype(t.switchDecl)
	for _, i := range t.cases.cases {
		res += declToNameGotype(i.body)
	}
	if t.cases.def != nil {
		res += declToNameGotype(t.cases.def)
	}
	res += "}"
	return res
}

type typeIdent struct {
	n string
}

func (t typeIdent) goType() string { return t.n }

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
	switch v := val.(type) {
	case declVoid:
		// Nothing to emit.
		return

	case declName:
		fmt.Printf("type %s %s\n", v.n, v.t.goType())
	}
}

func emitEnum(ident string, val []enumItem) {
}

func emitStruct(ident string, val []decl) {
}

func emitUnion(ident string, val typeUnion) {
}
