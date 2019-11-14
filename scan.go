package main

//go:generate goyacc -o xdr.go -p xdr xdr.y

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type lexer struct {
	s scanner.Scanner
}

const eof = 0

func (l *lexer) Lex(lval *xdrSymType) int {
	pos, tok, lit := l.s.Scan()
	if tok == token.EOF {
		return eof
	}

	// fmt.Printf("pos=%v, tok=%v, lit=%v\n", pos, tok, lit)

	switch tok {
	case token.CONST:
		return KWCONST

	case token.STRUCT:
		return KWSTRUCT

	case token.TYPE:
		lval.str = "type"
		return IDENT

	case token.SWITCH:
		return KWSWITCH

	case token.CASE:
		return KWCASE

	case token.DEFAULT:
		return KWDEFAULT

	case token.IDENT:
		switch lit {
		case "typedef":
			return KWTYPEDEF

		case "enum":
			return KWENUM

		case "union":
			return KWUNION

		case "void":
			return KWVOID

		case "string":
			return KWSTRING

		case "opaque":
			return KWOPAQUE

		case "unsigned":
			return KWUNSIGNED

		case "int":
			return KWINT

		case "hyper":
			return KWHYPER

		case "float":
			return KWFLOAT

		case "double":
			return KWDOUBLE

		case "quadruple":
			return KWQUADRUPLE

		case "bool":
			return KWBOOL

		case "program":
			return KWPROGRAM

		case "version":
			return KWVERSION

		default:
			lval.str = lit
			return IDENT
		}

	case token.ASSIGN:
		return '='

	case token.SEMICOLON:
		return ';'

	case token.COLON:
		return ':'

	case token.LSS:
		return '<'

	case token.GTR:
		return '>'

	case token.LBRACK:
		return '['

	case token.RBRACK:
		return ']'

	case token.LBRACE:
		return '{'

	case token.RBRACE:
		return '}'

	case token.COMMA:
		return ','

	case token.LPAREN:
		return '('

	case token.RPAREN:
		return ')'

	case token.MUL:
		return '*'

	case token.INT:
		lval.str = lit
		return CONST

	default:
		fmt.Printf("NOT HANDLED: pos=%v, tok=%v, lit=%v\n", pos, tok, lit)
		panic("token not handled")
	}
}

func (l *lexer) Error(e string) {
	panic(e)
}
