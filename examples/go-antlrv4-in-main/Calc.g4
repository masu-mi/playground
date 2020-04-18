grammar Calc;

@header {
// in @header
}

@members {
// in @members
}

prog:   expr NEWLINE*;
expr:   '(' expr_=expr ')'
    |   left=expr op=('*'|'/') right=expr
    |   left=expr op=('+'|'-') right=expr
    |   atom=INT
    ;

NEWLINE : [\r\n]+ ;
INT     : [0-9]+ ;
