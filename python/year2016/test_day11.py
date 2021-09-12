# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day11
import unittest


class Test_Day11(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day11.part1("X"), "not implemented")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day11.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
