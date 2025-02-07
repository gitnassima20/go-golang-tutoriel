package main
import "fmt"

//long assignment declaration
var smsSendingLimit int
var costPerSMS float64
var hasPermission bool
var username string


func printVariables() {
	// Short assignment declaration: most used and only works whitin a func
	totalMessages := 0
	textMsg := "This is a short assignment"
	temperature := 0.0

    //multiple variables declaration:
    averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"

	// Print values
	fmt.Println("Total Messages:", totalMessages)
	fmt.Println("Text Message:", textMsg)
	fmt.Println("Temperature:", temperature)

    fmt.Println(averageOpenRate, displayMessage)
}

