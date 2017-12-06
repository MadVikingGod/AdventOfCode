import pytest
from AoCDay1 import part1, part2
"""
1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
1111 produces 4 because each digit (all 1) matches the next.
1234 produces 0 because no digit matches the next.
91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
"""

@pytest.mark.parametrize('test_input,expected', [
    ("1122", 3),
    ("1111", 4),
    ("1234", 0),
    ("91212129", 9),
])
def test_day1_part1(test_input, expected):
    assert part1(test_input) == expected

    pass

"""
1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
    1221 produces 0, because every comparison is between a 1 and a 2.
    123425 produces 4, because both 2s match each other, but no other digit has a match.
    123123 produces 12.
    12131415 produces 4.
"""
@pytest.mark.parametrize('test_input,expected', [
    ("1221", 0),
    ("123425", 4),
    ("123123", 12),
    ("12131415", 4),
])
def test_day1_part2(test_input, expected):
    assert part2(test_input) == expected

    pass