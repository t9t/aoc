# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day1
import unittest


class Test_Day1(unittest.TestCase):
    def test_part1_1(self):
        self.assertEqual(day1.part1("R2, L3"), 5)

    def test_part1_2(self):
        self.assertEqual(day1.part1("R2, R2, R2"), 2)

    def test_part1_3(self):
        self.assertEqual(day1.part1("R5, L5, R5, R3"), 12)


if __name__ == '__main__':
    unittest.main()
