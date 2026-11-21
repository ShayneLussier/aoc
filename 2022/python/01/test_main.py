import pytest

from main import parseElves, sumTop

EXAMPLE = """1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
"""

@pytest.mark.parametrize(
    "text, want",
    [
        (EXAMPLE, [24000, 11000, 10000, 6000, 4000]),  # sorted biggest first
        ("100\n", [100]),                              # one elf, one item
        ("1\n2\n3\n", [6]),                            # one elf, items summed
        ("5\n\n7", [7, 5]),                            # no trailing newline
    ],
)

def testParseElves(tmp_path, text, want):
    path = tmp_path / "input.txt"
    path.write_text(text)
    assert parseElves(str(path)) == want

@pytest.fixture
def elves(tmp_path):
    path = tmp_path / "input.txt"
    path.write_text(EXAMPLE)
    return parseElves(str(path))

@pytest.mark.parametrize(
    "n, want",
    [
        (0, 0),
        (1, 24000),
        (3, 45000),
        (10, 55000),
    ],
)
def testSumTop(elves, n, want):
    assert sumTop(elves, n) == want
