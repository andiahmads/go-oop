package main

import "fmt"

//implementasikan interface

type LivingThings interface {
	Walk(length int64)
	Run(length int64, speed float64)
}

type Human struct {
	Age          int64
	Name         string
	HansomeLavel int64
}

//buat contract dan impl interface
func (h Human) Walk(length int64) {
	fmt.Println(h.Name, "Walk for ", length, "meters")
}

func (h Human) Run(length int64, speed float64) {
	fmt.Println(h.Name, "run for", length, "meter with speed: ", speed)
}

type Dog struct {
	Age  int64
	Name string
}

func (h Dog) Bark() {
	fmt.Println("guk guk guk auuuum")
}

func (d Dog) Walk(length int64) {
	fmt.Println(d.Name, " walking for", length, "meters")
}

func (d Dog) Run(length int64, speed float64) {
	fmt.Println(d.Name, " run for", length, "meter with speed: ", speed)
}

/*
Apakah untuk membuat function , saya harus membuat interface ?
Tentu saja tidak, kamu bisa bikin struct tanpa harus ada interface.
Tetapi jika kamu butuh injection dan operasi yang mengharuskan memiliki interface,
maka kamu harus buat interfacenya.
*/

// contoh penggunaan injection pada golang
func AcceptLivingsOnly(an LivingThings) {
	an.Run(20, 12)
	an.Walk(20)
}

func main() {
	andi := Human{
		Age:          10,
		Name:         "Andi",
		HansomeLavel: 21,
	}

	dog := Dog{
		Age:  10,
		Name: "caca",
	}
	//tanpa injection
	// andi.Walk(10)
	// andi.Run(78, 90)

	//dengan injection
	AcceptLivingsOnly(andi)
	AcceptLivingsOnly(dog)
	dog.Bark()
	// Dengan interface, kamu bisa memasukkan jenis yang berbeda beda pada satu parameter
}
