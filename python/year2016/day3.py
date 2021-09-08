
def part1(input):
    lines = [s for s in input.splitlines() if s]
    triangles = [[int(s.strip()) for s in line.split(" ") if s] for line in lines]
    return count_maybe_possible_triangles(triangles)


def part2(input):
    triangles = extract_numbers_columnwise(input)
    return count_maybe_possible_triangles(triangles)


def count_maybe_possible_triangles(l):
    possible = 0
    for s in l:
        if is_maybe_possible_triangle(s[0], s[1], s[2]):
            possible += 1

    return possible


def is_maybe_possible_triangle(one, two, three):
    if one+two <= three:
        return False
    if one+three <= two:
        return False
    if two+three <= one:
        return False
    return True


def extract_numbers_columnwise(input):
    lines = [s for s in input.splitlines() if s]
    lists = [[], [], []]
    n = 0
    for line in lines:
        nums = [int(s.strip()) for s in line.split(" ") if s]
        l = lists[n]
        for i in nums:
            l.append(i)
        n = n+1 if n < 2 else 0

    out = list()
    for i in range(len(lists[0])):
        out.append([lists[0][i], lists[1][i], lists[2][i]])
    return out
