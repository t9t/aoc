import re


def part1(input: str):
    return get_decompressed_length(input.strip())


def part2(input: str):
    return "not implemented"


regex_marker = r"\((\d+)x(\d+)\)"


def get_decompressed_length(s: str) -> int:
    total_length = 0
    i = 0
    while i < len(s):
        c = s[i]
        if c != '(':
            total_length += 1
            i += 1
            continue
        match = re.search(regex_marker, s[i:])
        marker_len = match.end()
        [data_len, data_times] = [int(s) for s in match.groups()]
        total_length += data_len * data_times
        i += marker_len+data_len

    return total_length
