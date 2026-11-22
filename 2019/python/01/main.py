# https://adventofcode.com/2019/day/1

from pathlib import Path

def fuelSum(modules: list[str]) -> tuple[int, int]:
    simple = recursive = 0

    for m in modules:
        mass = int(m)
        simple += (mass // 3) - 2
        recursive += fuelRequirement(mass)

    return simple, recursive

def fuelRequirement(mass: int) -> int:
    fuel = (mass // 3) - 2

    # base case
    if fuel <= 0:
        return 0

    return fuel + fuelRequirement(fuel)

def main() -> None:
    modules = Path("test.txt").read_text().splitlines()

    p1, p2 = fuelSum(modules)
    print(f"Part 1 = {p1}.") # 3442987
    print(f"Part 2 = {p2}.") # 5161601


if __name__ == "__main__":
    main()