/*
package main

import "fmt"

type Book struct {
	title     string
	author    string
	publisher Publisher
	price     float64
}

type Publisher struct {
	name           string
	location       string
	booksPublished int
}

func (b Book) printDetails() {
	fmt.Println(b.title, b.author, b.publisher.name, b.publisher.location, b.price)
}

func (b *Book) applyDiscount(discountPercentage float64) {
	discount := b.price * (discountPercentage / 100)
	b.price = b.price - discount
}

func (b *Book) updatePublisher(newPublisher string) {
	b.publisher.name = newPublisher
}

func (p Publisher) printInfo() {
	fmt.Println(p.name, p.location, p.booksPublished)
}

func (p *Publisher) publishBook() {
	p.booksPublished = p.booksPublished + 1
}

func (p *Publisher) changeLocation(newLocation string) {
	p.location = newLocation
}

func main() {
	book := Book{
		title:  "Echoes of a Distant Star",
		author: "Marcus Davenport",
		price:  20,
		publisher: Publisher{
			name:           "Horizon Books",
			location:       "New York",
			booksPublished: 10,
		},
	}

	book.printDetails()
	fmt.Println()

	book.publisher.printInfo()


	book.applyDiscount(50)
	fmt.Println(book.price)

	book.updatePublisher("New Publisher")
	fmt.Println("")
	book.printDetails()

	book.publisher.publishBook()
	book.publisher.changeLocation("Mostar")
	fmt.Println("")
	book.publisher.printInfo()
}



package main

import "fmt"

func main(){

	age := 12

    if age < 12 {

		fmt.Println("Child")

	} else if age >= 12 && age < 18 {

		fmt.Println("Teenager")

	} else if age >= 18 && age < 60 {

		fmt.Println("Adult")

	} else {

		fmt.Println("Senior")
	}

}
*/
package main

import ("fmt"
        "math/rand")

func main(){

	a := rand.Intn(4)
	switch {

	case a == 1 :
		fmt.Println("Broj je 1")
	
	case a == 2 :
		fmt.Println("Broj je 2")
	
	case a == 3 :
		fmt.Println("Broj je 3")

	default:
		fmt.Println("Broj je 4")

	}

}
