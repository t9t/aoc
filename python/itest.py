#!/usr/bin/env python3

from year2016 import *
import unittest
import sys

import year2016


class Test_All(unittest.TestCase):
    def test_2016(self):
        with open("../input/2016/results.txt") as f:
            results = f.read()
        lines = results.split("\n")
        for line in lines:
            if line == "":
                continue

            [id, expected] = line.split(":")
            [year, day, part] = id.split("-")
            with self.subTest(id):
                day_module = getattr(year2016, "day" + str(day))
                part_func = getattr(day_module, "part" + str(part))
                with open("../input/{0}/{1}.txt".format(year, day)) as f:
                    input = f.read()
                actual = str(part_func(input.strip()))
                self.assertEqual(expected.strip(), actual)


if __name__ == '__main__':
    unittest.main()
