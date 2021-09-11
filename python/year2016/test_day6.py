# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day6
import unittest


class Test_Day6(unittest.TestCase):
    _input = """ eedadn
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

    def test_part1(self):
        self.assertEqual(day6.part1(self._input), "easter")

    def test_part2(self):
        self.assertEqual(day6.part2(self._input), "advent")

    def test_get_min_or_max_min(self):
        c = {'a': 2, 'm': 9, 'x': 1}
        self.assertEqual(day6.get_min_or_max(c, min=True), 'x')

    def test_get_min_or_max_max(self):
        c = {'a': 2, 'm': 9, 'x': 1}
        self.assertEqual(day6.get_min_or_max(c, min=False), 'm')


if __name__ == '__main__':
    unittest.main()
