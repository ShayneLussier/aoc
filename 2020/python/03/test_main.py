import pytest

from main import Slope, parseHill, parseSlope

SAMPLE = """
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
"""

@pytest.mark.parametrize(
    "input, want",
    [
        (SAMPLE, 7)
    ],
)

def testParseSlope(tmp_path, input, want):
    path = tmp_path / "input.txt"
    path.write_text(input)
    assert parseSlope(slope=Slope(x=3, y=1), hill=parseHill(path)) == want