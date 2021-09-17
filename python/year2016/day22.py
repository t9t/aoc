import re


def part1(input: str):
    regex = r"(.+?)\s+(?:\d+)T\s+(\d+)T\s+(\d+)T"
    matches = list(re.finditer(regex, input.strip(), re.MULTILINE))
    drives = [(g[0], int(g[1]), int(g[2])) for g in (m.groups() for m in matches)]
    count = 0
    for a_name, a_used, a_avail in drives:
        if a_used == 0:
            continue
        for b_name, b_used, b_avail in drives:
            if a_name == b_name:
                continue
            b_used, b_avail = int(b_used), int(b_avail)
            if b_avail >= a_used:
                count += 1

    return count


def part2(input: str):
    return "not implemented"
