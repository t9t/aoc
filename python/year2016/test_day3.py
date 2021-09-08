# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day3
import unittest


class Test_Day3(unittest.TestCase):
    def test_is_maybe_possible_triangle(self):
        cases = {
            "1 2 3": False,
            "5 10 25": False,
            "3 4 5": True,
            "36 77 85": True,
        }
        for [input, expected] in cases.items():
            [one, two, three] = [int(s) for s in input.split(" ")]
            with self.subTest(input):
                self.assertEqual(day3.is_maybe_possible_triangle(one, two, three), expected)

    def test_extract_numbers_columnwise(self):
        input = """ 101 301 501
                    102 302 502
                    103 303 503
                    201 401 601
                    202 402 602
                    203 403 603"""
        actual = day3.extract_numbers_columnwise(input)
        expected = [[101, 102, 103], [301, 302, 303], [501, 502, 503],
                    [201, 202, 203], [401, 402, 403], [601, 602, 603]]
        self.assertEqual(actual, expected)

    def test_part1(self):
        input = """  1     2     3 
                     5    10    25 
                     3     4     5 
                    36    77    85 """
        self.assertEqual(day3.part1(input), 2)

    def test_part2(self):
        input = """  1   5   3 
                     2  10   4
                     3  25   5
                    36   6   9
                    77   9  40
                    85  30  41"""
        self.assertEqual(day3.part2(input), 3)


if __name__ == '__main__':
    unittest.main()
