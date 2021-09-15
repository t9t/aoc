
from collections import deque


def part1(input: str):
    ranges = list()
    for line in input.strip().splitlines():
        low, high = line.split("-")
        ranges.append((int(low), int(high)))
    ranges.sort()

    lowest = 0
    for low, high in ranges:
        if lowest >= low and lowest <= high:
            lowest = high+1
    return lowest


def part2(input: str):
    return number_of_allowed_ips(input, 4294967295)


def number_of_allowed_ips(input: str, max_ip: int):
    ranges = list()
    for line in input.strip().splitlines():
        low, high = line.split("-")
        ranges.append((int(low), int(high)))

    ranges = consolidate(sorted(ranges))

    allowed = 0
    for i in range(len(ranges)-1):
        allowed += (ranges[i+1][0] - ranges[i][1] - 1)

    allowed += (max_ip - ranges[len(ranges) - 1][1])

    return allowed


def consolidate(ranges):
    combined = list()

    while len(ranges) > 0:
        first = ranges[0]
        no_overlap = list()
        for other in ranges[1:]:
            if first[1]+1 >= other[0]:
                first = (min(first[0], other[0]), max(first[1], other[1]))
            else:
                no_overlap.append(other)
        combined.append(first)
        ranges = no_overlap
    return sorted(combined)
