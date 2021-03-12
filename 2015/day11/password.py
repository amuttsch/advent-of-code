initial_password = 'vzbxkghb'

def validate_password(pw):
    def increasing():
        for i, c in enumerate(pw[:-2]):
            if ord(c)+1 == ord(pw[i+1]) and ord(c)+2 == ord(pw[i+2]):
                return True

        return False

    def invalid_chars():
        if 'i' in pw or 'o' in pw or 'l' in pw:
            return False
        else:
            return True

    def overlapping():
        overlapping_chars = set()
        for i, c in enumerate(pw[:-1]):
            if c == pw[i+1]:
                overlapping_chars.add(c)

        if len(overlapping_chars) >= 2:
            return True
        else:
            return False

    return increasing() and invalid_chars() and overlapping()

def generate_next_password(pw):
    pw_ascii = list(map(lambda x: ord(x), pw))
    for i, c in enumerate(pw_ascii[::-1]):
        n = (c-97 + 1) % 26 + 97
        pw_ascii[len(pw_ascii)-i-1] = n
        if c >= 122:
            pw_ascii[len(pw_ascii)-i-2] += 1
        else:
            break

    return ''.join(list(map(lambda x: chr(x), pw_ascii)))

next_password = initial_password
while not validate_password(next_password):
    next_password = generate_next_password(next_password)

print('Part 1: {}'.format(next_password))
assert 'vzbxxyzz' == next_password

next_password = generate_next_password(next_password)
while not validate_password(next_password):
    next_password = generate_next_password(next_password)

print('Part 2: {}'.format(next_password))
assert 'vzcaabcc' == next_password

# Test for password validator
assert False == validate_password('hijklmmn')
assert False == validate_password('abbceffg')
assert False == validate_password('abbcegjk')
assert True == validate_password('abcdffaa')
assert True == validate_password('ghjaabcc')
assert True == validate_password('bbaabc')
assert False == validate_password('abcc')
assert generate_next_password('azz') == 'baa'
assert generate_next_password('a') == 'b'
assert generate_next_password('ayz') == 'aza'
