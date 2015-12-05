import os
import importlib

puzzles = [{'t': '--- Day 1: Not Quite Lisp ---', 'm': 'santa'},
           {'t': '--- Day 2: I Was Told There Would Be No Math ---', 'm': 'wrapping'},
           {'t': '--- Day 3: Perfectly Spherical Houses in a Vacuum ---', 'm': 'delivery'},
           {'t': '--- Day 4: The Ideal Stocking Stuffer ---', 'm': 'adventcoins'},
           {'t': '--- Day 5: Doesn\'t He Have Intern-Elves For This? ---', 'm': 'nice'},
          ]
          
print('Advent of Code 2015\n')
          
for i, puzzle in enumerate(puzzles):
    print(puzzle['t'])
    day = 'day{:02d}'.format(i+1)
    os.chdir(day)
    importlib.import_module(day + '.' + puzzle['m'])
    os.chdir('..')
    
    print()

