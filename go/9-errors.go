package main

import (
	"errors"
	"fmt"
)

// Handling errors

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

// Formatting Error messages

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %2.f to be sent to %v can not be sent", cost, recipient)
}

// Creating custom error
type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can't divide %v by zero", de.dividend)
}

func dividing(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func divideUsingErrorPackage(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("Cannot divide %v by zero", x)
	}
	return x / y, nil
}
