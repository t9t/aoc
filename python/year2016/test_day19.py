# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day19
import unittest


class Test_Day19(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day19.part1("5"), 3)

    def test_part2(self):
        self.assertEqual(day19.part2("5"), 2)


if __name__ == '__main__':
    unittest.main()
