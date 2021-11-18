package lib

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinAdj(adj [][]int) (Min int) {
	Min = 0
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(adj[i]); j++ {
			if Min > adj[i][j] {
				Min = adj[i][j]
			}
		}
	}
	return Min
}

func MinDijkstra(start, finish int, adj [][]int) []int {
	var path []int = make([]int, len(adj))

	// path initialize
	for i := 0; i < len(path); i++ {
		path[i] = start
	}

	// visits
	S := make(map[int]struct{})
	// Minimal vertex distance
	D := make([]int, len(adj))

	Min := getMinAdj(adj)

	// current vertex
	var currentVertex int = start
	for {
		S[currentVertex] = struct{}{}
		for j := 0; j < len(adj[currentVertex]); j++ {
			if adj[currentVertex][j] > Min {
				// can go to vertex
				if D[j] == 0 {
					D[j] = D[currentVertex] + adj[currentVertex][j]
					path[j] = currentVertex
				} else {
					if D[j] > D[currentVertex]+adj[currentVertex][j] {
						path[j] = currentVertex
					}
					D[j] = min(D[j], D[currentVertex]+adj[currentVertex][j])
				}
			}
		}

		previousCurrentVertex := currentVertex
		// Find the next vertex index
		var nextVertexValue int = 1 << 31

		for ind, v := range D {
			_, exist := S[ind]
			if exist || v == 0 {
				continue
			}
			if nextVertexValue > v {
				nextVertexValue = v
				currentVertex = ind
			}
		}
		if previousCurrentVertex == currentVertex {
			break
		}
	}

	var pathStack []int = make([]int, 0)
	currentPathIndex := finish
	pathStack = append(pathStack, finish)
	for {
		if currentPathIndex != start {
			pathStack = append(pathStack, path[currentPathIndex])
			currentPathIndex = path[currentPathIndex]
		} else {
			break
		}
	}
	for i := 0; i < len(pathStack)/2; i++ {
		pathStack[i], pathStack[len(pathStack)-i-1] = pathStack[len(pathStack)-i-1], pathStack[i]
	}
	return pathStack
}
