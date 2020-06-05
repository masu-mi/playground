#encoding:utf-8

from pulp import *

size = [21, 11, 15, 9, 34, 25, 41, 52]
weight = [22, 12, 16, 10, 35, 26, 42, 53]
capacity = 100
r = range(len(size))

m = LpProblem(sense=LpMaximize) # 数理モデル
x = [LpVariable('x%d'%i, cat=LpBinary) for i in r] # 変数
m += lpDot(weight, x) # 目的関数
m += lpDot(size, x) <= capacity # 制約
m.solve()
print((value(m.objective), [i for i in r if value(x[i]) > 0.5]))
