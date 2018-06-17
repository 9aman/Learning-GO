package word

import (
	"testing"
	"fmt"
)

func TestCount1(t *testing.T) {
	s := "Name of the game"
	v := Count1(s)

	if v != 4 {
		t.Error("Expected 4 -got ", v)
	}
}

func ExampleCount1() {
	s := "Name of the game"
	fmt.Println(Count1(s))
	// Output:
	// 4
}

func BenchmarkCount1(b *testing.B) {
	s := "Name of the game"
	for i:=0; i <b.N;i++ {
		Count1(s)
	}
}

func TestUseCount1(t *testing.T) {
	s := "Name of the game of"
	m1 := UseCount1(s)
	m2 := make(map[string]int)
	m2["Name"] = 1; m2["of"] = 2; m2["the"] = 1; m2["game"]= 1

	for k, v := range m1 {
		if m2[k] != v {
			t.Error("For",k,"Expected ", m2[k], "got", v)
		}
	}

	for k, v := range m2 {
		if m1[k] != v {
			t.Error("For",k,"Expected ", m2[k], "got", m1[k])
		}
	}
}

func BenchmarkUseCount1(b *testing.B) {
	s := "Name of the game of"
	for i := 0; i < b.N; i++ {
		UseCount1(s)
	}
}