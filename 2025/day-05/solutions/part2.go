package solutions

import (
	"sort"
)

type Range struct {
	Min int
	Max int
}

func PartTwo(min, max []int) (int, error) {
	ranges := make([]Range, len(min))

	for i := range min {
		if min[i] > max[i] {
			ranges[i] = Range{Min: max[i], Max: min[i]}
		} else {
			ranges[i] = Range{Min: min[i], Max: max[i]}
		}
	}

	if len(ranges) == 0 {
		return 0, nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		lastMerged := &merged[len(merged)-1]
		current := ranges[i]

		if current.Min <= lastMerged.Max+1 {
			if current.Max > lastMerged.Max {
				lastMerged.Max = current.Max
			}
		} else {
			merged = append(merged, current)
		}
	}

	totalUniqueCount := 0
	for _, r := range merged {
		totalUniqueCount += (r.Max - r.Min + 1)
	}

	return totalUniqueCount, nil
}
