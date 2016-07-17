package dt

import (
	"math"
	"reflect"
	"testing"
)

func Test_CopyReturnsADeepCopyOfAVector(t *testing.T) {
	vector := Vector{1, 2, 3, 4}
	copyVector := vector.Copy()

	if &copyVector == &vector || !reflect.DeepEqual(copyVector, vector) {
		t.Error("Test Failed, Copy function failed to create a deep copy of the given vector")
	}
}

func Test_WhenVectorsHaveEqualLength_DistanceReturnsTheExpectedDistance(t *testing.T) {
	vectorX := Vector{1, 2, 3, 4}
	vectorY := Vector{4, 3, 2, 1}

	expected := math.Pow(20, 0.5)
	actual := vectorX.Distance(vectorY)

	if expected != actual {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Test_WhenVectorsHaveUnequalLengths_DistanceReturnsNaN(t *testing.T) {
	vectorX := Vector{1, 2, 3, 3}
	vectorY := Vector{1, 2, 3}

	distance := vectorX.Distance(vectorY)

	if !math.IsNaN(distance) {
		t.Errorf("Test Failed, expected: NaN, got '%f'", distance)
	}
}

func Test_WhenVectorLengthIsOdd_MedianReturnsTheMiddleElement(t *testing.T) {
	vector := Vector{3, 2, 6, 5.2, 5.1}

	expected := 5.1
	actual := vector.Median()

	if expected != actual {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Test_WhenVectorLengthIsEven_MedianReturnsTheAverageOfMiddleElements(t *testing.T) {
	vector := Vector{3, 2, 6, 5}

	expected := 4.0
	actual := vector.Median()

	if expected != actual {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Benchmark_Copy(b *testing.B) {
	vector := Vector{1, 2, 3, 4}

	for i := 0; i < b.N; i++ {
		vector.Copy()
	}
}
