import hashlib


def part1(input):
    return find_door_password(input.strip())


def part2(input):
    return "not implemented"


def find_door_password(id: str) -> str:
    max = 100_000_000
    n = 0
    password = ""
    while True:
        hash = hashlib.md5((id + str(n)).encode("utf-8"))
        digest = hash.hexdigest()
        if digest.startswith("00000"):
            password += digest[5]
            if len(password) == 8:
                return password
        n += 1

        if n >= max:
            raise Exception("No door password found after {0} iterations".format(max))
