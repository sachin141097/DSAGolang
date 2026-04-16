package main

/*
Space Complexity: O(n^2)
Best for dense graphs
*/
import "fmt"

func main() {
	//n  vertices
	//m edges
	var n, m int
	fmt.Scan(&n, &m)
	//create adjancency matrix
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = make([]int, n+1)
	}
	//Input edges
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u][v] = 1
		adj[v][u] = 1

	}
	//Print matrix
	fmt.Println("Adjancency Matrix:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(adj[i][j], " ")

		}
		fmt.Println()

	}

}
