package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	antennaLocations := make(map[byte][]aoc.Coordinate)

	grid := input.Grid()

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			c := grid.Get(aoc.Coordinate{X: x, Y: y}, '~')
			if c == '.' {
				continue
			}

			if _, ok := antennaLocations[c]; !ok {
				antennaLocations[c] = make([]aoc.Coordinate, 0)
			}

			antennaLocations[c] = append(antennaLocations[c], aoc.Coordinate{X: x, Y: y})
		}
	}

	antinodeLocationsPart1 := make(map[aoc.Coordinate]struct{})
	antinodeLocationsPart2 := make(map[aoc.Coordinate]struct{})

	for c := range antennaLocations {
		locs := antennaLocations[c]

		for antenna1 := 0; antenna1 < len(locs); antenna1++ {
			for antenna2 := antenna1 + 1; antenna2 < len(locs); antenna2++ {
				diff := aoc.Coordinate{
					X: locs[antenna1].X - locs[antenna2].X,
					Y: locs[antenna1].Y - locs[antenna2].Y,
				}

				for pos, n := locs[antenna1], 1; grid.Inside(pos); pos, n = pos.Add(diff), n-1 {
					antinodeLocationsPart2[pos] = struct{}{}

					if n == 0 {
						antinodeLocationsPart1[pos] = struct{}{}
					}
				}

				for pos, n := locs[antenna2], 1; grid.Inside(pos); pos, n = pos.Subtract(diff), n-1 {
					antinodeLocationsPart2[pos] = struct{}{}

					if n == 0 {
						antinodeLocationsPart1[pos] = struct{}{}
					}

				}
			}
		}
	}

	return len(antinodeLocationsPart1), len(antinodeLocationsPart2)
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
