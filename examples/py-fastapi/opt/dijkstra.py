
from ortoolpy import knapsack

def exec():
    import networkx as nx
    g = nx.fast_gnp_random_graph(8, 0.26, 1)
    print(nx.dijkstra_path(g, 0, 2))

if __name__ == '__main__':
    exec()
