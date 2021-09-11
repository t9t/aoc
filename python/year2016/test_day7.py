# I don't understand this import, but I have to do it to make `pyton3 -m unittest` work
from year2016 import day7
import unittest


class Test_Day7(unittest.TestCase):
    def test_part2(self):
        cases = {
            "A": "not implemented",
            "b": "not implemented",
            "C": "not implemented",
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day7.part2(input), expected)

    def test_supports_tls(self):
        cases = {
            "abba[mnop]qrst": True,
            "abcd[bddb]xyyx": False,
            "aaaa[qwer]tyui": False,
            "ioxxoj[asdfgh]zxcvbn": True,
            "dfgasd[asdfgh]zrrz[ghjasd]fgahjsd": True,
            "dfgasd[asdfgh]ashdfgja[ghjasd]zukllkjh": True,
            "dfgasd[asdfgh]zrrz[pxaaxr]fgahjsd": False,
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day7.supports_tls(input), expected)

    def test_contains_abba(self):
        cases = {
            "abba": True,
            "mnop": False,
            "qrst": False,
            "abcd": False,
            "bddb": True,
            "xyyx": True,
            "aaaa": False,
            "qwer": False,
            "tyui": False,
            "ioxxoj": True,
            "asdfgh": False,
            "zxcvbn": False,
        }
        for [input, expected] in cases.items():
            with self.subTest(input):
                self.assertEqual(day7.contains_abba(input), expected)


if __name__ == '__main__':
    unittest.main()
