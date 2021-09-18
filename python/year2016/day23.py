

def part1(input: str):
    return process(input.strip(), a=7)


def part2(input: str):
    # return process(input.strip(), a=12) -> ~400 sec on my machine
    return sneaky_part2(input.strip())


def process(input: str, a: int) -> int:
    lines = input.strip().splitlines()

    instructions = list()
    for line in lines:
        cmd = line[0:3]
        vs = line[4:]
        func, args = None, None
        if cmd == "cpy":
            src, tgt = vs.split(" ")
            func = cpy
            args = (src, tgt)
        elif cmd == "inc":
            func = inc
            args = (vs)
        elif cmd == "dec":
            func = dec
            args = (vs)
        elif cmd == "jnz":
            src, jmp = vs.split(" ")
            func = jnz
            args = (src, jmp)
        elif cmd == "tgl":
            func = tgl
            args = vs

        instructions.append((func, args))

    regs = {'a': a, 'b': 0, 'c': 0, 'd': 0}

    ptr = 0
    l = len(instructions)
    while ptr < l:
        func, args = instructions[ptr]
        ptr += func(regs, args, ptr, instructions)

    return regs['a']


def cpy(regs, args, ptr, instructions):
    src, tgt = args
    if tgt not in regs:
        return 1

    if src in regs:
        value = regs[src]
    else:
        value = int(src)
    regs[tgt] = value
    return 1


def inc(regs, args, ptr, instructions):
    regs[args[0]] += 1
    return 1


def dec(regs, args, ptr, instructions):
    regs[args[0]] -= 1
    return 1


# jnz x y jumps to an instruction y away (positive means forward; negative means backward), but only if x is not zero.
def jnz(regs, args, ptr, instructions):
    src, jmp = args
    if src in regs:
        test = regs[src]
    else:
        test = int(src)
    if test == 0:
        return 1
    if jmp in regs:
        jmp = regs[jmp]
    return int(jmp)


def tgl(regs, args, ptr, instructions):
    n = regs[args[0]]
    t = ptr+n
    if t < 0 or t >= len(instructions):
        return 1
    target = instructions[t]
    func, fargs = target

    # For one-argument instructions, inc becomes dec, and all other one-argument instructions become inc.
    if func == dec or func == tgl:
        instructions[t] = (inc, fargs)
    elif func == inc:
        instructions[t] = (dec, fargs)
    # For two-argument instructions, jnz becomes cpy, and all other two-instructions become jnz.
    elif func == jnz:
        instructions[t] = (cpy, fargs)
    elif func == cpy:
        instructions[t] = (jnz, fargs)
    else:
        raise Exception(f"Unhandled instruction {target}")

    return 1


def sneaky_part2(input: str) -> int:
    # It looks like the pattern is 12 factorial, and then adding the product of the pair of 2-digit cpy and jnz instructions
    cpy, jnz = None, None
    for i, line in enumerate(lines := input.strip().splitlines()):
        if line.startswith("cpy ") and len(line) == 8:
            cpy = int(line[3:6])
            if i+1 < len(lines) and lines[i+1].startswith("jnz ") and len(lines[i+1]) == 8:
                jnz = int(lines[i+1][3:6])
    if not cpy or not jnz:
        raise Exception("Assumption violated, no 2-digit cpy+jnz pair found")

    # I know about math.factorial, but let's do SOME work ourselves shall we
    f, c = 11, 12
    while f > 1:
        c *= f
        f -= 1
    return c + (cpy*jnz)
