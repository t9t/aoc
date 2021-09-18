from itertools import permutations
from collections import deque


def part1(input: str):
    lines = input.strip().splitlines()
    lines = [line[1:len(line)-1] for line in lines[1:len(lines)-1]]
    locations = dict()
    pois = set()
    for y, line in enumerate(lines):
        for x, c in enumerate(line):
            if c != '#' and c != '.':
                locations[c] = (x, y)
                if c != '0':
                    pois.add(c)

    cache = dict()
    shortest = None
    for perm in permutations(pois):
        route = ('0',) + perm
        length = calculate_path(route, lines, locations, cache)
        if shortest is None or length < shortest:
            shortest = length

    return shortest


def part2(input: str):
    return "not implemented"


def calculate_path(route: tuple, lines: list, locations: dict, cache: dict) -> int:
    length = 0
    for i in range(0, len(route) - 1):
        start, end = route[i], route[i+1]
        if (start, end) in cache:
            l = cache[(start, end)]
        else:
            l = shortest_route_between(locations[start], locations[end], lines)
            cache[(start, end)] = l
        length += l
    return length


def shortest_route_between(start: tuple, end: tuple, lines: list) -> int:
    queue = deque([(start[0], start[1], 0)])
    visited = set([start])

    while queue:
        x, y, steps = queue.popleft()

        # up, down, left, right
        for dx, dy in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
            x2, y2 = x+dx, y+dy
            if x2 == end[0] and y2 == end[1]:
                return steps+1
            if (x2, y2) in visited:
                continue
            if y2 < 0 or y2 >= len(lines):
                continue
            line = lines[y2]
            if x2 < 0 or x2 >= len(line):
                continue
            if line[x2] != '#':
                queue.append((x2, y2, steps+1))
                visited.add((x2, y2))

    raise Exception("No solution found")
