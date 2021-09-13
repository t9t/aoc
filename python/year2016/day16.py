
def part1(input: str):
    return generate_disk_checksum(input, 272)


def part2(input: str):
    return generate_disk_checksum(input, 35651584)


def generate_disk_checksum(input: str, disk_size: int) -> str:
    data = generate_disk_data(disk_size, input.strip())
    return calculate_checksum(data)


def generate_disk_data(disk_size: int, input: str) -> str:
    a = input
    while len(a) < disk_size:
        a = generate_next_data(a)
    return a[:disk_size]


def generate_next_data(a: str) -> str:
    return a + "0" + "".join("1" if c == "0" else "0" for c in reversed(a))


def calculate_checksum(data: str) -> str:
    while len(data) % 2 == 0:
        data = "".join(checksum_chars(data))
    return data


def checksum_chars(data: str) -> str:
    i = iter(data)
    for c1 in i:
        c2 = next(i)
        yield "1" if c1 == c2 else "0"
