def validBirthYear(i: int) -> bool:
    return (1920 <= i <= 2002)

def validIssueYear(i: int) -> bool:
    return (2010 <= i <= 2020)

def validExpirationYear(i: int) -> bool:
    return (2020 <= i <= 2030)

def validHeight(s: str) -> bool:
    l = len(s)
    if l < 3:
        return False

    num = (s[:l-2])
    unit = s[l-2:]

    if not num.isdigit():
        return False

    h = int(num)

    match unit:
        case "in":
            if (59 <= h <= 76):
                return True
        case "cm":
            if (150 <= h <= 193):
                return True 

    return False

def validHairColor(s: str) -> bool:
    if len(s) != 7:
        return False

    if s[0] != '#':
        return False

    for char in s[1:]:
       isHexLetter = 'a' <= char <= 'f'
       if not char.isdigit() and not isHexLetter:
           return False

    return True

def validEyeColor(s: str) -> bool:
    return s in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

def validPassportID(s: str) -> bool:
    return len(s) == 9 and s.isdigit()