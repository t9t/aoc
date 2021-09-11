# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day10
import unittest


class Test_Day10(unittest.TestCase):
    _input = """value 5 goes to bot 2
        bot 2 gives low to bot 1 and high to bot 0
        value 3 goes to bot 1
        bot 1 gives low to output 1 and high to bot 0
        bot 0 gives low to output 2 and high to output 0
        value 2 goes to bot 2"""

    def test_process_bots_find(self):
        actual = day10.process_bots(self.input(), find=(2, 5))
        self.assertEqual(actual, 2)

    def test_part2(self):
        self.assertEqual(day10.part2(self.input()), 30)

    def input(self):
        return "\n".join([s.strip() for s in self._input.splitlines()])


if __name__ == '__main__':
    unittest.main()
