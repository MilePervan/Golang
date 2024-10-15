package main

import "fmt"
/*
type Person struct {
    firstName string
    lastName  string
    age       int
}

func main() {
    
    person1 := Person{"Mile", "Pervan", 25}

    person2 := Person{
        firstName: "Mile",
        lastName:  "Pervan",
        age:       25,
    }
    fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)
	fmt.Println(person1.age)
	fmt.Println(person2.firstName)
	fmt.Println(person2.lastName)
	fmt.Println(person2.age)

}
*/
type Book struct{
	title     string
	author    string
	year      int
	genre     string
	publisher Publisher
}

type Publisher struct{
	name    string
	address string
}

func main() {

	book := Book{
		title: "Echoes of a Distant Star",
		author: "Marcus Davenport",
		genre: "Science Fiction",
		year: 2020,
		publisher : Publisher{
			name: "Horizon Books",
			address: "115 S. Mitchell Street",
		},
	}

	fmt.Println(book)
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.year)
	fmt.Println(book.genre)
	fmt.Println(book.publisher)
	fmt.Println(book.publisher.name)
	fmt.Println(book.publisher.address)
}
