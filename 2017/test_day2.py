from AoCDay2 import part1, part2

"""

example spreadsheet:
5 1 9 5
7 5 3
2 4 6 8
result = 18
"""


def test_day2_part1():
    test_input = """5 1 9 5
7 5 3
2 4 6 8"""
    expected = 18
    assert part1(test_input) == expected

def test_day2_part2():
    test_input = """5 9 2 8
9 4 7 3
3 8 6 5"""
    expected = 9
    assert part2(test_input) == expected