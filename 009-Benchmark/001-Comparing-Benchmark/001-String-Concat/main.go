// Package stringsoncat performs concatenation using different methods.
package stringsoncat

import "strings"

//BruteForce returns the concatenated string using brute force
func BruteForces(xs []string) string{
	s := ""
	for _,v := range xs{
		s = s + v + " "
	}
	return s
}

// Joins returns the concatenated string using strings.Join(slice, limiter)
func Joins(xs []string) string {
	return strings.Join(xs," ")
}