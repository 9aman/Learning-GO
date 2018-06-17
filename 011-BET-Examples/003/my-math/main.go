package my_math

import "sort"

// Returns the average after removing the boundary elements.
func CenAvg(xi []int) float64 {
	if len(xi) < 3 {
		return 0
	}
	sort.Ints(xi)

	// Remove the boundaries from the slice
	xi = xi[1:(len(xi)-1)]

	var su float64

	for _,v := range xi{
		su = su + float64(v)
	}


	su = su / float64(len(xi))
	return su
}