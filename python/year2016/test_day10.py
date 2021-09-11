# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day10
import unittest


class Test_Day10(unittest.TestCase):
    def test_find_bot_comparing(self):
        input = """value 5 goes to bot 2
            bot 2 gives low to bot 1 and high to bot 0
            value 3 goes to bot 1
            bot 1 gives low to output 1 and high to bot 0
            bot 0 gives low to output 2 and high to output 0
            value 2 goes to bot 2"""
        actual = day10.find_bot_comparing("\n".join([s.strip() for s in input.splitlines()]), 2, 5)
        self.assertEqual(actual, 2)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day10.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
