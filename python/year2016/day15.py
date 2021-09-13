import re

import multiprocessing as mp

regex_discs = r"Disc #(?:\d) has (\d+) positions; at time=0, it is at position (\d)\."


def part1(input: str):
    discs = parse_discs(input)
    return first_time_I_can_press_the_button(discs)


def part2(input: str):
    discs = parse_discs(input)
    discs.append((11, 0))
    return first_time_I_can_press_the_button(discs)


def first_time_I_can_press_the_button(discs: list) -> int:
    max_time = 10_000_000
    for time in range(0, max_time):
        found = True
        for i, disc in enumerate(discs, start=1):
            here = time+i
            pos = disc[1] + here
            pos = pos % disc[0]
            if pos != 0:
                found = False
                break
        if found:
            return time
    raise Exception(f"No solution found after {max_time} iterations")


def parse_discs(input: str) -> list:
    return [tuple([int(s) for s in match.groups()]) for match in re.finditer(regex_discs, input, re.MULTILINE)]
