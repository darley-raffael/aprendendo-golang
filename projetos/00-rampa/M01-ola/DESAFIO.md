# M01: ola

## Contexto

O primeiro programa: uma CLI que cumprimenta quem for passado como argumento. O objetivo não é a lógica (que é trivial), e sim passar pelo ciclo completo de Go: criar o módulo, separar um pacote, exportar uma função, testar, formatar, compilar e rodar o binário.

## Requisitos

- [x] `./ola Maria` → `Olá, Maria!`
- [x] `./ola` (sem argumento) → `Olá, mundo!`
- [x] `./ola Maria João` → usa apenas o primeiro argumento: `Olá, Maria!`
- [ ] A regra da saudação mora num pacote separado, `saudacao`, numa função exportada `Ola(nome string) string`; o `main` só lê o argumento, chama `Ola` e imprime.

## Estrutura (andaime: o tutor define nesta etapa)

```
projetos/00-rampa/M01-ola/
├── go.mod                  ← você cria com go mod init
├── main.go                 ← você cria (package main)
└── saudacao/
    ├── saudacao.go         ← você cria (package saudacao, func Ola)
    └── saudacao_test.go    ← do tutor, não altere
```

Caminho do módulo: `github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M01-ola`

## Critérios de aceite

- [ ] `gofmt -l .` sai vazio
- [ ] `go vet ./...` sem avisos
- [ ] `go test ./...` passa
- [ ] `go build` gera o binário `ola` e os três exemplos acima funcionam rodando o binário
- [ ] Nada exportado além do necessário

## Conceitos em jogo

`toolchain-go`, `modules-e-go-mod`, `pacotes-e-visibilidade`, `fmt-verbos`

## Restrições didáticas

- Só a stdlib.
- Para ler os argumentos use `os.Args` (é um `[]string`; `os.Args[0]` é o próprio programa). Slices e `os.Args` serão aprofundados depois; aqui basta `len(os.Args)` e indexação.
