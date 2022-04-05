package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c *Customer) setName(name string) {
	c.Name = name
}

func main() {
	customer := Customer{
		Name: "Rafly",
	}
	// customer.Name = "Nagachi"
	fmt.Println(customer.Name)
	customer.setName("Nagachi")
	fmt.Println(customer.Name)
}
