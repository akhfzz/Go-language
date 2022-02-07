package main

import "fmt"

func changeValue(str *string) {
	*str = "kamu"
}
func changeValue2(str string) string {
	str = "berubahhhh"
	return str
}

func main() {
	b := 7
	c := &b
	fmt.Println(b, c)
	*c = 8
	fmt.Println(b, c)

	txt := "aku"
	changeValue(&txt)
	fmt.Println(txt)
	changeValue2(txt)
	fmt.Println(txt)
}
