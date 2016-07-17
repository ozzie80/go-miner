package kmeans

import (
	"testing"

	"github.com/ozzie80/go-miner/dt"
)

type TestDouble struct {
	clusters []Cluster
	dScore   float64
	dbScore  float64
}

func Test_WhenThereIsASingleCluster_DunnIndexGetScoreReturnsZero(t *testing.T) {
	testDouble := getTestDouble()
	dIndex := DunnIndex{}
	score := dIndex.GetScore(testDouble.clusters[:1])

	expected := 0.0
	actual := score

	if actual != expected {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Test_WhenThereAreKClusters_DunnIndexGetScoreReturnsTheExpectedResult(t *testing.T) {
	testDouble := getTestDouble()
	dIndex := DunnIndex{}
	score := dIndex.GetScore(testDouble.clusters)

	expected := testDouble.dScore
	actual := score

	if actual != expected {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Test_WhenThereIsASingleCluster_DaviesBouldinIndexGetScoreReturnsZero(t *testing.T) {
	testDouble := getTestDouble()
	dbIndex := DaviesBouldinIndex{}
	score := dbIndex.GetScore(testDouble.clusters[:1])

	expected := 0.0
	actual := score

	if actual != expected {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func Test_WhenThereAreKClusters_DaviesBouldinIndexGetScoreReturnsTheExpectedResult(t *testing.T) {
	testDouble := getTestDouble()
	dbIndex := DaviesBouldinIndex{}
	score := dbIndex.GetScore(testDouble.clusters)

	expected := testDouble.dbScore
	actual := score

	if actual != expected {
		t.Errorf("Test Failed, expected: '%f', got '%f'", expected, actual)
	}
}

func getTestDouble() TestDouble {
	clusterPoints := make(map[int][]dt.Vector)
	clusterPoints[0] = []dt.Vector{{1, 1}, {1, 2}, {2, 2}}
	clusterPoints[1] = []dt.Vector{{5, 2}, {5, 1}}
	clusterPoints[2] = []dt.Vector{{7, 5}, {7, 6}, {8, 8}}

	centroids := []dt.Vector{{1, 2}, {5, 1.5}, {7, 6}}
	clusters := make([]Cluster, len(clusterPoints))

	for i, points := range clusterPoints {
		clusters[i] = Cluster{i, centroids[i], points}
	}

	dScore := 1.8027756377319943
	dbScore := 0.3101936139590087
	testDouble := TestDouble{clusters, dScore, dbScore}

	return testDouble
}

func Benchmark_GetDunnIndex(b *testing.B) {
	clusters := getTestDouble().clusters
	dIndex := DunnIndex{}

	for i := 0; i < b.N; i++ {
		dIndex.GetScore(clusters)
	}
}

func Benchmark_GetDaviesBouldinIndex(b *testing.B) {
	clusters := getTestDouble().clusters
	dbIndex := DaviesBouldinIndex{}

	for i := 0; i < b.N; i++ {
		dbIndex.GetScore(clusters)
	}
}
