package main

import "fmt"

type Data struct {
	tempat  string
	tanggal int
	bulan   string
	tahun   int
}

type Identitas struct {
	nama string
	ttl  *Data
}

func main() {
	idn := Identitas{"faizal", &Data{"madiun", 23, "maret", 2000}}
	fmt.Println(idn.ttl)
}
