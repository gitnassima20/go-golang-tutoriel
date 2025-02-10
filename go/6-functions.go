package main
import "fmt"


/* could also be succinct
  func add(x, y int) int {
    return x + y
  }
*/
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func concatUsage() {
	fmt.Println(concat("Go", "is fantastic!"))
	fmt.Println(concat("The 5th scenario", "has been released to the audience"))
}

//Passing variables by values in go
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}