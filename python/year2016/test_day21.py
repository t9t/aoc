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

    def test_unscramble(self):
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
        self.assertEqual(day21.unscramble("decab", input.strip()), "abcde")

    ##### Scramble functions #####

    def test_password_swap_pos(self):
        self.assertEqual(day21.password_swap_pos("abcde", ('1', '3')), "adcbe")

    def test_password_swap_letter(self):
        self.assertEqual(day21.password_swap_letter("abcde", ('a', 'c')), "cbade")

    def test_password_rotate_lr_left(self):
        self.assertEqual(day21.password_rotate_lr("abcde", ("left", '2')), "cdeab")

    def test_password_rotate_lr_left_too_much(self):
        self.assertEqual(day21.password_rotate_lr("abcde", ("left", '8')), "deabc")

    def test_password_rotate_lr_right(self):
        self.assertEqual(day21.password_rotate_lr("abcde", ("right", '2')), "deabc")

    def test_password_rotate_lr_right_too_much(self):
        self.assertEqual(day21.password_rotate_lr("abcde", ("right", '8')), "cdeab")

    def test_password_rotate_pos_below_4(self):
        self.assertEqual(day21.password_rotate_pos("abdec", ('b')), "ecabd")

    def test_password_rotate_pos_4(self):
        self.assertEqual(day21.password_rotate_pos("ecabd", ('d')), "decab")

    def test_password_reverse(self):
        self.assertEqual(day21.password_reverse("abcdefg", ('2', '5')), "abfedcg")

    def test_password_move(self):
        self.assertEqual(day21.password_move("bcdea", ('1', '4')), "bdeac")

    def test_password_move_backwards(self):
        self.assertEqual(day21.password_move("bdeac", ('3', '0')), "abdec")

    ##### Unscramble functions #####

    def test_unpassword_swap_pos(self):
        self.assertEqual(day21.unpassword_swap_pos("adcbe", (1, 3)), "abcde")

    def test_unpassword_swap_letter(self):
        self.assertEqual(day21.unpassword_swap_letter("cbade", ('a', 'c')), "abcde")

    def test_unpassword_rotate_lr_left(self):
        self.assertEqual(day21.unpassword_rotate_lr("cdeab", ("left", 2)), "abcde")

    def test_unpassword_rotate_lr_left_too_much(self):
        self.assertEqual(day21.unpassword_rotate_lr("deabc", ("left", 8)), "abcde")

    def test_unpassword_rotate_lr_right(self):
        self.assertEqual(day21.unpassword_rotate_lr("deabc", ("right", 2)), "abcde")

    def test_unpassword_rotate_lr_right_too_much(self):
        self.assertEqual(day21.unpassword_rotate_lr("cdeab", ("right", 8)), "abcde")

    def test_unpassword_rotate_pos_below_4(self):
        self.assertEqual(day21.unpassword_rotate_pos("ecabd", ('b')), "abdec")

    def test_unpassword_rotate_pos_4(self):
        self.assertEqual(day21.unpassword_rotate_pos("decab", ('d')), "ecabd")

    def test_unpassword_reverse(self):
        self.assertEqual(day21.unpassword_reverse("abfedcg", (2, 5)), "abcdefg")

    def test_unpassword_move(self):
        self.assertEqual(day21.unpassword_move("bdeac", (1, 4)), "bcdea")

    def test_unpassword_move_backwards(self):
        self.assertEqual(day21.unpassword_move("abdec", (3, 0)), "bdeac")


if __name__ == '__main__':
    unittest.main()
