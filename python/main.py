#!/usr/bin/env python3

from year2016 import *
import sys
from datetime import datetime

if __name__ == "__main__":
    args = sys.argv[1:]
    if len(args) != 3:
        print("Usage:")
        print("\t{0} <year> <day> <part>\n".format(sys.argv[0]))
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

    print("Running Year: {0}; Day: {1}; Part: {1}".format(year, day, part))

    start = datetime.now()
    result = part_func(input)
    end = datetime.now()

    print("Result ({0}): {1}".format(end-start, result))
