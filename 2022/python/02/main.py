# https://adventofcode.com/2022/day/2

def parseGuide(path: str) -> list[str]:
    with open(path) as f:
        return [line.strip() for line in f if line.strip()]

def scorePart1(j, k: int) -> int:
    if j == 'A' and k == 'X':
        return 4
    elif j == 'A' and k == 'Y':
        return 8
    elif j == 'A' and k == 'Z':
        return 3
    elif j == 'B' and k == 'X':
        return 1
    elif j == 'B' and k == 'Y':
        return 5
    elif j == 'B' and k == 'Z':
        return 9
    elif j == 'C' and k == 'X':
        return 7
    elif j == 'C' and k == 'Y':
        return 2
    elif j == 'C' and k == 'Z':
        return 6

def scorePart2(j, k: int) -> int:
    if j == 'A' and k == 'X':
        return 3
    elif j == 'A' and k == 'Y':
        return 4
    elif j == 'A' and k == 'Z':
        return 8
    elif j == 'B' and k == 'X':
        return 1
    elif j == 'B' and k == 'Y':
        return 5
    elif j == 'B' and k == 'Z':
        return 9
    elif j == 'C' and k == 'X':
        return 2
    elif j == 'C' and k == 'Y':
        return 6
    elif j == 'C' and k == 'Z':
        return 7

def playBingo(guide: list[str]) -> tuple[int, int]:
    pts1 = pts2 = 0

    for line in guide: 
        j, k = line[0], line[-1]

        pts1 += scorePart1(j, k)
        pts2 += scorePart2(j, k)


    return pts1, pts2


def main()-> None:
    guide = parseGuide("input.txt")

    pts1, pts2 = playBingo(guide)

    print(f"Playing according to plan earns {pts1} points!") # 14264
    print(f"Playing as indicated earns {pts2} points!") # 12382

if __name__ == "__main__":
    main()