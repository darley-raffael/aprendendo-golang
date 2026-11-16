// Pacote conversao fornece um método para fazer a conversão de unidades de medidas
package conversao

import "fmt"

const (
	celsiusPorF = 32
	baseCalculo = 1.8
	kgPorLibra  = 0.45359237
	kmPorMilhas = 1.609344
)

// Converter faz a conversão de unidades de medidas
func Converter(valor float64, de, para string) (float64, error) {

	if de == "c" && para == "f" {
		return (valor * baseCalculo) + celsiusPorF, nil
	}
	if de == "f" && para == "c" {
		return (valor - celsiusPorF) / baseCalculo, nil
	}
	if de == "km" && para == "mi" {
		return valor / kmPorMilhas, nil
	}
	if de == "mi" && para == "km" {
		return valor * kmPorMilhas, nil
	}
	if de == "kg" && para == "lb" {
		return valor / kgPorLibra, nil
	}
	if de == "lb" && para == "kg" {
		return valor * kgPorLibra, nil
	}

	return 0, fmt.Errorf("falha ao tentar converter %q para %q", de, para)
}
