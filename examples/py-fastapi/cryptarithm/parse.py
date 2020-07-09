# coding: utf-8

from lexer import Token
from node import Op,EqExp,Var

"""
Parse
"""

class Parser:
    def __init__(self, tokenizer):
        self.tokenizer = tokenizer
        self.nextToken()

    def nextToken(self):
        self.head = self.tokenizer.__next__()

    def match(self, v):
        tk = self.head
        self.nextToken()
        if v in Token.operators:
            assert tk.type == Token.OP and v == tk.text
            return
        assert tk.type == v

    def parse(self):
        return self.sentence()

    def sentence(self):
        l = self.expr()
        self.match('=')
        r = self.expr()
        return EqExp(l, r)

    def expr(self):
        t = self.head
        self.match(Token.VARIABLE)
        ex = Var(t.text)
        while True:
            t = self.head
            if t.type != Token.OP:
                break
            if t.text == '+':
                self.match('+')
                r = self.head
                self.match(Token.VARIABLE)
                ex = Op('+', ex, Var(r.text))
            elif t.text == '-':
                self.match('-')
                r = self.head
                self.match(Token.VARIABLE)
                ex = Op('-', ex, Var(r.text))
            else:
                break
        return ex
