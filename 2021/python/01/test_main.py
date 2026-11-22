import pytest

from main import parseInput, countIncreases, countWindowIncreases

SAMPLE_1 = """
199
200
208
210
200
207
240
269
260
263
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE_1, 7)
    ],
)

def testCountIncreases(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert countIncreases(parseInput(str(path))) == want

SAMPLE_2 = """
199    
200   
208
210
200
207
240 
269
260
263
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE_2, 5)
    ],
)

def testCountWindowIncreases(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert countWindowIncreases(parseInput(str(path))) == want