package main

import (
	"strings"
	"fmt"
	"github.com/GoesToEleven/sampletemp/009-Benchmark/001-Comparing-Benchmark/001-String-Concat"
)

const s = "The is the and the and a an an no an and! asdasd"

func main() {
	xs := strings.Split(s, " ") // String to be split and delimiter.
	fmt.Println(xs) // return a slice of strings.

	fmt.Println(stringsoncat.BruteForces(xs))
	fmt.Println(stringsoncat.Joins(xs))
}
