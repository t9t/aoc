from year2016 import day1
import sys
from datetime import datetime


if __name__ == "__main__":
    args = sys.argv[1:]
    if len(args) != 3:
        print("Usage:")
        print("\t%s <year> <day> <part>\n", sys.argv[0])
        sys.exit(1)

    year = int(args[0])
    day = int(args[1])
    part = int(args[2])

    if year != 2016 or day != 1 or part != 1:
        print("Unsupported arguments, Year: {0}; Day: {1}; Part: {2}",
              year, day, part)
        sys.exit(2)

    with open("../input/{0}/{1}.txt".format(year, day)) as f:
        input = f.read()

    print("Running Year: {0}; Day: {1}; Part: {1}".format(year, day, part))

    start = datetime.now()
    result = day1.find_distance(input)
    end = datetime.now()

    print("Result ({0}): {1}".format(end-start, result))
