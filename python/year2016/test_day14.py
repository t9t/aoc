# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day14
import unittest


class Test_Day14(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day14.part1("abc"), 22728)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day14.part2(input), expected)

    def test_extract_first_triple(self):
        cases = {
            "dajuohfdjasd": None,
            "dazzzjuohfdjasd": "zzz",
            "dallljuohfdzzzjasd": "lll",
            "dajrrruohfdjasddd": "rrr",
            "dajuohfdjasdxxx": "xxx",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day14.extract_first_triple(input), expected)


if __name__ == '__main__':
    unittest.main()
