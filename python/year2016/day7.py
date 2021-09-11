import re

regex_hypernets = r"\[(\w+)\]"


def part1(input: str):
    return find_matches(input, supports_tls)


def part2(input: str):
    return find_matches(input, supports_ssl)


def find_matches(input: str, match_func) -> int:
    lines = [s.strip() for s in input.strip().splitlines()]
    c = 0
    for line in lines:
        if match_func(line):
            c += 1

    return c


def supports_tls(ip: str) -> bool:
    hypernets = re.findall(regex_hypernets, ip)
    for match in hypernets:
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


def supports_ssl(ip: str) -> bool:
    supernets = re.sub(regex_hypernets, ",", ip, 0).split(",")
    hypernets = re.findall(regex_hypernets, ip)

    for supernet in supernets:
        for i in range(0, len(supernet)-2):
            c1, c2, c3 = supernet[i], supernet[i+1], supernet[i+2]
            if c1 != c3 or c1 == c2:
                continue
            bab = c2+c1+c2
            for hypernet in hypernets:
                if bab in hypernet:
                    return True

    return False
