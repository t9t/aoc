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
        self.assertEqual(day22.part1(input.strip()), 3)

    def test_part2(self):
        input = """
Filesystem            Size  Used  Avail  Use%
/dev/grid/node-x0-y0   10T    8T     2T   80%
/dev/grid/node-x0-y1   11T    6T     5T   54%
/dev/grid/node-x0-y2   32T   28T     4T   87%
/dev/grid/node-x1-y0    9T    7T     2T   77%
/dev/grid/node-x1-y1    8T    0T     8T    0%
/dev/grid/node-x1-y2   11T    7T     4T   63%
/dev/grid/node-x2-y0   10T    6T     4T   60%
/dev/grid/node-x2-y1    9T    8T     1T   88%
/dev/grid/node-x2-y2    9T    6T     3T   66%
        """
        self.assertEqual(day22.part2(input.strip()), 7)


if __name__ == '__main__':
    unittest.main()
