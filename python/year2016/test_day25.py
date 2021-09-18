# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day25
import unittest


class Test_Day25(unittest.TestCase):
    def test_part1_simplest(self):
        input = """
out 0
out 1
jnz 1 -2
"""
        self.assertEqual(day25.part1_slower(input.strip()), 0)


if __name__ == '__main__':
    unittest.main()
