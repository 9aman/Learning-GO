package my_math

import (
	"testing"
	"fmt"
)

type inp struct {
	input []int
	output float64
}

func createInputs() []inp{
	w := []int{1,2,3,4}
	x := []int{2,3,4,5}
	y := []int{1,2,3}
	z := []int{1,2}
    v := []int{1}

    f := []inp{
    	{w, 2.5},
    	{x, 3.5},
    	{y, 2},
    	{z, 0},
    	{v, 0},
	}

	return f
}

func TestCenAvg(t *testing.T) {
	inps := createInputs()
	for _,v := range inps {
		if CenAvg(v.input) != v.output {
			t.Error("Expected ",v.output,"got",CenAvg(v.input))
		}
	}
}

func ExampleCenAvg() {
	inps := createInputs()
	for _, v := range inps {
		fmt.Println(CenAvg(v.input))
	}
	//Output:
	// 2.5
	// 3.5
	// 2
	// 0
	// 0
}

func BenchmarkCenAvg(b *testing.B) {
	inps := createInputs()

	for i:=0; i < b.N; i++ {
		for _,v := range inps{
			CenAvg(v.input)
		}
	}
}