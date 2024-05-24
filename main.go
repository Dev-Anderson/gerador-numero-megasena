package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

const maiorNumero = 60

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Declara as variaveis para os argumetnos de linha de comando
	var (
		jogos  int
		dezena int
	)

	// Define os argumentos de linha de comando e valores padrao
	flag.IntVar(&dezena, "d", 6, "Dezenas")
	flag.IntVar(&jogos, "j", 1, "Jogos")
	flag.Parse()

	// Loop para gerar e imprimir os numeros para cada jogo
	for i := 0; i < jogos; i++ {
		numeros, err := gera(dezena)
		if err != nil {
			log.Fatal(err)
		}
		// Imprimi os numeros gerados
		fmt.Println(numeros)
	}
}

func gera(n int) ([]int, error) {
	// verifica se n esta fora do intervalo
	if n < 0 || n > maiorNumero {
		return nil, errors.New(fmt.Sprintf("impossivel gerar %v numeros", n))
	}

	// cria um mapa para armazenar numeros unicos
	mapa := make(map[int]bool)
	for len(mapa) < n {
		dezena := rand.Intn(maiorNumero) + 1
		mapa[dezena] = true
	}

	// Converte o mapa para um slice
	numeros := make([]int, 0, n)
	for k := range mapa {
		numeros = append(numeros, k)
	}

	// Ordena o slice
	sort.Ints(numeros)

	return numeros, nil
}
