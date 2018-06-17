package main

import (
	"fmt"
	word2 "github.com/GoesToEleven/sampletemp/011-BET-Examples/002/word"
	"github.com/GoesToEleven/sampletemp/011-BET-Examples/002/quotes"
)

func main() {
	fmt.Println(word2.Count1(quotes.SunAlso))

	m := word2.UseCount1(quotes.SunAlso)

	for k,v := range m{
		fmt.Println(v,k)
	}
}
