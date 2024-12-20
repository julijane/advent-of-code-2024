package main

import (
	"fmt"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type Entry struct {
	pos  aoc.Coordinate
	cost int
}

func findPath(gridWidth, gridHeight int, bytes []aoc.Coordinate) int {
	grid := aoc.NewGrid(gridWidth, gridHeight, '.')

	for _, bytePos := range bytes {
		grid.Set(bytePos, '#')
	}

	startPos := aoc.Coordinate{X: 0, Y: 0}
	endPos := aoc.Coordinate{X: gridWidth - 1, Y: gridHeight - 1}

	toCheck := []Entry{{pos: startPos, cost: 0}}
	added := make(map[aoc.Coordinate]struct{})

	for len(toCheck) > 0 {
		toCheckNow := toCheck[0]
		toCheck = toCheck[1:]

		if toCheckNow.pos == endPos {
			return toCheckNow.cost
		}

		for _, dir := range aoc.DirsStraight {
			newPos := toCheckNow.pos.Add(dir)

			if grid.Get(newPos, '#') == '#' {
				continue
			}

			if _, ok := added[newPos]; ok {
				continue
			}

			toCheck = append(toCheck, Entry{pos: newPos, cost: toCheckNow.cost + 1})

			added[newPos] = struct{}{}
		}
	}

	return -1
}

func calc(input *aoc.Input, doPart1, doPart2 bool, param ...any) (any, any) {
	resultPart2 := ""

	params := param[0].([]interface{})

	gridWidth := params[0].(int)
	gridHeight := params[1].(int)
	numBytesPart1 := params[2].(int)

	bytePos := aoc.Coordinates{}

	for _, line := range input.PlainLines() {
		val := aoc.ExtractNumbers(line)
		bytePos = append(bytePos, aoc.Coordinate{X: val[0], Y: val[1]})
	}

	resultPart1 := findPath(gridWidth, gridHeight, bytePos[:numBytesPart1])

	for numBytes := numBytesPart1 + 1; numBytes < len(bytePos); numBytes++ {
		if findPath(gridWidth, gridHeight, bytePos[:numBytes]) == -1 {
			resultPart2 = fmt.Sprintf("%d,%d", bytePos[numBytes-1].X, bytePos[numBytes-1].Y)
			break
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true, 7, 7, 12)
	aoc.Run("input.txt", calc, true, true, 71, 71, 1024)
}
