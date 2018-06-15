// Package adder implements an adder
package adder

import (
	"testing"
)

type inout struct {
	x []int
	y int
}

// To take numbers and their sum and pack them into inout struct.
func makeStruct(x []int, y int) inout {
	v := inout{
		x: x,
		y: y,
	}
	return v
}

// This is known as table test i.e. testing our code against multiple inputs.

// TestAdd tests the Add function in package adder
func TestAdd(t *testing.T) {
	x1 := []int{1, 2, 3}
	y1 := 6

	x2 := []int{1, 2, 3, 4}
	y2 := 10

	x3 := []int{1, 2, 3, 4, 5}
	y3 := 15

	values := []inout{makeStruct(x1, y1),
		makeStruct(x2, y2),
		makeStruct(x3, y3),
	}

	// Even if the test fails, the execution continues.

	for i, v := range values {
		if Add((v.x)...) != v.y {
			t.Error("For test", i+1, "Expected", v.y, "got", Add((v.x)...))
		}
	}
}
