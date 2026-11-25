import pytest

from main import dirSizes, part1, part2

EXAMPLE: str = """\
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
"""

SIZES = {"/": 48381165, "//a": 94853, "//a/e": 584, "//d": 24933642}

def testDirSizes():
    assert dirSizes(EXAMPLE.splitlines()) == SIZES

@pytest.mark.parametrize(
    "part, want",
    [
        pytest.param(part1, 95437, id="part1"),
        pytest.param(part2, 24933642, id="part2"),
    ]
)

def testParts(part, want):
    assert part(SIZES) == want
