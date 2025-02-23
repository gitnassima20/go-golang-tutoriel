package main

import (
	"fmt"
	"testing"
	"time"
)

// 1. Interfaces are collections of method signatures

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// 2. Interface Implementation is implicit in Go

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// 3. Multiple interfaces

func (e email) format() string {
	return e.body
}

type expense interface {
	cost() float64
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

// 4. Type Assertion: take and an interface and cast it to its underlying type

func getExpenseReport(e expense) (string, float64) {
	em, okEmail := e.(email)
	if okEmail {
		return em.toAddress, em.cost()
	}

	sm, okSms := e.(sms)
	if okSms {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func TestGetExpenseReport(t *testing.T) {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}
	tests := []testCase{
		{
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// 5. Type Switches: make it easy to do several type assertions in a series

func getExpenseReport2(e expense) (string, float64) {
	switch expenseNature := e.(type) {
	case email:
		return expenseNature.toAddress, expenseNature.cost()
	case sms:
		return expenseNature.toPhoneNumber, expenseNature.cost()
	default:
		return "", 0.0
	}
}

// TO REMEMBER:
/* Interfaces should have no knoweldge of satisfying types,
   but some types have knoweldge of the interfaces they satisfy */
