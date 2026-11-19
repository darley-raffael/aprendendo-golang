package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M03-adivinha/jogo"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		if !novaPartida(sc) {
			break
		}

		if !jogarNovamente(sc) {
			break
		}

	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "erro de leitura:", err)
		os.Exit(1)
	}
}

func novaPartida(sc *bufio.Scanner) bool {
	segredo := rand.IntN(jogo.Maximo-jogo.Minimo+1) + jogo.Minimo
	tentativas := 0

	fmt.Printf("Pensei em um número de %d a %d.\n", jogo.Minimo, jogo.Maximo)

	for {
		fmt.Print("Seu palpite: ")
		if !sc.Scan() {
			return false
		}

		texto := sc.Text()

		palpite, err := jogo.LerPalpite(texto)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		dica, acertou := jogo.Avaliar(palpite, segredo)

		tentativas++
		if !acertou {
			fmt.Println(dica)
			continue
		}
		tentativaTexto := "tentativa"
		if tentativas > 1 {
			tentativaTexto = "tentativas"
		}

		fmt.Printf("Acertou em %d %s!\n", tentativas, tentativaTexto)
		return true
	}

}

func jogarNovamente(sc *bufio.Scanner) bool {
	for {
		fmt.Print("Jogar novamente? (s/n): ")
		if !sc.Scan() {
			return false
		}

		switch sc.Text() {
		case "s", "S":
			return true
		case "n", "N":
			return false
		}
	}

}
