package main

import "fmt"

// immutable and mutable
func main() {
	// immutable
	x := 6
	y := x
	y = 7
	fmt.Println(x, y)

	// mutable
	i := []int{5, 7, 8, 9}
	t := i
	t[0] = 20
	fmt.Println(i, t)
	dict := map[string]int{
		"satu": 1,
		"dua":  2,
		"tiga": 3,
	}
	z := dict
	z["satu"] = 4
	dict["satu"] = 5
	fmt.Println(dict, z)
}
