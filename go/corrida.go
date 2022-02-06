package main

import (
    "fmt"
    "time"
	"sync"
	"math/rand"
)

// Determina o tempo que cada correador levou para percorrer os 100m 
// mínimo é 9s e máximo é 11.99s (são corredores profissinais)
func randTime() float64 {
	i := 9 + rand.Intn(3)
	f := rand.Float64()
	return float64(i) + f
}

// função que presenta os corredores
// wg -> variável para sincronizar as threads
// runner -> número do participante
// bastao -> canal para verificar se recebeu o bastao e então pode correr
// nextRunner -> canal que presenta o próximo corredor, no qual tem que passar o bastao
func c1(wg *sync.WaitGroup, runner int, bastao chan bool, nextRunner chan<- bool) {
	defer wg.Done()

	// enquanto não receber o bastao, fica em espera
	correr := false
	for !correr {
		select {
		case correr = <- bastao:
		}
	}

	// Dados do corredor
	fmt.Printf("Corredor %d: ", runner)

	// 1 segundo, pois adicionar o tempo real demoraria para terminar de rodar o programa
	time.Sleep(1*time.Second)

	// imprimi o tempo do corredor
	fmt.Printf("%.2fs\n", randTime())

	// caso seja 4 então é o último e não precisa passar o bastão
	if (runner != 4) {
		nextRunner <- correr
	}
}


func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)

	// canais que representam os corredores
	runner1 := make(chan bool)
    runner2 := make(chan bool)
	runner3 := make(chan bool)
	runner4 := make(chan bool)


	// 1º corredor e o próximo é o 2º
	wg.Add(1)
	go c1(wg, 1, runner1, runner2)

	// 2º corredor e o próximo é o 3º
	wg.Add(1)
	go c1(wg, 2, runner2, runner3)

	// 3º corredor e o próximo é o 4º
	wg.Add(1)
	go c1(wg, 3, runner3, runner4)

	// 4º corredor
	// não tem próximo, mas a função exige um parâmetro
	wg.Add(1)
	go c1(wg, 4, runner4, runner4)

	// Determina o início da corrida
	runner1 <- true

	wg.Wait()
}