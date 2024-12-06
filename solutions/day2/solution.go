package day2

import (
	"log"
	"math"
	"strconv"
	"strings"
)

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
