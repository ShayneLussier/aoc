def parseElves(path: str) -> list[int]:
    elves: list[int] = []
    current: int = 0

    with open(path) as f:
        for line in f:
            if not line.strip(): # if the line is only whitespace
                elves.append(current)
                current = 0
                continue

            current += int(line)

    elves.append(current)
    elves.sort(reverse=True)
    return elves

def sumTop(elves: list[int], n: int) -> int:
    total: int = 0
    for elf in elves[:n]:
        total += elf

    return total
    
def main() -> None:
    elves = parseElves("input.txt")

    print(f"The top elf has {sumTop(elves=elves, n=1)} calories!") # 72478
    print(f"The top 3 elves have {sumTop(elves=elves, n=3)} calories!") # 210367

if __name__ == "__main__":
    main()