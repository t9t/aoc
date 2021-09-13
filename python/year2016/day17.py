from collections import deque
import hashlib


def part1(input: str):
    return find_shortest_path(input)


def part2(input: str):
    return "not implemented"


def find_shortest_path(passcode: str) -> str:
    queue = deque([(0, 0, "", determine_open_doors(passcode, ""))])
    while len(queue) > 0:
        x, y, path, open_doors = queue.popleft()

        for c in generate_candidates(x, y, path, open_doors, passcode):
            if c[0] == 3 and c[1] == 3:
                return c[2]
            queue.append(c)
    raise Exception("Did not find a solution")


def generate_candidates(x: int, y: int, path: str, open_doors: list, passcode: str):
    for d in open_doors:
        dx, dy = x + d[0], y + d[1]
        if dx < 0 or dy < 0 or dx > 3 or dy > 3:
            continue
        new_path = path + d[2]
        new_open_doors = determine_open_doors(passcode, new_path)
        yield (dx, dy, new_path, new_open_doors)


open_chars = {'b', 'c', 'd', 'e', 'f'}


def determine_open_doors(passcode: str, path: str) -> list:
    hash = md5(passcode+path)
    open = list()
    if hash[0] in open_chars:
        open.append((0, -1, 'U'))
    if hash[1] in open_chars:
        open.append((0, 1, 'D'))
    if hash[2] in open_chars:
        open.append((-1, 0, 'L'))
    if hash[3] in open_chars:
        open.append((1, 0, 'R'))
    return open


def md5(s: str) -> str:
    return hashlib.md5(s.encode("utf-8")).hexdigest()
