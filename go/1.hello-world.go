package main

import "fmt"

func main() {
	fmt.Println("Starting Textio Server")
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	printVariables()

	convertVariable()

	formattingTypes()

	conditionalMsgFunc()

	concatUsage()

	getProductMessage("basic")

	yearsUntilEvents(22)

	divide(2, 0)

	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)

	bootUp()
}
