package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input_03.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var banks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		banks = append(banks, strings.TrimSpace(scanner.Text()))
	}

	result1 := 0
	result2 := 0

	// First
	for _, bank := range banks {
		maxJoltage := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				joltage := atoi(bank[i:i+1] + bank[j:j+1])
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		result1 += maxJoltage
	}

	// Second
	for _, bank := range banks {
		selected := make([]byte, 0, 12)
		startPos := 0
		needed := 12
		for step := 0; step < 12; step++ {
			endSearch := len(bank) - needed + 1
			bestPos := startPos
			bestDigit := bank[startPos]
			for pos := startPos; pos < endSearch; pos++ {
				if bank[pos] > bestDigit {
					bestDigit = bank[pos]
					bestPos = pos
				}
			}
			selected = append(selected, bestDigit)
			startPos = bestPos + 1
			needed--
		}
		joltageStr := string(selected)
		joltage := atoi(joltageStr)
		result2 += joltage
	}

	fmt.Printf("first: %d\n", result1)
	fmt.Printf("second: %d\n", result2)
}

func atoi(s string) int {
	var n int
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
