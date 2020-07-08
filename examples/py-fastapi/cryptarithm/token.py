# coding: utf-8


class Token:

    EOF = -1
    VARIABLE = 1
    OP = 2
    operators = {'+', '-', '='}

    def __init__(self, token):
        self.text = token
        self.type = self.VARIABLE
        if token in self.operators:
            self.type = self.OP
        elif token == '':
            self.type = self.EOF
        return

    @classmethod
    def tokenize(cls, raw):
        for token in raw.split(" "):
            if token != '':
                yield Token(token)
        yield Token('')
