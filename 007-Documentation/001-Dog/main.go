// Package Dog converts human years into dog years.
package Dog

// Year takes human years as input and converts them into dog years.
func Year(hyears int) int{
	dyears := hyears*7
	return dyears
}
