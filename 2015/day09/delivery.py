cities = {}

def read_distances():
    for line in open('input.txt'):
        s = line.split(' ')
        cities.setdefault(s[0], {})[s[2]] = int(s[4])
        cities.setdefault(s[2], {})[s[0]] = int(s[4])

def brute_force_distance(f):
    total_length = 0

    def calc_path(city, visited):
        length = 0

        visited.append(city)

        for to_city in cities[city]:
            if to_city not in visited:
                l = cities[city][to_city] + calc_path(to_city, list(visited))
                if length == 0 or f(l, length):
                    length = l

        return length

    for city in cities:
        l = calc_path(city, [])
        if total_length == 0 or f(l, total_length):
            total_length = l

    return total_length

read_distances()

part_1 = brute_force_distance(lambda x, y: x < y)
part_2 = brute_force_distance(lambda x, y: x > y)
print('Part 1: Shortest distance = {}'.format(part_1))
print('Part 2: Longest distance = {}'.format(part_2))

assert part_1 == 141
assert part_2 == 736
