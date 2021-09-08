import main
import unittest


class Test_Main(unittest.TestCase):
    def test_find_distance_1(self):
        self.assertEqual(main.find_distance("R2, L3"), 5)

    def test_find_distance_2(self):
        self.assertEqual(main.find_distance("R2, R2, R2"), 2)

    def test_find_distance_3(self):
        self.assertEqual(main.find_distance("R5, L5, R5, R3"), 12)


if __name__ == '__main__':
    unittest.main()
