import hashlib
import multiprocessing as mp


def part1(input):
    return find_door_password(input.strip())


def part2(input):
    return "not implemented"


def calculate_hash(s: str) -> str:
    return hashlib.md5(s.encode("utf-8")).hexdigest()


def is_matching_id(id: str, n: int) -> str:
    hash = calculate_hash(id+str(n))
    if hash.startswith("00000"):
        return hash[5]
    return None


def find_door_password(id: str) -> str:
    with mp.Pool(mp.cpu_count()) as pool:
        max = 100_000_000
        batch_size = 500_000
        n = 0
        password = ""
        while True:
            results = pool.starmap(is_matching_id, [(id, i) for i in range(n, n+batch_size)])
            for r in results:
                if not r:
                    continue
                password += r
                if len(password) == 8:
                    return password
            n += batch_size

            if n >= max:
                raise Exception("No door password found after {0} iterations".format(max))
