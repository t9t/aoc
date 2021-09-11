import re

regex_hypernets = r"\[(\w+)\]"


def part1(input: str):
    lines = [s.strip() for s in input.strip().splitlines()]
    c = 0
    for line in lines:
        tls = supports_tls(line)
        if tls:
            c += 1

    return c


def part2(input: str):
    return "not implemented"


def supports_tls(ip: str) -> bool:
    hypernet_matches = re.findall(regex_hypernets, ip)
    for match in hypernet_matches:
        if contains_abba(match):
            return False

    supernets = re.sub(regex_hypernets, ",", ip, 0).split(",")
    for supernet in supernets:
        if contains_abba(supernet):
            return True

    return False


def contains_abba(s: str) -> bool:
    for i in range(0, len(s)-1):
        c1, c2 = s[i], s[i+1]
        if c1 == c2:
            continue
        abba = c1 + c2 + c2 + c1
        if abba in s:
            return True
    return False
