package main

import "fmt"

func main() {
	fmt.Println("Starting Textio Server")
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	printVariables()

	convertVariable()

	formattingTypes()

	conditionalMsgFunc()

	/* functions sections */
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

	/* structs sections */

	output := canSendMessage(messageToSend{
		sender: user{
			name: "Textio",
		},
		recipient: user{
			name: "You",
		},
		message: "Welcome to Textio",
	})
	fmt.Println(output)

	auth := authenticationInfo{
		username: "myUser",
		password: "myPass",
	}

	fmt.Println(auth.getBasicAuth())

	/* loops */
	fmt.Println(bulkSend(10))
}
