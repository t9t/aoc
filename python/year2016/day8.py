
def part1(input: str):
    display = create_display(50, 6)
    process_input(input, display)
    return count_on_pixels(display)


def part2(input: str):
    return "not implemented"


def process_input(input: str, display: list):
    process_all(input.strip().splitlines(), display)


def process_all(instructions: list, display: list):
    for instruction in instructions:
        process(instruction, display)


def process(instruction: str, display: list):
    if instruction.startswith("rect "):
        rect(instruction[len("rect "):], display)
    elif instruction.startswith("rotate row y="):
        rotate_row(instruction[len("rotate row y="):], display)
    elif instruction.startswith("rotate column x="):
        rotate_column(instruction[len("rotate column x="):], display)
    else:
        raise Exception("Invalid instruction: \"{0}\"".format(instruction))


def rect(size: str, display: list):
    [w, h] = [int(s) for s in size.split("x")]
    for y in range(0, h):
        for x in range(0, w):
            display[y][x] = True


def rotate_row(ybyn: str, display: list):
    [y, n] = [int(s) for s in ybyn.split(" by ")]
    rowlen = len(display[0])
    prev = [display[y][x] for x in range(0, rowlen)]

    for x in range(0, rowlen):
        fromx = x - (n % rowlen)
        display[y][x] = prev[fromx]


def rotate_column(xbyn: str, display: list):
    [x, n] = [int(s) for s in xbyn.split(" by ")]
    prev = [display[y][x] for y in range(0, len(display))]

    for y in range(0, len(display)):
        fromy = y - (n % len(display))
        display[y][x] = prev[fromy]


def render_display(display: list) -> str:
    s = ""
    for row in display:
        for pixel in row:
            s += '#' if pixel else '.'
        s += '\n'
    return s


def create_display(width: int, height: int) -> list:
    return [([False] * width) for _ in range(0, height)]


def count_on_pixels(display: list) -> int:
    c = 0
    for row in display:
        for pixel in row:
            if pixel:
                c += 1
    return c
