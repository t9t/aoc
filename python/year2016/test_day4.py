# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day4
import unittest


class Test_Day4(unittest.TestCase):
    def test_part1(self):
        input = """aaaaa-bbb-z-y-x-123[abxyz]
                   a-b-c-d-e-f-g-h-987[abcde]
                   not-a-real-room-404[oarel]
                   totally-real-room-200[decoy]"""
        self.assertEqual(day4.part1(input), 1514)

    def test_part2(self):
        input = """aaaaa-bbb-z-y-x-123[abxyz]
                   a-b-c-d-e-f-g-h-987[abcde]
                   bcfhvdczs-cpxsqh-ghcfous-480[chsfb]
                   not-a-real-room-404[oarel]
                   totally-real-room-200[decoy]"""
        self.assertEqual(day4.part2(input), 480)

    def test_is_real_room(self):
        cases = {
            ("aaaaa-bbb-z-y-x", "abxyz"): True,
            ("a-b-c-d-e-f-g-h", "abcde"): True,
            ("not-a-real-room", "oarel"): True,
            ("totally-real-room", "decoy"): False,
        }
        for [(name, checksum), expected] in cases.items():
            with self.subTest(name):
                self.assertEqual(day4.is_real_room(name, checksum), expected)

    def test_shift_very_securely(self):
        self.assertEqual(day4.shift_very_securely("qzmt-zixmtkozy-ivhz", 343), "very encrypted name")


if __name__ == '__main__':
    unittest.main()
