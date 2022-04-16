package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestLogJson(t *testing.T) {
	logJson("Nagachi")
	logJson(1)
	logJson(true)
	logJson([]string{"Rafly", "Rigan", "Nagachi"})
}

type Customer struct {
	// json only accept capitalize
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		Firstname:  "Rafly",
		Middlename: "Rigan",
		Lastname:   "Nagachi",
		Age:        21,
		Hobbies:    []string{"watch, code, and coffee"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJson(t *testing.T) {
	jsonString := `{"Firstname":"Rafly","Middlename":"Rigan","Lastname":"Nagachi","Age":21,"Hobbies":["watch, code, and coffee"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname:  "Rafly",
		Middlename: "Rigan",
		Lastname:   "Nagachi",
		Age:        21,
		Hobbies:    []string{"watch, code, and coffee"},
		Addresses: []Address{
			{
				Street:     "Jl. Kemerdekaan",
				Country:    "Indonesia",
				PostalCode: "32381",
			},
			{
				Street:     "Jl. Kebangsaan",
				Country:    "Indonesia",
				PostalCode: "32431",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"Rafly","Middlename":"Rigan","Lastname":"Nagachi","Age":21,"Hobbies":["watch, code, and coffee"],"Addresses":[{"Street":"Jl. Kemerdekaan","Country":"Indonesia","PostalCode":"32381"},{"Street":"Jl. Kebangsaan","Country":"Indonesia","PostalCode":"32431"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestOnlyJsonArray(t *testing.T) {
	jsonString := `[{"Street":"Jl. Kemerdekaan","Country":"Indonesia","PostalCode":"32381"},{"Street":"Jl. Kebangsaan","Country":"Indonesia","PostalCode":"32431"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)

	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}
