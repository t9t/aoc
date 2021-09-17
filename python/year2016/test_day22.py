# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day22
import unittest


class Test_Day22(unittest.TestCase):
    def test_part1(self):
        input = """
root@ebhq-gridcenter# df -h
Filesystem              Size  Used  Avail  Use%
/dev/grid/node-x1-y1     35T    0T    35T    0%
/dev/grid/node-x1-y2     35T   10T    15T    0%
/dev/grid/node-x1-y3     35T   20T    15T    0%
"""
        # 1,1 -> 1,2: 1 is empty, not considered
        # 1,1 -> 1,3: 1 is empty, not considered

        # 1,2 -> 1,1: 10T fits in 35T
        # 1,2 -> 1,3: 10T fits in 15T

        # 1,3 -> 1,1: 20T fits in 35T
        # 1,3 -> 1,2: 20T does NOT fit in 35T
        self.assertEqual(day22.part1(input), 3)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day22.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
