// Package word is used to returns word count for the sentence as well as for that particular word.
package word

import "strings"

//UseCount1 returns a map which contains how many times a word occurs in a string.
func UseCount1(s string) map[string]int{
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count1 returns the word count.
func Count1(s string) int {
	return len(strings.Fields(s))
}