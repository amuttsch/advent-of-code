from functools import reduce
from itertools import product
import random

class Wizard:
    spells = {'Magic Missile':  {'cost': 53,  'rounds': 0, 'dmg': 4, 'armor': 0, 'heal': 0, 'mana': 0},
              'Drain':          {'cost': 73,  'rounds': 0, 'dmg': 2, 'armor': 0, 'heal': 2, 'mana': 0},
              'Shield':         {'cost': 113, 'rounds': 6, 'dmg': 0, 'armor': 7, 'heal': 0, 'mana': 0},
              'Poison':         {'cost': 173, 'rounds': 6, 'dmg': 3, 'armor': 0, 'heal': 0, 'mana': 0},
              'Recharge':       {'cost': 229, 'rounds': 5, 'dmg': 0, 'armor': 0, 'heal': 0, 'mana': 101}}

    def __init(self):
        self.init()

    def init(self, difficulty='normal'):
        self.boss = {'hp': 71, 'dmg': 10}
        self.player = {'hp': 50, 'mana': 500, 'dmg': 0, 'armor': 0}
        self.active_spells = {}
        self.mana_spent = 0
        self.spells_cast = []
        self.min_mana = 229
        self.difficulty = difficulty

    def apply_spell(self, spell):
        self.boss['hp'] -= spell['dmg']
        self.player['hp'] += spell['heal']
        self.player['mana'] += spell['mana']
        self.player['armor'] += spell['armor']

    def apply_active_spells(self):
        self.player['armor'] = 0
        to_remove = []
        for active_spell in self.active_spells:
            self.apply_spell(self.spells[active_spell])

            self.active_spells[active_spell]['rounds_active'] -= 1
            if self.active_spells[active_spell]['rounds_active'] <= 0:
                to_remove.append(active_spell)

        for spell in to_remove:
            del self.active_spells[spell]

    def can_cast_spell(self, spell):
        if spell in self.active_spells:
            return False
        return self.player['mana'] >= self.spells[spell]['cost']

    def next_turn(self, spell_to_cast):
        if self.difficulty == 'difficult':
            self.player['hp'] -= 1
            if self.player['hp'] <= 0:
                return -1

        self.apply_active_spells()
        if self.boss['hp'] <= 0:
            return 1

        if not self.can_cast_spell(spell_to_cast):
            return -2

        self.player_turn(spell_to_cast)
        if self.boss['hp'] <= 0:
            return 1

        self.apply_active_spells()
        if self.boss['hp'] <= 0:
            return 1

        self.boss_turn()
        if self.player['hp'] <= 0:
            return -1

        return 0

    def boss_turn(self):
        self.player['hp'] -= max(1, self.boss['dmg'] - self.player['armor'])

    def player_turn(self, spell_to_cast):
        self.boss['hp'] -= min(1, self.player['dmg'])
        if self.spells[spell_to_cast]['rounds'] == 0:
            self.apply_spell(self.spells[spell_to_cast])
        else:
            self.active_spells[spell_to_cast] = {'rounds_active': self.spells[spell_to_cast]['rounds']}
        self.player['mana'] -= self.spells[spell_to_cast]['cost']
        self.mana_spent += self.spells[spell_to_cast]['cost']

        self.spells_cast.append(spell_to_cast)

def simulate(difficulty='normal'):
    min_mana_spent = 9999
    spell_combinations = [[]]
    w = Wizard()

    while True:
        cast_this_round = []

        for sc in spell_combinations:
            for s in Wizard.spells:
                cast_this_round.append(sc + [s])

        spell_combinations = cast_this_round
        cast_this_round = []

        for sc in spell_combinations:
            w.init(difficulty)
            ret = 0
            for s in sc:
                ret = w.next_turn(s)
                if ret != 0:
                    break

            if ret == 1:
                min_mana_spent = min(min_mana_spent, w.mana_spent)
            elif ret == 0:
                cast_this_round.append(sc)

        spell_combinations = cast_this_round

        if min_mana_spent != 9999:
            break

    return min_mana_spent

part_1 = simulate()
part_2 = simulate('difficult')
print('Part 1: Player can win with spending a minimum of {} mana'.format(part_1))
print('Part 2: Player can win with spending a minimum of {} mana'.format(part_2))

assert part_1 == 1824
