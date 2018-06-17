package main

import (
	"fmt"
	"github.com/GoesToEleven/sampletemp/011-BET-Examples/003/my-math"
)

func main() {
	xxs := createSlices()
	for _,v := range xxs {
		fmt.Println(my_math.CenAvg(v))
	}
}

func createSlices() [][]int{

	w := []int{1,2,3,4}
	x := []int{2,3,4,5}
	y := []int{1,2,3}
	z := []int{1,2}
    v := []int{1}
	f := [][]int{v,w,x,y,z}
	return f
}
