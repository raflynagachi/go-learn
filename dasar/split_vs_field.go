package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "                  Rafly       Rigan Nagachi \nadalah Rafly atau Nagachi    \t"
	fmt.Println(word)
	var arr []string = strings.Split(word, " ")
	var arr2 []string = strings.Fields(word)
	var arr3 []string = strings.FieldsFunc(word, func(r rune) bool {
		if r == ' ' {
			return true
		} else {
			return false
		}
	})
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
