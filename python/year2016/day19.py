
def part1(input: str):
    # https://www.youtube.com/watch?v=uCsD3ZGzMgE
    # The Josephus Problem - Numberphile
    b = bin(int(input.strip()))
    return int(b[3:] + '1', 2)


def part2(input: str):
    winner = 1
    elf_count = int(input.strip())
    for i in range(1, elf_count):
        winner = (winner % i)+1
        if winner > (i+1)//2:
            winner += 1
    return winner
