package main

import "fmt"

/*
There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. Any computer can reach any other computer directly or indirectly through the network.

You are given an initial computer network connections. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected.

Return the minimum number of times you need to do this in order to make all the computers connected. If it is not possible, return -1.

Input: n = 4, connections = [[0,1],[0,2],[1,2]]
Output: 1
Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
*/

/*
Time Complexity: O(N+M)
where N->Number of vertices
M-> Number of connections

Space Complexity: O(N)
*/
func makeConnected(n int, connections [][]int) int {
	//Not enough cables to connect all computers
	if len(connections) < n-1 {
		return -1
	}
	ds := NewDisJointSet(n)
	for _, connection := range connections {
		ds.unionByRank(connection[0], connection[1])
	}
	components := 0
	for i := 0; i < n; i++ {
		if ds.findUltimateParent(i) == i {
			components++
		}
	}
	return components - 1

}

type DisJointSet struct {
	parent []int
	rank   []int
}

func NewDisJointSet(n int) *DisJointSet {
	ds := &DisJointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
	return ds
}
func (ds *DisJointSet) findUltimateParent(node int) int {
	if node == ds.parent[node] {
		return node
	}
	ds.parent[node] = ds.findUltimateParent(ds.parent[node])
	return ds.parent[node]
}
func (ds *DisJointSet) unionByRank(u, v int) {
	rootU := ds.findUltimateParent(u)
	rootV := ds.findUltimateParent(v)
	if rootU == rootV {
		return
	}
	if ds.rank[rootU] < ds.rank[rootV] {
		ds.parent[rootU] = rootV
	} else if ds.rank[rootV] < ds.rank[rootU] {
		ds.parent[rootV] = rootU
	} else {
		ds.parent[rootV] = rootU
		ds.rank[rootU]++
	}
}
func main() {
	n := 4
	connections := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}
	fmt.Println(makeConnected(n, connections))
}
