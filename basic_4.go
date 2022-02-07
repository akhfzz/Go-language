package main

import (
	"bufio"
	"fmt"
	"os"
)

func loop(yz int64) {
	x := []int{2, 3, 10, 5, 19}
	for i := 0; i <= len(x); i++ {
		fmt.Println(int64(x[i]) + yz)
	}
}

func test() {
	fmt.Println("hai")
}

func main() {
	data := bufio.NewScanner(os.Stdin)
	fmt.Printf("Angka berapa: ")
	data.Scan()

	// angka, _ := strconv.ParseInt(data.Text(), 10, 64)
	// loop(angka)
	d := test
	d()
}
