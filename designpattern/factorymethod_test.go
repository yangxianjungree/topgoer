package main

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var p Provider = NewEnvironmentProvider()
	fmt.Println("name: ", p.Name())

	var s SecretReader = NewEnvironmentProvider()
	res, b := s.GetSecret("name")
	if !b {
		fmt.Println("get secret failed.")
	} else {
		fmt.Println("get secret: ", res)
	}
}
