package main

// Time: O(n + m), Space: O(n)
// Each node is visited once and each edge is explored once (no repeated work due to visited[])
import "fmt"

func bfsTraversal(start int, adj [][]int, n int) {
	visited := make([]bool, n+1)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}

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
	bfsTraversal(start, adj, n)

}
