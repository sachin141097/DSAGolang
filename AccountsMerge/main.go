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
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	ds := NewDisJointSet(n)
	emailToAccount := make(map[string]int)
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if previousAccount, exists := emailToAccount[email]; exists {
				ds.unionByRank(i, previousAccount)
			} else {
				emailToAccount[email] = i
			}
		}
	}
	mergedEmails := make(map[int][]string)
	for email, accountIndex := range emailToAccount {
		root := ds.findUltimateParent(accountIndex)
		mergedEmails[root] = append(mergedEmails[root], email)
	}
	result := [][]string{}
	for root, emails := range mergedEmails {
		sort.Strings(emails)
		mergedAccount := []string{}
		mergedAccount = append(mergedAccount, accounts[root][0])
		mergedAccount = append(mergedAccount, emails...)
		result = append(result, mergedAccount)

	}
	return result
}

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnybravo@mail.com"},
	}
	result := accountsMerge(accounts)
	for _, account := range result {
		fmt.Println(account)
	}
}
