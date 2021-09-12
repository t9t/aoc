
def part1(input: str):
    return process(input, c=0)


def part2(input: str):
    return process(input, c=1)


def process(input: str, c: int) -> int:
    lines = input.strip().splitlines()

    reg_nums = {'a': 0, 'b': 1, 'c': 2, 'd': 3}

    instructions = list()
    for line in lines:
        cmd = line[0:3]
        vs = line[4:]
        func, args = None, None
        if cmd == "cpy":
            src, tgt = vs.split(" ")
            tgt_reg = reg_nums[tgt]

            if src in reg_nums:
                func = cpy_reg
                args = (tgt_reg, reg_nums[src])
            else:
                func = cpy_val
                args = (tgt_reg, int(src))
        elif cmd == "inc":
            func = inc
            args = (reg_nums[vs])
        elif cmd == "dec":
            func = dec
            args = (reg_nums[vs])
        elif cmd == "jnz":
            src, jmp = vs.split(" ")
            jmp = int(jmp)

            if src == "0":
                func = jnz_zero
                args = None
            elif src in reg_nums:
                func = jnz_reg
                args = (reg_nums[src], jmp)
            else:
                func = jnz_nonzero_val
                args = (jmp)
        instructions.append((func, args))

    regs = [0, 0, c, 0]

    ptr = 0
    l = len(instructions)
    while ptr < l:
        func, args = instructions[ptr]
        ptr += func(regs, args)

    return regs[0]


def cpy_val(regs, args):
    regs[args[0]] = args[1]
    return 1


def cpy_reg(regs, args):
    regs[args[0]] = regs[args[1]]
    return 1


def inc(regs, num):
    regs[num] += 1
    return 1


def dec(regs, num):
    regs[num] -= 1
    return 1


def jnz_zero(regs, args):
    return 1


def jnz_nonzero_val(regs, args):
    return args


def jnz_reg(regs, args):
    if regs[args[0]] != 0:
        return args[1]
    return 1
