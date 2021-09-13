# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day18
import unittest


class Test_Day18(unittest.TestCase):
    def test_part_generate_room_small(self):
        input = "..^^."
        expected = """
..^^.
.^^^^
^^..^
""".lstrip()
        self.assertEqual(day18.generate_room(input, 3), expected)

    def test_part_generate_room_large(self):
        input = ".^^.^.^^^^"
        self.assertEqual(day18.generate_room(input, 10), self._big_room)

    def test_count_safe_tiles(self):
        self.assertEqual(day18.count_safe_tiles(self._big_room), 38)

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day18.part2(input), expected)

    _big_room = """
.^^.^.^^^^
^^^...^..^
^.^^.^.^^.
..^^...^^^
.^^^^.^^.^
^^..^.^^..
^^^^..^^^.
^..^^^^.^^
.^^^..^.^^
^^.^^^..^^
""".lstrip()


if __name__ == '__main__':
    unittest.main()
