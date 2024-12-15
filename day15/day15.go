package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func doCalc(gridText []string, instructions []byte, isPart2 bool) int {
	if isPart2 {
		r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

		for i := 0; i < len(gridText); i++ {
			gridText[i] = r.Replace(gridText[i])
		}
	}

	grid := aoc.NewGridFromStrings(gridText)

	robotPos := grid.Find('@')

instLoop:
	for _, instruction := range instructions {
		moveVal := aoc.DirsStraight[instruction]

		allPosToMove := []aoc.Coordinate{robotPos}
		currentFieldContent := make(map[aoc.Coordinate]byte)

		for len(allPosToMove) > 0 {
			nowToMove := allPosToMove[0]
			allPosToMove = allPosToMove[1:]

			if _, ok := currentFieldContent[nowToMove]; ok {
				// we have been here before
				continue
			}

			currentFieldContent[nowToMove] = grid.Get(nowToMove, '#')

			nextPos := nowToMove.Add(moveVal)
			switch grid.Get(nextPos, '#') {
			case '#':
				continue instLoop
			case '[':
				rightOfNextPos := nextPos.Add(aoc.DirR)
				allPosToMove = append(allPosToMove, nextPos)
				allPosToMove = append(allPosToMove, rightOfNextPos)
			case ']':
				leftOfNextPos := nextPos.Add(aoc.DirL)
				allPosToMove = append(allPosToMove, nextPos)
				allPosToMove = append(allPosToMove, leftOfNextPos)
			case 'O':
				allPosToMove = append(allPosToMove, nextPos)
			}
		}

		for pos := range currentFieldContent {
			grid.Set(pos, '.')
		}

		for pos := range currentFieldContent {
			grid.Set(pos.Add(moveVal), currentFieldContent[pos])
		}

		grid.Set(robotPos, '.')
		robotPos = robotPos.Add(moveVal)
		grid.Set(robotPos, '@')
	}

	result := 0
	for _, multiplePos := range grid.FindMultipleAll("O[") {
		for _, pos := range multiplePos {
			result += 100*pos.Y + pos.X
		}
	}

	return result
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	blocks := input.TextBlocks()

	instructions := []byte{}
	possibleInstruction := "^>v<"

	for _, line := range blocks[1] {
		for _, char := range line {
			instructions = append(instructions, byte(strings.Index(possibleInstruction, string(char))))
		}
	}

	var sumPart1 int
	if doPart1 {
		sumPart1 = doCalc(blocks[0], instructions, false)
	}

	var sumPart2 int
	if doPart2 {
		sumPart2 = doCalc(blocks[0], instructions, true)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
