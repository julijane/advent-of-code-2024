package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func isCorrect(targetSum int, numbers []int, mutation int) (bool, bool) {
	sum := numbers[0]

	notPart1 := false
	for _, number := range numbers[1:] {
		switch mutation % 3 {
		case 0:
			sum += number
		case 1:
			sum *= number
		case 2:
			notPart1 = true

			maxValue := 1
			for number/maxValue > 0 {
				maxValue *= 10
			}

			sum = sum*maxValue + number

		}

		if sum > targetSum {
			return false, false
		}

		mutation /= 3
	}

	return sum == targetSum, notPart1
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

lineLoop:
	for _, line := range input.PlainLines() {
		numbers := aoc.ExtractNumbers(line)

		result := numbers[0]
		numbers = numbers[1:]

		numMutations := 3
		for i := 0; i < len(numbers)-2; i++ {
			numMutations *= 3
		}

		for mutation := 0; mutation < numMutations; mutation++ {
			if ok, notPart1 := isCorrect(result, numbers, mutation); ok {
				sumPart2 += result
				if !notPart1 {
					sumPart1 += result
				}

				continue lineLoop
			}
		}

	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
