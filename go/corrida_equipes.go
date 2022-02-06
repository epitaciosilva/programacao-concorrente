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

// função que presenta a corrida de cada corredor
// wg -> variável para sincronizar as threads
// runner -> número do participante
// bastao -> canal para verificar se recebeu o bastao e então pode correr
// nextRunner -> canal que presenta o próximo corredor, no qual tem que passar o bastao
func run(wg *sync.WaitGroup, runner int, team int, bastao chan bool, nextRunner chan<- bool) {
	defer wg.Done()

	// enquanto não receber o bastao, fica em espera
	correr := <- bastao
	
	// tempo do corredor
	runner_time := randTime()

	// Dados do corredor
	fmt.Printf("Equipe %d - Corredor %d: %.2fs\n", team, runner, runner_time)
	
	// 1 segundo, pois adicionar o tempo real demoraria para terminar de rodar o programa
	time.Sleep(1*time.Second)

	// // somando o tempo do corredor para o do time
	// team_time_r := <- team_time
	// team_time_r = team_time_r + runner_time
	// team_time <- team_time_r

	// caso seja 4 então é o último e não precisa passar o bastão
	if (runner != 4) {
		nextRunner <- correr
	}
}

func equip(wg *sync.WaitGroup, team int) {
	defer wg.Done()

	// wg := new(sync.WaitGroup)
	runner1 := make(chan bool)
	runner2 := make(chan bool)
	runner3 := make(chan bool)
	runner4 := make(chan bool)

	// team_time := make(chan float64)

	// 1º corredor e o próximo é o 2º
	wg.Add(1)
	go run(wg, 1, team, runner1, runner2)

	// 2º corredor e o próximo é o 3º
	wg.Add(1)
	go run(wg, 2, team, runner2, runner3)

	// 3º corredor e o próximo é o 4º
	wg.Add(1)
	go run(wg, 3, team, runner3, runner4)

	// 4º corredor
	// não tem próximo, mas a função exige um parâmetro
	wg.Add(1)
	go run(wg, 4, team, runner4, runner4)

	runner1 <- true
	// team_time <- 0.0

	// wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)

	for i := 1; i < 5; i++ {
		wg.Add(1)
		equip(wg, i)
	}

	wg.Wait()
}
