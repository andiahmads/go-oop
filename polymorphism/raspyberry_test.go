package polymorphism

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type raspyberryService interface {
	getFullInformation() string
}

type termometer struct {
	merk  string
	value int
}

func (t termometer) getName() string {
	return t.merk
}
func (t termometer) getValue() int {
	return t.value
}

func (t termometer) getFullInformation() string {
	return fmt.Sprintf("%s dengan temperatur =  %d ", t.getName(), t.getValue())
}

func NewTermometer(merk string, value int) raspyberryService {
	return termometer{merk: merk, value: value}
}

func SendInformation1(channel chan<- interface{}) {

	randomValue := rand.Intn(40 - 20 + 1)
	termometer1 := NewTermometer("termometer 1", randomValue)

	time.Sleep(2 * time.Second)

	channel <- termometer1.getFullInformation()

	fmt.Println("send info from termometer", termometer1.getFullInformation())

}

func SendInformation2(channel chan<- interface{}) {

	randomValue := rand.Intn(40 - 20 + 1)
	termometer1 := NewTermometer("termometer 2", randomValue)

	time.Sleep(2 * time.Second)

	channel <- termometer1.getFullInformation()

	fmt.Println("send info from termometer", termometer1.getFullInformation())

}

func ReceipInformation(channel <-chan interface{}) {
	time.Sleep(5 * time.Second)

	data := <-channel
	fmt.Println("receipt full information from", data)
}

func TestReadInfo(t *testing.T) {

	termometer1 := make(chan interface{})
	termometer2 := make(chan interface{})
	defer close(termometer1)
	defer close(termometer2)

	for {
		time.Sleep(2 * time.Second)
		go SendInformation1(termometer1)
		go SendInformation2(termometer2)

		go ReceipInformation(termometer1)
		go ReceipInformation(termometer2)
	}
}
