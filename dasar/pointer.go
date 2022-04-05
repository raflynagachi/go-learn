package main

import "fmt"

type Nums struct {
	num1  []int
	total int
}

func (nums *Nums) calculate() {
	for _, val := range nums.num1 {
		nums.total += val
	}
	fmt.Println("in func: ", nums.total)
}

func multiplyByTen(base *int) {
	*base *= 10
}

func main() {
	num1 := 60
	var num2 *int = &num1
	fmt.Println(num1)
	fmt.Println(*num2)

	num3 := 30
	num1 = 40
	num2 = &num3
	fmt.Println(num1)
	fmt.Println(*num2)

	var num4 = new(int)
	fmt.Println(num4)
	num4 = &num1
	fmt.Println("================")
	multiplyByTen(num2)
	fmt.Println(num1)
	fmt.Println(*num2)
	fmt.Println(num3)
	fmt.Println(*num4)
	fmt.Println("================")

	// pointer method
	var nums Nums
	nums.num1 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums.total)
	nums.calculate()
	fmt.Println(nums.total)
}
