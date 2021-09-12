
def part1(input: str):
    instructions = input.strip().splitlines()
    regs = {'a': 0, 'b': 0, 'c': 0, 'd': 0, }

    ptr = 0
    while ptr < len(instructions):
        instruction = instructions[ptr]
        cmd = instruction[0:3]
        vs = instruction[4:]
        jump = 1

        if cmd == "cpy":
            v, reg = vs.split(" ")
            regs[reg] = regs[v] if v in regs else int(v)
        elif cmd == "inc":
            regs[vs] += 1
        elif cmd == "dec":
            regs[vs] -= 1
        elif cmd == "jnz":
            v, n = vs.split(" ")
            v = regs[v] if v in regs else int(v)
            if v != 0:
                jump = int(n)

        ptr += jump

    return regs['a']


def part2(input: str):
    return "not implemented"
