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
    |   left=expr { fmt.Printf("%v\n", $left.text) } op=('+'|'-') right=expr { fmt.Printf("ctx: %v, start_line: %v, left: %v, right_start_token: %v\n", $ctx, $start.GetLine(), $left.text, $right.start) }
    |   atom=INT
    ;

NEWLINE : [\r\n]+ ;
INT     : [0-9]+ ;
