# go-miner #
Data Mining Algorithms in GoLang.

## Installation
$ go get github.com/ozzie80/go-miner

## Algorithms 

### k-means 

Description From [Wikipedia](https://en.wikipedia.org/wiki/K-means_clustering): 
> k-means clustering aims to partition n observations into k clusters in which each observation belongs to the cluster with the nearest mean, serving as a prototype of the cluster. This results in a partitioning of the data space into Voronoi cells.

The k-means implementation of go-miner uses the [k-means++](https://en.wikipedia.org/wiki/K-means%2B%2B "k-means++") algorithm for choosing the initial cluster centroids. The implementation provides internal quality indexes, [Dunn Index](https://en.wikipedia.org/wiki/Dunn_index) and [Davies-Bouldin Index](https://en.wikipedia.org/wiki/Davies%E2%80%93Bouldin_index), for evaluating clustering results. 


Usage Example

    // Create points or read from a .csv file
    points := []dt.Vector{{1.0, 2.0, 3.0}, {5.1, 6.2, 7.3}, {2.0, 3.5, 5.0}}
    
    // Specify Parameters: K, Points, MaxIter (optional) 
    params := kmeans.Params{2, points, math.MaxInt64}

	// Run k-means
	clusters, err := kmeans.Run(params)

	// Get quality index score
	index := kmeans.DunnIndex{} // DaviesBouldinIndex{}
	score := index.GetScore(clusters)

To Do

- Concurrent version
- Cuda/GPU version


### To Be Added
- Apriori
- *K*NN
- Naive Bayes
- PageRank
- SVM
- ...

## License
go-miner is MIT License.