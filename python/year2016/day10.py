import re

regex_value = r"value (\d+) goes to bot (\d+)"
regex_action = r"bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)"


def part1(input: str):
    return process_bots(input, find=(17, 61))


def part2(input: str):
    return process_bots(input, find=None)


def process_bots(input: str, find) -> int:
    bots = {}
    outputs = {0: [], 1: [], 2: []}

    def ensure(d, num):
        if num not in d:
            d[num] = list()
        return d[num]

    def ensure_bot(num):
        return ensure(bots, num)

    def apply(type, num, value):
        (ensure_bot(num) if type == "bot" else ensure(outputs, num)).append(value)

    matches = re.finditer(regex_value, input, re.MULTILINE)
    for match in matches:
        value, num = [int(s) for s in match.groups()]
        ensure_bot(num).append(value)

    matches = list(re.finditer(regex_action, input, re.MULTILINE))
    while True:
        for match in matches:
            src_bot_num, low_out_type, low_out_num, high_out_type, high_out_num = match.groups()
            src_bot_num, low_out_num, high_out_num = int(src_bot_num), int(low_out_num), int(high_out_num)

            src_bot = ensure_bot(src_bot_num)
            if len(src_bot) != 2:
                continue

            v1, v2 = src_bot
            src_bot.clear()
            low, high = min(v1, v2), max(v1, v2)
            if find != None and low == find[0] and high == find[1]:
                return src_bot_num

            apply(low_out_type, low_out_num, low)
            apply(high_out_type, high_out_num, high)

            if not find:
                one, two, three = outputs[0], outputs[1], outputs[2]
                if len(one) > 0 and len(two) > 0 and len(three) > 0:
                    return one[0] * two[0] * three[0]
