package main

import "fmt"

/*Prvi zadatak

func main(){

	var firstNumber int 
	var secondNumber int
	firstNumber = 20
	secondNumber = 30
	firstNumber, secondNumber = secondNumber, firstNumber
	

	fmt.Println(firstNumber,secondNumber)
}

Drugi zadatak */

func main(){

	var firstName = "Mile"
	var lastName = "Pervan"
	var fullName string

	fullName = firstName + " " + lastName

	fmt.Println(fullName)
}