import math
from collections import defaultdict

min_number_presents = 34000000
houses = defaultdict(int)

def part1():
    for i in range(1, min_number_presents):
        for j in range(i, min_number_presents // 34, i):
            houses[j] += i * 10
            if houses[j] >= min_number_presents:
                return j

part_1 = part1()
print('Part 1: house {} gets at first more than {} presents'.format(part_1, min_number_presents))

assert part_1 == 786240
