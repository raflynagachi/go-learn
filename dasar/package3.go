package main

import (
	"fmt"
	"reflect"
	"regexp"
)

type User struct {
	Name string `required:"true" max:"10"`
	Age  int
}

func isValid(user User) bool {
	userType := reflect.TypeOf(user)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(user).Field(i).Interface() != ""
		}
	}
	return true
}

type T struct {
	A int
	B string
}

func main() {
	//////////////////////REFLEC////////////////////////
	user := User{Name: "Nagachi", Age: 21}
	var userType reflect.Type = reflect.TypeOf(user)

	fmt.Println(user)
	fmt.Println(userType)
	fmt.Println("NumField: ", userType.NumField())
	fmt.Println(userType.Field(0).Name)
	fmt.Println(userType.Field(1).Name)
	fmt.Println("====TAG====")
	fmt.Println(userType.Field(0).Tag)
	fmt.Println(userType.Field(0).Tag.Get("required"))
	fmt.Println(userType.Field(0).Tag.Get("max"))
	fmt.Println(userType.Field(0).Tag.Get("min")) // blank
	fmt.Println(userType.Field(1).Tag)            // blank

	fmt.Println("====isValid====")
	fmt.Println(isValid(user))

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
	// output
	// 0: A int = 23
	// 1: B string = skidoo
	// t is now {77 Sunset Strip}

	/////////////////////REGEXP////////////////////////
	var regex *regexp.Regexp = regexp.MustCompile("([a-zA-Z0-9._]+)@([a-z.]+)\\.([a-z]{2,4})")
	fmt.Println(regex.MatchString("rafly.rigannagachi@yahoo.com"))
	fmt.Println(regex.MatchString("rafly.rigannagachi@yahoo.co.id"))
	fmt.Println(regex.MatchString("rafly.yahoo.com"))
	fmt.Println(regex.FindAllString("rafly.rigannagachi@yahoo.com nagachi@gmail.com ragachi yahu", -1))
}
