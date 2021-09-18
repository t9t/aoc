# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day24
import unittest


class Test_Day24(unittest.TestCase):
    def test_part1(self):
        input = """
###########
#0.1.....2#
#.#######.#
#4.......3#
###########
"""
        self.assertEqual(day24.part1(input.strip()), 14)

    def test_part2(self):
        input = """
###########
#0.1.....2#
#.#######.#
#4.......3#
###########
"""
        self.assertEqual(day24.part2(input.strip()), 20)


if __name__ == '__main__':
    unittest.main()
