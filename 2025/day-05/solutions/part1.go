package solutions

func PartOne(min, max, input []int) (int, error) {
	res := 0

	for i := range input {
		for j := range min {
			if input[i] >= min[j] && input[i] <= max[j] {
				res++
				break
			}
		}
	}
	return res, nil
}
