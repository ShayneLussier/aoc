import pytest

from main import parseInput, checkSum, threeCheckSum

SAMPLE = """
1721
979
366
299
675
1456
"""

@pytest.mark.parametrize(
    "input, target, want",
    [
        (SAMPLE, 2020, 514579)
    ],
)

def testCheckSum(tmp_path, input, target, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert checkSum(expenses=parseInput(path), target=target) == want

@pytest.mark.parametrize(
    "input, target, want",
    [
        (SAMPLE, 2020, 241861950)
    ],
)

def testThreeCheckSum(tmp_path, input, target, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert threeCheckSum(expenses=parseInput(path), target=target) == want