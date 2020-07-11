# coding: utf-8

from typing import Iterator

class Token:
    EOF = -1
    VARIABLE = 1
    OP = 2
    operators = {'+', '-', '='}

    def __init__(self, token: str) -> None:
        self.text = token
        self.type = self.VARIABLE
        if token in self.operators:
            self.type = self.OP
        elif token == '':
            self.type = self.EOF
        return

    def __eq__(self, other: object) -> bool:
        if other is None or type(self) != type(other): return False
        return self.__dict__ == other.__dict__
    def __ne__(self, other: object) -> bool:
        return not self.__eq__(other)

def tokenize(raw: str) -> Iterator[Token]:
    internal = _split(raw)
    for token in internal:
        if token != '':
            yield Token(token)
    yield Token('')

def _split(raw: str) -> Iterator[str]:
    last = -1
    for i in range(len(raw)):
        if raw[i] == " ":
            yield raw[last+1:i]
            last = i
            continue
        if raw[i] in Token.operators:
            yield raw[last+1:i]
            yield raw[i]
            last = i
    yield raw[last+1:len(raw)]
