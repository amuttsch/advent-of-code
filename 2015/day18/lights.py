from itertools import starmap, product

def run_animation(stuck_lights=False):
    def neighbor_values(x, y):
        n = list(starmap(lambda a, b: (x+a, y+b), product((0, 1, -1), (0, 1, -1))))
        n.remove((x, y))
        return sum(map(lambda a: lights[a[1]][a[0]] if -1 not in a and 100 not in a else 0, n))

    lights = [[] for _ in range(100)]

    for i, line in enumerate(open('input.txt')):
        lights[i] = list(map(lambda x: 1 if x == '#' else 0, line.strip('\n')))

    if stuck_lights:
        lights[0][0] = lights[0][99] = lights[99][0] = lights[99][99] = 1

    for _ in range(100):
        switch = []

        for x, y in product(range(100), range(100)):
            neighbor_lights_on = neighbor_values(x, y)
            if lights[y][x] == 1 and neighbor_lights_on not in [2, 3]:
                switch.append((x, y, 0))
            elif lights[y][x] == 0 and neighbor_lights_on == 3:
                switch.append((x, y, 1))

        for x, y, a in switch:
            if not stuck_lights or (x, y) not in [(0, 0), (0, 99), (99, 0), (99, 99)]:
                lights[y][x] = a

    return sum(map(sum, lights))

total_lights_on = run_animation()
total_lights_on_with_stuck = run_animation(True)

print('Part 1: {} lights are on after 100 animations'.format(total_lights_on))
print('Part 2: {} lights are on after 100 animations with stuck edge lights'.format(total_lights_on_with_stuck))

assert total_lights_on == 821
assert total_lights_on_with_stuck == 886
