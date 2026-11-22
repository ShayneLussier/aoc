# https://adventofcode.com/2020/day/3

from dataclasses import dataclass

@dataclass
class Slope:
    x: int
    y: int

def parseHill(path: str) -> list[str]:
    with open(path) as f:
        data = f.read().strip().splitlines()

    return data
            
def parseSlope(slope: Slope, hill: list[str]) -> int:
    trees: int = 0
    x: int = 0

    y = 0
    while y < len(hill):
        if hill[y][x] == '#':
            trees += 1

        y += slope.y
        x += slope.x

        x = x % len(hill[0])

    return trees

def main() -> None:
    hill = parseHill("input.txt")

    print(f"We encounter {parseSlope(slope=Slope(x=3, y=1), hill=hill)} when traversing at x=1 and y=3.") # 286

    slopes = [Slope(1, 1), Slope(3, 1), Slope(5, 1), Slope(7, 1), Slope(1, 2)]
    product = 1
    for s in slopes:
        product *= parseSlope(s, hill)
    print(f"Part 2 = {product}") # 3638606400

if __name__ == "__main__":
    main()