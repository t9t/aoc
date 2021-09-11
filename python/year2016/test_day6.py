# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day6
import unittest


class Test_Day6(unittest.TestCase):
    def test_part1(self):
        input = """ eedadn
                    drvtee
                    eandsr
                    raavrd
                    atevrs
                    tsrnev
                    sdttsa
                    rasrtv
                    nssdts
                    ntnada
                    svetve
                    tesnvt
                    vntsnd
                    vrdear
                    dvrsen
                    enarar"""
        self.assertEqual(day6.part1(input), "easter")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "b": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day6.part2(input), expected)

    def test_get_max(self):
        c = {'a': 2, 'm': 9, 'x': 1}
        self.assertEqual(day6.get_max(c), 'm')


if __name__ == '__main__':
    unittest.main()
