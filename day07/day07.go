package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

/*

This was my first solution. It takes about 3 times as much runtime as the new solution

func mutationIsCorrect(targetSum int, numbers []int, mutation int) (bool, bool) {
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

func checkAllMutation(targetSum int, numbers []int) (bool, bool) {
	numMutations := 3
	for i := 0; i < len(numbers)-2; i++ {
		numMutations *= 3
	}

	for mutation := 0; mutation < numMutations; mutation++ {
		if ok, notPart1 := mutationIsCorrect(targetSum, numbers, mutation); ok {
			return true, notPart1
		}
	}

	return false, false
}
*/

// New solution: Start at the endresult and the end of list and use the reverse operations
// This allows early break for example if the division would result in a fraction or if
// the deconcatenation is not possible. Even though we use recursion here, this is much faster.

func checkWithRecursion(targetSum int, numbers []int, canBePart1 bool) (bool, bool) {
	if len(numbers) == 1 {
		return numbers[0] == targetSum, canBePart1
	}

	if targetSum <= 0 {
		return false, false
	}

	lastNumber := numbers[len(numbers)-1]

	ok, cbp1 := checkWithRecursion(targetSum-lastNumber, numbers[:len(numbers)-1], canBePart1)
	if ok {
		return true, cbp1
	}

	if targetSum%lastNumber == 0 {
		ok, cbp1 := checkWithRecursion(targetSum/lastNumber, numbers[:len(numbers)-1], canBePart1)
		if ok {
			return true, cbp1
		}
	}

	maxValue := 1
	for lastNumber/maxValue > 0 {
		maxValue *= 10
	}

	if targetSum%maxValue == lastNumber {
		ok, _ := checkWithRecursion(targetSum/maxValue, numbers[:len(numbers)-1], false)
		if ok {
			return true, false
		}
	}

	return false, false
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	for _, line := range input.PlainLines() {
		numbers := aoc.ExtractNumbers(line)

		result := numbers[0]
		numbers = numbers[1:]

		// if ok, notPart1 := checkAllMutation(result, numbers); ok {
		if ok, notPart1 := checkWithRecursion(result, numbers, true); ok {
			sumPart2 += result
			if !notPart1 {
				sumPart1 += result
			}
		}

	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
