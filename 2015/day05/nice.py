import re

regexes = [r'(?=(\w*[aeiou]\w*){3,})(?=\w*([a-zA-Z])\2\w*)(?!\w*(ab|cd|pq|xy)\w*)',  # Part 1
           r'(?=\w*(\w\w)\w*\1\w*)(?=\w*(\w)\w\2\w*)']                               # Part 2
asserts = [236, 51]
for i, regex in enumerate(regexes):
    nice = 0
    naughty = 0

    with open('input.txt', 'r') as file:
        for line in file:
            r = re.compile(regex)
            if re.match(r, line):
                nice += 1
            else:
                naughty += 1

    print('Part {}: {} nice strings, {} naughty strings'.format(i+1, nice, naughty))
    assert nice == asserts[i]
