package main

import (
    "fmt"
    "time"
	"sync"
	"math/rand"
	"sort"
)

// Determina o tempo que cada correador levou para percorrer os 100m 
// mínimo é 9s e máximo é 11.99s (são corredores profissinais)
func randTime() float64 {
	i := 9 + rand.Intn(3)
	f := rand.Float64()
	return float64(i) + f
}

// função que representa a corrida de cada corredor
// wg -> variável para sincronizar a execução da thread principal com as demais
// runner -> número do participante
// number_team -> número da equipe
// bastao -> canal para verificar se recebeu o bastao e então pode correr
// nextRunner -> canal que presenta o próximo corredor, no qual tem que passar o bastao
// team_time -> canal com os tempos de cada membro da equipe
func run(wg *sync.WaitGroup, runner int, number_team int, bastao chan bool, nextRunner chan<- bool, team_time chan<- float64) {
	defer wg.Done()

	// enquanto não receber o bastao, fica em espera
	correr := <- bastao

	// tempo do corredor
	runner_time := randTime()

	// Dados do corredor
	fmt.Printf("Equipe %d - Corredor %d: %.2fs\n", number_team, runner, runner_time)

	// 1 segundo, pois adicionar o tempo real demoraria para terminar de rodar o programa
	time.Sleep(1*time.Second)

	// // somando o tempo do corredor para o do time
	// team_time_r := <- team_time
	// team_time_r = team_time_r + runner_time
	team_time <- runner_time

	// caso seja 4 então é o último e não precisa passar o bastão
	if (runner != 4) {
		nextRunner <- correr
	} else {
		// fecha o canal
		close(team_time)
	}
}

// função que representa a corrida de cada equipe
// wg -> variável para sincronizar a execução da thread principal com as demais
// number_team -> número da equipe
// teams_time -> slice para armazenar o tempo total cada equipe
func team(wg *sync.WaitGroup, number_team int, teams_time []float64) {
	defer wg.Done()

	// canais que representam os corredores
	runner1 := make(chan bool)
	runner2 := make(chan bool)
	runner3 := make(chan bool)
	runner4 := make(chan bool)

	team_time := make(chan float64)

	wg.Add(1)
	// 1º corredor e o próximo é o 2º
	go run(wg, 1, number_team, runner1, runner2, team_time)

	wg.Add(1)
	// 2º corredor e o próximo é o 3º
	go run(wg, 2, number_team, runner2, runner3, team_time)

	wg.Add(1)
	// 3º corredor e o próximo é o 4º
	go run(wg, 3, number_team, runner3, runner4, team_time)

	wg.Add(1)
	// 4º corredor
	// não tem próximo, mas a função exige um parâmetro
	go run(wg, 4, number_team, runner4, runner4, team_time)

	runner1 <- true

	// contagem do tempo total da equipe
	// lê do buffer do canal e realiza a soma
	soma := 0.0
	for t := range team_time {
		soma = soma + t
	}
	// adiciona a soma ao slice, cada posição representa uma equipe
	teams_time[number_team-1] = soma
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)

	teams_time := make([]float64, 4)
	teams_time_sorted := make([]float64, 4)

	// inicializando as equipes
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go team(wg, i, teams_time)
	}

	// espera todos os corredores (threads) terminarem de correr
	wg.Wait()

	// feito a copia pois o tempo de cada equipe é determinado pelo index
	// e ao ordernar, isto é perdido
	copy(teams_time_sorted, teams_time)

	sort.Float64s(teams_time_sorted)
	fmt.Println("\n---------- Ranking ----------")

	for idx, ts := range teams_time_sorted {
		fmt.Printf("%dº - Equipe ", idx+1)
		// verificando a qual equipe pertence o tempo em questão
		for team_idx, t := range teams_time {
			if (ts == t) {
				fmt.Printf("%d - Tempo: %.2f\n", team_idx+1, t)
			}
		}
	}
}
