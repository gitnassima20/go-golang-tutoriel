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

	// Print values
	fmt.Println("SMS Limit:", smsSendingLimit)
	fmt.Println("Cost per SMS:", costPerSMS)
	fmt.Println("Has Permission:", hasPermission)
	fmt.Println("Username:", username)
	fmt.Println("Total Messages:", totalMessages)
	fmt.Println("Text Message:", textMsg)
	fmt.Println("Temperature:", temperature)
}

