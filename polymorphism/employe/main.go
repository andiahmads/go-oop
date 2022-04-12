package main

import "fmt"

type employee interface {
	develop() int
	name() string
}

type team1 struct {
	totalapp_1 int
	name_1     string
}

func (t1 team1) develop() int {
	return t1.totalapp_1
}
func (t1 team1) name() string {
	return t1.name_1
}

type team2 struct {
	totalapp_2 int
	name_2     string
}

func (t2 team2) develop() int {
	return t2.totalapp_2
}

func (t2 team2) name() string {
	return t2.name_2
}

func findaDevelope(i []employee) {
	totalProject := 0
	for _, ele := range i {
		fmt.Printf("\nProject environment = %s\n", ele.name())
		fmt.Printf("\nTotal Project environment = %d\n", ele.develop())
		totalProject += ele.develop()
	}
}

func main() {
	res1 := team1{
		totalapp_1: 20,
		name_1:     "android",
	}
	res2 := team2{
		totalapp_2: 40,
		name_2:     "IOS",
	}

	final := []employee{res1, res2}
	findaDevelope(final)
}
