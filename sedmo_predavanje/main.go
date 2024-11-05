package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) deposit(amount float64) {

	if amount > 0 {
		b.Balance += amount
		fmt.Println("Owner:", b.Owner, "Amount:", amount, "Balance:", b.Balance)
	} else {
		fmt.Println("amount can't be negative or zero")
	}
}

func (b *BankAccount) withdraw(amount float64) error {

	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	if amount > 0 {
		b.Balance -= amount
		fmt.Println("Owner:", b.Owner, "Amount:", amount, "Balance:", b.Balance)
		return nil
	}
	return errors.New("withdraw amount must be positive")
}

func main() {

	account := BankAccount{Owner: "Mile", Balance: 100}

	account.deposit(100)

	err := account.withdraw(300)
	if err != nil {
		fmt.Println(err)
	}

	err = account.withdraw(10)
	if err == nil {
		fmt.Println("Withdrawal successful.")
	}
}

/*

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
		fmt.Println("Quantity:", b.Quantity)
	} else {
		fmt.Println("count must be positive")
	}
}

func (b *Book) RemoveBooks(count int) error {

	if count > b.Quantity {
		return errors.New("insufficient stocks")
	}
	if count > 0 {
		b.Quantity -= count
		fmt.Println("Quantity:", b.Quantity)
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

	err := book.RemoveBooks(2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	book.PrintBookInfo()
}
*/