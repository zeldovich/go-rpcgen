%{
package main
%}

%union {
	num uint64;
	ident string;
}

%token KWCONST
%token KWTYPEDEF
%token KWENUM
%token KWSTRUCT
%token KWUNION
%token KWSWITCH
%token KWCASE
%token KWDEFAULT
%token KWVOID
%token KWOPAQUE
%token KWSTRING
%token KWUNSIGNED
%token KWINT
%token KWHYPER
%token KWFLOAT
%token KWDOUBLE
%token KWQUADRUPLE
%token KWBOOL
%token <num> CONST
%token <ident> IDENT
%token '='
%token ';'
%token '<'
%token '>'
%token '['
%token ']'
%token '{'
%token '}'
%token ','
%token ':'
%token '*'

%%

spec: | spec defn

defn: typedef | constdef

decl: typespec IDENT
| typespec IDENT '[' val ']'
| typespec IDENT varlen
| KWOPAQUE IDENT '[' val ']'
| KWOPAQUE IDENT varlen
| KWSTRING IDENT varlen
| typespec '*' IDENT
| KWVOID

varlen: '<' '>'
| '<' val '>'

val: CONST | IDENT

typespec: maybeunsig KWINT
| maybeunsig KWHYPER
| KWFLOAT
| KWDOUBLE
| KWQUADRUPLE
| KWBOOL
| enumtypespec
| structtypespec
| uniontypespec
| IDENT

maybeunsig: | KWUNSIGNED

enumtypespec: KWENUM enumbody

enumbody: '{' enumitems '}'

enumitems: enumitem | enumitems ',' enumitem

enumitem: IDENT '=' val

structtypespec: KWSTRUCT structbody

structbody: '{' structdecls '}'

structdecls: | structdecls decl ';'

uniontypespec: KWUNION unionbody

unionbody: KWSWITCH '(' decl ')' '{' unioncasesdef '}'

unioncasesdef: unioncases | unioncases KWDEFAULT ':' decl ';'

unioncases: | unioncases unioncase

unioncase: caselist decl ';'

caselist: KWCASE val ':' | caselist KWCASE val ':'

constdef: KWCONST IDENT '=' CONST ';'
{
	emitConst($2, $4)
}

typedef: KWTYPEDEF decl ';'
| KWENUM IDENT enumbody ';'
| KWSTRUCT IDENT structbody ';'
| KWUNION IDENT unionbody ';'

%%

