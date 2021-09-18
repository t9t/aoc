#!/usr/bin/env python3

from year2016 import *
import sys
from datetime import datetime
from datetime import timedelta


def main(name, args):
    if len(args) == 1 and (args[0] == "benchmark" or args[0] == "all"):
        run_all(args[0] == "benchmark")
        return

    if len(args) != 3:
        print("Usage:")
        print("\t{0} <year> <day> <part>\n".format(name))
        sys.exit(1)

    year = int(args[0])
    day = int(args[1])
    part = int(args[2])

    try:
        day_module = getattr(sys.modules["year" + str(year)], "day" + str(day))
        part_func = getattr(day_module, "part" + str(part))
    except:
        print("Unsupported arguments, Year: {0}; Day: {1}; Part: {2}".format(
            year, day, part))
        sys.exit(2)

    with open("../input/{0}/{1}.txt".format(year, day)) as f:
        input = f.read()

    print("Running Year: {0}; Day: {1}; Part: {2}".format(year, day, part))

    start = datetime.now()
    result = part_func(input.strip())
    end = datetime.now()

    print("Result ({0}): {1}".format(format_duration(end-start), result))


def run_all(benchmark_mode):
    year = 2016
    total = 49
    begin = datetime.now()
    total_run_time = timedelta()
    results = list()

    print("| Year | Day | Part | Output             | Run time   |")
    print("|------|-----|------|--------------------|------------|")

    for day in range(1, 26):
        with open("../input/{0}/{1}.txt".format(year, day)) as f:
            input = f.read()
            day_module = getattr(sys.modules["year" + str(year)], "day" + str(day))

            for part in (1, 2):
                if day == 25 and part == 2:
                    # Day 25 only has 1 part
                    continue
                print(f"{clearLine()}{(day-1)*2+part:2}/{total:2}; "
                      f"{format_duration(datetime.now()-begin)}; day: {day}; part: {part}", end="")

                func = getattr(day_module, "part" + str(part))
                start = datetime.now()
                result = func(input.strip())
                run_time = datetime.now()-start
                total_run_time += run_time
                if benchmark_mode:
                    results.append((year, day, part, result, run_time))
                else:
                    print(clearLine(), end="")
                    print(f"| {year:4} | {day:3} | {part:4} | {result:>18} | {format_duration(run_time):>10} |")

    if benchmark_mode:
        print(clearLine(), end="")
        results.sort(key=lambda i: i[4])
        for r in results:
            year, day, part, result, run_time = r
            print(f"| {year:4} | {day:3} | {part:4} | {result:>18} | {format_duration(run_time):>10} |")
    print(f"\nTotal run time: {format_duration(total_run_time)}")


def format_duration(d: timedelta) -> str:
    s = d.total_seconds()
    if s < 1e-03:
        return str(round(s * 1_000_000, 3)) + "μs"
    if s < 1:
        return str(round(s * 1000, 3)) + "ms"
    return str(s) + "s"


def clearLine() -> str:
    # https://unix.stackexchange.com/a/26592
    return "\033[1K\r"


if __name__ == "__main__":
    main(sys.argv[0], sys.argv[1:])
