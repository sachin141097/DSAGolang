package main

import "fmt"

// Time: O(n + m), Space: O(n)
// DFS visits each node once and explores each edge once (no repeated work due to visited[])

func dfsTraversal(node int, adj [][]int, visited []bool) {
	visited[node] = true
	fmt.Print(node, " ")
	for _, neighbor := range adj[node] {
		if !visited[neighbor] {
			dfsTraversal(neighbor, adj, visited)
		}
	}
	fmt.Println()

}
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	fmt.Println("Enter Start vertex:")
	var start int
	fmt.Scan(&start)
	visited := make([]bool, n+1)
	dfsTraversal(start, adj, visited)

}
