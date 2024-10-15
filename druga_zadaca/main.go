/*
1. Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu. Kreiraj metode za izračunavanje površine i opsega pravokutnika,
metode moraju biti vezane za novo kreiranu strukturu Pravokutnika. U main funkciji inicijalizirajte jedan pravokutnik, 
te pozovite iznad kreirane metode za računanje površine i opsega.

package main 

import "fmt"

type Pravokutnik struct{

	sirina int 
	visina int
}

func povrsina(p Pravokutnik)int{

	return p.sirina * p.visina

}

func opseg(p Pravokutnik)int{

	return 2 * (p.sirina + p.visina)

}

func main(){
	pravokutnik := Pravokutnik{
		sirina : 5,
		visina : 6,
	}
	povrsinaPravokutnika := povrsina(pravokutnik)
	opsegPravokutnika := opseg(pravokutnik)
	
	fmt.Println(pravokutnik)
	fmt.Println(povrsinaPravokutnika)
	fmt.Println(opsegPravokutnika)
}
*/

/*2. Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu.
Zatim napiši strukturu koja predstavlja osobu, koja sadrži ime, godine i adresu.
Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu. 
(Sva polja) u formatu: Ime Prezime, 20 godina, živi u Grad, Ulica.
*/

package main 

import "fmt"

type Adresa struct{

	grad  string
	ulica string
}

type Osoba struct{

	ime    string
	godine int
	adresa Adresa
}

func opis(o Osoba)string{

	return fmt.Sprintf("%s, %d godina, živi u %s, %s.", o.ime, o.godine, o.adresa.grad, o.adresa.ulica)

}

func main(){

	osoba := Osoba{
		
		ime : "Mile Pervan",
		godine : 25,
		adresa : Adresa{
			grad : "Mostar",
			ulica : "Rudarska 126",
		},

	}
	opisOsobe := opis(osoba)

	fmt.Println(opisOsobe)
}