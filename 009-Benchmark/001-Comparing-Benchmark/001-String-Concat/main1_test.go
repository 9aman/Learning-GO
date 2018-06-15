package stringsoncat

import (
	"testing"
	"fmt"
)

func TestJoins(t *testing.T) {
	vs := []string{"Aman", "KC", "is"}
	v := Joins(vs)

	if v != "Aman KC is"{
		t.Error("expectec Aman KC is -got", v)
	}
}

func ExampleJoins() {
	vs := []string{"Aman", "KC", "is"}
	fmt.Println(Joins(vs))
	// Output :
	// Aman KC is
}

func BenchmarkJoins(b *testing.B) {
	vs := []string{"Aman", "KC", "is"}
	for i:=0; i <b.N;i++ {
		Joins(vs)
	}
}
