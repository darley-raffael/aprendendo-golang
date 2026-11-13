package main

import (
	"fmt"
	"os"

	"github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M01-ola/saudacao"
)

func main() {
	args := os.Args
	var name string

	if len(args) > 1 {
		name = args[1]
	}

	fmt.Println(saudacao.Ola(name))
}
