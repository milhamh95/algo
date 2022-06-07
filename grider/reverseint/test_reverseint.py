from unittest import TestCase
from reverseint import reverseint

class TestSolution(TestCase):
    def test_reverseint_3(self):
        res = reverseint(3)
        self.assertEqual(res, 3)

    def test_reverseint_negative_13(self):
        res = reverseint(-13)
        self.assertEqual(res, -31)

if __name__ == "__main__":
    unittest.main()