import hashlib
import multiprocessing as mp


def part1(input):
    return find_door_password(input.strip(), is_next_matching_id)


def part2(input):
    return find_door_password(input.strip(), is_positionally_matching_id)


def calculate_hash(s: str) -> str:
    return hashlib.md5(s.encode("utf-8")).hexdigest()


def is_next_matching_id(id: str, n: int):
    hash = calculate_hash(id+str(n))
    if hash.startswith("00000"):
        return (None, hash[5])
    return None


def is_positionally_matching_id(id: str, n: int):
    hash = calculate_hash(id+str(n))
    if hash.startswith("00000"):
        pos = hash[5]
        if pos >= '0' and pos <= '7':
            return (int(pos), hash[6])
    return None


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
