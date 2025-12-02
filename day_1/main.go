package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DialDirection string;

const (
	Left  DialDirection = "L"
	Right DialDirection = "R"
)

const STARTING_POS = 50
const FILE_PATH = "input/aoc2025_day1_input.txt"

func turnDial(steps int, dir DialDirection, startPos int) int {
	if dir == Left {
		for i := 0; i < steps; i++ {
			startPos--

			if startPos < 0 {
				startPos = 99
			}
		}

		return startPos
	}
	
	if dir == Right {
		for i := 0; i < steps; i++ {
			startPos++

			if startPos > 99 {
				startPos = 0
			}
		}

		return startPos
	}

	return startPos
}

func main() {
	file, err := os.Open(FILE_PATH)

	if err != nil {
		fmt.Printf("Error opening file: %v \n", err)
	}

	defer file.Close()
	
	currentPos := STARTING_POS
	finalPasscode := 0;

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		dir := DialDirection(line[0:1])
		steps, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Printf("Error converting steps to integer: %v \n", err)
			break
		}

		output := turnDial(steps, dir, currentPos)
		currentPos = output

		if output == 0 {
			finalPasscode++
		}

		fmt.Printf("Turned dial %s %d steps to position %d \n", dir, steps, output)
	}

	if 	err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}
	fmt.Println("-----")
	fmt.Printf("Final passcode: %d \n", finalPasscode)
}