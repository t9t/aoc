# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day13
import unittest


class Test_Day13(unittest.TestCase):
    def test_min_step_to_reach_location(self):
        self.assertEqual(day13.min_step_to_reach_location(10, 7, 4), 11)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day13.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
