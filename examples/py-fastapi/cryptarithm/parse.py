# coding: utf-8

import lexer
from node import Op,EqExp,Var
from typing import Iterator,Union

"""
Parse
"""

class Parser:
    def __init__(self, tokenizer: Iterator[lexer.Token]) -> None:
        self.tokenizer = tokenizer
        self.nextToken()

    def nextToken(self) -> None:
        self.head = self.tokenizer.__next__()

    def match(self, v: Union[str, int]):
        tk = self.head
        self.nextToken()
        if type(v) is str:
            assert tk.type == lexer.Token.OP and v == tk.text
            return
        assert tk.type == v

    def parse(self) -> EqExp:
        return self.sentence()

    def sentence(self) -> EqExp:
        l = self.expr()
        self.match('=')
        r = self.expr()
        return EqExp(l, r)

    def expr(self) -> Union[Op, Var]:
        t = self.head
        self.match(lexer.Token.VARIABLE)
        ex: Union[Op, Var] = Var(t.text)
        while True:
            t = self.head
            if t.type != lexer.Token.OP:
                break
            if t.text == '+':
                self.match('+')
                r = self.head
                self.match(lexer.Token.VARIABLE)
                ex = Op('+', ex, Var(r.text))
            elif t.text == '-':
                self.match('-')
                r = self.head
                self.match(lexer.Token.VARIABLE)
                ex = Op('-', ex, Var(r.text))
            else:
                break
        return ex
