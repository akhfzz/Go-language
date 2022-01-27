package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := bufio.NewScanner(os.Stdin)
	fmt.Printf("Mau berapa: ")
	x.Scan()
	input, _ := strconv.ParseInt(x.Text(), 10, 64)
	y := int64(0)
	//* conditional *
	if input < y {
		fmt.Println("x lebih kecil dari y")
	} else if input == y {
		fmt.Println("sama saja")
	} else {
		fmt.Println("x lebih besar dari y")
	}
	// * looping *
	for ; y < input; y++ { //or for y < x
		fmt.Println("Masih belum")
		// y++
	}
}
