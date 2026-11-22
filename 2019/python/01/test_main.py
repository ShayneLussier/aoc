import pytest

from main import fuelSum

@pytest.mark.parametrize(
    "input, want",
    [
        (["100756"], (33583, 50346)),
    ],
)

def testFuelSum(input, want):
    assert fuelSum(input) == want
