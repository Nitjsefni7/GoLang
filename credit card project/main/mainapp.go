package main

import (
	"creditcard"
	"log"
	"fmt"
)

func main() {
	log.SetPrefix("credit card project: ")
	log.SetFlags(0)


	ccNum, err := creditcard.TakeCCNumber()
	if err != nil {
		log.Fatal(err)
	}
	if isCorrect := creditcard.LuhnSumAlg(ccNum); isCorrect {
		fmt.Println("The credit card number you gave is valid")
	} else {
		fmt.Println("The credit card number you gave is not valid")
	}
}