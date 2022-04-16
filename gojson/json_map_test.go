package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJson(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Pasta Gigi",
		"price": 4500,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestMapJsonDecode(t *testing.T) {
	jsonString := `{"Firstname":"Rafly","Middlename":"Rigan","Lastname":"Nagachi","Age":21,"Hobbies":["watch, code, and coffee"],"Addresses":[{"Street":"Jl. Kemerdekaan","Country":"Indonesia","PostalCode":"32381"},{"Street":"Jl. Kebangsaan","Country":"Indonesia","PostalCode":"32431"}]}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}
