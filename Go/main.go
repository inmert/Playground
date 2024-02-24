package main

import (
	"fmt"
)

func main() {
	array := [...]int{2, 4, 5}
	slice := []int{1, 2, 3}
	slice = append(slice, array[:]...)
	fmt.Printf("Array: %v\nSlice: %v\n", array, slice)
	dict := map[int]string{
		1: "A",
		2: "B",
		4: "4",
	}
	delete(dict, 1)
	fmt.Printf("Dict: %v\n", dict)
	sampleSlice := []Sample{}

	for e := 0; e < 10; e++ {

		a := Sample{}
		a.name = "Name"
		a.value = e
		sampleSlice = append(sampleSlice, a)

	}
	fmt.Printf("Sample: %v\n", sampleSlice)
}

type Sample struct {
	name  string
	value int
}

func add(x, y int) int {
	return x + y
}
