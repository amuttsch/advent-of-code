import os

print('--- Day 1: Not Quite Lisp ---')

from day01 import santa

print()
print('--- Day 2: I Was Told There Would Be No Math ---')

os.chdir('day02')
from day02 import wrapping
os.chdir('..')

print()
print('--- Day 3: Perfectly Spherical Houses in a Vacuum ---')

os.chdir('day03')
from day03 import delivery
os.chdir('..')

print()
print('--- Day 4: The Ideal Stocking Stuffer ---')

os.chdir('day04')
from day04 import adventcoins
os.chdir('..')

print()
print('--- Day 5: Doesn\'t He Have Intern-Elves For This? ---')

os.chdir('day05')
from day05 import nice
os.chdir('..')