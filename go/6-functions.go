package main

import (
	"errors"
	"fmt"
)

/*
	could also be succinct

	func add(x, y int) int {
	  return x + y
	}
*/
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func concatUsage() {
	fmt.Println(concat("Go", "is fantastic!"))
	fmt.Println(concat("The 5th scenario", "has been released to the audience"))
}

// Passing variables by values in go
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

// Ignoring return values
func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

//Naming return values: Used for short functions
/*
func getCoords() (x, y int) {
  //x and y are intialized with zero values
  return //automatically returns x and y.
}
*/

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	//naked return
	return
}

// Guard Clauses OR Early returns
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

// Anonymous functions: Used when the func is only being used once
func printReports(intro, body, outro string) {
	// Call printCostReport for each message with the correct cost function
	printCostReport(func(message string) int {
		return 2 * len(message) // Intro cost: 2x length
	}, intro)

	printCostReport(func(message string) int {
		return 3 * len(message) // Body cost: 3x length
	}, body)

	printCostReport(func(message string) int {
		return 4 * len(message) // Outro cost: 4x length
	}, outro)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

//Defer
/*Deferred functions are typically used to clean up resources that are no
longer being used. Often to close database connections,
 file handlers and the like.*/

func connectToDB() bool {
	fmt.Println("Connecting to DB...")
	return false
}

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to Payment Provider...")
	return true
}

func bootUp() {
	defer fmt.Println("Textio BootUp Done!") // Ensures this message is always printed
	ok := connectToDB()
	if !ok {
		return // Early return if DB connection fails
	}
	ok = connectToPaymentProvider()
	if !ok {
		return // Early return if Payment Provider connection fails
	}
	fmt.Println("All systems ready!") // Only prints if both connections succeed
}
