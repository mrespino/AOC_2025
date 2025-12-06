// Advent of Code 2025 Day 6 - Cephalopod Math Worksheet (Go)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input_06.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("answer: %d\n", trySquidMath(lines))
}

func trySquidMath(lines []string) int {
	if len(lines) == 0 {
		return 0
	}
	ops := lines[len(lines)-1]
	nums := lines[:len(lines)-1]
	// Pad all rows to the same length
	maxLen := len(ops)
	for i := range nums {
		if len(nums[i]) < maxLen {
			nums[i] += strings.Repeat(" ", maxLen-len(nums[i]))
		}
	}
	// Transpose columns
	cols := make([][]byte, maxLen)
	for i := 0; i < maxLen; i++ {
		col := make([]byte, len(nums)+1)
		for j := 0; j < len(nums); j++ {
			col[j] = nums[j][i]
		}
		col[len(nums)] = ops[i]
		cols[i] = col
	}
	// Find problem ranges
	problemRanges := [][2]int{}
	inProblem := false
	start := 0
	for i, col := range cols {
		allSpace := true
		for j := 0; j < len(col); j++ {
			if col[j] != ' ' {
				allSpace = false
				break
			}
		}
		if allSpace {
			if inProblem {
				problemRanges = append(problemRanges, [2]int{start, i})
				inProblem = false
			}
		} else {
			if !inProblem {
				inProblem = true
				start = i
			}
		}
	}
	if inProblem {
		problemRanges = append(problemRanges, [2]int{start, maxLen})
	}
	answer := 0
	for _, pr := range problemRanges {
		start, end := pr[0], pr[1]
		op := strings.TrimSpace(ops[start:end])
		numDigits := make([][]byte, end-start)
		for col := end - 1; col >= start; col-- {
			for row := 0; row < len(nums); row++ {
				digit := nums[row][col]
				if digit != ' ' {
					numDigits[end-1-col] = append(numDigits[end-1-col], digit)
				}
			}
		}
		numbers := []int{}
		for _, digits := range numDigits {
			if len(digits) > 0 {
				numStr := string(digits)
				num := 0
				fmt.Sscanf(numStr, "%d", &num)
				numbers = append(numbers, num)
			}
		}
		if op == "+" {
			sum := 0
			for _, n := range numbers {
				sum += n
			}
			answer += sum
		} else if op == "*" {
			prod := 1
			for _, n := range numbers {
				prod *= n
			}
			answer += prod
		} else {
			panic(fmt.Sprintf("bullshit op: %s", op))
		}
	}
	return answer
}
