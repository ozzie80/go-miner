package kmeans

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ozzie80/go-miner/dt"
)

func Test_KmeansWithPointsFromInputFiles(t *testing.T) {
	inputFiles := []string{
		"..\\data\\kmeans\\input_1k.csv",
		"..\\data\\kmeans\\input_10k.csv",
		"..\\data\\kmeans\\input_100k.csv"}

	for _, inputFile := range inputFiles {
		runTest(inputFile, t)
	}

}

func runTest(inputFile string, t *testing.T) {
	points := readPointsFromFile(inputFile)

	start := time.Now()
	k := 5
	params := Params{K: k, Points: points}
	clusters, _ := Run(params)
	timeTrack(start, "Run")

	minExpected := 1
	maxExpected := k
	actual := len(clusters)

	if actual > maxExpected || actual < minExpected {
		t.Errorf("Test Failed, expected(min, max): '%d, %d', got '%d'", minExpected, maxExpected, actual)
	}

	reportQualityIndexes(clusters)
}

func readPointsFromFile(filePath string) []dt.Vector {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	data, _ := reader.ReadAll()
	points := make([]dt.Vector, len(data))
	dimSize := len(data[0])

	for i, line := range data {
		vector := make(dt.Vector, dimSize)
		for j, field := range line {
			val, _ := strconv.ParseFloat(field, 64)
			vector[j] = val
		}
		points[i] = vector
	}

	return points
}

func reportQualityIndexes(clusters []Cluster) {
	qualityIndexes := []QualityIndex{DunnIndex{}, DaviesBouldinIndex{}}

	for _, index := range qualityIndexes {
		reportQualityIndex(index, clusters)
	}
}

func reportQualityIndex(index QualityIndex, clusters []Cluster) {
	start := time.Now()
	score := index.GetScore(clusters)
	fmt.Printf("%s Score: %f\n", index.Name(), score)
	timeTrack(start, index.Name())
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
