import re

# group 1: command, 2,3 xy lower left, 4,5 xy upper right
regex = re.compile('(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)')
lights = [[0 for x in range(1000)] for x in range(1000)] # Light grid
brightness = [[0 for x in range(1000)] for x in range(1000)] # brightness grid

light = {'turn on': lambda x: 1,
         'turn off': lambda x: 0,
         'toggle': lambda x: (x+1)%2}

bright = {'turn on': lambda x: x+1,
          'turn off': lambda x: x-1 if x>0 else 0,
          'toggle': lambda x: x+2}

with open('input.txt', 'r') as file:
    for line in file:
        [op, x0, y0, x1, y1] = regex.match(line).groups()
        for x in range(int(x0), int(x1)+1):
            for y in range(int(y0), int(y1)+1):
                lights[x][y] = light[op](lights[x][y])
                brightness[x][y] = bright[op](brightness[x][y])

active_lights = sum(map(sum, lights))
total_brightness = sum(map(sum, brightness))

print('Active lights: {}'.format(active_lights))
print('Total brightness: {}'.format(total_brightness))

assert active_lights == 569999 # Part 1
assert total_brightness == 17836115 # Part 2
