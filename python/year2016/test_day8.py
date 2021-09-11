# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day8
import unittest


class Test_Day8(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(day8.part1("X"), "not implemented")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "b": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day8.part2(input), expected)

    def test_rect(self):
        display = day8.create_display(7, 3)
        day8.rect("3x2", display)
        expected = """###....
                      ###....
                      ......."""

        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_rotate_column1(self):
        input = """###....
                   ###....
                   ......."""
        display = self.parse_display(input)
        day8.rotate_column("1 by 1", display)
        expected = """#.#....
                      ###....
                      .#....."""

        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_rotate_column2(self):
        display = day8.create_display(1, 3)
        display[0][0] = True
        day8.rotate_column("0 by 10", display)
        expected = """.
                      #
                      ."""

        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_rotate_row1(self):
        input = """#.#....
                   ###....
                   .#....."""
        display = self.parse_display(input)
        day8.rotate_row("0 by 4", display)
        expected = """....#.#
                      ###....
                      .#....."""

        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_rotate_row2(self):
        display = day8.create_display(3, 1)
        display[0][0] = True
        day8.rotate_row("0 by 10", display)
        expected = ".#."

        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_combined(self):
        display = day8.create_display(7, 3)
        day8.rect("3x2", display)
        day8.rotate_column("1 by 1", display)
        day8.rotate_row("0 by 4", display)
        day8.rotate_column("1 by 1", display)

        expected = """.#..#.#
                      #.#....
                      .#....."""
        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_process_input(self):
        display = day8.create_display(7, 3)
        input = """rect 3x2
                   rotate column x=1 by 1
                   rotate row y=0 by 4
                   rotate column x=1 by 1"""
        day8.process_input(self.strip_lines(input), display)

        expected = """.#..#.#
                      #.#....
                      .#....."""
        self.assertEqual(day8.render_display(display), self.strip_lines(expected))

    def test_count_on_pixels_all_off(self):
        self.assertEqual(day8.count_on_pixels(day8.create_display(7, 3)), 0)

    def test_count_on_pixels_some_on(self):
        input = """.#..#.#
                   #.#....
                   .#....."""
        display = self.parse_display(input)
        self.assertEqual(day8.count_on_pixels(display), 6)

    def strip_lines(self, s: str) -> str:
        return "\n".join([line.strip() for line in s.split("\n")]) + "\n"

    def parse_display(self, s: str) -> list:
        out = list()
        for line in s.splitlines():
            row = list()
            for c in line.strip():
                row.append(True if c == '#' else False)
            out.append(row)
        return out


if __name__ == '__main__':
    unittest.main()
