package solutions

type Calculation struct {
	Numbers    [][]int
	Operations []string
}

func PartOne(numericData Calculation) int {
	totalValue := 0

	numColumns := len(numericData.Numbers[0])

	for columnIndex := range numColumns {
		nums := make([]int, 0)

		for i := range numericData.Numbers {
			nums = append(nums, numericData.Numbers[i][columnIndex])
		}
		totalValue += calculateValue(nums, numericData.Operations[columnIndex])
	}

	return totalValue
}

func calculateValue(numbers []int, operation string) int {
	result := 0

	for i := 0; i < len(numbers)-3; i++ {
		switch operation {
		case "+":
			result += (numbers[i] + numbers[i+1] + numbers[i+2] + numbers[i+3])
		case "*":
			result += (numbers[i] * numbers[i+1] * numbers[i+2] * numbers[i+3])
		}
	}
	return result
}
