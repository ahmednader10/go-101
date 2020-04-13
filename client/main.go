package main

import (
	"awesomeProject2/payment"
	"fmt"
)

type MathExpr = string

const (
	AddExpr = MathExpr("add")
	SubtExpr = MathExpr("subtract")
	MulExpr = MathExpr("multiply")
)

func main() {
	credit := payment.CreateCreditAccount("ahmed", "12345", 12, 2020, 124)
	fmt.Printf("Owner Name: %v\n", credit.OwnerName())
	fmt.Printf("Card Number: %v\n", credit.CardNumber())
	credit.ExposedProcessPayment(500)

	err := credit.SetSecurityCode(1)
	if err != nil {
		fmt.Printf("error while setting code %v\n", err)
	}

	fmt.Printf("Available Credit: %v\n", credit.AvailableCredit())

	fmt.Println("***************************************")

	chargeChan := make(chan float32)
	payment.CreateCreditAccountWithChannel(chargeChan, "ahmed", "12345", 12, 2020, 124)

	chargeChan <- 500

	//To make sure main method does not close before the message is passed to the channel
	//var a string
	//fmt.Scanln(&a)

	fmt.Println("***************************************")

	ca := payment.CompositionAccount{}
	ca.ProcessPayment(30)

	fmt.Println("***************************************")

	addExpr := mathExpression(AddExpr)
	fmt.Println(addExpr(2,3))

	fmt.Printf("%f", double(3, 2, mathExpression(AddExpr)))
}

func mathExpression(op MathExpr) func(float64, float64) float64 {
	switch op {
	case AddExpr: return func(f float64, f2 float64) float64 {
		return f + f2
	}
	case SubtExpr: return func(f float64, f2 float64) float64 {
		return f - f2
	}
	case MulExpr: return func(f float64, f2 float64) float64 {
		return f * f2
	}
	default:
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}
}

func double (f1, f2 float64, mathExpr func(float64, float64) float64) float64{
	return 2 * mathExpr(f1, f2)
}
