from collections import deque


def part1(input: str):
    return find_something(int(input.strip()), target=(31, 39))


def part2(input: str):
    return find_something(int(input.strip()), max_steps=50)


def min_step_to_reach_location(magic_number: int, target_x: int, target_y: int):
    visited = set((1, 1))
    queue = deque([(1, 1, 1)])

    while len(queue) > 0:
        steps, x, y = queue.popleft()
        for nx, ny in get_neighbors(magic_number, x, y):
            if nx == target_x and ny == target_y:
                return steps

            if (nx, ny) not in visited:
                queue.append((steps+1, nx, ny))
                visited.add((nx, ny))

    raise Exception("No solution found")


def find_something(magic_number: int, target: tuple = None, max_steps: int = None):
    visited = set([(1, 1)])
    queue = deque([(1, 1, 1)])
    within_max_steps = set([(1, 1)])

    while len(queue) > 0:
        steps, x, y = queue.popleft()
        if max_steps != None and steps > max_steps:
            continue

        for nx, ny in get_neighbors(magic_number, x, y):
            if target != None and nx == target[0] and ny == target[1]:
                return steps

            if (nx, ny) not in visited:
                queue.append((steps+1, nx, ny))
                visited.add((nx, ny))
                if max_steps != None and steps <= max_steps:
                    within_max_steps.add((nx, ny))

    if within_max_steps != None:
        return len(within_max_steps)
    else:
        raise Exception("No solution found")


def get_neighbors(n: int, x: int, y: int):
    for dx, dy in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
        nx, ny = x+dx, y+dy
        if nx >= 0 and ny >= 0 and not is_wall(n, nx, ny):
            yield (nx, ny)


def is_wall(n: int, x: int, y: int):
    return ("{0:b}".format((x*x + 3*x + 2*x*y + y + y*y) + n)).count("1") % 2 == 1
