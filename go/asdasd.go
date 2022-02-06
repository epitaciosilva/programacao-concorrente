package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(randomString(10))
}

func randomString(l int) string {
    bytes := make([]byte, l)
	t := -1
	i := 0
    for t != 0 {
		t = randInt(65, 90)
        bytes[i] = byte(t)
		i++
		fmt.Println(t)
    }
    return string(bytes)
}

func randInt(min int, max int) int {
	t := rand.Intn(max-min)
    return t
}