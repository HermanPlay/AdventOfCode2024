package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 1. Read the lists
// 2. Sort the lists
// 3. Calculate distance
func Part1(input string) int {
	first_heap := NewHeap(1000)
	second_heap := NewHeap(1000)
	pairs := strings.Split(input, "\n")
	for i, pair := range pairs {
		p := strings.Split(pair, "   ")
		if len(p) != 2 {
			break
		}
		first, err := strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			fmt.Printf("Pair %d, %s\n", i, p[0])
			panic("failed to parse int")
		}
		second, err := strconv.ParseInt(p[1], 10, 64)
		if err != nil {
			fmt.Printf("Pair %d, '%s'\n", i, p[1])
			panic("failed to parse int" + err.Error())
		}
		first_heap.Insert(int(first))
		second_heap.Insert(int(second))
	}
	diff := 0
	for range 1000 {
		first := first_heap.Pop()
		second := second_heap.Pop()
		difference := int(math.Abs(float64(first - second)))
		fmt.Printf("%d - %d = %d\n", first, second, difference)
		diff += difference
	}

	return diff

}

func Part2(input string) int64 {
	first_map := make(map[int64]int64)
	second_map := make(map[int64]int64)
	pairs := strings.Split(input, "\n")
	var total int64 = 0
	for i, pair := range pairs {
		p := strings.Split(pair, "   ")
		if len(p) != 2 {
			break
		}
		first, err := strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			fmt.Printf("Pair %d, %s\n", i, p[0])
			panic("failed to parse int")
		}
		second, err := strconv.ParseInt(p[1], 10, 64)
		if err != nil {
			fmt.Printf("Pair %d, %s\n", i, p[1])
			panic("failed to parse int")
		}
		total += first_map[second] * second
		second_map[second] += 1
		first_map[first] += 1
		total += first * second_map[first]
	}

	return total
}
