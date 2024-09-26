package algorithms

import (
	"/utils/"
	"container/list"
)

func BFS(graph *utils.Graph) map[*utils.Vertex]*utils.Vertex {
	queue := list.New()
	queue.PushBack(graph.Start)

	// Stores parent of each vertex to reconstruct the path
	parents := make(map[*utils.Vertex]*utils.Vertex)
	visited := make(map[*utils.Vertex]bool)
	visited[graph.Start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(*utils.Vertex)

		if current == graph.End {
			break // We reached the end, exit BFS
		}

		for _, neighbor := range graph.Edges[current.Name] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parents[neighbor] = current
				queue.PushBack(neighbor)
			}
		}
	}
	return parents
}

// Reconstruct the path from the BFS result
func GetPath(parents map[*utils.Vertex]*utils.Vertex, end *utils.Vertex) []string {
	var path []string
	for v := end; v != nil; v = parents[v] {
		path = append([]string{v.Name}, path...)
	}
	return path
}
