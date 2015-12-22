from collections import defaultdict
import re
from itertools import product
import math

shop_re = re.compile(r'(\w+(.\+\d{1})*)\s*(\d+).*(\d+).*(\d+)')

boss = {'hp': 109, 'damage': 8, 'armor': 2}
player = {'hp': 100, 'damage': 0, 'armor': 0}
shop = defaultdict(dict)

def simulate_fight(damage, armor):
    dmg_boss_per_round = max(1, player['damage'] + damage - boss['armor'])
    dmg_player_per_round = max(1, boss['damage'] - player['armor'] - armor)

    rounds_to_kill_boss = math.ceil(boss['hp'] / dmg_boss_per_round)
    rounds_to_kill_player = math.ceil(player['hp'] / dmg_player_per_round)

    return rounds_to_kill_boss <= rounds_to_kill_player

item_type = ''
for line in open('shop.txt'):
    t = line.split(':')
    if len(t) > 1:
        item_type = t[0]
        shop[item_type]['none'] = {'cost': 0,
                                    'damage': 0,
                                    'armor': 0}
        continue
    if line == '\n':
        continue

    item = shop_re.match(line).groups()
    if len(item) >= 4:
        shop[item_type][item[0]] = {'cost': int(item[-3]),
                                    'damage': int(item[-2]),
                                    'armor': int(item[-1])}

lowest_cost_to_win = 999
most_cost_to_loose = 0

for w, a, r1, r2 in product(shop['Weapons'], shop['Armor'], shop['Rings'], shop['Rings']):
    if r1 == r2 or w == 'none':
        continue
    cost = shop['Weapons'][w]['cost'] + shop['Armor'][a]['cost'] + shop['Rings'][r1]['cost'] + shop['Rings'][r2]['cost']
    damage = shop['Weapons'][w]['damage'] + shop['Armor'][a]['damage'] + shop['Rings'][r1]['damage'] + shop['Rings'][r2]['damage']
    armor = shop['Weapons'][w]['armor'] + shop['Armor'][a]['armor'] + shop['Rings'][r1]['armor'] + shop['Rings'][r2]['armor']

    if simulate_fight(damage, armor):
        lowest_cost_to_win = min(lowest_cost_to_win, cost)
    else:
        most_cost_to_loose = max(most_cost_to_loose, cost)

print('Part 1: Player must spend at least {} gold to win the fight'.format(lowest_cost_to_win))
print('Part 2: Most cost to still loose the fight = {} gold'.format(most_cost_to_loose))

assert lowest_cost_to_win == 111
assert most_cost_to_loose == 188
