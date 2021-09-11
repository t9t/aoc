import re


def part1(input: str):
    return get_decompressed_length(input.strip(), recurse=False)


def part2(input: str):
    return get_decompressed_length(input.strip(), recurse=True)


regex_marker = r"\((\d+)x(\d+)\)"


def get_decompressed_length(s: str, recurse: bool) -> int:
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
        data_start = i+marker_len
        data = s[data_start:data_start+data_len]
        actual_data_len = data_len
        if recurse and '(' in data:
            actual_data_len = get_decompressed_length(data, recurse=True)
        total_length += actual_data_len * data_times
        i += marker_len+data_len

    return total_length
