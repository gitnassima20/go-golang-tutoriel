package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costSMSCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costSMSSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costSMSCustomer + costSMSSpouse, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}
