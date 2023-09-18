package main

import "fmt"

type association string

func (a association) CanAdd() bool {
	fmt.Println("contact-task association can be added.")
	return true
}

// exported
var Association association
