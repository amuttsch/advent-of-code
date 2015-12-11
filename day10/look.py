input_sequence = '1113222113'

def look_and_say(times):
    result = sequence = input_sequence
    for _ in range(times):
        sequence = result
        current = result = ''
        n = 0
        for c in sequence:
            if current != c and n > 0:
                result += str(n) + current
                current = c
                n = 1
            else:
                current = c
                n += 1

        result += str(n) + current

    return result

part_1 = len(look_and_say(40))
part_2 = len(look_and_say(50))
print('Part 1: Length = {}'.format(part_1))
print('Part 2: Length = {}'.format(part_2))

assert 252594 == part_1
assert 3579328 == part_2
