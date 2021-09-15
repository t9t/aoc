# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day19
import unittest


class Test_Day19(unittest.TestCase):
    def test_who_gets_all_the_presents(self):
        self.assertEqual(day19.who_gets_all_the_presents(5), 3)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day19.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
