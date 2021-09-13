
def part1(input: str):
    room = generate_room(input.strip(), 40)
    return count_safe_tiles(room)


def part2(input: str):
    return "not implemented"


def generate_room(input: str, rows: int) -> str:
    room = input + "\n"
    previous = input
    for _ in range(1, rows):
        this = ""
        for x in range(len(input)):
            left = previous[x-1] if x > 0 else "."
            center = previous[x]
            right = previous[x+1] if x < len(input)-1 else "."

            traps = {"^^.", ".^^", "^..", "..^"}
            if left+center+right in traps:
                this += "^"
            else:
                this += "."

        room += this + "\n"
        previous = this

    return room


def count_safe_tiles(room: str) -> int:
    return room.count(".")
