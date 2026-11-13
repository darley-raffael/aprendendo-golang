package saudacao

import "testing"

// Especificação do M01 (escrita pelo tutor). Não altere este arquivo:
// faça os testes passarem implementando Ola em saudacao.go.
func TestOla(t *testing.T) {
	casos := []struct {
		nome     string
		entrada  string
		esperado string
	}{
		{"com nome", "Maria", "Olá, Maria!"},
		{"nome com acento", "José", "Olá, José!"},
		{"sem nome usa mundo", "", "Olá, mundo!"},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			obtido := Ola(c.entrada)
			if obtido != c.esperado {
				t.Errorf("Ola(%q) = %q; esperado %q", c.entrada, obtido, c.esperado)
			}
		})
	}
}
