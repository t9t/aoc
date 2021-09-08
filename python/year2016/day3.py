
def part1(input):
    lines = [s for s in input.splitlines() if s]
    possible = 0
    for line in lines:
        [one, two, three] = [int(s.strip()) for s in line.split(" ") if s]
        if is_maybe_possible_triangle(one, two, three):
            possible += 1

    return possible


def part2(input):
    return "not implemented"


def is_maybe_possible_triangle(one, two, three):
    if one+two <= three:
        return False
    if one+three <= two:
        return False
    if two+three <= one:
        return False
    return True
