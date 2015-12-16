from collections import defaultdict
import re
import math

ticker_tape = {
    "children": 3,
    "cats": 7, # greater than
    "samoyeds": 2,
    "pomeranians": 3, # fewer than
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5, # fewer than
    "trees": 3, # greater than
    "cars": 2,
    "perfumes": 1,
}

def MFCSAM_range(reading, aunt):
    if reading in ['cats', 'trees']:
        return 0 if aunt[reading] > ticker_tape[reading] else 1
    elif reading in ['pomeranians', 'goldfish']:
        return 0 if aunt[reading] < ticker_tape[reading] else 1
    else:
        return math.fabs(aunt[reading] - ticker_tape[reading])

data = open('input.txt').read()
l = re.findall('\w* \d+: (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)', data)
aunts = map(lambda x: {x[0]: int(x[1]), x[2]: int(x[3]), x[4]: int(x[5])}, l)

aunt_part_1 = aunt_part_2 = 0
for i, a in enumerate(aunts):
    score_part_1 = sum(map(lambda x: math.fabs(a[x] - ticker_tape[x]), a))
    score_part_2 = sum(map(lambda x: MFCSAM_range(x, a), a))
    aunt_part_1 = i+1 if score_part_1 == 0 else aunt_part_1
    aunt_part_2 = i+1 if score_part_2 == 0 else aunt_part_2

print('Part 1: Aunt {} sent the present'.format(aunt_part_1))
print('Part 2: Aunt {} sent the present'.format(aunt_part_2))

assert aunt_part_1 == 373
assert aunt_part_2 == 260
