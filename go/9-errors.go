package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	msgCus, msgSpo, err := sendSMS(msgToCustomer), sendSMS(msgToSpouse)
	if err != nil {
		fmt.Println(err)
	}
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
