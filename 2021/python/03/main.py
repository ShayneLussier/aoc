# https://adventofcode.com/2021/day/3

from pathlib import Path

def countBits(lines: list[str]) -> tuple[list[int], list[int]]:
    size = len(lines[0])
    ones = [0] * size
    zeros = [0] * size

    for line in lines:
        for i, char in enumerate(line):
            if char == '1':
                ones[i] += 1

    for i in range(len(zeros)):
        zeros[i] = len(lines) - ones[i]

    return ones, zeros

def powerConsumption(ones, zeros: list[int]) -> int:
    gamma = epsilon = ""

    for o, z in zip(ones, zeros):
        if o > z:
            gamma += '1'
            epsilon += '0'
        else:
            gamma += '0'
            epsilon += '1'

    return int(gamma, 2) * int(epsilon, 2)

def verifyRating(lines: list[str], criteria: int) -> int:
    for i in range(len(lines)):
        if len(lines) == 1:
            break

        ones, zeros = countBits(lines)

        want: int

        match criteria:
            # oxygen rating, most common
            case 1:
                if ones[i] >= zeros[i]:
                    want = 1
                else:
                    want = 0
            #  CO2 rating, least common
            case 0:
                if ones[i] >= zeros[i]:
                    want = 0
                else:
                    want = 1

        # keep only the valid lines
        kept: list[str] = []
        for line in lines:
            if int(line[i]) == want:
                kept.append(line)

        lines = kept

    return int(lines[0], 2)

def main() -> None:
    lines = Path("input.txt").read_text().strip().splitlines()

    print(f"The power consumption of the submarine is {powerConsumption(*countBits(lines))}!") # 3923414
    
    oxygenRating = verifyRating(lines=lines, criteria=1)
    scrubberRating = verifyRating(lines=lines, criteria=0)
    print(f"The life support rating of the submarine is {oxygenRating * scrubberRating}!") # 5852595

if __name__ == "__main__":
    main()