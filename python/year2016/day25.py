
def part1(input: str):
    return sneaky_part1(input)


def part1_slower(input: str):
    for a in range(1000):
        # 64 bytes ought to be enough for anyone
        if is_probably_zero_one_repeating(input.strip(), a=a, max_out=64):
            return a
    raise Exception("No solution found")


# There is no part 2

def is_probably_zero_one_repeating(input: str, a: int, max_out: int) -> int:
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
            args = (vs,)
        elif cmd == "dec":
            func = dec
            args = (vs,)
        elif cmd == "jnz":
            src, jmp = vs.split(" ")
            func = jnz
            args = (src, jmp)
        else:
            func = out
            args = (vs,)

        instructions.append((func, args))

    regs = {'a': a, 'b': 0, 'c': 0, 'd': 0}
    l = len(instructions)
    ptr, expect_out, out_count = 0, 0, 0
    while ptr < l:
        func, args = instructions[ptr]
        jmp, outval = func(regs, args)
        ptr += jmp
        if outval == None:
            continue
        if outval != expect_out:
            return False
        out_count += 1
        if out_count >= max_out:
            return True
        expect_out = 1 if expect_out == 0 else 0

    raise Exception("Reached end of execution")


def cpy(regs, args):
    src, tgt = args
    if tgt not in regs:
        return 1

    if src in regs:
        value = regs[src]
    else:
        value = int(src)
    regs[tgt] = value
    return 1, None


def inc(regs, args):
    regs[args[0]] += 1
    return 1, None


def dec(regs, args):
    regs[args[0]] -= 1
    return 1, None


def jnz(regs, args):
    src, jmp = args
    if src in regs:
        test = regs[src]
    else:
        test = int(src)
    if test == 0:
        return 1, None
    if jmp in regs:
        jmp = regs[jmp]
    return int(jmp), None


def out(regs, args):
    tgt = args[0]
    if tgt in regs:
        value = regs[tgt]
    else:
        value = int(tgt)
    return 1, value


def sneaky_part1(input: str) -> int:
    # it turns out the program repeats the binary representation of a+(cpy X c)*(cpy Y b), so all we have to do is find
    # the first "a" where the binary representation of a+(X*Y) repeats 01.. (from the least significant bit)
    lines = input.strip().splitlines()
    t = int(lines[1][3:5]) * int(lines[2][3:7])
    for a in range(1000):
        b = bin(a+t)[2:]
        if len(b) % 2 != 0:
            continue
        m = True
        for i in range(len(b)-1, 0, -2):
            if b[i] != '0' or b[i-1] != '1':
                m = False
        if m:
            return a
    raise Exception("No solution found")
