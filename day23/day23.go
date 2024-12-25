package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func calc(input *aoc.Input, _, doPart2 bool, param ...any) (any, any) {
	resultPart1 := 0

	connectionsFrom := make(map[string][]string)

	for _, c := range input.Lines {
		from := c.Data[0:2]
		to := c.Data[3:5]

		if _, ok := connectionsFrom[from]; !ok {
			connectionsFrom[from] = []string{}
		}
		connectionsFrom[from] = append(connectionsFrom[from], to)

		if _, ok := connectionsFrom[to]; !ok {
			connectionsFrom[to] = []string{}
		}
		connectionsFrom[to] = append(connectionsFrom[to], from)
	}

	seenThreeGroups := make(map[string]struct{})

	for from := range connectionsFrom {
		for _, to1 := range connectionsFrom[from] {
			// search:
			for _, to2 := range connectionsFrom[from] {
				if to1 == to2 {
					continue
				}

				to1Conn := connectionsFrom[to1]

				if slices.Contains(to1Conn, to2) {
					l := []string{from, to1, to2}
					slices.Sort(l)
					lS := fmt.Sprintf("%s,%s,%s", l[0], l[1], l[2])

					if _, ok := seenThreeGroups[lS]; ok {
						continue
					}

					seenThreeGroups[lS] = struct{}{}

					if l[0][0] == 't' || l[1][0] == 't' || l[2][0] == 't' {
						resultPart1++
					}
				}
			}
		}
	}

	groups := [][]string{}
	for from, connectionsTo := range connectionsFrom {
		found := false

	groupSearch:
		for ig, g := range groups {
			if slices.Contains(g, from) {
				continue
			}

			for _, to := range g {
				if !slices.Contains(connectionsFrom[to], from) {
					continue groupSearch
				}
			}

			groups[ig] = append(groups[ig], from)
			found = true
		}

		if !found {
			for _, to := range connectionsTo {
				groups = append(groups, []string{from, to})
			}
		}
	}

	resultPart2 := ""

	seenGroup := make(map[string]struct{})

	for _, group := range groups {
		slices.Sort(group)
		groupString := strings.Join(group, ",")

		if _, ok := seenGroup[groupString]; ok {
			continue
		}
		seenGroup[groupString] = struct{}{}

		if len(groupString) > len(resultPart2) {
			resultPart2 = groupString
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	// aoc.Run("sample2.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
