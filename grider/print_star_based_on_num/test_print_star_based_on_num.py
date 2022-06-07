from unittest import TestCase
from print_star_based_on_num import printstarbasedonnum


class Test(TestCase):
    def test_printstarbasedonnum_num3(self):
        res = printstarbasedonnum(3)
        self.assertEqual("1 * 3", res)

    def test_printstarbasedonnum_num5(self):
        res = printstarbasedonnum(5)
        self.assertEqual("1 * 3 4 *", res)

    def test_printstarbasedonnum_num10(self):
        res = printstarbasedonnum(10)
        self.assertEqual("1 * 3 4 * 6 * 8 9 *", res)

if __name__ == "__main__":
    unittest.main()