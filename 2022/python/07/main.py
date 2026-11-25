# https://adventofcode.com/2022/day/7

def dirSizes(commands: list[str]) -> dict[str, int]:
    sizes: dict[str, int] = {}
    currentDir: list[str] = []

    for c in commands:
        match c.split():
            case ["$", "cd", ".."]:
                currentDir.pop()
            case ["$", "cd", name]:
                currentDir.append(name)
            case [size, _]:
                if size.isdigit():
                    for i in range(len(currentDir)):
                        path = "/".join(currentDir[:i+1])
                        sizes[path] = sizes.get(path, 0) + int(size)
    return sizes

def part1(sizes: dict[str, int]) -> int:
    total: int = 0

    for v in sizes.values():
        if v < 100_000:
            total += v

    return total

def part2(sizes: dict[str, int]) -> int:
    needed: int = 30_000_000 - (70_000_000 - sizes["/"])
    eligible: list[int] = []

    for v in sizes.values():
        if v >= needed:
            eligible.append(v)

    return min(eligible)

def main() -> None:
    with open("input.txt") as f:
        commands = f.read().splitlines()

    print(f"Part 1 = {part1(dirSizes(commands))}!") # 1325919
    print(f"Part 2 = {part2(dirSizes(commands))}!") # 2050735


if __name__ == "__main__":
    main()