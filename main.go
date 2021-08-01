package main

import (
	"aoc/registry"
	"aoc/year2015"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	registry.Register(2015, 1, 1, year2015.Day1Part1)
	registry.Register(2015, 1, 2, year2015.Day1Part2)
	registry.Register(2015, 2, 1, year2015.Day2Part1)
	registry.Register(2015, 2, 2, year2015.Day2Part2)
	registry.Register(2015, 3, 1, year2015.Day3Part1)

	args := os.Args[1:]

	if len(args) != 3 {
		fatalUsage("")
	}

	year := mustParseIntArg("year", args[0])
	day := mustParseIntArg("day", args[1])
	part := mustParseIntArg("part", args[2])

	execution, found := registry.Get(year, day, part)
	if !found {
		fmt.Printf("Unsupported arguments, Year: %d; Day: %d; Part: %d\n", year, day, part)
		os.Exit(2)
	}
	inputFile := fmt.Sprintf("input/%d-%d.txt", year, day)
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read input from %q: %v\n", inputFile, err)
		os.Exit(3)
	}

	fmt.Printf("Running Year: %d; Day: %d; Part: %d\n", year, day, part)
	result, err := execution(strings.TrimSpace(string(inputData)))
	if err != nil {
		fmt.Printf("Error running: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("Result: %v\n", result)
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
