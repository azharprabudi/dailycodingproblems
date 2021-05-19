package productsummaries

// O(N^2)
// func ProductSummaries(numbers []int) []int {
// 	var summaries []int
// 	for index := range numbers {
// 		total := 1
// 		for indexChild := range numbers {
// 			if index == indexChild {
// 				continue
// 			}
//
// 			total = total * numbers[indexChild]
// 		}
// 		summaries = append(summaries, total)
// 	}
//
// 	return results
// }

// O(N) Division
// func ProductSummaries(numbers []int) []int {
// 	var total int = 1
// 	var summaries []int
//
// 	for _, number := range numbers {
// 		total = total * number
// 	}
//
// 	for _, number := range numbers {
// 		div := sum / number
// 		summaries = append(summaries, div)
// 	}
//
// 	return summaries
// }
//

// O(N)
func ProductSummaries(numbers []int) []int {
	var summaries []int
	var summariesLeft []int
	var summariesRight []int

	for i, number := range numbers {
		if i == 0 {
			summariesLeft = append(summariesLeft, number)
			continue
		}

		multiple := summariesLeft[len(summariesLeft)-1] * number
		summariesLeft = append(summariesLeft, multiple)
	}

	for i := range numbers {
		if len(numbers)-i == len(numbers) {
			summariesRight = append(summariesRight, numbers[len(numbers)-1])
			continue
		}

		multiple := numbers[len(numbers)-1-i] * summariesRight[len(summariesRight)-1]
		summariesRight = append(summariesRight, multiple)
	}

	var reversedSummariesRight []int
	for i := range summariesRight {
		reversedSummariesRight = append(reversedSummariesRight, summariesRight[len(summariesRight)-1-i])
	}

	for i := range numbers {
		if i == 0 {
			summaries = append(summaries, reversedSummariesRight[i+1])
			continue
		}

		if i == len(numbers)-1 {
			summaries = append(summaries, summariesLeft[i-1])
			continue
		}

		summaries = append(summaries, summariesLeft[i-1]*reversedSummariesRight[i+1])
	}

	return summaries
}
