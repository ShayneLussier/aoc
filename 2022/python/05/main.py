# https://adventofcode.com/2022/day/5

from dataclasses import dataclass
from copy import deepcopy

@dataclass
class Instruction:
    qty: int
    src: int
    dst: int

type Stack = list[str]

def parseStacks(data: str) -> list[Stack]:
    lines: list[str] = data.splitlines()
    stacks: list[Stack] = []

    # add n amount of empty stacks to the list
    stackCount = len(lines[-1].split())
    for _ in range(stackCount):
        stacks.append([])

    for line in lines:
        i = 0
        j = 0
        while j < len(line):
            symbol = line[j+1]
            if symbol.isalpha():
                stacks[i].append(symbol)
            i += 1
            j += 4

    for s in stacks:
        s.reverse()

    return stacks

def parseInstructions(data: str) -> list[Instruction]:
    lines: list[str] = data.splitlines()
    instructions: list[Instruction] = []

    for l in lines:
        qty = int(l.split(" ")[1])
        src = int(l.split(" ")[3])
        dst = int(l.split(" ")[5])

        instructions.append(Instruction(qty, src, dst))

    return instructions

def parseInput(path: str) -> tuple[list[Stack], list[Instruction]]:
    with open(path) as f:
        first, second = f.read().split("\n\n")

    stacks = parseStacks(first)
    instructions = parseInstructions(second)

    return stacks, instructions

def crateMover9000(stacks: list[Stack], instructions: list[Instruction]) -> str:
    for i in instructions:
        for _ in range(i.qty):
            crate = stacks[i.src-1].pop()
            stacks[i.dst-1].append(crate)

    topCrates = ""
    for s in stacks:
        if s: # if stack has at least 1 crate
            topCrates += s[-1]

    return topCrates

def crateMover9001(stacks: list[Stack], instructions: list[Instruction]) -> str:
    for i in instructions:
        crane: list[str] = []
        for _ in range(i.qty):
            crate = stacks[i.src-1].pop()
            crane.append(crate)

        crane.reverse()

        for crate in crane:
            stacks[i.dst-1].append(crate)

    topCrate = ""
    for s in stacks:
        if s:
            topCrate += s[-1]

    return topCrate

def main() -> None:
    stacks, instructions = parseInput("input.txt")
    stacksClone = deepcopy(stacks)

    print(f"The crate of the top of every stack makes code = {crateMover9000(stacks, instructions)}!") # GFTNRBZPF
    print(f"The crate of the top of every stack makes code = {crateMover9001(stacksClone, instructions)}!") # VRQWPDSGP

if __name__ == "__main__":
    main()