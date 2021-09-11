# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import template
import unittest


class Test_Template(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(template.part1("X"), "not implemented")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(template.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
