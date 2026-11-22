import pytest

from main import countBits, verifyRating

SAMPLE = """
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, ([7, 5, 8, 7, 5], [5, 7, 4, 5, 7]))
    ],
)

def testCountBits(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    lines = path.read_text().strip().splitlines()
    assert countBits(lines=lines) == want

@pytest.mark.parametrize(
    "input, criteria, want",
    [
        (SAMPLE, 1, 23),
        (SAMPLE, 0, 10),
    ],
)

def testVerifyRating(tmp_path, input, criteria, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    lines = path.read_text().strip().splitlines()
    assert verifyRating(lines=lines, criteria=criteria) == want