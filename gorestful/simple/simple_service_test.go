package simple

import (
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := InitializeService()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success create service")
	}
	fmt.Println(simpleService)
}
