// Advent of Code 2025 Day 5 - Cafeteria Inventory (Go)package dec05

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	Start int64
	End   int64
}

func main() {
	data, err := os.ReadFile("input_05.txt")
	if err != nil {
		panic(err)
	}
	sections := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	// Parse ranges
	var freshRanges []Interval
	for _, line := range strings.Split(sections[0], "\n") {
		parts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		freshRanges = append(freshRanges, Interval{start, end})
	}

	// Parse IDs
	var availableIDs []int64
	scanner := bufio.NewScanner(strings.NewReader(sections[1]))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		id, _ := strconv.ParseInt(line, 10, 64)
		availableIDs = append(availableIDs, id)
	}

	// Part I: Count fresh available IDs
	result1 := 0
	for _, id := range availableIDs {
		for _, rng := range freshRanges {
			if rng.Start <= id && id <= rng.End {
				result1++
				break
			}
		}
	}

	// Part II: Merge intervals
	sorted := make([]Interval, len(freshRanges))
	copy(sorted, freshRanges)
	// insertion sort
	for i := 1; i < len(sorted); i++ {
		j := i
		for j > 0 && sorted[j].Start < sorted[j-1].Start {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
	}
	merged := []Interval{}
	for _, rng := range sorted {
		if len(merged) == 0 || rng.Start > merged[len(merged)-1].End+1 {
			merged = append(merged, rng)
		} else {
			if rng.End > merged[len(merged)-1].End {
				merged[len(merged)-1].End = rng.End
			}
		}
	}
	result2 := int64(0)
	for _, rng := range merged {
		result2 += rng.End - rng.Start + 1
	}

	fmt.Printf("First: %d\n", result1)
	fmt.Printf("Second: %d\n", result2)
}
