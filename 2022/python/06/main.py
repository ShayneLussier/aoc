# https://adventofcode.com/2022/day/6

def parseStream(stream: str, n: int) -> int:
    i = 0
    while i < len(stream):
        seen = set()
        for char in stream[i:i+n]:
            seen.add(char)

        if len(seen) == n:
            return i+n

        i += 1

    return 0

def main() -> None:
    with open("input.txt") as f:
        stream = f.read()

    print(f"{parseStream(stream, 4)} characters processed by encountering the start-of-packet marker.") # 1343
    print(f"{parseStream(stream, 14)} characters processed by encountering the start-of-message marker.") # 2193

if __name__ == "__main__":
    main()