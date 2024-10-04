package main

import (
	"Lemin/src"
	"sort"

	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: invalid number of arguments: please enter 'go run ./main/ filename.txt'")
		return
	}
	Colonie, err := src.HandleFormat(args[0])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	pathsWithDfs := Colonie.DFS([]string{}, Colonie.StartRoom, map[string]bool{}, [][]string{})
	// uniquePaths := src.GetShortestPaths(pathsWithDfs)
	fmt.Println(pathsWithDfs)
	sort.Slice(pathsWithDfs, func(i, j int) bool {
		return len(pathsWithDfs[i]) < len(pathsWithDfs[j])
	})
	uniquePaths2 := src.GetShortestPaths(pathsWithDfs)
	// if len(uniquePaths) > len(uniquePaths2) {
	// 	uniquePaths2 = uniquePaths
	// 	sort.Slice(uniquePaths2, func(i, j int) bool {
	// 		return len(uniquePaths2[i]) < len(uniquePaths2[j])
	// 	})
	// }
	src.PrintResult(uniquePaths2, Colonie.Ants)
	// Initialize graphs
	// Run BFS to find the shortest path
	// parents := algorithms.BFS(mainGraph)

	// Reconstruct and display the shortest path
	// path := algorithms.GetPath(parents, mainGraph.End)
	// fmt.Println("Shortest path:", path)

	// maybe Run Edmonds-Karp for maximum flow optimization
	// maxFlow := algorithms.EdmondsKarp(mainGraph)
	// fmt.Println("Maximum flow:", maxFlow)

}
