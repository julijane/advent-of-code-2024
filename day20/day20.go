package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func countCheats(path aoc.Coordinates, timeToEndpos map[aoc.Coordinate]int, maxCheatLength, minSave int) int {
	maxTime := len(path) - 1

	numCheats := 0

	for time, pos := range path {
		candidatePos := make(map[aoc.Coordinate]struct{})

		for xDiff := 0; xDiff <= maxCheatLength; xDiff++ {
			for yDiff := 0; yDiff <= maxCheatLength-xDiff; yDiff++ {
				if xDiff+yDiff < 2 {
					continue
				}

				candidatePos[pos.Add(aoc.Coordinate{X: xDiff, Y: yDiff})] = struct{}{}
				candidatePos[pos.Add(aoc.Coordinate{X: xDiff, Y: -yDiff})] = struct{}{}
				candidatePos[pos.Add(aoc.Coordinate{X: -xDiff, Y: yDiff})] = struct{}{}
				candidatePos[pos.Add(aoc.Coordinate{X: -xDiff, Y: -yDiff})] = struct{}{}
			}
		}

		for candidate := range candidatePos {
			candidateTimeToEnd, ok := timeToEndpos[candidate]
			if !ok {
				continue
			}

			cheatedTime := time + candidateTimeToEnd + aoc.AbsInt(candidate.X-pos.X) + aoc.AbsInt(candidate.Y-pos.Y)
			savedTime := maxTime - cheatedTime

			if savedTime >= minSave {
				numCheats++
			}
		}
	}

	return numCheats
}

func calc(input *aoc.Input, _, _ bool, param ...any) (any, any) {
	resultPart1 := 0
	resultPart2 := 0

	params := param[0].([]interface{})
	minSave := params[0].(int)

	grid := input.Grid()
	startPos := grid.Find('S')
	endPos := grid.Find('E')

	path := aoc.Coordinates{startPos}

	// There is only a single path from start to the end and there are no deadends
	curPos := startPos
	lastPos := startPos
	for curPos != endPos {
		nextPos := curPos

		for _, dir := range aoc.DirsStraight {
			tryNextPos := curPos.Add(dir)

			if tryNextPos == lastPos {
				continue
			}

			if grid.Get(tryNextPos, '#') != '#' {
				nextPos = tryNextPos
				break
			}
		}

		if nextPos == curPos {
			panic("This should not happen")
		}

		path = append(path, nextPos)

		lastPos = curPos
		curPos = nextPos
	}

	timeToEndpos := make(map[aoc.Coordinate]int)
	for i, pos := range path {
		timeToEndpos[pos] = len(path) - i - 1
	}

	resultPart1 = countCheats(path, timeToEndpos, 2, minSave)
	resultPart2 = countCheats(path, timeToEndpos, 20, minSave)

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true, 50)
	aoc.Run("input.txt", calc, true, true, 100)
}
