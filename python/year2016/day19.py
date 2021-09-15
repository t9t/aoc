
def part1(input: str):
    return who_gets_all_the_presents(int(input.strip()))


def part2(input: str):
    return "not implemented"


def who_gets_all_the_presents(elf_count: int) -> int:
    elves = dict()
    for i in range(elf_count):
        elves[i] = 1
    while True:
        for i in range(elf_count):
            if elves[i] == 0:
                continue
            next_i = i+1 if i < elf_count-1 else 0
            while elves[next_i] == 0:
                next_i = next_i+1 if next_i < elf_count-1 else 0
            elves[i] += elves[next_i]
            elves[next_i] = 0

            if elves[i] == elf_count:
                return i+1
