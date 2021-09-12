# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day13
import unittest


class Test_Day13(unittest.TestCase):
    def test_min_step_to_reach_location(self):
        self.assertEqual(day13.find_something(10, target=(7, 4)), 11)

    def test_bur(self):
        self.assertEqual(day13.find_something(10, max_steps=50), 151)


if __name__ == '__main__':
    unittest.main()
