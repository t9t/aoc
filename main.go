package main

import (
	"aoc/year2015"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fatalUsage("")
	}

	year := mustParseIntArg("year", args[0])
	day := mustParseIntArg("day", args[1])
	part := mustParseIntArg("part", args[2])

	if year == 2015 && day == 1 && part == 1 {
		fmt.Printf("Running Year: %d; Day: %d; Part: %d\n", year, day, part)
		data, err := os.ReadFile("input/2015-1-1.txt")
		if err != nil {
			fmt.Printf("Unable to read input file: %v\n", err)
			os.Exit(3)
		}
		i, err := year2015.Day1Part1(string(data))
		if err != nil {
			fmt.Printf("Error running: %v\n", err)
			os.Exit(4)
		}
		fmt.Printf("Result: %d\n", i)
	} else {
		fmt.Printf("Unsupported Year: %d; Day: %d; Part: %d\n", year, day, part)
		os.Exit(2)
	}
}

func mustParseIntArg(argName, v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		fatalUsage(fmt.Sprintf("Argument %s is not an integer: %v", argName, err))
	}
	return i
}

func fatalUsage(errorMessage string) {
	if errorMessage != "" {
		fmt.Println(errorMessage)
		fmt.Println()
	}
	fmt.Println("Usage:")
	fmt.Printf("\t%s <year> <day> <part>\n", os.Args[0])

	os.Exit(1)
}
