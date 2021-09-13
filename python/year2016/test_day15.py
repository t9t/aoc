# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day15
import unittest


class Test_Day15(unittest.TestCase):
    def test_parse_discs(self):
        input = """
Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.
"""
        expected = [(5, 4), (2, 1)]
        self.assertEqual(day15.parse_discs(input.strip()), expected)

    def test_first_time_I_can_press_the_button(self):
        self.assertEqual(day15.first_time_I_can_press_the_button([(5, 4), (2, 1)]), 5)


if __name__ == '__main__':
    unittest.main()
