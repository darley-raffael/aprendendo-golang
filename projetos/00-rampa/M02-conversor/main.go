package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M02-conversor/conversao"
)

func main() {

	var sufixo string

	args := os.Args

	if len(args) != 4 {
		fmt.Fprintln(os.Stderr, "uso: conversor <valor> <de> <para>")
		os.Exit(2)
	}

	valor, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "valor inválido:", err)
		os.Exit(1)
	}
	de := args[2]
	para := args[3]

	resultado, err := conversao.Converter(valor, de, para)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if para == "f" {
		sufixo = "°F"
	} else if para == "c" {
		sufixo = "°C"
	} else {
		sufixo = para
	}

	fmt.Printf("%.2f %s\n", resultado, sufixo)

}
