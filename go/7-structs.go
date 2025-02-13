package main

//Structs: Represent structured data
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

//Anonymous structs: Used when the struct is only being used once
/*myCar := struct {
	Make string
	Model string
} {
	Make: "Toyota",
	Model: "Corolla",
}*/

//Embedded structs: When a struct contains another struct as a field
