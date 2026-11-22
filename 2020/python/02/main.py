# https://adventofcode.com/2020/day/2

from dataclasses import dataclass

@dataclass(frozen=True)
class Policy:
    min: int
    max: int
    val: str

def parseInput(path: str) -> tuple[list[Policy], list[str]]:
    policies: list[Policy] = []
    passwords: list[str] = []

    with open(path) as f:
        for line in f:
            if line.strip():
                data = line.strip().split()
                min, max = data[0].split("-")

                p = Policy(
                    min=int(min),
                    max=int(max),
                    val=data[1][0]
                    )
                policies.append(p)
                passwords.append(data[-1])

    return policies, passwords

def validatePasswords(policies: list[Policy], passwords: list[str]) -> int:
    validPasswords: int = 0

    for i in range(len(passwords)):
        charCount = 0
        for char in passwords[i]:
            if char == policies[i].val:
                charCount += 1

        if charCount >= policies[i].min and charCount <= policies[i].max:
            validPasswords += 1

    return validPasswords

def validatePasswords2(policies: list[Policy], passwords: list[str]) -> int:
    validPasswords = 0

    for i in range(len(passwords)):
        posCount = 0
        if passwords[i][policies[i].min-1] == policies[i].val:
            posCount += 1
        if passwords[i][policies[i].max-1] == policies[i].val:
            posCount += 1

        if posCount == 1:
            validPasswords += 1

    return validPasswords

def main() -> None:
    policies, passwords = parseInput("input.txt")

    print(f"There are {validatePasswords(policies=policies, passwords=passwords)} passwords.")
    print(f"There are {validatePasswords2(policies=policies, passwords=passwords)} passwords in part 2.")


if __name__ == "__main__":
    main()