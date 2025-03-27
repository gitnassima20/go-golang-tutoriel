package main

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
