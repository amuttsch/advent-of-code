import math
from collections import defaultdict

min_number_presents = 34000000

def get_house(presents_per_house=10, max_houses=0):
    houses = defaultdict(int)
    hits = []
    for elf in range(1, min_number_presents // 34):
        m = (max_houses * elf) + 1 if max_houses > 0 else min_number_presents // 34
        for j in range(elf, m, elf):
            houses[j] += elf * presents_per_house
            if houses[j] >= min_number_presents:
                hits.append(j)
                if max_houses == 0:
                    return j

    return min(hits)

part_1 = get_house()
part_2 = get_house(11, 50)
print('Part 1: house {} gets at first more than {} presents'.format(part_1, min_number_presents))
print('Part 2: house {} gets at first more than {} presents'.format(part_2, min_number_presents))

assert part_1 == 786240
assert part_2 == 831600
