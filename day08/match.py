import re

part_1 = 0
part_2 = 0

with open('input.txt') as file:
    for line in file:
        line = line[:-1] # Slice newline
        total = len(line)

        dec = line[1:-1]
        dec = dec.replace('\\\\', '\\').replace('\\"', '"')
        dec = re.sub(r'\\x[0-9a-f]{2}', '.', dec)
        part_1 += total - len(dec)

        enc = '"{}"'.format(line.replace('\\', r'\\').replace('"', r'\"'))
        part_2 += len(enc) - total

print('Part 1: {}'.format(part_1))
print('Part 2: {}'.format(part_2))

assert part_1 == 1333
assert part_2 == 2046
