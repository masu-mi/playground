# coding: utf-8

class Op():
  def __init__(self, op, l, r):
    self.l = l
    self.r = r
    self.op = op

class EqExp(Op):
  def __init__(self, l, r):
      super().__init__('=', l, r)

class Var():
  def __init__(self, token):
    self.name = token
