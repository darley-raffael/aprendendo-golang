# M03: adivinha

## Contexto

Um jogo de terminal: o programa pensa num número de 1 a 100 e você tenta adivinhar. Parece trivial, mas é o primeiro programa **interativo** da trilha: um loop que lê o stdin até uma condição, decide entre vários caminhos e precisa lidar com entrada inválida e com o usuário fechando a entrada (Ctrl+D) a qualquer momento.

## Requisitos

### Sessão de exemplo (o texto dos prompts é livre; o comportamento, não)

```
$ ./adivinha
Pensei num número de 1 a 100.
Seu palpite: 50
maior
Seu palpite: 75
menor
Seu palpite: abc
palpite inválido: ...            ← stderr
Seu palpite: 62
Acertou em 3 tentativas!
Jogar de novo? (s/n): s
Pensei num número de 1 a 100.
Seu palpite: 40
Acertou em 1 tentativa!
Jogar de novo? (s/n): n
$ echo $?
0
```

### Comportamento

- [x] Cada partida sorteia um número **novo** entre 1 e 100 (inclusive os dois).
- [x] Palpite menor que o segredo → `maior`; maior que o segredo → `menor`.
- [x] Palpite inválido (não inteiro, vazio, fora de 1..100) → mensagem clara no **stderr**; **não conta** como tentativa; o jogo continua.
- [x] Espaços ao redor do palpite são aceitos (`"  42 "` vale 42).
- [x] No acerto: `Acertou em N tentativas!`, com **singular** quando N = 1 (`1 tentativa`).
- [x] Depois do acerto: `Jogar de novo? (s/n)`. `s` ou `S` → nova partida (contador zerado); `n` ou `N` → encerra com exit code `0`; qualquer outra resposta → pergunta de novo.
- [x] **EOF** (Ctrl+D) em **qualquer** prompt → encerra de forma limpa, exit code `0`, sem panic.
- [x] Erro de leitura do stdin (diferente de EOF) → mensagem no stderr, exit code `1`.

Para testar sem digitar: `printf '50\n75\nn\n' | ./adivinha` (o resultado depende do sorteio, mas o programa precisa terminar sem travar).

## Estrutura (andaime: o tutor define nesta etapa)

```
projetos/00-rampa/M03-adivinha/
├── go.mod               ← você cria
├── main.go              ← você cria: sorteio, loop do jogo, leitura, saída, exit code
└── jogo/
    ├── jogo.go          ← você cria:
    │                       func LerPalpite(texto string) (int, error)
    │                       func Avaliar(palpite, segredo int) (dica string, acertou bool)
    └── jogo_test.go     ← do tutor, não altere
```

- `LerPalpite`: converte o texto num palpite válido (1..100) ou devolve erro, com o zero value no valor.
- `Avaliar`: devolve `"maior"`/`"menor"` e `false`, ou `""` e `true` no acerto (idioma comma-ok).

Caminho do módulo: `github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M03-adivinha`

## Critérios de aceite

- [x] `gofmt -l .` sai vazio
- [x] `go vet ./...` e `staticcheck ./...` sem avisos
- [x] `go test ./...` passa
- [x] Todos os comportamentos acima conferidos no binário, incluindo Ctrl+D no meio da partida e no "jogar de novo"
- [x] Nenhum erro ignorado; `sc.Err()` conferido
- [x] Um único `bufio.Scanner` no programa inteiro
- [x] Pelo menos um `switch` usado onde ele substitui uma cadeia de `if/else if`
- [x] Limites 1 e 100 como **constantes nomeadas**, usadas tanto no sorteio quanto na validação
- [x] O binário `adivinha` não é versionado
- ~~Doc comments~~ dispensados aqui: o M03b `docs` vai documentar este pacote

## Conceitos em jogo

Novos: `for-unico-loop`, `switch`, `funcoes-multiplos-retornos`, `bufio-scanner`, `math-rand`
Aplicação: `erros-como-valores`, `if-com-init`, `strconv`, `constantes`

## Restrições didáticas

- Só a stdlib. `strings.TrimSpace` está liberado (o pacote `strings` em si é o M06).
- O pacote `jogo` **não** lê o stdin, **não** imprime, **não** sorteia e **não** chama `os.Exit`: ele só decide. Entrada, saída, sorteio e encerramento ficam no `main`.
- `map` ainda não (M05).
