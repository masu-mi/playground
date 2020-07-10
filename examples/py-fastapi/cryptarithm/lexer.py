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

    def __eq__(self, other):
        return self.text == other.text and self.type == other.type
    def __ne__(self, other):
        return not self.__eq__(other)

    @classmethod
    def tokenize(cls, raw):
        internal = cls._split(raw)
        for token in internal:
            if token != '':
                yield Token(token)

    @classmethod
    def _split(cls, raw):
        last = -1
        for i in range(len(raw)):
            if raw[i] == " ":
                yield raw[last+1:i]
                last = i
                continue
            if raw[i] in cls.operators:
                yield raw[last+1:i]
                yield raw[i]
                last = i
        yield raw[last+1:len(raw)]
