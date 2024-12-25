package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type LogicGate struct {
	reg1      string
	reg2      string
	op        string
	regTarget string
}

type Logic []LogicGate

func (l *Logic) find(reg1, reg2, op string) (string, bool) {
	if reg2 < reg1 {
		reg1, reg2 = reg2, reg1
	}

	for _, lg := range *l {
		if lg.reg1 == reg1 && lg.reg2 == reg2 && lg.op == op {
			return lg.regTarget, true
		}
	}

	return "", false
}

func calc(input *aoc.Input, _, doPart2 bool, param ...any) (any, any) {
	blocks := input.TextBlocks()

	registers := make(map[string]bool)
	for _, startReg := range blocks[0] {
		regName := startReg[0:3]
		value := startReg[5:6] == "1"
		registers[regName] = value
	}

	logic := Logic{}
	for _, logicLine := range blocks[1] {
		el := strings.Split(logicLine, " ")

		reg1 := el[0]
		reg2 := el[2]

		if reg2 < reg1 {
			reg1, reg2 = reg2, reg1
		}

		logic = append(logic, LogicGate{
			reg1:      reg1,
			reg2:      reg2,
			op:        el[1],
			regTarget: el[4],
		})
	}

	// for _, logic := range logic {
	// 	fmt.Println(logic.reg1, " ", logic.reg2, " ", logic.op, " -> ", logic.regTarget)
	// }

	// os.Exit(1)
	// Part 1

	allResolved := false
	for !allResolved {
		allResolved = true

		for _, logic := range logic {

			reg1Val, ok1 := registers[logic.reg1]
			reg2Val, ok2 := registers[logic.reg2]

			if !ok1 || !ok2 {
				allResolved = false
				continue
			}

			switch logic.op {
			case "AND":
				registers[logic.regTarget] = reg1Val && reg2Val
			case "OR":
				registers[logic.regTarget] = reg1Val || reg2Val
			case "XOR":
				registers[logic.regTarget] = reg1Val != reg2Val
			}
		}
	}

	bit := 0
	resultPart1 := 0
	for {
		regName := fmt.Sprintf("z%02d", bit)
		val, ok := registers[regName]
		if !ok {
			break
		}

		if val {
			resultPart1 = resultPart1 | (1 << bit)
		}

		bit++
	}

	if !doPart2 {
		return resultPart1, ""
	}

	// Part 2

	regLastCarry := ""

	swapped := []string{}

	for bit := range 45 {
		// first halfadder takes the input bits
		regXn := fmt.Sprintf("x%02d", bit)
		regYn := fmt.Sprintf("y%02d", bit)

		// outputs of the first halfadder
		regHA1result, _ := logic.find(regXn, regYn, "XOR")
		regHA1carry, _ := logic.find(regXn, regYn, "AND")

		// the logic for input bit 0 has only the half adder
		if bit == 0 {
			regLastCarry = regHA1carry
			continue
		}

		// the other input bits have a full adder (= second half adder plus an OR)

		// check if the output wires of the first half adder are swapped
		if _, ok := logic.find(regLastCarry, regHA1result, "AND"); !ok {
			regHA1carry, regHA1result = regHA1result, regHA1carry
			swapped = append(swapped, regHA1carry, regHA1result)
		}

		regHA2carry, _ := logic.find(regLastCarry, regHA1result, "AND")
		regHA2result, _ := logic.find(regLastCarry, regHA1result, "XOR")

		if regHA1result[0] == 'z' {
			// the output of the first half adder was miswired to be the output of the full adder
			regHA1result, regHA2result = regHA2result, regHA1result
			swapped = append(swapped, regHA1result, regHA2result)
		} else if regHA1carry[0] == 'z' {
			// the carry of the first half adder was miswired to be the output of the full adder
			regHA1carry, regHA2result = regHA2result, regHA1carry
			swapped = append(swapped, regHA1carry, regHA2result)
		} else if regHA2carry[0] == 'z' {
			// the carry of the second half adder was miswired to be the output of the full adder
			regHA2carry, regHA2result = regHA2result, regHA2carry
			swapped = append(swapped, regHA2carry, regHA2result)
		}

		// the final carry of the full adder
		regFACarry, _ := logic.find(regHA2carry, regHA1carry, "OR")

		if regFACarry[0] == 'z' && regFACarry != "z45" {
			// the carry of the full adder was miswired to be the output of the full adder
			// (the carry of the last bit goes into the final z bit, z45)
			regFACarry, regHA2result = regHA2result, regFACarry
			swapped = append(swapped, regFACarry, regHA2result)
		}

		regLastCarry = regFACarry
	}

	// Sort and join swapped wires
	sort.Strings(swapped)
	resultPart2 := strings.Join(swapped, ",")

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, true, false)
	aoc.Run("input.txt", calc, true, true)
}
