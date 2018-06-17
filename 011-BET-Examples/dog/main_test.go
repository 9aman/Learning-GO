package dog

import (
	"testing"
	"fmt"
)

func TestYearsConv(t *testing.T) {
	v := YearsConv(10)
	if v != 70 {
		t.Error("Expected 70 -got ", v)
	}
}

func ExampleYearsConv() {
	fmt.Println(YearsConv(10))
	// Output:
	// 70
}

func BenchmarkYearsConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsConv(10)
	}
}

func TestYearsConv1(t *testing.T) {
	v := YearsConv1(10)
	if v != 70 {
		t.Error("Expected 70 -got ", v)
	}
}

func ExampleYearsConv1() {
	fmt.Println(YearsConv1(10))
	// Output:
	// 70
}

func BenchmarkYearsConv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsConv1(10)
	}
}
