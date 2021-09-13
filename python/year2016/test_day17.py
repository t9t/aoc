# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day17
import unittest


class Test_Template(unittest.TestCase):
    def test_part1(self):
        cases = {
            "ihgpwlah": "DDRRRD",
            "kglvqrro": "DDUDRLRRUDRD",
            "ulqzkmiv": "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day17.part1(input), expected)

    def test_part2(self):
        cases = {
            "ihgpwlah": 370,
            "kglvqrro": 492,
            "ulqzkmiv": 830,
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day17.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
