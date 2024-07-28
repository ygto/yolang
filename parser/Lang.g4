grammar Lang;

// Programın başlangıcı, ifade listelerini tanır
prog:   expr+ EOF;

// İfadeler, fonksiyon çağrıları veya değişkenlere atama içerebilir
expr:   assignExpr
    |   funcExpr
    |   callFuncExpr
    |   functions
    |   valExpr
    |   varExpr
    |   returnExpr
    ;
// özel fonksiyonlar
reservedFunctions:   'print'
    
    ;
functions:   reservedFunctions LPAREN valExpr RPAREN
    ;



types:   'int'
      | 'varchar'
      ;

valExpr: variableExpr
    |    callFuncExpr
    |    staticValues
    ;

variableExpr: '$' ID ;

varExpr: 'var' variableExpr types ;

returnExpr: 'return' valExpr ;


assignExpr: variableExpr ASSIGN valExpr ;

block: LSCOPE expr+ RSCOPE
    | LSCOPE RSCOPE
;


callFuncExpr: '#' ID '(' variableExpr ')';

funcExpr: '#' ID '(' types variableExpr ')' types block ;




// Tip tanımları için int, float, varchar tiplerini tanır
staticValues : INT
    | VARCHAR
    | NULL
    ;

// Virgülle ayrılmış ifade listesi
exprList : expr (',' expr)* ;

// Token tanımlamaları
INT : [0-9]+ ;
VARCHAR : '\'' .*? '\'' ; // Tek tırnak içindeki herhangi bir karakter dizisi
ID : [a-zA-Z_][a-zA-Z_0-9]* ;
WS : [ \t\r\n]+ -> skip ; // Boşlukları yoksay
NULL : 'null' ; // Boşlukları yoksay

// Operatör tokenları
ASSIGN : '=' ;

// Parantez tokenları
LPAREN : '(' ;
RPAREN : ')' ;
LSCOPE : '{' ;
RSCOPE : '}' ;
