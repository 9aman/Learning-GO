package sayings

import (
	"testing"
	"fmt"
)

func TestGreet(t *testing.T) {
	s := Greet("Aman")

	if s != "Greet Aman" {
		t.Error("Expected Greet Aman - got ", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Aman"))
	// Output: Greet Aman
}

func BenchmarkGreet(b *testing.B) {
	for i:=0; i< b.N;i++{
		Greet("Aman")
	}
}
// Use go test -bench .  to run all the benchmark codes. Here . signifies that all benchmarks should be run.