import enum
import functools
import re


def part1(input):
    matches = re.finditer(r"(.+)-(\d+)\[(\w{5})\]", input, re.MULTILINE)
    sum = 0
    for match in matches:
        [name, sector, checksum] = match.groups()
        if is_real_room(name.strip(), checksum):
            sum += int(sector)
    return sum


def part2(input):
    return "not implemented"


def is_real_room(name, checksum):
    counts = dict()
    for c in name:
        if c == "-":
            continue
        if c in counts:
            counts[c] += 1
        else:
            counts[c] = 1
    s = sorted(list(counts.items()), key=functools.cmp_to_key(compare_counts))
    c = s[0][0] + s[1][0] + s[2][0] + s[3][0] + s[4][0]
    return c == checksum


def compare_counts(l, r):
    if l[1] == r[1]:
        return 0 if l[0] == r[0] else 1 if l[0] > r[0] else -1
    return 1 if l[1] < r[1] else -1
