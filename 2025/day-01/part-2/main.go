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
	res := 0
	cur := 50
	prev := 0

	for i := range records {
		line := records[i]
		val, _ := strconv.Atoi(line[1:])

		res += val / 100

		val %= 100
		if line[0] == 'L' {
			val = -val
		}

		cur += val

		if cur > 99 {
			cur -= 100
		} else if cur < 0 {
			cur += 100
		}

		switch {
		case cur == 0:
			res++
		case cur >= prev && line[0] == 'L' && prev != 0:
			res++
		case cur <= prev && line[0] == 'R' && prev != 0:
			res++
		}
		prev = cur
	}

	return res, nil
}
