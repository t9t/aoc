package main

import (
	"aoc/registry"
	"aoc/year2015"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	year2015.RegisterAll()

	args := os.Args[1:]

	if len(args) == 1 && args[0] == "benchmark" {
		err := runBenchmark()
		if err != nil {
			fmt.Printf("\nError running benchmark: %v\n", err)
			os.Exit(4)
		}
		return
	}

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

	input, err := readInputData(year, day)
	if err != nil {
		fmt.Printf("Error reading input data: %v\n", err)
		os.Exit(3)
	}

	fmt.Printf("Running Year: %d; Day: %d; Part: %d\n", year, day, part)
	start := time.Now()
	result, err := execution(input)
	if err != nil {
		fmt.Printf("Error running: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("Result (%v): %v\n", time.Since(start), result)
}

func runBenchmark() error {
	sortedSelectors := registry.AllSelectorsSorted()

	begin := time.Now()
	runTimes := make(map[registry.Selector]time.Duration)
	for i, selector := range sortedSelectors {
		fmt.Printf("%s%3d/%3d; %v; running year: %d; day: %d; part: %d", clearLine(),
			i, len(sortedSelectors), time.Since(begin),
			selector.Year, selector.Day, selector.Part)

		input, err := readInputData(selector.Year, selector.Day)
		if err != nil {
			return err
		}

		start := time.Now()
		if _, err := registry.Map[selector](input); err != nil {
			return fmt.Errorf("error executing %d/%d/%d: %w", selector.Year, selector.Day, selector.Part, err)
		}
		runTime := time.Since(start)
		runTimes[selector] = runTime
	}

	sort.Slice(sortedSelectors, func(i, j int) bool {
		l, r := sortedSelectors[i], sortedSelectors[j]
		lt, rt := runTimes[l], runTimes[r]
		return lt < rt
	})

	fmt.Print(clearLine())
	fmt.Println("| Year | Day | Part | Run time   |")
	fmt.Println("|------|-----|------|------------|")
	var totalRunTime time.Duration
	for _, selector := range sortedSelectors {
		runTime := runTimes[selector]
		totalRunTime += runTime
		fmt.Printf("| %4d | %3d | %4d | %10v | \n", selector.Year, selector.Day, selector.Part, runTime)
	}
	fmt.Printf("\nTotal run time: %v\n", totalRunTime)
	return nil
}

func clearLine() string {
	// https://unix.stackexchange.com/a/26592
	return "\033[1K\r"
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

func readInputData(year, day int) (string, error) {
	filename := fmt.Sprintf("input/%d/%d.txt", year, day)
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot read data for %d/%d: %w", year, day, err)
	}
	return strings.TrimSpace(string(data)), nil
}
