package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading from file, %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rotations []string

	for scanner.Scan() {
		line := scanner.Text()

		rotations = append(rotations, line)
	}

	p, err := findPassword(rotations)
	if err != nil {
		fmt.Printf("Error acquiring password, %v\n", err)
		return
	}

	fmt.Printf("Password: %d\n", p)
}

func findPassword(rotations []string) (int, error) {
	if len(rotations) == 0 {
		return 0, fmt.Errorf("no rotations")
	}

	password := 0
	currentPos := 50

	for i := range rotations {
		line := rotations[i]
		dir := line[0]
		numStr := line[1:]
		num, _ := strconv.Atoi(numStr)

		num %= 100
		if dir == 'L' {
			num = -num
		}

		currentPos += num

		if currentPos > 99 {
			currentPos -= 100
		} else if currentPos < 0 {
			currentPos += 100
		}
		if currentPos == 0 {
			password++
		}
	}

	return password, nil
}
