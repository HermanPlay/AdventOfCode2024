package day2

import (
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
		log.Printf("Diff is out of range (%d %d)", last, num)
		return false
	}
	if (increasing && diff <= 0) || (decreasing && diff >= 0) {
		log.Printf("Not increaing nor decreasing (%d %d)", last, num)
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
			log.Printf("Report %s, is safe\n", report)
			safe += 1
		} else {
			log.Printf("Report %s, is not safe\n", report)
		}

		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter text: ")
		// reader.ReadString('\n')
	}
	return safe
}

func Part2(input string) int {
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
				if count > 1 {
					log.Print("Second error, stoping check for this report")
					correct = false
					break
				}
				if i == 1 {
					// If it first element, we can just drop it, and try again
					if verify(numbersStr[i], numbersStr[i+1], false, false) || verify(numbersStr[i-1], numbersStr[i+1], false, false) {
						// We can drop i==0
						log.Print("Dropping i==0 or i==1 when i==1")
						count++
						next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
						if err != nil {
							panic("parse next")
						}
						// But we need to calculate increasing decreasing
						increasing, decreasing = calculateIncDec(numbersStr[i], numbersStr[i+1])
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
					log.Print("Dropping last")
					count++
					continue
				} else if i == 2 {
					// 0 2
					if verify(numbersStr[i-1], numbersStr[i+1], increasing, decreasing) {
						// We can drop current element
						log.Printf("Dropping %d[%d] when i==2", num, i)
						count++
						next, err := strconv.ParseInt(numbersStr[i+1], 10, 64)
						if err != nil {
							panic("parse next")
						}
						i++
						last = next
						continue
					} else if verify(numbersStr[i-2], numbersStr[i], false, false) {
						// We can remove i==1. This can only happen when we 1 and 3 have oposite inequality to 1 and 2. Because otherwise, we would just be out of range
						log.Printf("dropping %d[%d] when i==2", last, i-1)
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
						log.Printf("Dropping %d[%d] when i is in the middle", num, i)
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
			log.Printf("Report %s, is safe\n", report)
			safe += 1
		} else {
			log.Printf("Report %s, is not safe\n", report)
		}
	}

	return safe
}
