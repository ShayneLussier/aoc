import pytest

from main import parseGuide, playBingo

EXAMPLE = """
A Y
B X
C Z
"""
@pytest.mark.parametrize(
    "input, want",
    [
        (EXAMPLE, (15, 12)),
        ("A X", (4, 3)),
        ("A Z", (3, 8)),
        ("B Y", (5, 5)),
        ("B Z", (9, 9)),
        ("C X", (7, 2)),
        ("C Y", (2, 6)),
        ("", (0, 0)),
    ],
)

def testPlayBingo(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert playBingo(parseGuide(str(path))) == want