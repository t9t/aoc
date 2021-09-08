
def part1(input):
    return traverse(input, False)


def part2(input):
    return traverse(input, True)


def traverse(input, first_double_visited):
    visited = set()
    directions = input.split(", ")
    x, y, face = 0, 0, 0
    for dir in directions:
        turn = 1 if dir[0] == "R" else -1
        face += turn
        if face == -1:
            face = 3
        elif face == 4:
            face = 0

        walk = int(dir[1:])
        for _ in range(walk):
            if face == 0:  # North
                y -= 1
            elif face == 1:  # East
                x += 1
            elif face == 2:  # South
                y += 1
            elif face == 3:  # West
                x -= 1

            if first_double_visited:
                key = "{0},{1}".format(x, y)
                if key in visited:
                    return abs(x)+abs(y)
                else:
                    visited.add(key)

    return abs(x)+abs(y)
