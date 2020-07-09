# coding: utf-8

from __future__ import print_function
from ortools.sat.python import cp_model

class VarArraySolutionPrinter(cp_model.CpSolverSolutionCallback):
    """Print intermediate solutions."""

    def __init__(self, variables):
        cp_model.CpSolverSolutionCallback.__init__(self)
        self.__variables = variables
        self.__solution_count = 0
        self.solutions = []

    def on_solution_callback(self):
        self.__solution_count += 1
        answer = {}
        for v in self.__variables:
            answer[str(v)] = self.Value(v)
            print('%s=%i' % (v, self.Value(v)), end=' ')
        self.solutions.append(answer)
        print()

    def solution_count(self):
        return self.__solution_count

