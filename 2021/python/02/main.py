# https://adventofcode.com/2021/day/2

from dataclasses import dataclass

@dataclass
class Move:
    dir: str
    n: int

def parseMoves(path: str) -> list[Move]:
    moves: list[Move] = []

    with open(path) as f:
        for line in f:
            if line.strip():
                parts = line.split()
                moves.append(Move(dir=parts[0], n=int(parts[1])))

    return moves

def calcDepth(moves: list[Move]) -> int:
    h = d = 0

    for move in moves:
        match move.dir:
            case "forward":
                h += move.n
            case "down":
                d += move.n
            case "up":
                d -= move.n

    return h * d

def calcAim(moves: list[Move]) -> int:
    h = d = a = 0

    for move in moves:
        match move.dir:
            case "forward":
                h += move.n
                d += a * move.n
            case "down":
                a += move.n
            case "up":
                a -= move.n

    return h * d

def main() -> None:
    moves = parseMoves("input.txt")

    print(f"Multiplying the final position and depth gives {calcDepth(moves)}.") # 1868935
    print(f"Considering the aim, the final value is {calcAim(moves)}.") # 1965970888

if __name__ == "__main__":
    main()