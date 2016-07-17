package dt

import (
	"math"
	"sort"
)

type Vector []float64

func (vector Vector) Copy() Vector {
	return append(vector)
}

func (vector Vector) Distance(aVector Vector) float64 {
	if len(vector) != len(aVector) {
		return math.NaN()
	}

	sum := 0.0
	for i, val := range vector {
		diff := val - aVector[i]
		sum += diff * diff
	}

	return math.Pow(sum, 0.5)
}

func (vector Vector) Median() float64 {
	copyVector := vector.Copy()
	sort.Float64s(copyVector)

	midIndex := len(copyVector) / 2
	median := copyVector[midIndex]

	if len(copyVector)%2 == 0 {
		median = (median + copyVector[midIndex-1]) / 2.0
	}

	return median
}
