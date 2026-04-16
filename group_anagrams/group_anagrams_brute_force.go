/*
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:

Input: strs = ["x"]

Output: [["x"]]
Example 3:

Input: strs = [""]

Output: [[""]]

Time: O(n² × k) where k is string length
Space: O(n)
*/
package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	visited := make([]bool, len(strs))
	for i := 0; i < len(strs); i++ {
		if visited[i] {
			continue
		}
		group := []string{strs[i]}
		for j := i + 1; j < len(strs); j++ {
			if !visited[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				visited[j] = true
			}
		}
		result = append(result, group)
		visited[i] = true
	}
	return result
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}
func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))
}
