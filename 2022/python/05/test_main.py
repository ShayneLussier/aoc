import pytest
from copy import deepcopy

from main import parseInput, parseStacks, parseInstructions, crateMover9000, crateMover9001
from main import Instruction

EXAMPLE = """\
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
"""

STACKS = [["Z", "N"], ["M", "C", "D"], ["P"]]
INSTRUCTIONS = [Instruction(1, 2, 1), Instruction(3, 1, 3), Instruction(2, 2, 1), Instruction(1, 1, 2)]

def testParseStacks():
    assert parseStacks(EXAMPLE.split("\n\n")[0]) == STACKS

def testParseInstructions():
    assert parseInstructions(EXAMPLE.split("\n\n")[1]) == INSTRUCTIONS

def testParseInput(tmp_path):
    path = tmp_path / "input.txt"
    path.write_text(EXAMPLE)
    assert parseInput(path) == (STACKS, INSTRUCTIONS)

@pytest.mark.parametrize(
    "mover, want",
    [
        pytest.param(crateMover9000, "CMZ", id="crateMover9000"),
        pytest.param(crateMover9001, "MCD", id="crateMover9001"),
    ]
)

def testCrateMovers(mover, want):
    stacks = deepcopy(STACKS)  # movers mutate stacks
    assert mover(stacks, INSTRUCTIONS) == want
