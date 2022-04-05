package helpers

import "fmt"

var name string

// init called while package initialized
func init(){
	name = "Rafly Rigan Nagachi"
}

func GetName() string{
	return name
}

func SayHello() {
	fmt.Println("Halo!!")
}
