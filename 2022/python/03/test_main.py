import pytest

from main import parseItems, parseBadges

SAMPLE = """
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 157)
    ],
)

def testParseItems(tmp_path, input, want):
    assert parseItems(input.strip().splitlines()) == want

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 70)
    ],
)

def testParseBadges(tmp_path, input, want):
    assert parseBadges(input.strip().splitlines()) == want