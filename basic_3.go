package main

import "fmt"

func main() {
	data := []int{2, 4, 5, 6, 8, 10}
	for i := 0; i < len(data); i++ {
		fmt.Printf("Angka tersebut %d \n", data[i])
	}
	for i, element := range data {
		fmt.Println("index: ", i, "element: ", element)
	}

}
