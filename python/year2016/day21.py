import re


def part1(input: str):
    return scramble("abcdefgh", input)


def part2(input: str):
    return unscramble("fbgdceah", input)


def scramble(password: str, input: str) -> str:
    return scramble_or_unscramble(password, parse_list(input), all_operations)


def unscramble(password: str, input: str) -> str:
    return scramble_or_unscramble(password, reverse(parse_list(input)), all_unoperations)


def scramble_or_unscramble(password: str, operations: list, _operations: list) -> str:
    for operation in operations:
        match = None
        for r, func in _operations:
            if match := re.search(r, operation):
                args = match.groups()
                out = func(password, args)
                password = out
                break
        if not match:
            raise Exception(f"Unsupported operation: {operation}")
    return password


def parse_list(input: str) -> list:
    return input.strip().splitlines()


##### Scramble functions #####

def password_swap_pos(password: str, args: tuple) -> str:
    # swap position (\d+) with position (\d+)
    i, j = int(args[0]), int(args[1])
    password = [c for c in password]
    password[i], password[j] = password[j], password[i]
    return "".join(password)


def password_swap_letter(password: str, args: tuple) -> str:
    # swap letter (\w) with letter (\w)
    password = [c for c in password]
    i, j = password.index(args[0]), password.index(args[1])
    password[i], password[j] = password[j], password[i]
    return "".join(password)


def password_rotate_lr(password: str, args: tuple) -> str:
    # rotate (left|right) (\d+) steps?
    steps = (int(args[1]) % len(password)) * (1 if args[0] == "left" else -1)
    return password[steps:] + password[:steps]


def password_rotate_pos(password: str, args: tuple) -> str:
    # rotate based on position of letter (\w)
    pos = password.index(args[0])
    return password_rotate_lr(password, ("right", pos+(2 if pos >= 4 else 1)))


def password_reverse(password: str, args: tuple) -> str:
    # reverse positions (\d+) through (\d+)
    i, j = int(args[0]), int(args[1])+1
    return password[:i] + reverse(password[i:j]) + password[j:]


def password_move(password: str, args: tuple) -> str:
    # move position (\d+) to position (\d+)
    i, j = int(args[0]), int(args[1])
    c = password[i]
    password = password[:i]+password[i+1:]
    return password[:j] + c + password[j:]


##### Unscramble functions #####

def unpassword_swap_pos(password: str, args: tuple) -> str:
    # swap position (\d+) with position (\d+)
    return password_swap_pos(password, args)


def unpassword_swap_letter(password: str, args: tuple) -> str:
    # swap letter (\w) with letter (\w)
    return password_swap_letter(password, args)


def unpassword_rotate_lr(password: str, args: tuple) -> str:
    # rotate (left|right) (\d+) steps?
    return password_rotate_lr(password, (args[0], -int(args[1])))


def unpassword_rotate_pos(password: str, args: tuple) -> str:
    # rotate based on position of letter (\w)
    for i in range(len(password)):
        unrotated = password_rotate_lr(password, ("left", i))
        scrambled = password_rotate_pos(unrotated, args)
        if scrambled == password:
            return unrotated


def unpassword_reverse(password: str, args: tuple) -> str:
    # reverse positions (\d+) through (\d+)
    return password_reverse(password, args)


def unpassword_move(password: str, args: tuple) -> str:
    # move position (\d+) to position (\d+)
    return password_move(password, (args[1], args[0]))


def reverse(s: str) -> str:
    return s[::-1]


all_operations = [
    (r"swap position (\d+) with position (\d+)", password_swap_pos),
    (r"swap letter (\w) with letter (\w)", password_swap_letter),
    (r"rotate (left|right) (\d+) steps?", password_rotate_lr),
    (r"rotate based on position of letter (\w)", password_rotate_pos),
    (r"reverse positions (\d+) through (\d+)", password_reverse),
    (r"move position (\d+) to position (\d+)", password_move),
]


all_unoperations = [
    (r"swap position (\d+) with position (\d+)", unpassword_swap_pos),
    (r"swap letter (\w) with letter (\w)", unpassword_swap_letter),
    (r"rotate (left|right) (\d+) steps?", unpassword_rotate_lr),
    (r"rotate based on position of letter (\w)", unpassword_rotate_pos),
    (r"reverse positions (\d+) through (\d+)", unpassword_reverse),
    (r"move position (\d+) to position (\d+)", unpassword_move),
]
