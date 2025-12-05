package solutions

// Didn't have time to finish Part Two, So got a Solution to view and learn from.

func deepCopyGrid(original Grid) Grid {
	newGrid := make(Grid, len(original))
	for i := range original {
		newGrid[i] = make([]rune, len(original[i]))
		copy(newGrid[i], original[i])
	}
	return newGrid
}

func PartTwo(records []string) (int, error) {

	var initialGrid Grid
	for _, row := range records {
		initialGrid = append(initialGrid, []rune(row))
	}

	currentGrid := deepCopyGrid(initialGrid)
	totalRemovedCount := 0

	for {
		removals := make([][2]int, 0)

		R := len(currentGrid)
		if R == 0 {
			break
		}
		C := len(currentGrid[0])

		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {

				if currentGrid[r][c] == '@' {

					neighbors := getNeighbors(currentGrid, r, c)

					adjacentRollsCount := 0
					for _, neighbor := range neighbors {
						if neighbor == '@' {
							adjacentRollsCount++
						}
					}

					if adjacentRollsCount < 4 {
						removals = append(removals, [2]int{r, c})
					}
				}
			}
		}

		if len(removals) == 0 {
			break
		}

		for _, pos := range removals {
			r, c := pos[0], pos[1]
			currentGrid[r][c] = '.'
		}

		totalRemovedCount += len(removals)
	}

	return totalRemovedCount, nil
}
