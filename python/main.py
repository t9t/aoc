#!/usr/bin/env python3

from year2016 import *
import sys
from datetime import datetime
from datetime import timedelta
from time import sleep


def main(name, args):
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
    result = part_func(input)
    end = datetime.now()

    print("Result ({0}): {1}".format(format_duration(end-start), result))


def format_duration(d: timedelta):
    s = d.total_seconds()
    if s < 1e-03:
        return str(round(s * 1_000_000, 3)) + "μs"
    if s < 1:
        return str(round(s * 1000, 3)) + "ms"
    return str(s) + "s"


if __name__ == "__main__":
    main(sys.argv[0], sys.argv[1:])
