import re


def part1(input: str):
    regex = r"(.+?)\s+(?:\d+)T\s+(\d+)T\s+(\d+)T"
    matches = list(re.finditer(regex, input.strip(), re.MULTILINE))
    drives = [(g[0], int(g[1]), int(g[2])) for g in (m.groups() for m in matches)]
    count = 0
    for a_name, a_used, a_avail in drives:
        if a_used == 0:
            continue
        for b_name, b_used, b_avail in drives:
            if a_name == b_name:
                continue
            b_used, b_avail = int(b_used), int(b_avail)
            if b_avail >= a_used:
                count += 1

    return count


def part2(input: str):
    grid = dict()
    xmax, ymax = 0, 0
    empty_node = None
    for node in parse_nodes(input):
        if not (y := node[1]) in grid:
            grid[y] = dict()
        grid[y][(x := node[0])] = node[2:]
        if node[3] == 0:
            empty_node = node
        if x > xmax:
            xmax = x
        if y > ymax:
            ymax = y

    # Step 1: "move" empty node "up" to the top row
    empty_x, empty_y, _, _, _ = empty_node
    empty = empty_node[2:]
    moves = 0
    while empty_y > 0:
        moves += 1
        empty_size, _, _ = empty
        y_above = empty_y - 1
        node = grid[y_above][empty_x]
        _, above_used, _ = node
        if empty_size >= above_used:
            # The node above's data can fit the empty node, move data down (or "move empty node up")
            empty = node
            empty_y = y_above
            continue
        # The node above's data cannot fit on empty node, move left (assuming this is valid for all inputs)
        empty_x -= 1
        empty = grid[empty_x][empty_y]

    # Step 2: "move" empty node "to the right"; we assume no obstacles here, so we just add the steps to get to the right
    steps_right = xmax - empty_x

    # Step 3: "move" empty node "to the left"; assuming no obstacles, moving one spot takes 5 steps and we're already 1 step towards the left
    steps_left = (xmax-1)*5
    total = moves + steps_right+steps_left
    return total


def parse_nodes(input: str) -> list:
    # 0 = x; 1 = y; 2 = Size; 3 = Used; 4 = Avail
    regex = r"/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T"
    matches = re.finditer(regex, input.strip(), re.MULTILINE)
    nodes = [tuple(int(i) for i in g) for g in (m.groups() for m in matches)]
    return nodes
