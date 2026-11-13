// saudacao fornce uma funçào que imprime uma mensagem para o usuário
package saudacao

import (
	"strings"
)

// Ola retorna uma mensagem de saudacao para o name informado
func Ola(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Olá, mundo!"
	}

	return "Olá, " + name + "!"
}
