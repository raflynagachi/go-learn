package main

import "fmt"

type User interface {
	GetRole() string
	SetRole(name string)
}

func sayHi(user User) {
	fmt.Println("Hi!!, " + user.GetRole())
}

type Admin struct {
	Name, Role string
}

func (admin Admin) GetRole() string {
	return admin.Role
}

func (admin Admin) SetRole(x string) {
	admin.Role = x
}

func random() interface{} {
	return 10
}

func main() {
	var admin Admin
	admin.SetRole("Admin")
	fmt.Println(admin)
	sayHi(admin)

	input := random()

	fmt.Printf("%T \n", input)
}
