package main

import "fmt"

/**
  %v - Interpolate the default representation
  %s - Interpolate a string
  %d - Interpolate an integer in decimal form
  %f - Interpolate a decimal
**/
func formattingTypes() {
	const name = "Kim Dokja"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %v, your open rate is %.1f", name, openRate)

	fmt.Println(msg)
}
