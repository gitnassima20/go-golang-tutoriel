package main
import "fmt"

func conditionalMsgFunc() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length", messageLen, "and a max length of", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

//Savety hack: if variable was intended to be used only in the if statement
/* if INITIAL_STATEMENT, CONDTION {
     
  }
*/

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}