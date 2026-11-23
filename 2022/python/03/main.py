# https://adventofcode.com/2022/day/3

from pathlib import Path

def parseItems(rucksacks: list[str]) -> int:
    priority: int = 0

    for s in rucksacks:
        l = int(len(s) / 2)
        c1 = s[:l]
        c2 = s[l:]

        intersection = list(set(c1) & set(c2)) # intersect operator '&' only works on sets
        item = intersection[0]

        if ('a' <= item <= 'z'):
            priority += ord(item) % 96
        elif ('A' <= item <= 'Z'):
            priority += (ord(item) % 64) + 26
        else:
            continue

    return priority

def parseBadges(rucksacks: list[str]) -> int:
    priority: int = 0

    i = 0
    while i < len(rucksacks):
        s1 = set(rucksacks[i])
        s2 = set(rucksacks[i+1])
        s3 = set(rucksacks[i+2])

        intersection = list(s1 & s2 & s3)
        item = intersection[0]

        if ('a' <= item <= 'z'):
            priority += ord(item) % 96
        elif ('A' <= item <= 'Z'):
            priority += (ord(item) % 64) + 26
        else:
            continue

        i += 3

    return priority

def main() -> None:
    rucksacks = Path("input.txt").read_text().strip().splitlines()

    print(f"Part 1 = {parseItems(rucksacks)}.") # 8139
    print(f"Part 2 = {parseBadges(rucksacks)}.") # 2668

if __name__ == "__main__":
    main()