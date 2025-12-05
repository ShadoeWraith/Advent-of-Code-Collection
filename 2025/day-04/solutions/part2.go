package solutions

// Was late to starting so didn't do part 2

func PartTwo(records []string) (int, error) {

	var grid Grid

	for _, row := range records {
		grid = append(grid, []rune(row))
	}

	accessibleRollsCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '@' {

				neighbors := getNeighbors(grid, i, j)

				adjacentRollsCount := 0
				for _, neighbor := range neighbors {
					if neighbor == '@' {
						adjacentRollsCount++
					}
				}

				if adjacentRollsCount < 4 {
					accessibleRollsCount++
				}
			}
		}
	}

	return accessibleRollsCount, nil
}
