package main

import "fmt"

/*
func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println(i)
	}
}


func main() {

	numbers := []int{10,20,30,40}

	for index, value := range numbers {

		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}


func main() {

	broj := 5

	faktorijal := 1

	for i := 1; i <= broj; i++ {

		faktorijal = faktorijal * i
	}
	fmt.Println(broj, faktorijal)


func main() {

    number := 5

	if number < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
	} else {
		fmt.Printf("Factorial of %d is %d\n", number, factorial(number))
	}
}

func factorial(number int) int {

		if number == 1 {
			return 1
		}

		factorial := number * factorial(number - 1)

		return factorial
	}



func main(){

	array_numbers := [5]int{10,20,30,40,50}
	sum := 0
	avg := 0


	for _, num:= range(array_numbers) {

		sum += num
	}
	avg = sum / len(numbers)

	fmt.Println("Brojevi su", array_numbers, "Prosjek im je", avg, "Suma im je", sum)

}


func main(){

	slice_numbers := []int{22, 1133, 818, 181, 718}

	fmt.Println(len(slice_numbers), cap(slice_numbers))

	slice_numbers = append(slice_numbers, 145, 323, 434)

	fmt.Println(slice_numbers)

	fmt.Println(len(slice_numbers), cap(slice_numbers))
}


func main() {

    phones := map[string]int { 

	  "phone1" : 300,
	  "phone2" : 500,
	  "phone3" : 120,

	 }

	fmt.Println(phones)

	if price, exists := phones["phone3"]; exists {
		fmt.Printf("Price is %v", price)
	} else {
		fmt.Println("phone doesn't exist")
	}
}
*/


func main() {
		
		osobe := map[string]int{

			"Mile": 25,
			"Marko": 30,
			"Filip": 35,
		}

	fmt.Println(osobe)

		osobe["Luka"] = 20
		osobe["Mile"] = 30

	fmt.Println(osobe)
	
}
