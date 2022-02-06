package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	var pointsJ1 = 0
	var pointsJ2 = 0
	var bola int = -1

	for {
		if (pointsJ1 >= 4 && pointsJ1 - 2 >= pointsJ2) {
			fmt.Println("Jogador 1 venceu! %d x %d", pointsJ1, pointsJ2)
			return
		} else if (pointsJ2 >= 4 && pointsJ2 - 2 >= pointsJ1) {
			fmt.Println("Jogador 2 venceu! %d x %d", pointsJ1, pointsJ2)
			return
		}
		j1(&bola, &pointsJ1)
		time.Sleep(1*time.Second)
		j2(&bola, &pointsJ2)
	}
}

func j1(bola *int, points *int) {
	rand.Seed(time.Now().UnixNano())
	if (*bola == 0 || *bola == 5) {
		(*points)++
	}
	*bola = rand.Intn(6)
	fmt.Println("1: ",*bola)
}

func j2(bola *int, points *int) {
	rand.Seed(time.Now().UnixNano())
	if (*bola == 0 || *bola == 5) {
		(*points)++
	}
	*bola = rand.Intn(6)
	fmt.Println("2: ",*bola)
}
