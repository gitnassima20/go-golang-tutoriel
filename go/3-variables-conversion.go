package main
import "fmt"

//To stick with(unles you have an explicit reason to care about the size):
// int, uint32, float64, complex128, byte, rune, bool, string

func convertVariable() {
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	
	fmt.Println("Your account has excited for", accountAgeInt, "years")

}
	
