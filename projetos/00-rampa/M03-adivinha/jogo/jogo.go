package jogo

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Minimo = 1
	Maximo = 100
)

func LerPalpite(texto string) (int, error) {
	palpite, err := strconv.Atoi(strings.TrimSpace(texto))
	if err != nil {
		return 0, fmt.Errorf("palpite %q inválido: escolha um valor numérico entre %d e %d", texto, Minimo, Maximo)
	}

	if palpite < Minimo || palpite > Maximo {
		return 0, fmt.Errorf("palpite %q inválido: escolha entre %d e %d", texto, Minimo, Maximo)
	}

	return palpite, nil
}

func Avaliar(palpite, segredo int) (dica string, acertou bool) {
	switch {
	case palpite > segredo:
		dica = "menor"
	case palpite < segredo:
		dica = "maior"
	default:
		acertou = true
	}

	return dica, acertou
}
