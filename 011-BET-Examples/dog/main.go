// Package dog converts human year into dog year.
package dog

// YearsConv Converts human year into dog year using multiplication
func YearsConv(n int) int {
	return n*7
}

// YearsConv1 converts human year into dog year using addition
func YearsConv1(n int) int {
	count := 0
	for i := 0; i < n;i++{
		count += 7
	}
	return count
}
