# https://adventofcode.com/2021/day/1

def parseInput(path: str) -> list[int]:
    depths: list[int] = []

    with open(path) as f:
        for line in f:
            if line.strip():
                depths.append(int(line))

    return depths

def countIncreases(depths: list[int]) -> int:
    increases: int = 0

    for i in range(len(depths)-1):
        if depths[i+1] > depths[i]:
            increases += 1
    
    return increases

def countWindowIncreases(depths: list[int]) -> int:
    increases: int = 0

    for i in range(len(depths)-3):
        if sum(depths[i:i+3]) < sum(depths[i+1:i+4]):
            increases += 1

    return increases

def main():
    depths = parseInput("input.txt")

    print(f"There are {countIncreases(depths=depths)} increases.") # 1564
    print(f"There are {countWindowIncreases(depths=depths)} sliding window increases.") # 1611

if __name__ == "__main__":
    main()