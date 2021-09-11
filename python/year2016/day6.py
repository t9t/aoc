
def part1(input: str):
    return decode_message(input, min=False)


def part2(input: str):
    return decode_message(input, min=True)


def decode_message(input: str, min: bool) -> str:
    lines = [s.strip() for s in input.strip().splitlines()]
    counts = [dict() for _ in lines[0]]
    for line in lines:
        for (i, c) in enumerate(line):
            pos = counts[i]
            if c not in pos:
                pos[c] = 1
            else:
                pos[c] += 1

    message = ""
    for c in counts:
        message += get_min_or_max(c, min)
    return message


def get_min_or_max(c: dict, min: bool) -> str:
    f = 1 if min else -1
    return sorted(c.items(), key=lambda i: f*i[1])[0][0]
