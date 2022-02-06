package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(3)
	f := rand.Float64()
	t := float64(9 + i) + f
	fmt.Println(i, f)
	fmt.Printf("%.2f\n", t)
}
