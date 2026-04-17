package main

import "fmt"

/*
Given an undirected graph with V vertices. Two vertices u and v belong to a single province if there is a path from u to v or v to u. Find the number of provinces. The graph is given as an n x n matrix adj where adj[i][j] = 1 if the ith city and the jth city are directly connected, and adj[i][j] = 0 otherwise.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

Input: adj=[ [1, 0, 0, 1], [0, 1, 1, 0], [0, 1, 1, 0], [1, 0, 0, 1] ]

Output: 2

Explanation:In this graph, there are two provinces: [1, 4] and [2, 3]. City 1 and city 4 have a path between them, and city 2 and city 3 also have a path between them. There is no path between any city in province 1 and any city in province 2.
*/
func dfsTraversal(node int, adj [][]int, visited []bool) {
	visited[node] = true
	for _, neighbor := range adj[node] {
		if !visited[neighbor] {
			dfsTraversal(neighbor, adj, visited)
		}
	}
}
func countProvinces(adj [][]int) int {
	/* Adjacency matrix
		1 2 3 4
	1	1 0 0 1
	2	0 1 1 0
	3	0 1 1 0
	4	1 0 0 1

	Adjacency List
	1->[4]
	2->[3]
	3->[2]
	4->[1]
	*/
	adjList := make([][]int, len(adj))
	for i := range adj {
		for j := range adj[i] {
			if i != j && adj[i][j] == 1 {
				adjList[i] = append(adjList[i], j)
			}

		}
	}
	visited := make([]bool, len(adj))
	var provinces int
	for i := 0; i < len(adj); i++ {
		if !visited[i] {
			provinces++
			dfsTraversal(i, adjList, visited)
		}
	}
	return provinces
}
func main() {
	adj := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}}
	fmt.Println(countProvinces(adj))
}
