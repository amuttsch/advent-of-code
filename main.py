#!/usr/bin/python3

import os
import importlib
import sys

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
           {'t': '--- Day 12: JSAbacusFramework.io ---', 'm': 'accounting'},
           {'t': '--- Day 13: Knights of the Dinner Table ---', 'm': 'knights'}
          ]

def execute_day(day):
    print(puzzles[day-1]['t'])
    day_dir = 'day{:02d}'.format(day)
    os.chdir(day_dir)
    importlib.import_module(day_dir + '.' + puzzles[day-1]['m'])
    os.chdir('..')

    print()

print('\n *** Advent of Code 2015 Solutions ***\n')

if len(sys.argv) == 1:
    for i in range(len(puzzles)):
        execute_day(i+1)
elif len(sys.argv) == 2:
    try:
        day = int(sys.argv[1])
        execute_day(day)
    except ValueError:
        print('Error: Parameter day must be numeric!')
    except IndexError:
        print('Error: Day not found, available days: 1-{}'.format(len(puzzles)))
else:
    print('Usage: main.py <day> (optional)')

print()
