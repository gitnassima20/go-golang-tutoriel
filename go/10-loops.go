package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (1.0 * float64(i))
	}
	return totalCost
}

// Omitting Conditions from a for Loop
func maxMessages(tresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + (i - 1)
		if totalCost > tresh {
			return i - 1
		}
	}
}

// Writing a while loop in Go's style
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	return maxMessagesToSend
}

// Modulo Operator
func fizzbuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fuzzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
