package day2

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func verify(a, b string, increasing, decreasing bool) bool {
	num, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}
	last, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}

	diff := num - last
	diffAbs := int64(math.Abs(float64(diff)))

	// Verify difference is in range [1, 3]
	if diffAbs < 1 || diffAbs > 3 {
		return false
	}
	if (increasing && diff <= 0) || (decreasing && diff >= 0) {
		return false
	}

	return true
}

func calculateIncDec(a, b string) (bool, bool) {
	increasing, decreasing := false, false

	num, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}
	last, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	diff := num - last

	if diff > 0 {
		increasing = true
	} else if diff < 0 {
		decreasing = true
	} else {
		panic("should not happen")
	}
	if !increasing && !decreasing {
		panic("should not happen")
	}
	return increasing, decreasing
}

func Part1(input string) int {
	reports := strings.Split(input, "\n")
	safe := 0
	for _, report := range reports {
		if len(report) == 0 {
			println("empty report")
		}
		numbersStr := strings.Split(report, " ")
		var last int64 = -1
		correct := true
		increasing := false
		decreasing := false
		for i, numberStr := range numbersStr {
			num, err := strconv.ParseInt(numberStr, 10, 64)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				last = num
				continue
			}

			diff := num - last
			diffAbs := int64(math.Abs(float64(diff)))

			// Verify difference is in range [1, 3]
			if diffAbs < 1 || diffAbs > 3 {
				correct = false
				break
			}

			if i == 1 {
				if diff > 0 {
					increasing = true
				} else if diff < 0 {
					decreasing = true
				} else {
					correct = false
					break
				}
			} else {
				// Ensure consistent trend
				if (increasing && diff <= 0) || (decreasing && diff >= 0) {
					correct = false
					break
				}
			}

			last = num
		}
		if correct {
			fmt.Printf("Report %s, is safe\n", report)
			safe += 1
		} else {
			fmt.Printf("Report %s, is not safe\n", report)
		}

	}
	return safe
}

func isSafe(report []string, excludeIdx int) bool {
	var last int64 = -1
	increasing, decreasing := false, false
	for i := 0; i < len(report); i++ {
		if i == excludeIdx {
			continue
		}

		num, err := strconv.ParseInt(report[i], 10, 64)
		if err != nil {
			panic(err)
		}
		if last == -1 {
			last = num
			continue
		}

		diff := num - last
		diffAbs := int64(math.Abs(float64(diff)))

		// Ensuire consistency and difference in range
		if diffAbs < 1 || diffAbs > 3 || (increasing && diff <= 0) || (decreasing && diff >= 0) {
			// fmt.Printf("%v is not safe diff=%d increasing=%v decreasing=%v", report, diff, increasing, decreasing)
			return false
		}

		if !increasing && !decreasing {
			if diff > 0 {
				increasing = true
			} else if diff < 0 {
				decreasing = true
			} else {
				panic("should not happen")
			}
		}
		last = num
	}
	return true

}

func Part2BruteForce(input string) int {
	reports := strings.Split(input, "\n")
	safe := 0
	for _, report := range reports {
		if len(report) == 0 {
			println("empty report")
		}
		numbersStr := strings.Split(report, " ")
		is_safe := false
		if isSafe(numbersStr, -1) {
			safe++
			fmt.Printf("%s is safe\n", report)
			is_safe = true
		} else {
			// fmt.Printf("%s is not safe naturally", report)
			for i := range numbersStr {
				if isSafe(numbersStr, i) {
					// Without idx i the report is safe
					safe++
					is_safe = true
					fmt.Printf("%s is safe after remove\n", report)
					break
				}
			}
		}
		if !is_safe {
			fmt.Printf("%s is not safe\n", report)
		}
	}
	return safe
}

func Part2Linear(input string) int {
	reports := strings.Split(input, "\n")
	safe := 0
	for _, report := range reports {
		if len(report) == 0 {
			println("empty report")
		}
		numbersStr := strings.Split(report, " ")
		var last int64 = -1
		correct := true
		increasing := false
		decreasing := false
		count := 0
		for i := 0; i < len(numbersStr); i++ {
			numberStr := numbersStr[i]

			num, err := strconv.ParseInt(numberStr, 10, 64)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				last = num
				continue
			}

			diff := num - last
			diffAbs := int64(math.Abs(float64(diff)))

			// Ensuire consistency and difference in range
			if diffAbs < 1 || diffAbs > 3 || (increasing && diff <= 0) || (decreasing && diff >= 0) {
				if count > 0 {
					correct = false
					count++
					break
				}
				if i == 1 {
					// If it first element, we can just drop it, and try again
					if verify(numbersStr[i], numbersStr[i+1], false, false) {
						new_inc, new_dec := calculateIncDec(numbersStr[i], numbersStr[i+1])
						// Verify that indeed we need to drop the i==0, because we may need to drop the i==1
						if verify(numbersStr[i+1], numbersStr[i+2], new_inc, new_dec) {
							// We can drop i==0
							log.Print("Removed i==0")
							count++
							next, err := strconv.ParseInt(numbersStr[i+2], 10, 64)
							if err != nil {
								panic("parse next")
							}
							// But we need to calculate increasing decreasing
							increasing, decreasing = new_inc, new_dec
							// Skip next element, cause we just checked it
							i += 2
							last = next
							continue
						}
					}
					if verify(numbersStr[i-1], numbersStr[i+1], false, false) {
						// We can drop i==1
						count++
						next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
						if err != nil {
							panic("parse next")
						}
						// But we need to calculate increasing decreasing
						increasing, decreasing = calculateIncDec(numbersStr[i-1], numbersStr[i+1])
						// Skip next element, cause we just checked it
						i++
						last = next
						continue
					} else {
						// Removing current of prev will not help
						correct = false
						break
					}
				} else if i == len(numbersStr)-1 {
					// Last element can just be removed
					count++
					continue
				} else if i == 2 {
					if verify(numbersStr[i-1], numbersStr[i], false, false) {
						new_inc, new_dec := calculateIncDec(numbersStr[i-1], numbersStr[i])
						if verify(numbersStr[i], numbersStr[i+1], new_inc, new_dec) {
							// We can drop first element
							count++
							next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
							if err != nil {
								panic("parse next")
							}
							i++
							last = next
							increasing = new_inc
							decreasing = new_dec
							continue
						}
					}
					if verify(numbersStr[i-1], numbersStr[i+1], increasing, decreasing) {
						log.Print("Removed i==2")
						count++
						next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
						if err != nil {
							panic("parse next")
						}
						i++
						last = next
						continue
					} else if verify(numbersStr[i-2], numbersStr[i], false, false) {
						// We can remove i==1. Monotonicity may change
						count++
						last = num
						increasing, decreasing = calculateIncDec(numbersStr[i-2], numbersStr[i])
						continue
					} else {
						correct = false
						break
					}
				} else {
					// i in [3, len-2]
					if verify(numbersStr[i-1], numbersStr[i+1], increasing, decreasing) {
						// We can drop current element
						count++
						next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
						if err != nil {
							panic("parse next")
						}
						i++
						last = next
						continue
					} else {
						// Dropping current will not solve the problem, hence nothing will help
						correct = false
						break
					}
				}
			} else {
				if i == 1 {
					if diff > 0 {
						increasing = true
					} else if diff < 0 {
						decreasing = true
					}
				}
				last = num
			}

		}
		if correct || !correct && count == 1 {
			if count == 1 {
				fmt.Printf("%s is safe after remove\n", report)
			} else {

				fmt.Printf("%s is safe\n", report)
			}
			safe += 1
		} else {
			fmt.Printf("%s is not safe\n", report)
		}
	}

	return safe
}
