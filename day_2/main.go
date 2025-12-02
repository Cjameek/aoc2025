package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FILE_PATH = "input/aoc2025_day2_input.txt"

func hasRepeatingHalves(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	
	if l%2 != 0 {
		return false
	}
	
	mid := l / 2
	return s[:mid] == s[mid:]
}


func getTotalFromRanges(ranges []string) int {
	total := 0
	m := make(map[int]bool)

	for _, r := range ranges {
		split := strings.Split(r, "-")

		min, err := strconv.Atoi(split[0])

		if err != nil {
			fmt.Printf("Could not convert %d into min number \n", min)
			continue
		}

		max, err := strconv.Atoi(split[1])

		if err != nil {
			fmt.Printf("Could not convert %d into max number \n", max)
			continue
		}

		for i := min; i <= max; i++ {
			if !m[i] && hasRepeatingHalves(i) {
				m[i] = true
				total += i
			}
		}
	}

	return total
}

// Read the input file of comma separated ranges and split by ','
// Then return a sorted list of ranges
func readInputFile(file *os.File) []string {
	ranges := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		r := strings.Split(line, ",")
		ranges = append(ranges, r...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}

	sort.Strings(ranges)

	return ranges
}

func main() {
	file, err := os.Open(FILE_PATH)

	if err != nil {
		fmt.Printf("Error opening file: %v \n", err)
	}

	defer file.Close()

	ranges := readInputFile(file)

	result := getTotalFromRanges(ranges)
	
	fmt.Println(result)
}