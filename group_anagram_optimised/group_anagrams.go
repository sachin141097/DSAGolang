package main

/*
Time Complexity: O(N*klogk)
*/
import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		key := sortedKey(s)
		groups[key] = append(groups[key], s)
	}
	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
func sortedKey(s string) string {
	b := []rune(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))
}
