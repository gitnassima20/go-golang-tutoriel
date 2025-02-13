package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseCost := calculateDiscountRate(numMessages)
	discountRate := calculateBaseBill(costPerMessage, numMessages)
	finalCost := baseCost * (1 - discountRate)
	return finalCost
}

func calculateDiscountRate(messagesSent int) float64 {
	switch {
	case messagesSent > 5000:
		return .2
	case messagesSent > 1000:
		return .1
	default:
		return 0
	}
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
