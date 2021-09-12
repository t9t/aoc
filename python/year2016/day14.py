import hashlib
import math
import multiprocessing as mp


def part1(input: str):
    return find_key_index(input, stretching=False)


def part2(input: str):
    return find_key_index(input, stretching=True)


def find_key_index(input: str, stretching: bool) -> int:
    keys = 0
    max_keys = 25_000  # increase on KeyError
    cache = gen_keys(input, 0, max_keys, stretching)

    for i in range(0, max_keys):
        triple = extract_first_triple(cache[i])
        if triple is None:
            continue

        quintuple = triple[1]*5
        for j in range(i+1, i+1001):
            if quintuple in cache[j]:
                keys += 1
                if keys == 64:
                    return i
                break

    raise Exception("No solution found after {} iterations".format(max_keys))


def gen_keys(input: str, start: int, total: int, stretching: bool):
    cpus = mp.cpu_count()
    batch_size = int(math.ceil(total/cpus))

    with mp.Pool(cpus) as pool:
        batches = list()
        for _ in range(cpus):
            batches.append((input, start, start+batch_size, stretching))
            start += batch_size
        out = dict()
        results = pool.starmap(gen_keys_between, batches)
        for d in results:
            out.update(d)
        return out


def gen_keys_between(input: str, min: int, max: int, stretching: bool) -> dict:
    out = dict()
    for i in range(min, max):
        out[i] = gen_key(input, i, stretching)
    return out


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
