package main

import (
	"fmt"
)

func main() {
	// old
	var data [3]int
	data[0] = 10
	data[1] = 20
	data[2] = 30
	fmt.Println(data)

	// new
	// arr := [4]string{"cinta", "kamu", "love", "you"}
	twoD_arr := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(twoD_arr)

	//dict
	dictio := map[string]string{
		"nama":     "izal",
		"prodi":    "informatika",
		"angkatan": "2019",
	}
	fmt.Println(dictio)

	//slice
	x := []int{10, 18, 21, 23, 30}
	z := x[1:3]
	fmt.Println(cap(x))
	fmt.Println(z)
	b := append(x, 15)
	fmt.Println(b)
	a := make([]int, 5)
	fmt.Println(a)
}
