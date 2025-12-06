package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromFile() []string {
	const filePath = "input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
		return []string{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		line = strings.Join(fields, " ")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return []string{}
	}

	return lines
}
