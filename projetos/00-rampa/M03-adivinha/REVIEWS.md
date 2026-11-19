# Reviews: M03 adivinha

## 2026-11-19: Review 1

### Ferramentas
- `gofmt -l .`: vazio ✅
- `go vet ./...`: limpo ✅
- `staticcheck ./...`: limpo ✅
- `go test -cover ./...`: passa; `jogo` 100% ✅

### Critérios de aceite
- ✅ Sorteio novo por partida: duas partidas com `seq 1 100` acertaram em 26 e 31 (segredo diferente, contador zerado)
- ✅ `maior`/`menor` corretos
- ⚠️ Palpite inválido vai ao stderr, não conta e o jogo continua, **mas a mensagem não é clara**: `palpite inválido: strconv.Atoi: parsing "abc": invalid syntax`, `palpite inválido: 0 < 1`
- ✅ Espaços aceitos (testes)
- ✅ Singular/plural
- ✅ Jogar de novo: `x` e `50` repetem a pergunta (60 prompts para 59 números + `x`); `S` reinicia; `n` sai com 0
- ✅ EOF em qualquer prompt → exit 0
- ✅ Erro de leitura (`./adivinha < .`) → stderr, exit 1
- ✅ Um único Scanner; `sc.Err()` conferido
- ✅ `switch` substituindo `if/else if` (`LerPalpite`, `Avaliar`, replay com `case "s", "S"`)
- ⚠️ Constantes: `Maximo` usado no sorteio, `Minimo` não (o `+ 1` é o mínimo escrito à mão)
- ✅ Binário ignorado pelo `.gitignore`

### Itens
- 🟡 `jogo/jogo.go:17,22,24`: mensagens de erro. O erro cru do `strconv` vaza para o usuário, e `0 < 1` não diz o que está errado. Já apontado no parecer parcial (item 2) e não tratado.
- 🟡 `main.go:22-27`: dois `if` com condições opostas e consecutivas (`!resetar` / `resetar`).
- 🟡 `main.go:18,38`: sorteio duplicado e com o mínimo hardcoded (`+ 1`); `jogo.Minimo` não participa.
- 🔵 `main.go:20,34-46,68`: um único `for` alterna entre dois modos via flag `resetar` (parecer item 6, não tratado). Funciona, mas o fluxo só se entende simulando o flag.
- 🔵 `main.go:16`: "Pensei em um número..." só aparece na primeira partida.
- 🔵 `main.go:41`: `os.Exit(0)` dentro do loop, enquanto o EOF sai por `break` → dois caminhos de saída para o mesmo resultado.
- ⚪ Linhas em branco no início/fim de `main` e na linha 45.
- ⚪ No EOF, o shell volta na mesma linha do prompt (falta uma quebra de linha).
- ⚪ `.tool-versions` só neste projeto.

### Bom
- `Avaliar` com `switch` sem expressão e comma-ok bem aplicado; retorno explícito depois do parecer.
- Replay com `switch texto { case "s", "S": ... }`: forma certa, vários valores num case.
- Leitura robusta: um Scanner, EOF limpo, `sc.Err()` com exit 1.

### Veredito
🔁 **Refatorar**
1. Mensagens de erro do `LerPalpite` claras para o usuário, sem prefixo duplicado com o `main`.
2. Os dois `if` opostos (linhas 22-27).
3. Sorteio usando `Minimo` e `Maximo`, sem duplicar a expressão.

## 2026-11-19: Review 2

### Ferramentas
- `gofmt`, `go vet`, `staticcheck`: limpos ✅ · `go test -cover`: passa, `jogo` 100% ✅

### Itens do review 1
1. 🟡 Mensagens de erro: **parcial**. Agora citam o texto (`%q`) e o intervalo (`escolha entre 1 e 100`) ✅. Mas a falha de parse ainda mostra `strconv.Atoi: parsing "abc": invalid syntax` ao jogador (via `%w`), e há erros de digitação: `escolhar`, `númerico` (`jogo/jogo.go:17`).
2. 🟡 Dois `if` opostos (`main.go:20-26`): **não tratado**.
3. 🟡 Sorteio duplicado com `+ 1` (`main.go:16,37`): **não tratado**. Além disso, `main.go:21` escreve "1 a 100" à mão.

### Novo
- 🔴 Regressão (`main.go:21`): "Pensei em um número de 1 a 100." foi para dentro do loop e agora sai **antes de cada palpite** (63 vezes numa partida de 63 tentativas, inclusive após palpite inválido).

### Veredito
🔁 **Refatorar**
1. Regressão do "Pensei em um número..." (uma vez por partida).
2. Mensagem de parse sem detalhes internos do `strconv` para o jogador; corrigir a digitação.
3. Os dois `if` opostos.
4. Sorteio sem duplicação, usando `Minimo` e `Maximo`; o texto "1 a 100" também derivado das constantes.

## 2026-11-19: Review 3 (após dicas N1 e N2 sobre a estrutura dos loops)

### Ferramentas
- `gofmt`, `go vet`, `staticcheck`: limpos ✅ · `go test -cover`: passa, `jogo` 100% ✅

### Itens do review 2
1. ✅ Regressão: "Pensei em um número..." uma vez por partida (duas partidas → duas vezes).
2. ✅ Mensagens: `palpite "abc" inválido: escolha um valor numérico entre 1 e 100`; sem detalhe do `strconv`, sem erros de digitação.
3. ✅ Flag `resetar` eliminado: um `for` por nível de repetição, partida e replay em funções.
4. ✅ Sorteio num só lugar com `IntN(Maximo-Minimo+1) + Minimo`; texto derivado das constantes.

### Critérios
- ✅ Todos, exceto: ⚠️ EOF no meio da partida (`printf '50\n' | ./adivinha`) imprime `Jogar novamente? (s/n):` depois que a entrada acabou; o mesmo antes de `erro de leitura`.

### Itens
- 🟡 `main.go:32,40-41,17-19`: `iniciarPartida` não devolve nada, então o `main` não sabe que a entrada acabou e chama `jogarNovamente` mesmo assim (funciona só porque o `Scan` continua falso). O sinal de "fim de entrada" é engolido.
- 🔵 `main.go:32`: `iniciarPartida` joga a partida inteira; o nome promete só o começo.
- ⚪ `main.go:81-82`: `default: continue` é redundante no fim do corpo do `for`.
- ⚪ `main.go:68-69`: falta linha em branco entre as funções; linhas em branco no início/fim de `main`.

### Bom
- A reestruturação: `main` com 6 linhas de lógica que se leem como a descrição do jogo; variáveis da partida nascem zeradas a cada chamada.
- `jogarNovamente` como loop próprio com `switch` de vários valores.
- Fórmula do intervalo com as duas constantes.

### Veredito
🔁 **Refatorar** (1 item)
1. Propagar o fim da entrada da partida para o `main` (nenhum prompt depois de EOF).

## 2026-11-19: Review 4

### Ferramentas
- `gofmt`, `go vet`, `staticcheck`: limpos ✅ · `go test -cover`: passa, `jogo` 100% ✅

### Item do review 3
1. ✅ `novaPartida` devolve `bool` e o `main` encerra no EOF: `printf '50\n' | ./adivinha` termina em `Seu palpite:` com exit 0, sem perguntar "Jogar novamente?". Erro de leitura → exit 1 sem prompt extra.

### Critérios
- ✅ Todos conferidos no binário (duas partidas com sorteio novo e contador zerado; inválidos no stderr; replay repete em resposta inválida; EOF em qualquer prompt → exit 0; erro de leitura → exit 1).

### Para depois
- 🔵 `main.go:32`: o `bool` de `novaPartida` significa "a entrada continua", mas isso não está escrito em lugar nenhum; candidato a doc comment no M03b.
- ⚪ `main.go:15-24`: dá para expressar o loop do `main` só com a condição do `for` (`for a() && b() {}`); questão de gosto.
- ⚪ Linhas em branco antes de `}` (23, 67, 84).

### Bom
- Os dois loops agora sinalizam o fim da entrada do mesmo jeito (`bool`), e o `main` decide o encerramento num lugar só.
- `jogarNovamente` sem o `default` redundante.
- Renomeou `iniciarPartida` → `novaPartida`.

### Veredito
✅ **Aprovado**
