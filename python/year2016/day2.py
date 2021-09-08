
def part1(input):
    keypad = {
        0: {0: "1", 1: "2", 2: "3"},
        1: {0: "4", 1: "5", 2: "6"},
        2: {0: "7", 1: "8", 2: "9"}}
    return get_keypad_code(input, keypad, 1, 1)


def part2(input):
    keypad = {
        0: {2: "1"},
        1: {1: "2", 2: "3", 3: "4"},
        2: {0: "5", 1: "6", 2: "7", 3: "8", 4: "9"},
        3: {1: "A", 2: "B", 3: "C"},
        4: {2: "D"}}
    return get_keypad_code(input, keypad, 0, 2)


def get_keypad_code(input, keypad, startx, starty):
    lines = input.split("\n")
    result = ""
    x, y = startx, starty
    for line in lines:
        line = line.strip()
        if line == "":
            continue
        for c in line:
            dx, dy = 0, 0
            if c == "U":  # Up
                dy = -1
            elif c == "R":  # Right
                dx = 1
            elif c == "D":  # Down
                dy = 1
            elif c == "L":  # Left
                dx = -1

            newx, newy = x+dx, y+dy
            if newy in keypad:
                if newx in keypad[newy]:
                    y = newy
                    x = newx
        result += keypad[y][x]

    return result
