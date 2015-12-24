from functools import reduce
from itertools import combinations
import sys
import operator

presents = [int(x) for x in open('input.txt')]

weight_per_group_3 = sum(presents) // 3
weight_per_group_4 = sum(presents) // 4

min_qe_3 = sys.maxsize
min_qe_4 = sys.maxsize
for i in range(1, len(presents)):
    for c in combinations(presents, i):
        min_qe_3 = min(min_qe_3, reduce(operator.mul, c)) if sum(c) == weight_per_group_3 else min_qe_3
        min_qe_4 = min(min_qe_4, reduce(operator.mul, c)) if sum(c) == weight_per_group_4 else min_qe_4

    if min_qe_3 != sys.maxsize and min_qe_4 != sys.maxsize:
        break

print('Part 1: {} is the mininum quantum entanglement for 3 groups'.format(min_qe_3))
print('Part 2: {} is the mininum quantum entanglement for 4 groups'.format(min_qe_4))

assert min_qe_3 == 11266889531
assert min_qe_4 == 77387711
