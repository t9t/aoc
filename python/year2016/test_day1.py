# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day1
import unittest


class Test_Day1(unittest.TestCase):
    def test_part1(self):
        cases = {
            "R2, L3": 5,
            "R2, R2, R2": 2,
            "R5, L5, R5, R3": 12
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day1.part1(input), expected)

    def test_part2(self):
        self.assertEqual(day1.part2("R8, R4, R4, R8"), 4)


if __name__ == '__main__':
    unittest.main()
