
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
    return "not implemented"
