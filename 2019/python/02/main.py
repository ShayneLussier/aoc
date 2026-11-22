# https://adventofcode.com/2019/day/2

def parseIntcode(path: str) -> list[int]:
    program: list[int] = []

    with open(path) as f:
        contents = f.read().split(",")
        for c in contents:
            program.append(int(c.strip()))

    return program

def parseCode(codes: list[int], noun, verb: int) -> int:
    codes = codes.copy() # cloning the list allows use to transform it without touching the original
    codes[1] = noun
    codes[2] = verb

    i = 0
    while i < len(codes):
        match codes[i]:
            case 1:
                codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
            case 2:
                codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
            case 99:
                return codes[0]
        i += 4
            
    return codes[0]

def parseInputs(codes: list[int]) -> int:
    for noun in range(100):
        for verb in range(100):
            if parseCode(codes, noun, verb) == 19690720:
                return 100 * noun + verb

    return 0

def main() -> None:
    intcode = parseIntcode("input.txt")

    print(f"Part 1 = {parseCode(intcode, 12, 2)}.") # 3101878
    print(f"Part 2 = {parseInputs(intcode)}.") # 8444

if __name__ == "__main__":
    main()