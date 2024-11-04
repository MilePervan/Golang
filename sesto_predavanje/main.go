package main

import (
	"errors"
	"fmt"
)

type Book struct {
	Title    string
	Author   string
	Quantity int
}

func (b *Book) AddBooks(count int) {

	if count > 0 {
		b.Quantity += count
		fmt.Println("Count:", count, "Quantity:", b.Quantity)
	} else {
		fmt.Println("Count to add must be positive.")
	}
}

func (b *Book) RemoveBooks(count int) error {

	if count > b.Quantity {
		return errors.New("insufficient stocks")
	}
	if count > 0 {
		b.Quantity -= count
		fmt.Println("Count:", count, "Quantity:", b.Quantity)
		return nil
	}
	return errors.New("count must be positive")
}

func (b *Book) PrintBookInfo() {

	fmt.Println("Title:", b.Title, "Author:", b.Author,"Quantity:", b.Quantity)
}

func main() {
	book := Book{Title:"Some Book", Author:"Mile", Quantity: 10}

	book.PrintBookInfo()

	book.AddBooks(5)

	err := book.RemoveBooks(12)
	if err != nil {
		fmt.Println("Error:", err)
	}
	book.PrintBookInfo()
}

/*
func main() {

	fmt.Println(greet("CodeCamper"))
}

func greet(name string) string {

	return fmt.Sprintf("Greetings %s", name)
}


func main() {

	names := []string{"Codecamper", "MoreNames", "Third Name"}
	greet(names...)
}

func greet(names ...string) {

	for _, name := range names {
		fmt.Printf("Greetings %s!\n", name)
	}
}


type Person struct {
	Name     string
	Lastname string
}

func main() {

	p := Person{
		Name: "John", 
		Lastname: "Doe",
	}
 
	name, lastname := greet(p)

	fmt.Println(name)
	fmt.Println(lastname)
}

func greet(person Person) (string, string) {

	return person.Name, person.Lastname
}


type Person struct {
	Name     string
	Lastname string
}

func main() {

	p := Person{
		Name: "John", 
		Lastname: "Doe",
	}
 
	name, lastname := greet(p)

	fmt.Println(name)
	fmt.Println(lastname)
}

func greet(person Person) (name string, lastname string) {

	name = person.Name
	lastname = person.Lastname

	return 
}

func calculator(firstNumber float64, secondNumber float64, operator string) float64 {

	switch operator {

	case "+":
		return firstNumber + secondNumber

	case "-":
		return firstNumber - secondNumber

	case "*":
		return firstNumber * secondNumber

	case "/":
		return firstNumber / secondNumber

	default:
		return 0 
	}
}
func main() {

	result := calculator(2, 0, "/")

	fmt.Println(result)
}




type DivisionError struct {
	FirstNumber  int
	SecondNumber int
	Message      string
}

func (de *DivisionError) Error() string {
	return fmt.Sprintf("division error: %d / %d, err: %s", de.FirstNumber, de.SecondNumber, de.Message)
}

func main() {
	r, err := divide(5, 0)
	if err != nil {
		fmt.Printf("error while dividing two numbers: %v", err.Error())
		return
	}
	fmt.Println(r)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{a, b, "cannot divide by zero"}
	}

	return a / b, nil
}
*/