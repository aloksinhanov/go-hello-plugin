package main

import "fmt"

type association string

func (a association) CanAdd() bool {
	fmt.Println("contact-device association cannot be added.")
	return false
}

// exported
var Association association
