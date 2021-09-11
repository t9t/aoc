
def part1(input: str):
    display = create_display(50, 6)
    process_input(input, display)
    return count_on_pixels(display)


def part2(input: str):
    display = create_display(50, 6)
    process_input(input, display)
    parts = split_horizontally(display, 5)
    return "".join([parse_display_letter(part) for part in parts])


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


def split_horizontally(display: list, width: int) -> list:
    if len(display[0]) % width != 0:
        raise Exception("Display width {0} not divisible by len {1}".format(len(display[0], width)))
    out = list()
    for part in range(0, int(len(display[0]) / width)):
        start = part*width
        nd = create_display(width, len(display))
        for y in range(0, len(display)):
            for x in range(0, width):
                nd[y][x] = display[y][start+x]
        out.append(nd)
    return out


def parse_display_letter(display: list) -> str:
    rendered = render_display(display)
    letter = parse_letter(rendered)
    return letter


def parse_letter(s: str) -> str:
    s = s.strip()
    if s not in letter_map:
        raise Exception("Unknown letter: {0}".format(s))
    return letter_map[s]


letter_map = {
    ".##..\n#..#.\n#..#.\n####.\n#..#.\n#..#.": 'A',
    "###..\n#..#.\n###..\n#..#.\n#..#.\n###..": 'B',
    ".##..\n#..#.\n#....\n#....\n#..#.\n.##..": 'C',
    "###..\n#..#.\n#..#.\n#..#.\n#..#.\n###..": 'D',  # guessed
    "####.\n#....\n###..\n#....\n#....\n####.": 'E',
    "####.\n#....\n###..\n#....\n#....\n#....": 'F',
    ".##..\n#..#.\n#....\n#.##.\n#..#.\n.###.": 'G',
    "#..#.\n#..#.\n####.\n#..#.\n#..#.\n#..#.": 'H',
    ".###.\n..#..\n..#..\n..#..\n..#..\n.###.": 'I',
    "..##.\n...#.\n...#.\n...#.\n#..#.\n.##..": 'J',
    "#..#.\n#.#..\n##...\n#.#..\n#.#..\n#..#.": 'K',
    "#....\n#....\n#....\n#....\n#....\n####.": 'L',
    # M missing
    "#...#\n##..#\n#.#.#\n#.#.#\n#..##\n#...#": 'N',  # guessed
    ".##..\n#..#.\n#..#.\n#..#.\n#..#.\n.##..": 'O',
    "###..\n#..#.\n#..#.\n###..\n#....\n#....": 'P',
    # Q missing
    "###..\n#..#.\n#..#.\n###..\n#.#..\n#..#.": 'R',
    ".###.\n#....\n#....\n.##..\n...#.\n###..": 'S',
    "#####\n..#..\n..#..\n..#..\n..#..\n..#..": 'T',  # guessed
    "#..#.\n#..#.\n#..#.\n#..#.\n#..#.\n.##..": 'U',
    "#...#\n#...#\n#...#\n.#.#.\n.#.#.\n..#..": 'V',  # guessed
    # W missing
    # X missing
    "#...#\n#...#\n.#.#.\n..#..\n..#..\n..#..": 'Y',
    "####.\n...#.\n..#..\n.#...\n#....\n####.": 'Z'
}
