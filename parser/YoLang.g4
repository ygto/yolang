grammar YoLang;

// Programın başlangıcı, ifade listelerini tanır
prog:   expr+ EOF;


expr:   importFile
  | variableDefinition
  | assign
  | functionDefinition
  | call
  | structDecl
    ;


primitiveTypes: 'int'
  | 'varchar'
  ;

types: primitiveTypes
  | ID
  ;

 value : staticValue
   | variable
   | call
   ;

parameters : value (',' value)* ;

block: LSCOPE expr+ RSCOPE
    | LSCOPE RSCOPE
;

variable : '$' ID;
variableDefinition : types variable
  | varcharDecl;

structDecl: 'type' ID 'struct' LSCOPE variableDefinition+ RSCOPE;
structVariable : variable '.' ID ('.' ID)*;

varcharDecl : 'varchar' LPAREN INT RPAREN variable;

functionDefinition : ID LPAREN types variable RPAREN types block;
call : ID LPAREN variable RPAREN;
importFile : '#import'+ LPAREN VARCHAR RPAREN ;


assign: variable ASSIGN value
  | structVariable ASSIGN value
  ;


staticValue : INT
  | VARCHAR
  | NULL
  ;

// Token tanımlamaları
INT : [0-9]+ ;
VARCHAR : '\'' .*? '\'' ; // Tek tırnak içindeki herhangi bir karakter dizisi
ID : [a-zA-Z_][a-zA-Z_0-9]* ;
WS : [ \t\r\n]+ -> skip ; // Boşlukları yoksay
NULL : 'null' ;

// Operatör tokenları
ASSIGN : '=' ;

// Parantez tokenları
LPAREN : '(' ;
RPAREN : ')' ;
LSCOPE : '{' ;
RSCOPE : '}' ;

