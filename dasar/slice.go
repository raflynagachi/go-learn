package main

import "fmt"

func main() {
	iniArray := [3]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	// slice & array method
	var lengthSize = len(iniSlice)
	var capacitySize = cap(iniArray)

	fmt.Println(iniSlice)
	fmt.Println("length:", lengthSize)
	fmt.Println("capacity:", capacitySize)

	// slice only method
	// create slice with length(2) and capacity(5)
	var names = make([]string, 2, 5)

	// append from last element in slice
	names = append(names, "Patrick")

	// copy slice
	var names2 = make([]string, len(names), cap(names))
	copy(names2, names)

	fmt.Println(names)      // [  Patrick]
	fmt.Println(len(names)) // 3

	fmt.Println(names2)      // [  Patrick]
	fmt.Println(len(names2)) // 3
}
