import pytest

from main import parseInput, validatePasswords, validatePasswords2

SAMPLE = """
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 2)
    ],
)

def testValidatePasswords(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert validatePasswords(*parseInput(path)) == want

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 1)
    ],
)

def testValidatePasswords2(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert validatePasswords2(*parseInput(path)) == want