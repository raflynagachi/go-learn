package main

import "fmt"

func main() {
	var names = "Rafly Rigan Nagachi"
	// if short statement
	// length only in conditional statement
	if length := len(names); length < 20 {
		fmt.Println("Panjang nama kurang dari 20")
	} else if length < 10 {
		fmt.Println("Panjang nama kurang dari 10")
	} else {
		fmt.Println("Panjang nama terlalu pendek")
	}

	// fmt.Println(length) is error

	// short statement and switch expression
	switch length := len(names); length {
	case 19:
		fmt.Println("Rafly Rigan Nagachi?")
	case 20:
		fmt.Println("Hanya mencoba")
	}

	// short statement and no expression
	switch length := len(names); {
	case length < 20:
		fmt.Println("Nama terlalu panjang")
	case length < 10:
		fmt.Println("Nama cukup panjang")
	default:
		fmt.Println("Nama terlalu pendek")
	}
}
