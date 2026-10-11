package main

import (
	"fmt"
	"sort"
)

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
		ds.rank[i] = 0
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
func (ds *DisJointSet) find(u, v int) bool {
	return ds.findUltimateParent(u) == ds.findUltimateParent(v)
}
func (ds *DisJointSet) unionByRank(u, v int) {
	ultimateParentU := ds.findUltimateParent(u)
	ultimateParentV := ds.findUltimateParent(v)
	if ultimateParentU == ultimateParentV {
		return
	}
	if ds.rank[ultimateParentU] < ds.rank[ultimateParentV] {
		ds.parent[ultimateParentU] = ultimateParentV
	} else if ds.rank[ultimateParentV] < ds.rank[ultimateParentU] {
		ds.parent[ultimateParentV] = ultimateParentU
	} else {
		ds.parent[ultimateParentV] = ultimateParentU
		ds.rank[ultimateParentU]++
	}
}

type Edge struct {
	src, dst, weight int
}

func kruskal(n int, edges []Edge) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	ds := NewDisJointSet(n)
	minimumSpanningTreeWeight := 0
	edgesUsed := 0
	for _, edge := range edges {
		if !ds.find(edge.src, edge.dst) {
			minimumSpanningTreeWeight += edge.weight
			ds.unionByRank(edge.src, edge.dst)
			edgesUsed++
			if edgesUsed == n-1 {
				break
			}
		}
	}
	if edgesUsed != n-1 {
		fmt.Println("MST doesn't exist graph is disconnected")
	} else {
		fmt.Println("Total MST weight:", minimumSpanningTreeWeight)
	}
}
func main() {
	n := 4
	edges := []Edge{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}
	kruskal(n, edges)
}
