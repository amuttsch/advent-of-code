from itertools import combinations

LITER_EGGNOG = 150

containers = sorted([int(line) for line in open('input.txt')])

def get_min_container_count(l):
    amount = 0
    for i, c in enumerate(l):
        amount += c
        if amount > LITER_EGGNOG:
            return i + 1

min_container_count = get_min_container_count(containers[::-1])
max_container_count = get_min_container_count(containers)

sum_container_hold_eggnog = 0
sum_min_containers_hold_eggnog = 0
min_containers_hold_eggnog = len(containers)

for i in range(min_container_count, max_container_count):
    for c in combinations(containers, i):
        if sum(c) == LITER_EGGNOG:
            min_containers_hold_eggnog = min(min_containers_hold_eggnog, i)
            sum_min_containers_hold_eggnog += 1 if i == min_containers_hold_eggnog else 0
            sum_container_hold_eggnog += 1

print('Part 1: Container combinations holding exactly {} liters of eggnog = {}'.format(LITER_EGGNOG, sum_container_hold_eggnog))
print('Part 2: Number of combinations with the minimum container count = {}'.format(sum_min_containers_hold_eggnog))

assert sum_container_hold_eggnog == 1304
assert sum_min_containers_hold_eggnog == 18
