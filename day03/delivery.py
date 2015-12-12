def get_directions():
    with open('input.txt', 'r') as file:
        return file.read()

def get_houses(directions):
    current = (0, 0)
    houses = set()
    houses.add(current)
    move = lambda x, y, d: {'^': (x, y+1), 'v': (x, y-1), '<': (x-1, y), '>': (x+1, y)}.get(d, (x, y)) # switch

    for d in directions:
        current = move(current[0], current[1], d)
        houses.add(current)

    return houses

directions = get_directions()
part_1 = get_houses(directions)
part_2 = get_houses(directions[::2]) | get_houses(directions[1::2])

print('Part 1: Visited houses santa {}'.format(len(part_1)))
print('Part 2: Visited houses with robo {}'.format(len(part_2)))

assert len(part_1) == 2592 # Part 1
assert len(part_2) == 2360 # Part 2
