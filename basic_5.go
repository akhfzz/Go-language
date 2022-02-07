package main

import "fmt"

type Data struct {
	nama     string
	prodi    string
	angkatan int
}

func main() {
	dt := Data{nama: "izal", prodi: "informatika", angkatan: 2019}
	fmt.Println(dt)
}
