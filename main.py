#!/usr/bin/python3

import os
import importlib

puzzles = [{'t': '--- Day 1: Not Quite Lisp ---', 'm': 'santa'},
           {'t': '--- Day 2: I Was Told There Would Be No Math ---', 'm': 'wrapping'},
           {'t': '--- Day 3: Perfectly Spherical Houses in a Vacuum ---', 'm': 'delivery'},
           {'t': '--- Day 4: The Ideal Stocking Stuffer ---', 'm': 'adventcoins'},
           {'t': '--- Day 5: Doesn\'t He Have Intern-Elves For This? ---', 'm': 'nice'},
           {'t': '--- Day 6: Probably a Fire Hazard ---', 'm': 'lights'},
           {'t': '--- Day 7: Some Assembly Required ---', 'm': 'bit'},
           {'t': '--- Day 8: Matchsticks ---', 'm': 'match'},
           {'t': '--- Day 9: All in a Single Night ---', 'm': 'delivery'},
           {'t': '--- Day 10: Elves Look, Elves Say ---', 'm': 'look'},
           {'t': '--- Day 11: Corporate Policy ---', 'm': 'password'},
           {'t': '--- Day 12: JSAbacusFramework.io ---', 'm': 'accounting'}
          ]

print('\nAdvent of Code 2015 Solutions\n')

for i, puzzle in enumerate(puzzles):
    print(puzzle['t'])
    day = 'day{:02d}'.format(i+1)
    os.chdir(day)
    importlib.import_module(day + '.' + puzzle['m'])
    os.chdir('..')

    print()
