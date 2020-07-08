# coding: utf-8

from __future__ import print_function
from ortools.sat.python import cp_model

from ast import Op,EqExp,Var
from callbacks import VarArraySolutionPrinter


class Problem:

    def __init__(self, eqExp, base = 10):
        self.base = base
        self.ast = eqExp
        vals = analyze_values(eqExp)
        self.heads, self.chars = analyze_semantic(vals)

    def ortools_model(self):
        model = cp_model.CpModel()
        variables = self.get_model_variables(model)
        model.Add(self.build_expr(variables, self.ast))
        return (model, variables)

    def search_all_solution(self):
        solver = cp_model.CpSolver()
        model, variables = self.ortools_model()
        solution_printer = VarArraySolutionPrinter(variables.values())
        status = solver.SearchForAllSolutions(model, solution_printer)
        return status

    def build_expr(self, variables, node):
        if type(node) is Var:
            return self.build_value(variables, node.name)

        l = self.build_expr(variables, node.l)
        r = self.build_expr(variables, node.r)
        if type(node) is Op:
            if node.op == '+':
                return l+r
            elif node.op == '-':
                return l-r
        elif type(node) is EqExp:
            return l == r

    def build_value(self, variables, text):
        value = 0
        for c in text:
            v = None
            if c in self.chars:
                v = variables[c]
            else:
                v = int(c)
            value = value*self.base + v
        return value

    def get_model_variables(self, model):
        variables = {}
        for c in self.chars:
            if c in self.heads:
                df = model.NewIntVar(1, self.base-1, c)
                variables[c] = df
            else:
                df = model.NewIntVar(0, self.base-1, c)
                variables[c] = df
        return variables

def analyze_values(ast):
    v = []
    _values(ast, v)
    return v

def _values(ast, l):
    if issubclass(type(ast), Op):
        _values(ast.l, l)
        _values(ast.r, l)
    else:
        l.append(ast.name)

def analyze_semantic(variables):
    def is_digit(c):
        return c >= '0' and c <= '9'

    heads = set()
    chars = set()
    for variable in variables:
        c = variable[0]
        if not is_digit(c):
            heads.add(c)
        for c in variable:
            if not is_digit(c):
                chars.add(c)
    return (heads, chars)
