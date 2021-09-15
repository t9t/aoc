
def part1(input: str):
    return count_total_safe_tiles(input, 40)


def part2(input: str):
    return count_total_safe_tiles(input, 400_000)


def count_total_safe_tiles(input: str, rows: int) -> int:
    # return count_safe_tiles_slow(input, rows)
    return start_count_safe_tiles_fast(input, rows)


def count_safe_tiles_slow(input: str, rows: int) -> int:
    previous = input
    safe = input.count(".")
    for _ in range(1, rows):
        this = ""
        for x in range(len(input)):
            left = previous[x-1] if x > 0 else "."
            right = previous[x+1] if x < len(input)-1 else "."

            # While there are 3 rules, it seems that if the right and left tile are the same, the tile will be safe
            if left == right:
                this += "."
                safe += 1
            else:
                this += "^"

        previous = this

    return safe


# Borrowed from: https://old.reddit.com/r/adventofcode/comments/5iyp50/2016_day_18_solutions/dbc0l6j/
def start_count_safe_tiles_fast(input: str, rows: int) -> int:
    traps = 0
    mask = 0
    for c in input:
        traps = (traps << 1) | (c == '^')
        mask = (mask << 1) | 1
    return count_safe_tiles_fast(traps, mask, rows)


def count_safe_tiles_fast(traps, mask, rows) -> int:
    safe = 0
    for _ in range(rows):
        safe += bin(mask ^ traps).count('1')
        traps = (traps << 1) ^ (traps >> 1)
        traps &= mask
    return safe
