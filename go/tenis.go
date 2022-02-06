package main

import (
	"fmt"
	"time"
	"math/rand"
	// "sync"
)

// var wg sync.WaitGroup

// func j1(player1 chan string, player2 chan string, bola chan int, win chan bool, points chan int) {
// 	rand.Seed(time.Now().UnixNano())
// 	poinst_player := <-points
// 	fmt.Println(poinst_player)
// 	for i := 0; i< 1; i++ {
// 		select {
// 		case <-player1:
// 			fmt.Println("Jogador 1")
// 			// 0 -> bater na rede
// 			// 5 -> fora da quadra
// 			// 1...4 -> dentro da quadra
// 			bola_caiu := <-bola
// 			if bola_caiu == 0 || bola_caiu == 5 {
// 				poinst_player := poinst_player + 1
// 				points <- poinst_player
// 			} else {
// 				bola <- rand.Intn(6)
// 			}
// 			player2 <- "joga 2"
// 		}
// 	}
// }

// func j2(player1 chan string, player2 chan string, bola chan int, win chan bool, points chan int) {
// 	rand.Seed(time.Now().UnixNano())
// 	poinst_player := <-points

// 	for {
// 		select {
// 		case <-player2:
// 			// 0 -> bater na rede
// 			// 5 -> fora da quadra
// 			// 1...4 -> dentro da quadra
// 			bola_caiu := <-bola
// 			if bola_caiu == 0 || bola_caiu == 5 {
// 				poinst_player := poinst_player + 1
// 				points <- poinst_player
// 			} else {
// 				bola <- rand.Intn(6)
// 			}
// 			player1 <- "joga 1"
// 			win <- true
// 		case <-win:
// 			fmt.Println("Alguém ganhou 2!")
// 			return
// 		}
// 	}
// }

func main() {
	// player1 := make(chan string)
	// player2 := make(chan string)
	// pointsJ1 := make(chan int)
	// pointsJ2 := make(chan int)
	bola := make(chan int)
	win := make(chan bool)
	// for i := 0; i < 5; i++ {
		// }
	// wg.Add(2) // 2
	go func () {
		// player1 <- "joga"
		bola <- 0
		// wg.Done()
	}()
	go func () {
		// player1 <- "joga"
		bola <- 0
		// wg.Done()
	}()
	for i := 0; i< 2; i++ {
		var teste int = 0
		teste = i
		select {
		case bola_caiu := <-bola:
			fmt.Println("Jogador 1")
			// 0 -> bater na rede
			// 5 -> fora da quadra
			// 1...4 -> dentro da quadra
			if bola_caiu == 0 || bola_caiu == 5 {
				// poinst_player := poinst_player + 1
				// pointsJ1 <- 1
				// win <- true
			} else {
				bola <- rand.Intn(6)
			}
			// player2 <- "joga 2"
		}
		if teste == 1 {
			win <- true
			close(win)
		}
	}
	for i :=0; i < 1; i++ {
		fmt.Println(<-win)
	}
	// go j1(player1, player2, bola, win, pointsJ1)
		// go j2(player1, player2, bola, win, pointsJ2)

	// wg.Wait() // 4
	time.Sleep(1e9)
}