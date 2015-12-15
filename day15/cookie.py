from collections import defaultdict
from itertools import product

ingredients = defaultdict(dict)
teaspoons = 100
best_score = 0
best_score_with_500_cal = 0

for line in open('input.txt'):
    s = line.split(' ')
    ingredients[s[0][:-1]][s[1]] = int(s[2][:-1])
    ingredients[s[0][:-1]][s[3]] = int(s[4][:-1])
    ingredients[s[0][:-1]][s[5]] = int(s[6][:-1])
    ingredients[s[0][:-1]][s[7]] = int(s[8][:-1])
    ingredients[s[0][:-1]][s[9]] = int(s[10][:-1])

for p in product(range(teaspoons+1), repeat=4):
    if sum(p) == teaspoons:
        capacity = max(0, sum(map(lambda x, y: x['capacity']*y, ingredients.values(), p)))
        durability = max(0, sum(map(lambda x, y: x['durability']*y, ingredients.values(), p)))
        flavor = max(0, sum(map(lambda x, y: x['flavor']*y, ingredients.values(), p)))
        texture = max(0, sum(map(lambda x, y: x['texture']*y, ingredients.values(), p)))
        calories = max(0, sum(map(lambda x, y: x['calories']*y, ingredients.values(), p)))

        best_score = max(best_score, capacity * durability * flavor * texture)
        if calories == 500:
            best_score_with_500_cal = max(best_score_with_500_cal, capacity * durability * flavor * texture)

print('Part 1: Best score = {}'.format(best_score))
print('Part 2: Best score with exactly 500 cal = {}'.format(best_score_with_500_cal))

assert best_score == 222870
assert best_score_with_500_cal == 117936
