from math import floor

travel_time = 2503
reindeers = []

for line in open('input.txt'):
    r = line.split(' ')
    reindeer = {'reindeer': r[0],
                'speed': int(r[3]),
                'duration': int(r[6]),
                'rest': int(r[13]),
                'distance_travelled': 0,
                'score': 0}
    reindeers.append(reindeer)

    # Mathematical solution for part 1
    #d = floor(travel_time / (reindeer['duration'] + reindeer['rest']))
    #remaining = (travel_time - d * (reindeer['duration'] + reindeer['rest']))
    #travelled = d * reindeer['speed'] * reindeer['duration'] + min(remaining, reindeer['duration']) * reindeer['speed']
    #max_distance = max(max_distance, travelled)

for s in range(travel_time):
    for r in reindeers:
        t = r['duration'] + r['rest']
        f = floor(s / t)
        if f*t <= s <= f*t+r['duration']-1:
            r['distance_travelled'] += r['speed']

    reindeers = sorted(reindeers, key=lambda k: k['distance_travelled'], reverse=True)
    for r in reindeers:
        if r['distance_travelled'] == reindeers[0]['distance_travelled']:
            r['score'] += 1

part_1 = reindeers[0]
part_2 = reindeers = sorted(reindeers, key=lambda k: k['score'], reverse=True)[0]

print('Part 1: The fastest reindeer {} travelled {} km.'.format(part_1['reindeer'], part_1['distance_travelled']))
print('Part 2: The reindeer {} has the highest score of {}.'.format(part_2['reindeer'], part_2['score']))

assert part_1['distance_travelled'] == 2655
assert part_2['score'] == 1059
