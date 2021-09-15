# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day21
import unittest


class Test_Day21(unittest.TestCase):
    def test_scramble(self):
        input = """
swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
        """
        self.assertEqual(day21.scramble("abcde", input.strip()), "decab")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day21.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
