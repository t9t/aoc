from collections import deque
import hashlib


def part1(input: str):
    return find_shortest_path(input)


def part2(input: str):
    return len(find_longest_path(input))


def find_shortest_path(passcode: str) -> str:
    queue = deque([starting_pos(passcode)])
    while len(queue) > 0:
        room = queue.popleft()

        for c in generate_candidates(room, passcode):
            if c[0] == 3 and c[1] == 3:
                return c[2]
            queue.append(c)
    raise Exception("Did not find a solution")


def find_longest_path(passcode: str) -> str:
    return find_longest_path_from(starting_pos(passcode), passcode)


def starting_pos(passcode: str) -> tuple:
    return (0, 0, "", determine_open_doors(passcode, ""))


def find_longest_path_from(room, passcode: str) -> str:
    if room[0] == 3 and room[1] == 3:
        return room[2]

    longest = None
    candidates = list(generate_candidates(room, passcode))
    for c in candidates:
        that_longest = find_longest_path_from(c, passcode)
        if that_longest and (longest is None or len(that_longest) > len(longest)):
            longest = that_longest
    return longest


def generate_candidates(room: tuple, passcode: str):
    x, y, path, open_doors = room
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
