package test

import (
	"fmt"
	"gorestful/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success create service")
	}
	fmt.Println(simpleService)
}
