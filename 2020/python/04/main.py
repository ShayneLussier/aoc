# https://adventofcode.com/2020/day/4

from dataclasses import dataclass
import validators

@dataclass
class Passport:
    byr: int = 0
    iyr: int = 0
    eyr: int = 0
    hgt: str = ""
    hcl: str = ""
    ecl: str = ""
    pid: str = ""
    cid: int = 0

    def hasFields(self) -> bool:
        return (
            self.byr != 0
            and self.iyr != 0
            and self.eyr != 0
            and self.hgt != ""
            and self.hcl != ""
            and self.ecl != ""
            and self.pid != ""
        )

    def isValid(self) -> bool:
        return (
            validators.validBirthYear(self.byr)
            and validators.validIssueYear(self.iyr)
            and validators.validExpirationYear(self.eyr)
            and validators.validHeight(self.hgt)
            and validators.validHairColor(self.hcl)
            and validators.validEyeColor(self.ecl)
            and validators.validPassportID(self.pid)
        )

def parseInput(path: str) -> list[Passport]:
    passports: list[Passport] = []

    passport = Passport()
    with open(path) as f:
        for line in f:
            if not line.strip():
                passports.append(passport)
                passport = Passport()
                continue

            data = line.strip().split()
            for info in data:
                parts = info.split(":")
                match parts[0]:
                    case "byr":
                        passport.byr = int(parts[1])
                    case "iyr":
                        passport.iyr = int(parts[1])
                    case "eyr":
                        passport.eyr = int(parts[1])
                    case "hgt":
                        passport.hgt = parts[1]
                    case "hcl":
                        passport.hcl = parts[1]
                    case "ecl":
                        passport.ecl = parts[1]
                    case "pid":
                        passport.pid = parts[1]
                    case "cid":
                        passport.cid = int(parts[1])

        # handle the last passport
        passports.append(passport)

    return passports       

def validPassports(passports: list[Passport]) -> tuple[int, int]:
    part1 = part2 = 0

    for p in passports:
        if p.hasFields():
            part1 += 1

        if p.isValid():
            part2 += 1

    return part1, part2

def main() -> None:
    passports = parseInput("input.txt")

    p1, p2 = validPassports(passports=passports)
    print(f"{p1} passports have all the required fields.") # 204
    print(f"{p2} passports are valid.") # 179


if __name__ == "__main__":
    main()