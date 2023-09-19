package main

import (
	common "commom"
	"fmt"
)

type association string

func (a association) CanAdd(hhelper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	fmt.Println("contact-device association cannot be added.")
	return false
}

func (a association) CanUpdate(helper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	fmt.Println("contact-device association cannot be updated.")
	return false
}

func (a association) CanDelete(helper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	fmt.Println("contact-device association cannot be deleted.")
	return false
}

// exported
var Association association
