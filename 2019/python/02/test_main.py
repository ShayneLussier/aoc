from pathlib import Path

import pytest

from main import parseIntcode, parseCode, parseInputs

SAMPLE = """
1,9,10,3,
2,3,11,0,
99,
30,40,50
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 3500)
    ]
)

def testParseCodes(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert parseCode(parseIntcode(path), 12, 2) == want

def testParseInputs():
    # part 2 has no sample
    # the search only makes sense against the real puzzle input
    path = Path(__file__).parent / "input.txt"
    assert parseInputs(parseIntcode(str(path))) == 8444