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

    def test_process_c0(self):
        input = "inc c\ninc c\ncpy c a"
        self.assertEqual(day12.process(input, c=0), 2)

    def test_process_custom_c(self):
        input = """dec c\ninc c\ndec c\ncpy c a"""
        self.assertEqual(day12.process(input, c=1337), 1336)


if __name__ == '__main__':
    unittest.main()
