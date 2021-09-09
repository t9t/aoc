import hashlib
import multiprocessing as mp


def part1(input):
    return find_door_password1(input.strip())


def part2(input):
    return find_door_password(input.strip(), is_positionally_matching_id)


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


def find_door_password1(id):
    c = mp.cpu_count()
    bs = 500_000
    with mp.Pool(c) as pool:
        n = 0
        d = dict()
        while True:
            batches = list()
            for _ in range(0, c):
                batches.append((id, n, n+bs))
                n += bs

            for results in pool.starmap(find_hashes, batches):
                for (k, v) in results.items():
                    d[k] = v

            if len(d) >= 8:
                l = sorted(d.items(), key=lambda i: i[0])[:8]
                return "".join([i[1][5] for i in l])


def find_door_password(id: str, matching_fun) -> str:
    with mp.Pool(mp.cpu_count()) as pool:
        max = 100_000_000
        batch_size = 500_000
        n = 0
        i = 0
        password = list("________")
        while True:
            results = pool.starmap(matching_fun, [(id, i) for i in range(n, n+batch_size)])
            for r in results:
                if not r:
                    continue

                pos, char = r[0], r[1]
                if pos is None:
                    pos = i
                    i += 1

                if password[pos] != "_":
                    continue

                password[pos] = char
                if "_" not in password:
                    return "".join(password)

            n += batch_size

            if n >= max:
                raise Exception("No door password found after {0} iterations".format(max))
