package main

import (
	"advent-of-code-2024/solutions/day2"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const MODULE_NAME = "advent-of-code-2024"

func printUsage() {
	println("Usage: go run main.go -day=1")
	os.Exit(1)
}

func readInput(dayNumber int, example bool) string {
	var path string
	if example {
		path = fmt.Sprintf("solutions/day%d/example.txt", dayNumber)
	} else {
		path = fmt.Sprintf("solutions/day%d/input.txt", dayNumber)
	}
	fmt.Printf("Input path: %s\n", path)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	n := 1
	var input strings.Builder
	for n != 0 {
		bufferSize := 1000
		buffer := make([]byte, bufferSize)
		n, err = f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				panic(err)
			} else {
				break
			}
		}
		if n != bufferSize {
			buffer = buffer[0 : n-1]
		}
		input.Write(buffer)
	}

	return input.String()

}

func main() {
	// Code
	day := 0
	example := false
	flag.IntVar(&day, "day", 0, "the day of the advent of code")
	flag.BoolVar(&example, "example", false, "use example input")
	flag.Parse()
	if flag.NArg() != 0 {
		printUsage()
	}

	input := readInput(day, example)

	fmt.Printf("Part 1 result: %d\n", day2.Part1(input))
	fmt.Printf("Part 2 result: %d\n", day2.Part2Linear(input))

}
