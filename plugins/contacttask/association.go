package main

import (
	common "commom"
	"fmt"
)

type association string

// Assuming that contact-task allows only one-one
func (a association) CanAdd(helper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	policies, _ := helper.HasPolicy(partnerId, "", "", []string{"client-end-user", "hd-client-profile"})
	//lo.Contains(policies, "client-end-user")
	fmt.Printf("\nPolicies checked: %v", policies)

	features, _ := helper.HasEntitlement(partnerId, "", "", []string{"command", "manage"})
	fmt.Printf("\nFeatures checked: %v", features)

	exists, _, _ := helper.OneExists("contacttask", "123456", "1000")
	if !exists {
		fmt.Println("contact-task association can be added.")
		return true
	}
	fmt.Println("contact-task association can not be added.")
	return false
}

func (a association) CanUpdate(helper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	fmt.Println("contact-task association can be updated.")
	return false
}

func (a association) CanDelete(helper common.Helper, partnerId string, mapFromId string, mapToId string) bool {
	fmt.Println("contact-task association can be deleted.")
	return false
}

// exported
var Association association
