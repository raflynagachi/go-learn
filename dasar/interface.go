package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name, Species string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong
func printer(a ...interface{}) interface{} {
	for _, val := range a {
		if val == 1 {
			return "one"
		} else if val == 1.2 {
			return "float64"
		}
		fmt.Println("check..")
	}
	return -1
}

func main() {
	person := Person{}
	person.Name = "Nagachi"
	fmt.Println(person)

	SayHello(person)

	animal := Animal{}
	animal.Name = "Poppy"
	animal.Species = "Dog"
	SayHello(animal)

	input := []interface{}{4, 5, 6, 6}
	printer(input...)
}
