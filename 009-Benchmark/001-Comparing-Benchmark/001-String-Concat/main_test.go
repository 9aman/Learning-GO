package stringsoncat

import (
	"testing"
	"fmt"
)

func TestBruteForces(t *testing.T) {
	vs := []string{"Aman", "KC", "is"}
	v := BruteForces(vs)

	if v != "Aman KC is "{
		t.Error("expectec Aman KC is -got", v)
	}
}

func ExampleBruteForces() {
	vs := []string{"Aman", "KC", "is"}
	fmt.Println(BruteForces(vs))
	// Output :
	// Aman KC is
}

func BenchmarkBruteForces(b *testing.B) {
	vs := []string{"Aman", "KC", "is"}
	for i := 0;i< b.N;i++{
		BruteForces(vs)
	}
}
