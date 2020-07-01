from pulp import *
import networkx as nx

g = nx.fast_gnp_random_graph(8, 0.26, 1).to_directed()
source, sink = 0, 2 # 始点, 終点
r = list(enumerate(g.edges()))
m = LpProblem() # 数理モデル
x = [LpVariable('x%d'%k, lowBound=0, upBound=1) for k, (i, j) in r] # 変数(路に入るかどうか)
m += lpSum(x) # 目的関数
for nd in g.nodes():
    m += lpSum(x[k] for k, (i, j) in r if i == nd) \
      == lpSum(x[k] for k, (i, j) in r if j == nd) + {source:1, sink:-1}.get(nd, 0) # 制約
m.solve()
print([(i, j) for k, (i, j) in r if value(x[k]) > 0.5])
