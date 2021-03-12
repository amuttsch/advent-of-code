from itertools import permutations
import sys

distances = {}

for line in open('input.txt'):
    (frm, _, to, _, dist) = line.split(' ')
    distances.setdefault(frm, {})[to] = int(dist)
    distances.setdefault(to, {})[frm] = int(dist)

shortest = sys.maxsize
longest = 0
for i in permutations(distances.keys()):
    l = sum(map(lambda x, y: distances[x][y], i[:-1], i[1:]))
    shortest = min(l, shortest)
    longest = max(l, longest)

print('Part 1: Shortest distance = {}'.format(shortest))
print('Part 2: Longest distance = {}'.format(longest))

assert shortest == 141
assert longest == 736
