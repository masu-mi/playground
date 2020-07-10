# coding: utf-8

import lexer
from typing import Union

class Op():
    def __init__(self, op: str, l: Union[Op, Var], r: Union[Op, Var]) -> None:
        self.l = l
        self.r = r
        self.op = op

class EqExp(Op):
    def __init__(self, l: Union[Op, Var], r: Union[Op, Var]) -> None:
        super().__init__('=', l, r)

class Var():
    def __init__(self, token: str) -> None:
        self.name = token
