package main

import "fmt"

type Mahasiswa struct {
	nama     string
	semester []int
	umur     int
}

func (s *Mahasiswa) getUmur(age int) int {
	s.umur = age
	return s.umur
}

func (s Mahasiswa) averageSemester() int {
	// firstV := 0
	maxim := 0
	for _, v := range s.semester {
		if maxim < v {
			maxim = v
		}
		// firstV += v
	}
	return maxim
}

func main() {
	age := Mahasiswa{"faizal", []int{1, 2, 3, 4, 5, 6}, 22}
	fmt.Println(age)
	fmt.Println(age.getUmur(20))
	fmt.Println(age)
	average := age.averageSemester()
	fmt.Println(average)
}
