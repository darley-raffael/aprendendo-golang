package conversao

import (
	"math"
	"testing"
)

// Especificação do M02 (escrita pelo tutor). Não altere este arquivo:
// faça os testes passarem implementando Converter em conversao.go.

// tolerancia para comparar float64: nunca compare floats com == depois de contas.
const tolerancia = 1e-9

func TestConverter(t *testing.T) {
	casos := []struct {
		nome     string
		valor    float64
		de, para string
		esperado float64
	}{
		{"água ferve", 100, "c", "f", 212},
		{"água congela", 0, "c", "f", 32},
		{"-40 é igual nas duas escalas", -40, "c", "f", -40},
		{"fahrenheit para celsius", 212, "f", "c", 100},
		{"quilômetros para milhas", 10, "km", "mi", 6.213711922373339},
		{"milhas para quilômetros", 1, "mi", "km", 1.609344},
		{"quilos para libras", 1, "kg", "lb", 2.2046226218487757},
		{"libras para quilos", 1, "lb", "kg", 0.45359237},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido, err := Converter(c.valor, c.de, c.para)
			if err != nil {
				t.Fatalf("Converter(%v, %q, %q) retornou erro inesperado: %v", c.valor, c.de, c.para, err)
			}
			if math.Abs(obtido-c.esperado) > tolerancia {
				t.Errorf("Converter(%v, %q, %q) = %v; esperado %v", c.valor, c.de, c.para, obtido, c.esperado)
			}
		})
	}
}

func TestConverterErros(t *testing.T) {
	casos := []struct {
		nome     string
		de, para string
	}{
		{"unidade de origem desconhecida", "x", "f"},
		{"unidade de destino desconhecida", "c", "parsec"},
		{"temperatura para distância", "c", "km"},
		{"peso para temperatura", "kg", "c"},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido, err := Converter(10, c.de, c.para)
			if err == nil {
				t.Fatalf("Converter(10, %q, %q) deveria retornar erro", c.de, c.para)
			}
			if obtido != 0 {
				t.Errorf("em caso de erro o valor deve ser o zero value; obtido %v", obtido)
			}
		})
	}
}
