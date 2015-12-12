data = open('input.txt').read().replace('\n', '')

part_1 = sum(map(lambda x: 1 if x == '(' else -1, data))
part_2 = 0

floor = 0
for i, c in enumerate(data):
    floor += (1 if c == '(' else -1)
    if floor == -1:
        part_2 = i+1
        break

print('Part 1: Santa is on floor {}'.format(part_1))
print('Part 2: Santa is on basement on position {}'.format(part_2))

assert part_1 == 280
assert part_2 == 1797
