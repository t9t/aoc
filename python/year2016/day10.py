import re
from sys import hash_info

regex_value = r"value (\d+) goes to bot (\d+)"
regex_action = r"bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)"


def part1(input: str):
    return find_bot_comparing(input, 17, 61)


def find_bot_comparing(input: str, find_low: int, find_high: int) -> int:
    bots = {}

    def ensure_bot(num):
        if num not in bots:
            print("    ~Creating bot", num)
            bots[num] = Bot()
        return bots[num]

    print("Setting intial values")
    matches = re.finditer(regex_value, input, re.MULTILINE)
    for match in matches:
        value, num = [int(s) for s in match.groups()]
        print("  Setting bot num", num, "to value", value)
        ensure_bot(num).give(value)

    print("Initial bots:", bots)

    print("Processing instructions")

    matches = list(re.finditer(regex_action, input, re.MULTILINE))
    i, max = 1, 100
    while i <= max:
        print("Iteration", i)
        i += 1

        for match in matches:
            print("  Processing match:", [match])
            src_bot_num, low_out_type, low_out_num, high_out_type, high_out_num = match.groups()
            src_bot_num, low_out_num, high_out_num = int(src_bot_num), int(low_out_num), int(high_out_num)

            src_bot = ensure_bot(src_bot_num)
            if not src_bot.has_two():
                print("    Source bot", src_bot_num, "does not have 2 values:", src_bot)
                continue

            low, high = src_bot.get_and_clear()
            if low == find_low and high == find_high:
                print("    Found bot about to give away low", low, "and high", high, ":", src_bot_num)
                return src_bot_num

            print("    Bot", src_bot_num, "giving away low", low, "and high", high)

            if low_out_type == "bot":
                ensure_bot(low_out_num).give(low)
                print("      Gave low", low, "to bot", low_out_num, ":", bots[low_out_num])

            if high_out_type == "bot":
                ensure_bot(high_out_num).give(high)
                print("      Gave high", high, "to bot", high_out_num, ":", bots[high_out_num])

    raise Exception("Found no answer after {} iterations".format(max))


def part2(input: str):
    return "not implemented"


class Bot:
    def __init__(self) -> None:
        self.chips = []

    def has_two(self) -> bool:
        return len(self.chips) == 2

    def get_and_clear(self):
        if not self.has_two():
            raise Exception("get_and_clear called on {} when not having 2 values (but {})", self, len(self.chips))
        num1, num2 = self.chips
        low = min(num1, num2)
        high = max(num1, num2)
        self.chips = []
        return low, high

    def give(self, val: int):
        self.chips.append(val)

    def __repr__(self) -> str:
        return "Bot({0})".format(self.__dict__)
