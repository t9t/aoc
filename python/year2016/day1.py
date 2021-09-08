
def part1(input):
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
        if face == 0:  # North
            y -= walk
        elif face == 1:  # East
            x += walk
        elif face == 2:  # South
            y += walk
        elif face == 3:  # West
            x -= walk
    return abs(x)+abs(y)


def part2(input):
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

            key = "{0},{1}".format(x, y)
            if key in visited:
                return abs(int(x))+abs(int(y))
            else:
                visited.add(key)

    raise Exception("no location visited twice")
