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

type declType interface {
	goType() string
}

type declTypeTypespec struct {
	t typespec
}

func (t declTypeTypespec) goType() string {
	return t.t.goType()
}

type declTypeArray struct {
	t  typespec
	sz string
}

func (t declTypeArray) goType() string {
	return fmt.Sprintf("[%s]%s", t.sz, t.t.goType())
}

type declTypeVarArray struct {
	t  typespec
	sz string
}

func (t declTypeVarArray) goType() string {
	return fmt.Sprintf("[]%s", t.t.goType())
}

type declTypeOpaqueArray struct {
	sz string
}

func (t declTypeOpaqueArray) goType() string {
	return fmt.Sprintf("[%s]byte", t.sz)
}

type declTypeOpaqueVarArray struct {
	sz string
}

func (t declTypeOpaqueVarArray) goType() string {
	return "[]byte"
}

type declTypeString struct {
	sz string
}

func (t declTypeString) goType() string {
	return "string"
}

type declTypePtr struct {
	t typespec
}

func (t declTypePtr) goType() string {
	return fmt.Sprintf("*%s", t.t.goType())
}

type typespec interface {
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

type typeFloat struct{}

func (t typeFloat) goType() string { return "float32" }

type typeDouble struct{}

func (t typeDouble) goType() string { return "float64" }

type typeQuadruple struct{}

func (t typeQuadruple) goType() string { panic("quadruple") }

type typeBool struct{}

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
	cases      unionCasesDef
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
	val  string
}

type unionCasesDef struct {
	cases []unionCaseDecl
	def   decl
}

type unionCaseDecl struct {
	cases []string
	body  decl
}

func emitConst(ident string, val string) {
	fmt.Printf("const %s = %s\n", ident, val)
}

func emitTypedef(val decl) {
	switch v := val.(type) {
	case declName:
		fmt.Printf("type %s %s\n", v.n, v.t.goType())
	}
}

func emitEnum(ident string, val []enumItem) {
	fmt.Printf("type %s int\n", ident)

	for _, v := range val {
		fmt.Printf("const %s = %s\n", v.name, v.val)
	}
}

func emitStruct(ident string, val []decl) {
	fmt.Printf("type %s struct {\n", ident)

	for _, v := range val {
		switch v := v.(type) {
		case declName:
			fmt.Printf("  %s %s;\n", v.n, v.t.goType())
		}
	}

	fmt.Printf("}\n")
}

func emitUnion(ident string, val typeUnion) {
	fmt.Printf("type %s struct {\n", ident)

	switch v := val.switchDecl.(type) {
	case declName:
		fmt.Printf("  %s %s;\n", v.n, v.t.goType())
	}

	for _, c := range val.cases.cases {
		switch v := c.body.(type) {
		case declName:
			fmt.Printf("  %s %s;\n", v.n, v.t.goType())
		}
	}

	if val.cases.def != nil {
		switch v := val.cases.def.(type) {
		case declName:
			fmt.Printf("  %s %s;\n", v.n, v.t.goType())
		}
	}

	fmt.Printf("}\n")
}
