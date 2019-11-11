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
%token <num> NUM
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

commands: | commands command ';'

command: const | typedef | enum | struct | union

const: KWCONST IDENT '=' NUM

typedef: KWTYPEDEF typename varname

typename: IDENT

varname: '*' varname | IDENT | IDENT '<' '>' | IDENT '<' IDENT '>' | IDENT '[' IDENT ']'

enum: KWENUM IDENT '{' enumitems '}'

enumitems: enumitem | enumitems ',' enumitem

enumitem: IDENT | IDENT '=' NUM

struct: KWSTRUCT IDENT '{' structfields '}'

structfields: | structfields structfield ';'

structfield: typename varname | KWVOID

union: KWUNION IDENT KWSWITCH '(' typename varname ')' '{' unioncases '}'

unioncases: | unioncases unioncase

unioncase: unioncaseval ':' structfields

unioncaseval: KWCASE IDENT | KWDEFAULT

%%

