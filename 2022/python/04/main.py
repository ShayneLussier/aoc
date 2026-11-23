# https://adventofcode.com/2022/day/4

from dataclasses import dataclass

@dataclass
class Section:
    lo: int
    hi: int

type Pair = tuple[Section, Section]

def parseInput(path: str) -> list[Pair]:
    pairs: list[Pair] = []

    with open(path) as f:
        for line in f:
            sections = line.strip().split(",")

            first = Section(*map(int, sections[0].split("-")))
            second = Section(*map(int, sections[1].split("-")))

            pair: Pair = (first, second)
            pairs.append(pair)

    return pairs

def parseSections(pairs: list[Pair]) -> int:
    overlaps: int = 0

    for pair in pairs:
        if pair[0].lo >= pair[1].lo and pair[0].hi <= pair[1].hi:
            overlaps += 1
            continue
        if pair[1].lo >= pair[0].lo and pair[1].hi <= pair[0].hi:
            overlaps += 1
            continue

    return overlaps

def parseSections2(pairs: list[Pair]) -> int:
    overlaps: int = 0

    for pair in pairs:
        if pair[0].lo <= pair[1].hi and pair[1].lo <= pair[0].hi:
            overlaps += 1

    return overlaps


def main() -> None:
    pairs = parseInput("input.txt")

    print(f"Part 1 = {parseSections(pairs)}.") # 569
    print(f"Part 2 = {parseSections2(pairs)}.") # 936

if __name__ == "__main__":
    main()