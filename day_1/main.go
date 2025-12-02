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

func turnDial(steps int, dir DialDirection, startPos int) (int, int) {
	const dialSize = 100
	
	if dir == Left {
		finalPos := (startPos - steps) % dialSize

		if finalPos < 0 {
			finalPos += dialSize
		}
		
		iterations := 0

		if steps > 0 {
			stepsToZero := startPos

			if stepsToZero == 0 {
				stepsToZero = dialSize
			}
			
			if steps >= stepsToZero {
				iterations = 1 + (steps-stepsToZero)/dialSize
			}
		}
		
		return finalPos, iterations
	}
	
	if dir == Right {
		finalPos := (startPos + steps) % dialSize
		iterations := (startPos + steps) / dialSize
		
		return finalPos, iterations
	}
	
	return startPos, 0
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

		output, iteration := turnDial(steps, dir, currentPos)
		currentPos = output
		finalPasscode = finalPasscode + iteration
		

		fmt.Printf("Turned dial %s %d steps to position %d \n", dir, steps, output)
	}

	if 	err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}
	
	fmt.Println("-----")
	fmt.Printf("Final passcode: %d \n", finalPasscode)
}