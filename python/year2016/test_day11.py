# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day11
import unittest


class Test_Day11(unittest.TestCase):
    def test_part1(self):
        input = """The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
                    The second floor contains a hydrogen generator.
                    The third floor contains a lithium generator.
                    The fourth floor contains nothing relevant."""
        self.assertEqual(day11.part1(self.strip_lines(input)), 11)

    def test_part2(self):
        input = """The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
                    The second floor contains a hydrogen generator and a lithium generator.
                    The third floor contains nothing relevant.
                    The fourth floor contains nothing relevant."""
        self.assertEqual(day11.part2(self.strip_lines(input)), 35)

    def strip_lines(self, input):
        return "\n".join([s.strip() for s in input.strip().splitlines()])


if __name__ == '__main__':
    unittest.main()
