import re

row, column = map(int, re.match(r'\D*(\d+)\D*(\d+)', open('input.txt').read()).groups())

code_position = (row*(row+1))//2 - row + 1 + (column-1) * row - column-1 + (column*(column+1))//2 + 1

code = 20151125
for _ in range(2, code_position+1):
    code = (code * 252533) % 33554393

print('Part 1: The code for the snow maschine is {}'.format(code))

assert code == 9132360
