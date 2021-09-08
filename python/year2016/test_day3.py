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

    def test_part1(self):
        input = """  1     2     3 
                     5    10    25 
                     3     4     5 
                    36    77    85 """
        self.assertEqual(day3.part1(input), 2)


if __name__ == '__main__':
    unittest.main()
