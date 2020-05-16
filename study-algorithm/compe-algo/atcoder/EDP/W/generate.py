#!/usr/bin/env python3
import random
N = random.randint(1, 100000)
M = random.randint(100, 1000)
print(N, M)
for _ in range(M):
    l = random.randint(1, 30000)
    r = random.randint(l, 100000)
    a = random.randint(-10000, 20000)
    print(l, r, a)
