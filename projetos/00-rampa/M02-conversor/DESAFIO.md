# M02: conversor

## Contexto

Uma CLI de conversão de unidades. É pequena, mas tem tudo que uma CLI de verdade precisa: ler e validar a entrada, operar sobre tipos numéricos sem coerção implícita, devolver erros como valores e sair com o exit code correto.

## Requisitos

### Conversões suportadas (unidades em minúsculas)

| Categoria   | Pares                              |
| ----------- | ---------------------------------- |
| Temperatura | `c` ↔ `f`                          |
| Distância   | `km` ↔ `mi` (1 mi = 1,609344 km)   |
| Peso        | `kg` ↔ `lb` (1 lb = 0,45359237 kg) |

### Comportamento

- [x] `./conversor 100 c f` → `212.00 °F`
- [x] `./conversor 10 km mi` → `6.21 mi`
- [x] `./conversor 1 lb kg` → `0.45 kg`
- [x] `./conversor -40 c f` → `-40.00 °F`
- [x] `./conversor abc c f` → mensagem de erro clara no **stderr**, exit code `1`
- [x] `./conversor 100 c km` → mensagem de erro clara no **stderr**, exit code `1`
- [x] `./conversor` ou número errado de argumentos → `uso: conversor <valor> <de> <para>` no **stderr**, exit code `2`
- [x] Saída sempre com 2 casas decimais; temperaturas com `°C`/`°F`, demais com a própria unidade.

Para conferir o exit code: `./conversor abc c f; echo $?`
Para conferir que foi para o stderr: `./conversor abc c f 2>/dev/null` não deve imprimir nada.

## Estrutura (andaime: o tutor define nesta etapa)

```
projetos/00-rampa/M02-conversor/
├── go.mod                   ← você cria
├── main.go                  ← você cria: argumentos, parse, saída, exit code
└── conversao/
    ├── conversao.go         ← você cria: func Converter(valor float64, de, para string) (float64, error)
    └── conversao_test.go    ← do tutor, não altere
```

Caminho do módulo: `github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M02-conversor`

## Critérios de aceite

- [ ] `gofmt -l .` sai vazio
- [ ] `go vet ./...` e `staticcheck ./...` sem avisos
- [ ] `go test ./...` passa
- [ ] Todos os exemplos de comportamento acima conferidos no binário
- [ ] Nenhum erro ignorado (`_` no lugar de `err`)
- [ ] Fatores de conversão como **constantes nomeadas**, sem números mágicos espalhados
- ~~Doc comments em tudo que for exportado~~ **dispensado no M02** a pedido do aluno; o tema ganhou o micro-projeto M03b `docs`
- [ ] O binário `conversor` não é versionado

## Conceitos em jogo

`variaveis-e-zero-values`, `tipos-basicos`, `conversao-de-tipos-explicita`, `constantes`, `strconv`, `erros-como-valores`, `if-com-init`, `fmt-verbos` (aplicação)

## Restrições didáticas

- Só a stdlib.
- `switch` e `map` ainda não foram apresentados (M03 e M05): resolva com `if`/`else if`. Se o código ficar repetitivo, anote o incômodo: ele é a motivação do próximo conceito.
- `Converter` não imprime nada e não chama `os.Exit`: quem decide sobre saída e encerramento é o `main`.
