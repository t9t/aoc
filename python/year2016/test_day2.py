# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day2
import unittest


class Test_Day1(unittest.TestCase):
    def test_part1(self):
        input = """ULL
                   RRDDD
                   LURDL
                   UUUUD"""
        self.assertEqual(day2.part1(input), "1985")

    def test_part2(self):
        input = """ULL
                   RRDDD
                   LURDL
                   UUUUD"""
        self.assertEqual(day2.part2(input), "5DB3")


if __name__ == '__main__':
    unittest.main()
