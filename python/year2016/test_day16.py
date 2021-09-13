# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day16
import unittest


class Test_Day16(unittest.TestCase):
    def test_generate_disk_checksum(self):
        self.assertEqual(day16.generate_disk_checksum("10000", 20), "01100")

    def test_part2(self):
        cases = {
            "A": "not implemented",
            "B": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day16.part2(input), expected)

    def test_generate_data(self):
        cases = {
            "1": "100",
            "0": "001",
            "11111": "11111000000",
            "111100001010": "1111000010100101011110000",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day16.generate_next_data(input), expected)

    def test_calculate_checksum(self):
        self.assertEqual(day16.calculate_checksum("110010110100"), "100")

    def test_generate_disk_data(self):
        self.assertEqual(day16.generate_disk_data(20, "10000"), "10000011110010000111")


if __name__ == '__main__':
    unittest.main()
