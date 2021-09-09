# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day5
import unittest


class Test_Day5(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day5.part1("abc"), "18f47a30")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "b": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day5.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
