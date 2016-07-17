package kmeans

import (
	"reflect"
	"testing"

	"github.com/ozzie80/go-miner/dt"
)

var points = []dt.Vector{{1, 1}, {5, 2}, {1, 2}, {6, 1}, {6, 2}, {1, 3}}

func Test_SelectKRandomPointsReturnsKUniquePoints(t *testing.T) {
	numOfPoints := len(points)

	for k := 1; k <= numOfPoints; k++ {
		kPoints := selectKRandomPoints(k, points)
		observedPoints := make([]dt.Vector, k)

		for i := 0; i < k; i++ {
			expected := false
			actual := isObserved(kPoints[i], observedPoints)

			if expected != actual {
				t.Errorf("Test Failed, expected: '%b', got '%b'", expected, actual)
			}

			observedPoints[i] = kPoints[i]
		}
	}
}

func isObserved(point dt.Vector, observedPoints []dt.Vector) bool {
	for _, p := range observedPoints {
		if reflect.DeepEqual(p, point) {
			return true
		}
	}

	return false
}

func Test_GetIndexOfTheClosestCentroidReturnsTheExpectedIndex(t *testing.T) {
	centroids := points[3:]
	point := centroids[0]

	expected := 0
	actual := getIndexOfTheClosestCentroid(point, centroids)

	if expected != actual {
		t.Errorf("Test Failed, expected: '%d', got '%d'", expected, actual)
	}
}

func Test_WhenAllPointsAreTheSame_KMeansReturnsASingleCluster(t *testing.T) {
	equalPoints := []dt.Vector{{5, 2}, {5, 2}, {5, 2}, {5, 2}, {5, 2}, {5, 2}}
	k := 3
	params := Params{K: k, Points: equalPoints}
	clusters, _ := Run(params)

	expected := 1
	actual := len(clusters)

	if actual != expected {
		t.Errorf("Test Failed, expected: '%d', got '%d'", expected, actual)
	}
}

func Test_WhenKIsTwo_KMeansReturnsTwoOrLessClusters(t *testing.T) {
	k := 2
	params := Params{K: k, Points: points}
	clusters, _ := Run(params)

	minExpected := 1
	maxExpected := k
	actual := len(clusters)

	if actual > maxExpected || actual < minExpected {
		t.Errorf("Test Failed, expected(min, max): '%d, %d', got '%d'", minExpected, maxExpected, actual)
	}
}

func Test_WhenKIsThree_KMeansReturnsThreeOrLessClusters(t *testing.T) {
	k := 3
	params := Params{K: k, Points: points}
	clusters, _ := Run(params)

	minExpected := 1
	maxExpected := k
	actual := len(clusters)

	if actual > maxExpected || actual < minExpected {
		t.Errorf("Test Failed, expected(min, max): '%d, %d', got '%d'", minExpected, maxExpected, actual)
	}
}

func Benchmark_Run(b *testing.B) {
	k := 3

	for i := 0; i < b.N; i++ {
		params := Params{K: k, Points: points}
		Run(params)
	}
}
