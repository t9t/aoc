import hashlib


def part1(input: str):
    return find_key_index(input, stretching=False)


def part2(input: str):
    return find_key_index(input, stretching=True)


def find_key_index(input: str, stretching: bool) -> int:
    maxi = 100_000
    keys = 0
    cache = dict()

    for i in range(0, maxi):
        hash = cache.pop(i) if i in cache else gen_key(input, i, stretching)

        triple = extract_first_triple(hash)
        if triple is None:
            continue

        quintuple = triple[1]*5
        for j in range(i+1, i+1001):
            if j in cache:
                other = cache[j]
            else:
                other = gen_key(input, j, stretching)
                cache[j] = other

            if quintuple in other:
                keys += 1
                if keys == 64:
                    return i
                break

    raise Exception("No solution found after {} iterations".format(maxi))


def gen_key(input: str, i: int, stretching: bool) -> str:
    s = input+str(i)
    hash = md5(s)
    if not stretching:
        return hash
    return stretch(hash)


def md5(s: str) -> str:
    return hashlib.md5(s.encode("utf-8")).hexdigest()


def stretch(hash: str) -> str:
    for _ in range(2016):
        hash = md5(hash)
    return hash


def extract_first_triple(s: str) -> str:
    for i in range(0, len(s)-2):
        c1, c2, c3 = s[i], s[i+1], s[i+2]
        if c1 == c2 and c1 == c3:
            return c1+c2+c3
    return None
