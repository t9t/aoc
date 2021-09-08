
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
