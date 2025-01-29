package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
func groupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range words {
		sortedKey := sortString(word)
		groups[sortedKey] = append(groups[sortedKey], word)
	}
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
func main() {
	fmt.Println("Enter list of space separated strings")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)
	result := groupAnagrams(words)
	fmt.Println(result)

}
