import pytest

from main import parseMoves, calcDepth, calcAim

SAMPLE = """
forward 5
down 5
forward 8
up 3
down 8
forward 2
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 150)
    ],
)

def testCalcDepth(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert calcDepth(parseMoves(path)) == want

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 900)
    ],
)

def testCalcAim(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert calcAim(parseMoves(path)) == want