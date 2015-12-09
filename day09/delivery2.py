from itertools import permutations
import sys

cities = set()
distances = {}

def read_distances():
    for line in open('input.txt'):
        (frm, _, to, _, dist) = line.split(' ')
        distances.setdefault(frm, {})[to] = int(dist)
        distances.setdefault(to, {})[frm] = int(dist)
        cities.add(frm)
        cities.add(to)

read_distances()
shortest = sys.maxsize
longest = 0
for i in permutations(cities):
    l = sum(map(lambda x, y: distances[x][y], i[:-1], i[1:]))
    shortest = min(l, shortest)
    longest = max(l, longest)

print('Part 1: Shortest distance = {}'.format(shortest))
print('Part 2: Longest distance = {}'.format(longest))

assert shortest == 141
assert longest == 736
