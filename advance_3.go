package main

import "fmt"

type dataStruct struct {
	tanggal int
	bulan   string
	tahun   int
}

func changeV(pt dataStruct) {
	pt.tanggal = 20
	fmt.Println(pt)
}

func main() {
	//immutabel
	cd := dataStruct{tanggal: 23, bulan: "maret", tahun: 2000}
	x := cd.bulan
	cd.bulan = "april"
	z := cd.bulan
	if x != z {
		fmt.Println("immutable")
	} else {
		fmt.Println("mutable")
	}

	changeV(cd)
	fmt.Println(cd)

	// fmt.Println(cd)
}
