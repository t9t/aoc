
def part1(input):
    lines = input.split("\n")
    keypad = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    result = ""
    x, y = 1, 1
    for line in lines:
        line = line.strip()
        if line == "":
            continue
        for c in line:
            if c == "U" and y > 0:  # Up
                y -= 1
            elif c == "R" and x < 2:  # Right
                x += 1
            elif c == "D" and y < 2:  # Down
                y += 1
            elif c == "L" and x > 0:  # Left
                x -= 1
        result += str(keypad[y][x])

    return result


def part2(input):
    return "not implemented"
