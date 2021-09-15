# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day20
import unittest


class Test_Day20(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day20.part1("5-8\n0-2\n4-7\n"), 3)

    def test_number_of_allowed_ips(self):
        self.assertEqual(day20.number_of_allowed_ips("5-8\n0-2\n4-7\n", 9), 2)


if __name__ == '__main__':
    unittest.main()
