package kmeans

import (
	"math"
)

type QualityIndex interface {
	GetScore(clusters []Cluster) float64
	Name() string
}

type DunnIndex struct{}
type DaviesBouldinIndex struct{}

func (q DunnIndex) Name() string {
	return "Dunn Index"
}

func (q DaviesBouldinIndex) Name() string {
	return "Davies-Bouldin Index"
}

func (q DunnIndex) GetScore(clusters []Cluster) float64 {
	minInterDist := getMinInterClusterDistance(clusters)
	maxIntraDist := getMaxIntraClusterDistance(clusters)

	return minInterDist / maxIntraDist
}

func (q DaviesBouldinIndex) GetScore(clusters []Cluster) float64 {
	sum := 0.0
	for _, cluster := range clusters {
		sum += getMaxWithinToBetweenClusterDistanceRatio(cluster, clusters)
	}

	numOfClusters := len(clusters)
	return sum / float64(numOfClusters)
}

func getMinInterClusterDistance(clusters []Cluster) float64 {
	numOfClusters := len(clusters)
	if numOfClusters <= 1 {
		return 0.0
	}

	minDistance := math.MaxFloat64
	for i := 0; i < numOfClusters; i++ {
		for j := i + 1; j < numOfClusters; j++ {
			distance := getInterClusterDistance(clusters[i], clusters[j])
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

func getMaxIntraClusterDistance(clusters []Cluster) float64 {
	maxDistance := 0.0
	for _, cluster := range clusters {
		distance := getMaxDistanceFromCentroid(cluster)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func getMaxDistanceFromCentroid(cluster Cluster) float64 {
	maxDistance := 0.0
	for _, point := range cluster.Points {
		distance := point.Distance(cluster.Centroid)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func getMaxWithinToBetweenClusterDistanceRatio(cluster Cluster, clusters []Cluster) float64 {
	maxDistance := 0.0
	for _, iCluster := range clusters {
		if cluster.Id == iCluster.Id {
			continue
		}
		distance := getWithinToBetweenClusterDistanceRatio(cluster, iCluster)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func getWithinToBetweenClusterDistanceRatio(clusterX, clusterY Cluster) float64 {
	avgX := getAverageDistaceFromCentroid(clusterX)
	avgY := getAverageDistaceFromCentroid(clusterY)
	distance := getInterClusterDistance(clusterX, clusterY)

	return (avgX + avgY) / distance
}

func getAverageDistaceFromCentroid(cluster Cluster) float64 {
	sum := 0.0
	for _, point := range cluster.Points {
		sum += point.Distance(cluster.Centroid)
	}

	numOfPoints := len(cluster.Points)
	return sum / float64(numOfPoints)
}

func getInterClusterDistance(clusterX, clusterY Cluster) float64 {
	fromPoint := clusterX.Centroid
	toPoint := clusterY.Centroid

	return fromPoint.Distance(toPoint)
}
