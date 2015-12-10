from itertools import groupby

input_sequence = '1113222113'

def look_and_say(times):
    sequence = result = input_sequence
    for i in range(times):
        sequence = result
        result = ''
        for c, l in groupby(sequence):
            result += str(len(list(l))) + str(c)

    return result

part_1 = len(look_and_say(40))
part_2 = len(look_and_say(50))
print('Part 1: Length = {}'.format(part_1))
print('Part 2: Length = {}'.format(part_2))

assert 252594 == part_1
assert 3579328 == part_2
