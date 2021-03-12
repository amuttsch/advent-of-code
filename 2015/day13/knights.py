from itertools import permutations

def calculate_happiness(include_myself=False):
    happiness = {}

    for line in open('input.txt'):
        (person, _, gain_loose, amount, _, _, _, _, _, _, neighbor) = line[:-2].split(' ')
        happiness.setdefault(person, {})[neighbor] = int(amount) if gain_loose == 'gain' else int(amount) *-1
        if include_myself:
            happiness[person]['me'] = 0

    if include_myself:
        for neighbor in list(happiness.keys()):
            happiness.setdefault('me', {})[neighbor] = 0

    best_happiness = 0
    for seating in permutations(happiness.keys()):
        seating_happiness = 0
        for i, person in enumerate(seating):
            seating_happiness += happiness[person][seating[i-1]] + \
                                 happiness[person][seating[(i+1) % len(happiness.keys())]]

        best_happiness = max(best_happiness, seating_happiness)

    return best_happiness

part_1 = calculate_happiness()
part_2 = calculate_happiness(True)
print('Part 1: Best possible happiness = {}'.format(part_1))
print('Part 2: Best possible happiness with me = {}'.format(part_2))

assert part_1 == 733
