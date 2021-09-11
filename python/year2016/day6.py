
def part1(input: str):
    lines = [s.strip() for s in input.strip().splitlines()]
    counts = [dict() for c in lines[0]]
    for line in lines:
        for (i, c) in enumerate(line):
            pos = counts[i]
            if c not in pos:
                pos[c] = 1
            else:
                pos[c] += 1

    message = ""
    for c in counts:
        message += get_max(c)
    return message


def part2(input: str):
    return "not implemented"


def get_max(c: dict) -> str:
    return sorted(c.items(), key=lambda i: -1*i[1])[0][0]
