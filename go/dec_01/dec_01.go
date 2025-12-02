package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input_01.txt")
	if err != nil {
		fmt.Println("Error reading input", "error", err)
		return
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	first(scanner)

	input.Seek(0, 0)
	scanner = bufio.NewScanner(input)
	second(scanner)
}

func first(scanner *bufio.Scanner) {
	boundary := 100
	dial := 50
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		switch line[0] {
		case 'R':

		case 'L':
			steps = -steps
		default:

			continue
		}

		dial = ((dial+steps)%boundary + boundary) % boundary

		if dial == 0 {
			result++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanning input failed", "error", err)
		return
	}

	fmt.Println("First", result)
}

func second(scanner *bufio.Scanner) {
	dial := 50
	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("invalid step count", "line", line, "error", err)
			continue
		}

		step := -1
		if line[0] == 'R' {
			step = 1
		}

		for i := 0; i < steps; i++ {
			dial = (dial + step + 100) % 100
			if dial == 0 {
				result++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanning input failed", "error", err)
		return
	}

	fmt.Println("Second:", result)
}
