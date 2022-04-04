package main

/*
1. goroutines untuk wasit
2. wasit akan mengambil bola ketika salah satu pemain tidak dapat mengembalikan bola
3. permainan selesai, wasit menentukan pemenang.
*/
import (
	"log"
	"math/rand"
	"time"
)

func main() {
	table := make(chan *ball)
	done := make(chan *ball)
	go Player("andi", table, done)
	go Player("joni", table, done)

	wasit(table, done)

}

type ball struct {
	hits       int
	lastplayer string
}

func wasit(table chan *ball, done chan *ball) {
	table <- new(ball)

	for {
		select {
		case ball := <-done:
			log.Println("winner is", ball.lastplayer)

			return
		}
	}
}

func Player(name string, table chan *ball, done chan *ball) {
	for {

		// untuk fitur no 2
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")

				done <- ball
				return
			}
			ball.hits++
			ball.lastplayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return

		}

	}
}
