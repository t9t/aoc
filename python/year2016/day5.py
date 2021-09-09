import hashlib
import multiprocessing as mp


def part1(input):
    return find_door_password(input.strip(), part2=False)


def part2(input):
    return find_door_password(input.strip(), part2=True)


def calculate_hash(s: str) -> str:
    return hashlib.md5(s.encode("utf-8")).hexdigest()


def is_positionally_matching_id(id: str, n: int):
    hash = calculate_hash(id+str(n))
    if hash.startswith("00000"):
        pos = hash[5]
        if pos >= '0' and pos <= '7':
            return (int(pos), hash[6])
    return None


def find_hashes(id, start, end):
    results = dict()
    for n in range(start, end):
        hash = hashlib.md5((id + str(n)).encode("utf-8")).hexdigest()
        if hash.startswith("00000"):
            results[n] = hash
    return results


def find_door_password(id, part2):
    c = mp.cpu_count()
    bs = 500_000
    with mp.Pool(c) as pool:
        n = 0
        password = list("________")
        i = 0
        while True:
            batches = list()
            for _ in range(0, c):
                batches.append((id, n, n+bs))
                n += bs

            d = dict()
            for results in pool.starmap(find_hashes, batches):
                for (k, v) in results.items():
                    d[k] = v

            l = sorted(d.items(), key=lambda i: i[0])
            for (_, hash) in l:
                pos = i
                i += 1
                if part2:
                    pos = hash[5]
                    if pos >= '0' and pos <= '7':
                        pos = int(pos)
                    else:
                        continue

                if password[pos] == "_":
                    password[pos] = hash[6] if part2 else hash[5]
                    if "_" not in password:
                        return "".join(password)
