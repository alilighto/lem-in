package algorithms

import "/utils/"

func EdmondsKarp(graph *utils.Graph) int {
	maxFlow := 0
	for {
		// Use BFS to find an augmenting path
		parents := BFS(graph)
		if _, found := parents[graph.End]; !found {
			break // No more augmenting paths, exit
		}

		// Find bottleneck capacity (1 since edges are unweighted)
		pathFlow := 1

		// Update the capacities in the residual graph
		for v := graph.End; v != graph.Start; v = parents[v] {
			u := parents[v]
			// Remove capacity from u to v
			removeEdge(graph, u, v)
			// Add reverse edge to represent residual capacity
			graph.AddEdge(v, u)
		}

		maxFlow += pathFlow
	}
	return maxFlow
}

func removeEdge(graph *utils.Graph, from, to *utils.Vertex) {
	neighbors := graph.Edges[from.Name]
	for i, v := range neighbors {
		if v == to {
			graph.Edges[from.Name] = append(neighbors[:i], neighbors[i+1:]...)
			break
		}
	}
}
