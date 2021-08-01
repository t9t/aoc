package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	println("Hello, world")
	args := os.Args[1:]

	if len(args) != 3 {
		fatalUsage("")
	}

	fmt.Printf("Args: %v\n", args)
	year := mustParseIntArg("year", args[0])
	day := mustParseIntArg("day", args[1])
	part := mustParseIntArg("part", args[2])

	fmt.Printf("Year: %d; Day: %d; Part: %d\n", year, day, part)
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
