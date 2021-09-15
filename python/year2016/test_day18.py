# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day18
import unittest


class Test_Day18(unittest.TestCase):
    def test_count_safe_tiles_slow_small(self):
        input = "..^^."
        self.assertEqual(day18.count_safe_tiles_slow(input, 3), 6)

    def test_count_safe_tiles_slow_large(self):
        input = ".^^.^.^^^^"
        self.assertEqual(day18.count_safe_tiles_slow(input, 10), 38)

    def test_start_count_safe_tiles_fast_small(self):
        input = "..^^."
        self.assertEqual(day18.start_count_safe_tiles_fast(input, 3), 6)

    def test_start_count_safe_tiles_fast_large(self):
        input = ".^^.^.^^^^"
        self.assertEqual(day18.start_count_safe_tiles_fast(input, 10), 38)


if __name__ == '__main__':
    unittest.main()
