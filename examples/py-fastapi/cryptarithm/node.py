# coding: utf-8

from typing import Union

class Node():
    def __init__(self) -> None:
        pass

class Var(Node):
    def __init__(self, token: str) -> None:
        self.name = token

class Op(Node):
    def __init__(self, op: str, l: Node, r: Node) -> None:
        self.l = l
        self.r = r
        self.op = op

class EqExp(Op):
    def __init__(self, l: Node, r: Node) -> None:
        super().__init__('=', l, r)
