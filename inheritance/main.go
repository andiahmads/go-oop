package main

import "fmt"

/* Inheritance atau Pewarisan/Penurunan adalah konsep pemrograman
dimana sebuah class dapat ‘menurunkan’ property dan method yang dimilikinya  kepada class lain.
Konsep inheritance digunakan untuk memanfaatkan fitur ‘code reuse’ untuk menghindari duplikasi kode program
*/

type Comic struct {
	Universe string
}

func (comic Comic) ComicUniverse() string {
	return comic.Universe
}

//implementasikan inherintance
type Marvel struct {
	Comic
}

type Dc struct {
	Comic
}

func main() {
	c1 := Marvel{
		Comic{
			Universe: "MCU",
		},
	}
	fmt.Println(c1.ComicUniverse())

	c2 := Dc{
		Comic{
			Universe: "DC",
		},
	}
	fmt.Println(c2.ComicUniverse())

}
