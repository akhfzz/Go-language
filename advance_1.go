package main

import "fmt"

func test(myFunc func(int) int) {
	fmt.Println(myFunc(6))
}
func main() {
	variabel := func(x int) int {
		return x * -2
	}
	data := func(x string) string {
		return x
	}("siapa")
	test(variabel)
	x := data
	fmt.Printf(x)
}
