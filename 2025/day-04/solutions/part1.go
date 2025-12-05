package solutions

type Grid [][]rune

func getNeighbors(grid Grid, row int, col int) []rune {
	offsets := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	foundNeighbors := []rune{}

	R := len(grid)
	if R == 0 {
		return foundNeighbors
	}
	C := len(grid[0])

	for _, offset := range offsets {
		dr, dc := offset[0], offset[1]

		nRow := row + dr
		nCol := col + dc

		if nRow >= 0 && nRow < R && nCol >= 0 && nCol < C {
			neighborChar := grid[nRow][nCol]
			foundNeighbors = append(foundNeighbors, neighborChar)
		}
	}

	return foundNeighbors
}

func PartOne(records []string) (int, error) {

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
