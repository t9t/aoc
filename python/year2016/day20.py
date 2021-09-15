
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
    ranges.sort()

    ranges = consolidate(ranges)
    ranges.sort()

    allowed = 0
    for i in range(len(ranges)-1):
        between = ranges[i+1][0] - ranges[i][1] - 1
        allowed += between

    allowed += (max_ip - ranges[len(ranges) - 1][1])

    return allowed


def overlap_or_adjacent(range1, range2):
    return range2[0] < range1[1] or range1[1]+1 == range2[0]


def combine(range1, range2):
    return (min(range1[0], range2[0]), max(range1[1], range2[1]))


def consolidate(ranges):
    combined = list()

    while True:
        first = ranges[0]
        others = ranges[1:]
        no_overlap = list()
        for other in others:
            if overlap_or_adjacent(first, other):
                first = combine(first, other)
            else:
                no_overlap.append(other)
        combined.append(first)
        ranges = no_overlap
        if len(ranges) == 1:
            combined.append(ranges[0])
            return combined
        elif len(ranges) == 0:
            return combined
