package main

import (
	"fmt"
)

func main() {
	var score = map[string]int{
		"fisika":  90,
		"kimia":   70,
		"biologi": 75,
	}
	score["MTK"] = 87
	fmt.Println(score)
	fmt.Println(len(score))

	score2 := make(map[string]string)
	score2["physics"] = "90"
	score2["ups"] = "8888"
	fmt.Println(score2)

	delete(score2, "ups")
	fmt.Println(score2)
}
