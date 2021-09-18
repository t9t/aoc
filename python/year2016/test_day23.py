# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day23
import unittest


class Test_Day23(unittest.TestCase):
    def test_process(self):
        input = """
cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
        """
        self.assertEqual(day23.process(input.strip(), a=0), 3)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day23.part2(input), expected)


if __name__ == '__main__':
    unittest.main()
