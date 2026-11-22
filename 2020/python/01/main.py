# https://adventofcode.com/2020/day/1

def parseInput(path: str) -> list[int]:
    expenses: list[int] = []

    with open(path) as f:
        for line in f:
            if line.strip():
                expenses.append(int(line.strip()))

    return expenses

def checkSum(expenses: list[int], target: int) -> int:
    for i in range(len(expenses)):
        searching = target - expenses[i]
        for e in expenses[i+1:]:
            if e == searching:
                return expenses[i] * e

def threeCheckSum(expenses: list[int], target) -> int:
    for i in range(len(expenses)):
        searching = target - expenses[i]
        product = checkSum(expenses=expenses[i+1:], target=searching)
        if product is not None:
            return expenses[i] * product

def main() -> None:
    expenses = parseInput("input.txt")

    print(checkSum(expenses=expenses, target=2020)) # 326211
    print(threeCheckSum(expenses=expenses, target=2020)) # 131347190

if __name__ == "__main__":
    main()