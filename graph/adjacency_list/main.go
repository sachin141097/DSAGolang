package main

/*
Space Complexity: O(n+m)
*/
import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	//Create adjacecny list
	adj := make([][]int, n+1)

	//input edges
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		//Undirected graph
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	//Print adjacecny list
	fmt.Println("Adjacency list:")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d -> %v\n", i, adj[i])
	}
}
