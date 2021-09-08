# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import template
import unittest


class Test_Template(unittest.TestCase):
    def test_part1(self):
        cases = {
            "A": "not implemented",
            "b": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(template.part1(input), expected)

    def test_part2(self):
        self.assertEqual(template.part2("X"), "not implemented")


if __name__ == '__main__':
    unittest.main()
