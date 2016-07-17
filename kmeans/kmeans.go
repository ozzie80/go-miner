package kmeans

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ozzie80/go-miner/dt"
)

const MAX_ITERATION int = math.MaxInt64
const MEAN_DEV_MARGIN float64 = math.SmallestNonzeroFloat64

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type Cluster struct {
	Id       int
	Centroid dt.Vector
	Points   []dt.Vector
}

func (cluster *Cluster) String() string {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(fmt.Sprintf("Cluster Id: %d\n", cluster.Id))
	buffer.WriteString(fmt.Sprintf("Centroid: %v\n", cluster.Centroid))
	buffer.WriteString(fmt.Sprintf("Points: "))
	for _, point := range cluster.Points {
		buffer.WriteString(fmt.Sprintf("%v  ", point))
	}
	buffer.WriteString(fmt.Sprintf("\n\n"))

	return buffer.String()
}

type Params struct {
	K       int
	Points  []dt.Vector
	MaxIter int
}

func validateParameters(params Params) (Params, error) {
	if params.MaxIter <= 0 {
		params = Params{K: params.K, Points: params.Points, MaxIter: MAX_ITERATION}
	}

	if params.Points == nil || len(params.Points) == 0 {
		return params, errors.New("No input is provided.")
	}

	if params.K <= 0 || params.K > len(params.Points) {
		return params, errors.New("Invalid argument k.")
	}

	return params, nil
}

func Run(params Params) ([]Cluster, error) {
	params, err := validateParameters(params)
	if err != nil {
		return nil, err
	}

	centroids := selectKRandomPoints(params.K, params.Points)
	formerCentroids := make([]dt.Vector, params.K)
	meanDeviation := math.MaxFloat64
	var clusterPoints map[int][]dt.Vector

	for iter := 0; iter < params.MaxIter && meanDeviation > MEAN_DEV_MARGIN; iter++ {
		copy(formerCentroids, centroids)
		clusterPoints = getClusterPoints(centroids, params.Points)
		centroids = getCentroids(formerCentroids, clusterPoints)
		meanDeviation = getMeanCentroidDeviation(formerCentroids, centroids)
	}

	clusters := makeClusters(params.K, centroids, clusterPoints)

	return clusters, nil
}

func selectKRandomPoints(k int, points []dt.Vector) []dt.Vector {
	kPoints := make([]dt.Vector, k)
	numOfPoints := len(points)
	squaredDistances := make([]float64, numOfPoints)

	kPoints[0] = points[rand.Intn(numOfPoints)].Copy()

	for i := 1; i < k; i++ {
		sum := 0.0
		for j, point := range points {
			minDist := getMinDistanceToSelectedPoints(point, kPoints[:i])
			squaredDistances[j] = math.Pow(minDist, 2.0)
			sum += squaredDistances[j]
		}

		targetSum := rand.Float64() * sum
		index := 0
		for tempSum := squaredDistances[0]; tempSum < targetSum; tempSum += squaredDistances[index] {
			index++
		}

		kPoints[i] = points[index].Copy()
	}

	return kPoints
}

func getMinDistanceToSelectedPoints(point dt.Vector, points []dt.Vector) float64 {
	minDist := math.MaxFloat64

	for _, iPoint := range points {
		dist := point.Distance(iPoint)
		if dist < minDist {
			minDist = dist
		}
	}

	return minDist
}

func getClusterPoints(centroids, points []dt.Vector) map[int][]dt.Vector {
	clusterPoints := make(map[int][]dt.Vector)

	for _, point := range points {
		index := getIndexOfTheClosestCentroid(point, centroids)
		clusterPoints[index] = append(clusterPoints[index], point)
	}

	return clusterPoints
}

func getCentroids(formerCentroids []dt.Vector, clusterPoints map[int][]dt.Vector) []dt.Vector {
	numOfCentroids := len(formerCentroids)
	newCentroids := make([]dt.Vector, numOfCentroids)

	for i := 0; i < numOfCentroids; i++ {
		if len(clusterPoints[i]) == 0 {
			copy(newCentroids[i], formerCentroids[i])
		} else {
			newCentroids[i] = calculateCentroid(clusterPoints[i])
		}
	}

	return newCentroids
}

func calculateCentroid(points []dt.Vector) dt.Vector {
	numOfPoints := len(points)
	vectorDimSize := len(points[0])
	centroid := make(dt.Vector, vectorDimSize)
	column := make(dt.Vector, numOfPoints)

	for i := 0; i < vectorDimSize; i++ {
		for j := 0; j < numOfPoints; j++ {
			column[j] = points[j][i]
		}
		centroid[i] = column.Median()
	}

	return centroid
}

func getIndexOfTheClosestCentroid(point dt.Vector, centroids []dt.Vector) int {
	minDistance := math.MaxFloat64
	var minIndex int

	for i, centroid := range centroids {
		distance := point.Distance(centroid)
		if distance < minDistance {
			minDistance = distance
			minIndex = i
		}
	}

	return minIndex
}

func getMeanCentroidDeviation(formerCentroids, newCentroids []dt.Vector) float64 {
	numOfCentroids := len(formerCentroids)
	sum := 0.0

	for i := 0; i < numOfCentroids; i++ {
		sum += formerCentroids[i].Distance(newCentroids[i])
	}

	return sum / float64(numOfCentroids)
}

func makeClusters(k int, centroids []dt.Vector, clusterPoints map[int][]dt.Vector) []Cluster {
	clusters := make([]Cluster, 0)
	id := 0

	for i := 0; i < k; i++ {
		if len(clusterPoints[i]) > 0 {
			clusters = append(clusters, Cluster{id, centroids[i], clusterPoints[i]})
			id++
		}
	}

	return clusters
}
