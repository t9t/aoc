# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day9
import unittest


class Test_Day9(unittest.TestCase):
    def test_part1(self):
        cases = {
            "ADVENT": 6,
            "A(1x5)BC": 7,
            "(3x3)XYZ": 9,
            "A(2x2)BCD(2x2)EFG": 11,
            "(6x1)(1x3)A": 6,
            "X(8x2)(3x3)ABCY": 18,
            "(10x42)0123456789": 420,  # blaze it
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day9.part1(input), expected)

    def test_part2(self):
        cases = {
            "(3x3)XYZ": 9,
            "X(8x2)(3x3)ABCY": 20,
            "(27x12)(20x12)(13x14)(7x10)(1x12)A": 241920,
            "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN": 445,
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day9.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
