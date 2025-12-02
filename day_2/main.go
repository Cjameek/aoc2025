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

func getTotal(ids []string) int {
	total := 0

	for _, i := range ids {
		in, err := strconv.Atoi(i)

		if err != nil {
			fmt.Printf("Could not convert %d into number \n", in)
			break
		}

		total += in
	}

	return total
}

// Iterate through list of ranges and get the min / max of each range
// Then perform inner loop to find numbers in between ranges
func findInvalidIds(ids []string) []string {
	invalidIds := []string{}

	for _, id := range ids {
		split := strings.Split(id, "-")

		min, err := strconv.Atoi(split[0])

		if err != nil {
			fmt.Printf("Could not convert %d into min number \n", min)
			break
		}

		max, err := strconv.Atoi(split[1])

		if err != nil {
			fmt.Printf("Could not convert %d into max number \n", max)
			break
		}

		for i := min; i < max; i++ {
			s := strconv.Itoa(i)
			l := len(s)
			index := l / 2

			if(l > 1) {
				first := s[:index]
				second := s[index:]

				if(first == second) {
					invalidIds = append(invalidIds, s)
				}
			}
		}
	}

	return invalidIds
}

// Read the input file of comma separated ranges and split by ','
// Then return a sorted list of ranges
func readInputFile(file *os.File) []string {
	ranges := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
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

	ids := findInvalidIds(ranges)

	result := getTotal(ids)
	
	fmt.Println(result)
}