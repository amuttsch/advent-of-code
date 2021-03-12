from collections import defaultdict
from random import shuffle

molecule = open('molecule.txt').read().strip('\n')
replacements = defaultdict(list)
new_molecules = set()

for line in open('replacements.txt'):
    s = line.strip('\n').split(' ')
    replacements[s[0]] += [s[2]]

for m in replacements.keys():
    s = molecule.split(m)
    for i in range(len(s)-1):
        for r in replacements[m]:
            new_molecule = ''
            new_molecule += m.join(s[:i+1])
            new_molecule += r
            new_molecule += m.join(s[i+1:])

            new_molecules.add(new_molecule)

print('Part 1: {} distinct molecules can be created'.format(len(new_molecules)))

assert len(new_molecules) == 509

m = molecule
steps = 0
while m != 'e':
    t = m
    for a, b in replacements.items():
        for r in b:
            if r not in m:
                continue

            m = m.replace(r, a, 1)
            steps += 1

    if t == m:
        m = molecule
        steps = 0
        shuffle(replacements)

print('Part 2: Fewest number of steps to make the medicine = {}'.format(steps))

assert steps == 195
