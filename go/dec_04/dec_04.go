package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
	{0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func main() {
	file, err := os.Open("input_04.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	rows := len(lines)
	cols := 0
	if rows > 0 {
		cols = len(lines[0])
	}

	// Part One
	result1 := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lines[r][c] != '@' {
				continue
			}
			adj := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && lines[nr][nc] == '@' {
					adj++
					if adj >= 4 {
						break
					}
				}
			}
			if adj < 4 {
				result1++
			}
		}
	}

	// Part Two
	grid := make([][]byte, rows)
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	result2 := 0
	for {
		toRemove := [][2]int{}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] != '@' {
					continue
				}
				if adjCount(grid, r, c) < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, rc := range toRemove {
			grid[rc[0]][rc[1]] = '.'
		}
		result2 += len(toRemove)
	}

	fmt.Printf("first: %d\n", result1)
	fmt.Printf("second: %d\n", result2)
}

func adjCount(grid [][]byte, r, c int) int {
	rows := len(grid)
	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}
	cnt := 0
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
			cnt++
		}
	}
	return cnt
}
