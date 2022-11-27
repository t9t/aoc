package main

import (
	"aoc/registry"
	"aoc/year2015"
	_ "aoc/year2018"
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
	if len(args) != 2 && len(args) != 3 {
		fatalUsage("")
	}

	if len(args) == 2 {
		if args[0] == "benchmark" || args[0] == "all" || args[0] == "results" {
			year := mustParseIntArg("year", args[1])
			err := runAll(args[0], year)
			if err != nil {
				fmt.Printf("\nError running all: %v\n", err)
				os.Exit(4)
			}
			return
		} else {
			fatalUsage("")
		}
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

func runAll(mode string, year int) error {
	modeBenchmark := mode == "benchmark"
	modeResults := mode == "results"

	sortedSelectors := registry.AllSelectorsSorted(year)

	type result struct {
		output  string
		runTime time.Duration
	}

	if !modeResults {
		fmt.Println("| Year | Day | Part | Output                     | Run time     |")
		fmt.Println("|------|-----|------|----------------------------|--------------|")
	}

	printOutput := func(selector registry.Selector, output string, runTime time.Duration) {
		if modeResults {
			fmt.Printf("%d-%d-%d: %s\n", selector.Year, selector.Day, selector.Part, output)
		} else {
			fmt.Printf("%s| %4d | %3d | %4d | %26v | %12v |\n", clearLine(), selector.Year, selector.Day, selector.Part, output, runTime)
		}
	}

	totalRunTime := time.Duration(0)
	begin := time.Now()
	results := make(map[registry.Selector]result)
	for i, selector := range sortedSelectors {
		if !modeResults {
			fmt.Printf("%s%3d/%3d; %v; running year: %d; day: %d; part: %d", clearLine(),
				i, len(sortedSelectors), time.Since(begin),
				selector.Year, selector.Day, selector.Part)
		}

		input, err := readInputData(selector.Year, selector.Day)
		if err != nil {
			return err
		}

		start := time.Now()
		output, err := registry.Map[selector](input)
		runTime := time.Since(start)
		if err != nil {
			return fmt.Errorf("error executing %d/%d/%d: %w", selector.Year, selector.Day, selector.Part, err)
		}
		totalRunTime += runTime
		if !modeBenchmark {
			printOutput(selector, output, runTime)
		} else {
			results[selector] = result{output: output, runTime: runTime}
		}
	}

	if modeBenchmark {
		sort.Slice(sortedSelectors, func(i, j int) bool {
			l, r := sortedSelectors[i], sortedSelectors[j]
			return results[l].runTime < results[r].runTime
		})

		fmt.Print(clearLine())
		for _, selector := range sortedSelectors {
			result := results[selector]
			printOutput(selector, result.output, result.runTime)
		}
	}
	if !modeResults {
		fmt.Printf("\nTotal run time: %v\n", totalRunTime)
	}
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
	fmt.Println("Usage (one of):")
	fmt.Printf("\t%s <year> <day> <part>\n", os.Args[0])
	fmt.Printf("\t%s all <year>\n", os.Args[0])
	fmt.Printf("\t%s results <year>\n", os.Args[0])
	fmt.Printf("\t%s benchmark <year>\n", os.Args[0])

	os.Exit(1)
}

func readInputData(year, day int) (string, error) {
	filename := fmt.Sprintf("../input/%d/%d.txt", year, day)
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot read data for %d/%d: %w", year, day, err)
	}
	// 2018/13 requires leading & trailing spaces
	return strings.TrimRight(string(data), "\r\n"), nil
}
