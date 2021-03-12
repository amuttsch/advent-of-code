import json

def parse(elem, ignore_red):
    i = 0
    for k in (elem if isinstance(elem, dict) else range(len(elem))):
        if isinstance(elem[k], int):
            i += int(elem[k])
        elif isinstance(elem[k], str):
            if ignore_red and elem[k] == 'red' and isinstance(elem, dict):
                return 0
        else:
            i += parse(elem[k], ignore_red)

    return i

data = open('input.json').read()
json_data = json.loads(data)

part_1 = parse(json_data, False)
part_2 = parse(json_data, True)

print('Part 1: {}'.format(part_1))
print('Part 2: {}'.format(part_2))

assert 119433 == part_1
assert 68466 == part_2
