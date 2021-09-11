import re
from sys import hash_info

regex_value = r"value (\d+) goes to bot (\d+)"
regex_action = r"bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)"


def part1(input: str):
    return process_bots(input, find=(17, 61))


def part2(input: str):
    return process_bots(input, find=None)


def process_bots(input: str, find) -> int:
    bots = {}
    outputs = {0: [], 1: [], 2: []}

    def ensure_bot(num):
        if num not in bots:
            bots[num] = Bot()
        return bots[num]

    def ensure_output(num):
        if num not in outputs:
            outputs[num] = []
        return outputs[num]

    matches = re.finditer(regex_value, input, re.MULTILINE)
    for match in matches:
        value, num = [int(s) for s in match.groups()]
        ensure_bot(num).give(value)

    matches = list(re.finditer(regex_action, input, re.MULTILINE))
    i, max = 1, 100
    while i <= max:
        i += 1

        for match in matches:
            src_bot_num, low_out_type, low_out_num, high_out_type, high_out_num = match.groups()
            src_bot_num, low_out_num, high_out_num = int(src_bot_num), int(low_out_num), int(high_out_num)

            src_bot = ensure_bot(src_bot_num)
            if not src_bot.has_two():
                continue

            low, high = src_bot.get_and_clear()
            if find != None and low == find[0] and high == find[1]:
                return src_bot_num

            if low_out_type == "bot":
                ensure_bot(low_out_num).give(low)
            else:
                ensure_output(low_out_num).append(low)

            if high_out_type == "bot":
                ensure_bot(high_out_num).give(high)
            else:
                ensure_output(high_out_num).append(high)

            if not find:
                one, two, three = outputs[0], outputs[1], outputs[2]
                if len(one) > 0 and len(two) > 0 and len(three) > 0:
                    return one[0] * two[0] * three[0]

    raise Exception("Found no answer after {} iterations".format(max))


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
