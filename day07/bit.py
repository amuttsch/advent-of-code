wires = {}

f = lambda x: int(x) if x.isdigit() else int(wires.get(x, 0))
bin_not = lambda x: int(0b1111111111111111 ^ x)

def power_wires(b=0):
    change = True
    while change:
        change = False
        with open('input.txt', 'r') as file:
            for line in file:
                assignment = line.replace('\n', '').split(' -> ') # [ops, dest]
                dest = assignment[1]
                ops = assignment[0].split(' ')
                old = wires.get(dest, 0)

                if 'NOT' in ops:
                    wires[dest] = bin_not(f(ops[1]))
                elif 'AND' in ops:
                    wires[dest] = f(ops[0]) & f(ops[2])
                elif 'OR' in ops:
                    wires[dest] = f(ops[0]) | f(ops[2])
                elif 'LSHIFT' in ops:
                    wires[dest] = f(ops[0]) << f(ops[2])
                elif 'RSHIFT' in ops:
                    wires[dest] = f(ops[0]) >> f(ops[2])
                else:
                    # Simple assignment: 123 -> a
                    if dest == 'b' and b != 0:
                        wires[dest] = b
                    else:
                        wires[dest] = f(assignment[0])

                if not change:
                    change = old != wires[dest]

    return wires['a']

p1 = power_wires()
print('Part 1: a={}'.format(p1))

p2 = power_wires(p1)
print('Part 2: a={}'.format(p2))

assert p1 == 3176 # Part 1
assert p2 == 14710 # Part 2
