package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var s1, s2 [26]int
	for i := 0; i < len(s); i++ {
		s1[s[i]-'a']++
		s2[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func main() {
	//take two strings as input from user
	var s, t string
	fmt.Println("Enter the first string")
	fmt.Scan(&s)
	fmt.Println("Enter the second string")
	fmt.Scan(&t)
	ans := isAnagram(s, t)
	if ans {
		fmt.Println("The strings are anagrams")
	} else {
		fmt.Println("The strings are not anagrams")
	}
}
