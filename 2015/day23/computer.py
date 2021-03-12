class Computer():
    instruction = {
        'hlf': lambda x: x // 2,
        'tpl': lambda x: x * 3,
        'inc': lambda x: x + 1,
        'jmp': lambda r, c: c + r - 1,
        'jie': lambda r, o, c: c + o - 1 if r % 2 == 0 else c,
        'jio': lambda r, o, c: c + o - 1 if r == 1 else c
    }

    def __init__(self, a=0):
        self.instruction_list = []
        for line in open('instructions.txt'):
            self.instruction_list.append(line.strip('\n'))

        self.registers = {'a': a, 'b': 0}
        self.counter = 0

    def run(self):
        while True:
            i = self.instruction_list[self.counter].split(' ')

            if i[0] in ['hlf', 'tpl', 'inc']:
                self.registers[i[1]] = self.instruction[i[0]](self.registers[i[1]])
            elif i[0] in ['jmp']:
                self.counter = self.instruction[i[0]](int(i[1]), self.counter)
            elif i[0] in ['jie', 'jio']:
                self.counter = self.instruction[i[0]](self.registers[i[1].strip(',')], int(i[2]), self.counter)

            self.counter += 1
            if self.counter >= len(self.instruction_list):
                break

        return self.registers['b']

part_1 = Computer().run()
part_2 = Computer(a=1).run()
print('Part 1: b = {}'.format(part_1))
print('Part 2: b = {}'.format(part_2))

assert part_1 == 184
assert part_2 == 231
