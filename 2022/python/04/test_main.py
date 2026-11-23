import pytest

from main import parseInput, parseSections, parseSections2
from main import Section

@pytest.mark.parametrize(
    "input, want",
    [
        pytest.param("8-41,8-79", [(Section(8, 41), Section(8, 79))], id="exmaple1"),
        pytest.param("1-71,2-71", [(Section(1, 71), Section(2, 71))], id="exmaple2"),
        pytest.param("11-74,74-75", [(Section(11, 74), Section(74, 75))], id="exmaple3"),
    ]
)

def testParseInput(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert parseInput(path) == want

@pytest.mark.parametrize(
    "input, want",
    [
        pytest.param([(Section(8, 41), Section(8, 79))], 1, id="exmaple1"),
        pytest.param([(Section(1, 71), Section(2, 71))], 1, id="exmaple2"),
        pytest.param([(Section(11, 74), Section(74, 75))], 0, id="exmaple3"),
    ]
)

def testParseSections(input, want):
    assert parseSections(input) == want

@pytest.mark.parametrize(
    "input, want",
    [
        pytest.param([(Section(8, 41), Section(8, 79))], 1, id="exmaple1"),
        pytest.param([(Section(1, 71), Section(2, 71))], 1, id="exmaple2"),
        pytest.param([(Section(11, 74), Section(74, 75))], 1, id="exmaple3"),
    ]
)

def testParseSections2(input, want):
    assert parseSections2(input) == want