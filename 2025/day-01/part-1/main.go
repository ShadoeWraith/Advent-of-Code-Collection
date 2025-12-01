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
	var records []string

	for scanner.Scan() {
		line := scanner.Text()

		records = append(records, line)
	}

	p, _ := findPassword(records)

	fmt.Printf("Password: %d\n", p)
}

func findPassword(records []string) (int, error) {
	password := 0
	currentPos := 50

	for i := range records {
		line := records[i]
		num, _ := strconv.Atoi(line[1:])

		num %= 100
		if line[0] == 'L' {
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
