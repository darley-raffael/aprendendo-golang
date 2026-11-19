package jogo

import "testing"

// Especificação do M03 (escrita pelo tutor). Não altere este arquivo:
// faça os testes passarem implementando LerPalpite e Avaliar em jogo.go.

func TestLerPalpite(t *testing.T) {
	casos := []struct {
		nome     string
		entrada  string
		esperado int
	}{
		{"número comum", "50", 50},
		{"limite inferior", "1", 1},
		{"limite superior", "100", 100},
		{"espaços ao redor são ignorados", "  42  ", 42},
		{"tab e espaço", "\t7 ", 7},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido, err := LerPalpite(c.entrada)
			if err != nil {
				t.Fatalf("LerPalpite(%q) retornou erro inesperado: %v", c.entrada, err)
			}
			if obtido != c.esperado {
				t.Errorf("LerPalpite(%q) = %d; esperado %d", c.entrada, obtido, c.esperado)
			}
		})
	}
}

func TestLerPalpiteErros(t *testing.T) {
	casos := []struct {
		nome    string
		entrada string
	}{
		{"vazio", ""},
		{"só espaços", "   "},
		{"texto", "abc"},
		{"decimal", "3.5"},
		{"abaixo do mínimo", "0"},
		{"negativo", "-5"},
		{"acima do máximo", "101"},
		{"número com lixo", "42abc"},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido, err := LerPalpite(c.entrada)
			if err == nil {
				t.Fatalf("LerPalpite(%q) deveria retornar erro", c.entrada)
			}
			if obtido != 0 {
				t.Errorf("em caso de erro o valor deve ser o zero value; obtido %d", obtido)
			}
		})
	}
}

func TestAvaliar(t *testing.T) {
	casos := []struct {
		nome            string
		palpite         int
		segredo         int
		dicaEsperada    string
		acertouEsperado bool
	}{
		{"palpite baixo", 50, 70, "maior", false},
		{"palpite alto", 90, 70, "menor", false},
		{"acerto", 70, 70, "", true},
		{"acerto no mínimo", 1, 1, "", true},
		{"acerto no máximo", 100, 100, "", true},
		{"vizinho de baixo", 69, 70, "maior", false},
		{"vizinho de cima", 71, 70, "menor", false},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			dica, acertou := Avaliar(c.palpite, c.segredo)
			if dica != c.dicaEsperada || acertou != c.acertouEsperado {
				t.Errorf("Avaliar(%d, %d) = (%q, %t); esperado (%q, %t)",
					c.palpite, c.segredo, dica, acertou, c.dicaEsperada, c.acertouEsperado)
			}
		})
	}
}
