package main

import "fmt"

func logging() {
	message := recover() //recover function
	if message != nil {
		fmt.Println("Pesan panic:", message)
	}
	fmt.Println("program done")
}

func calculate(num int) {
	defer logging() //defer function
	if num == 0 {
		panic("0 tidak bisa menjadi pembagi") //panic function
	}
	result := 100 / num
	fmt.Println(result)
}

func main() {
	calc := calculate
	calc(0)
}
