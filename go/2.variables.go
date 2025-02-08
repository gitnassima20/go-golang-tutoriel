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

    //constants declaration: doesn't support short declaration syntax
    const premiumPlanName = "Premmium Plan"

	//computed constans
	const secondInMinute = 60
	const minutesInHour = 60
	const secondInHour = secondInMinute * minutesInHour

	// Print values
	fmt.Println("Total Messages:", totalMessages)
	fmt.Println("Text Message:", textMsg)
	fmt.Println("Temperature:", temperature)

    fmt.Println(averageOpenRate, displayMessage)

    fmt.Println("plan:", premiumPlanName)

	fmt.Println("number of seconds in an hour:", secondInHour)
}

