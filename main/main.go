package main

import (
	"/src/"
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
	// Initialize graphs
	mainGraph, _ := utils.SetGraphs(Colonie)

	// Run BFS to find the shortest path
	parents := algorithms.BFS(mainGraph)

	// Reconstruct and display the shortest path
	path := algorithms.GetPath(parents, mainGraph.End)
	fmt.Println("Shortest path:", path)

	// (Optional) Run Edmonds-Karp for maximum flow optimization
	// maxFlow := algorithms.EdmondsKarp(mainGraph)
	// fmt.Println("Maximum flow:", maxFlow)

}
