package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input_02.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read input_02.txt:", err)
		return
	}

	var result1 int64
	var result2 int64

	// iterate through ranges
	for _, token := range strings.Split(string(input), ",") {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		parts := strings.Split(token, "-")
		if len(parts) != 2 {
			// malformed token, skip
			continue
		}

		start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			// skip malformed range
			continue
		}

		// iterate through range
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			number := len(s)

			// check if symetrical
			if number%2 == 0 {
				mid := number / 2
				if s[:mid] == s[mid:] {
					result1 += int64(i)
				}
			}

			// second test: exists a block length L (1..n/2) that divides n
			// and s is the first-block repeated n/L times
			found := false
			for L := 1; L <= number/2; L++ {
				if number%L != 0 {
					continue
				}
				rep := number / L
				prefix := s[:L]
				if strings.Repeat(prefix, rep) == s {
					result2 += int64(i)
					found = true
					break
				}
			}
			_ = found
		}
	}

	fmt.Printf("first: %d\n", result1)
	fmt.Printf("second: %d\n", result2)
}
