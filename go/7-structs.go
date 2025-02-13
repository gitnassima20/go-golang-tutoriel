package main

import "fmt"

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

/* type sender struct {
	rateLimit int
	userEmbeded
}

type userEmbeded struct {
	name   string
	number int
}
*/

// Struct Methods in Go
type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		auth.username, auth.password,
	)
}
