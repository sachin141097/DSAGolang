package main

import "fmt"

func dfsTraversal(node int, adj [][]int, visited []bool) {
	visited[node] = true
	for _, neighbor := range adj[node] {
		if !visited[neighbor] {
			dfsTraversal(neighbor, adj, visited)
		}
	}
}

func findConnectedComponents(V int, edges [][]int) int {
	adj := make([][]int, V)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	visited := make([]bool, V)
	components := 0
	for i := 0; i < V; i++ {
		if !visited[i] {
			components++
			dfsTraversal(i, adj, visited)
		}
	}
	return components
}
func main() {
	V := 4
	edges := [][]int{{0, 1}, {1, 2}}
	components := findConnectedComponents(V, edges)
	fmt.Println(components)
}
