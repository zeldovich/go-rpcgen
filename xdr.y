%{
package main
%}

%union {
  decl decl;
  typespec typespec;
  str string;
  bool bool;
  enumItem enumItem;
  enumItems []enumItem;
  decls []decl;
  typeUnion typeUnion;
  unionCasesDef unionCasesDef;
  unionCaseDecls []unionCaseDecl;
  unionCaseDecl unionCaseDecl;
  strs []string;
  progCall progCall;
  progCalls []progCall;
  progVer progVer;
  progVers []progVer;
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
%token KWPROGRAM
%token KWVERSION
%token <str> CONST
%token <str> IDENT
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

%type <decl> decl
%type <typespec> typespec enumtypespec structtypespec uniontypespec
%type <str> val varlen
%type <bool> maybeunsig
%type <enumItem> enumitem
%type <enumItems> enumbody enumitems
%type <decls> structbody structdecls
%type <typeUnion> unionbody
%type <unionCasesDef> unioncasesdef
%type <unionCaseDecls> unioncases
%type <unionCaseDecl> unioncase
%type <strs> caselist
%type <str> progtype
%type <progCall> progcall
%type <progCalls> progcalls
%type <progVers> progvers
%type <progVer> progver

%%

spec: | spec defn

defn: typedef | constdef | progdef

decl: typespec IDENT
  { $$ = declName{declTypeTypespec{$1}, $2} }
| typespec IDENT '[' val ']'
  { $$ = declName{declTypeArray{$1, $4}, $2} }
| typespec IDENT varlen
  { $$ = declName{declTypeVarArray{$1, $3}, $2} }
| KWOPAQUE IDENT '[' val ']'
  { $$ = declName{declTypeOpaqueArray{$4}, $2} }
| KWOPAQUE IDENT varlen
  { $$ = declName{declTypeOpaqueVarArray{$3}, $2} }
| KWSTRING IDENT varlen
  { $$ = declName{declTypeString{$3}, $2} }
| typespec '*' IDENT
  { $$ = declName{declTypePtr{$1}, $3} }
| KWVOID
  { $$ = declVoid{} }

varlen: '<' '>'
  { $$ = "" }
| '<' val '>'
  { $$ = $2 }

val: CONST
  { $$ = $1 }
| IDENT
  { $$ = $1 }

typespec: maybeunsig KWINT
  { $$ = typeInt{$1} }
| maybeunsig KWHYPER
  { $$ = typeHyper{$1} }
| KWFLOAT
  { $$ = typeFloat{} }
| KWDOUBLE
  { $$ = typeDouble{} }
| KWQUADRUPLE
  { $$ = typeQuadruple{} }
| KWBOOL
  { $$ = typeBool{} }
| enumtypespec
  { $$ = $1 }
| structtypespec
  { $$ = $1 }
| uniontypespec
  { $$ = $1 }
| IDENT
  { $$ = typeIdent{$1} }

maybeunsig: { $$ = false } | KWUNSIGNED { $$ = true }

enumtypespec: KWENUM enumbody
  { $$ = typeEnum{$2} }

enumbody: '{' enumitems '}'
  { $$ = $2 }

enumitems: enumitem
  { $$ = []enumItem{$1} }
| enumitems ',' enumitem
  { $$ = append($1, $3) }

enumitem: IDENT '=' val
  { $$ = enumItem{$1, $3} }

structtypespec: KWSTRUCT structbody
  { $$ = typeStruct{$2} }

structbody: '{' structdecls '}'
  { $$ = $2 }

structdecls: { $$ = nil } | structdecls decl ';'
  { $$ = append($1, $2) }

uniontypespec: KWUNION unionbody
  { $$ = $2 }

unionbody: KWSWITCH '(' decl ')' '{' unioncasesdef '}'
  { $$ = typeUnion{switchDecl: $3, cases: $6} }

unioncasesdef: unioncases
  { $$ = unionCasesDef{$1, nil} }
| unioncases KWDEFAULT ':' decl ';'
  { $$ = unionCasesDef{$1, $4} }

unioncases: { $$ = nil } | unioncases unioncase
  { $$ = append($1, $2) }

unioncase: caselist decl ';'
  { $$ = unionCaseDecl{$1, $2} }

caselist: KWCASE val ':'
  { $$ = []string{$2} }
| caselist KWCASE val ':'
  { $$ = append($1, $3) }

constdef: KWCONST IDENT '=' CONST ';'
  { emitConst($2, $4) }

typedef: KWTYPEDEF decl ';'
  { emitTypedef($2) }
| KWENUM IDENT enumbody ';'
  { emitEnum($2, $3) }
| KWSTRUCT IDENT structbody ';'
  { emitStruct($2, $3) }
| KWUNION IDENT unionbody ';'
  { emitUnion($2, $3) }

progdef: KWPROGRAM IDENT '{' progvers '}' '=' CONST ';'
  { emitProg(progDef{$2, $4, $7}) }

progvers: { $$ = nil } | progvers progver
  { $$ = append($1, $2) }

progver: KWVERSION IDENT '{' progcalls '}' '=' CONST ';'
  { $$ = progVer{$2, $4, $7} }

progcalls: { $$ = nil } | progcalls progcall
  { $$ = append($1, $2) }

progcall: progtype IDENT '(' progtype ')' '=' CONST ';'
  { $$ = progCall{$2, $4, $1, $7} }

progtype: KWVOID
  { $$ = "" }
| IDENT
  { $$ = $1 }

%%

