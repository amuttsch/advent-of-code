import hashlib

input = 'ckczppom'
hash_5 = {}
hash_6 = {}
i=0

while True:
    i += 1
    key = input + str(i)
    m = hashlib.md5(str.encode(key))
    hash = m.hexdigest()

    if not 'hash' in hash_5 and hash[:5] == '0' * 5:
        hash_5['position'] = i
        hash_5['hash'] = hash

    if hash[:6] == '0' * 6:
        hash_6['position'] = i
        hash_6['hash'] = hash
        break


print('Part 1: Hash with 5 leading zeros {} at position {}'.format(hash_5['hash'], hash_5['position']))
print('Part 2: Hash with 6 leading zeros {} at position {}'.format(hash_6['hash'], hash_6['position']))

assert 117946 == hash_5['position']
assert 3938038 == hash_6['position']
