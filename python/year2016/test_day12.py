# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day12
import unittest


class Test_Day12(unittest.TestCase):
    def test_part1(self):
        input = """
cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
"""
        self.assertEqual(day12.part1(input), 42)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day12.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
