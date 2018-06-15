package adder

func Add(x ...int) int {
	su := 0
	for _, v := range x {
		su += v
	}
	return su
}
