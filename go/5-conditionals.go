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