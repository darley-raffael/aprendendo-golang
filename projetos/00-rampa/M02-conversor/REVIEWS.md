# Reviews: M02 conversor

## 2026-11-16, review 1: 🔁 refatorar

### Ferramentas
- `gofmt -l .`: vazio ✅
- `go vet ./...`: ok ✅
- `staticcheck ./...`: ok ✅
- `go test -cover ./...`: ok; `conversao` 100%

### Critérios de aceite
- ✅ `100 c f` → `212.00 °F`; `10 km mi` → `6.21 mi`; `1 lb kg` → `0.45 kg`; `-40 c f` → `-40.00 °F`
- ✅ `abc c f` → stderr, exit 1
- ✅ `100 c km` → stderr (`falha ao tentar converter "c" para "km"`), exit 1
- ❌ argumentos errados → exit 2 ✅, mas a mensagem não é `uso: conversor <valor> <de> <para>`
- ✅ 2 casas decimais e símbolos
- ✅ nenhum erro ignorado com `_`
- ✅ fatores como constantes nomeadas
- ➖ doc comments: dispensados no M02
- ⚠️ binário: não há binário na pasta, mas `conversor` não foi incluído no `.gitignore`

### Itens
1. 🟡 `main.go:22-29`: checagem de `err` separada da chamada por duas atribuições; a mensagem `Valor inválido` descarta o `err` e não mostra qual valor foi rejeitado.
2. 🟡 `main.go:18`: mensagem de uso fora do especificado (e com maiúscula/sem acento: "Numero").
3. 🟡 `.gitignore`: falta `projetos/00-rampa/M02-conversor/conversor`.
4. 🔵 `conversao/conversao.go:7-8`: `celsiusPorF = 32` é um deslocamento (ponto de congelamento em °F), `baseCalculo = 1.8` não diz o que é; `kmPorMilhas` × `kgPorLibra` (plural inconsistente).
5. 🔵 `main.go:13`: `sufixo` declarado no topo, longe do único uso (linha 37).
6. ⚪ `NaN c f` → `NaN °F` e `Inf c f` → `+Inf °F`: `ParseFloat` aceita esses literais. Faz sentido para um conversor?
7. ⚪ linhas em branco no início/fim dos corpos (`main.go:12,46`, `conversao.go:15`).
8. ⚪ `conversao.go:1`: em Go "método" é função com receiver; `Converter` é função.

### Pontos fortes
- `Converter` virou 6 `if` planos com early return, um por par direcionado: cada caso legível isoladamente.
- Erro criado no ponto da falha com `fmt.Errorf` + `%q`, minúsculo e com contexto.
- stderr + exit codes distintos (2 uso, 1 entrada) e staticcheck instalado e limpo.

### Veredito
🔁 Refatorar itens 1–3.

## 2026-11-16, review 2: 🔁 falta 1 ajuste

Ferramentas: gofmt, vet, staticcheck e testes ✅.

1. 🟡 parcial: a mensagem agora inclui o `err` ✅, mas a checagem continua separada da chamada por `de := ...` / `para := ...` (linhas 23-24). Bônus: `"valor invalido: ", err` com `Fprintln` gera espaço duplo (`invalido:  strconv...`) e falta acento.
2. ✅ mensagem de uso conforme o enunciado, exit 2.
3. ✅ binário ignorado (`git check-ignore` confirma `.gitignore:29`).

Veredito: 🔁 só o item 1.

## 2026-11-16, review 3: ✅ aprovado

Ferramentas: gofmt, vet, staticcheck e testes ✅.

- ✅ item 1: `err` checado na linha seguinte ao `ParseFloat`; mensagem `valor inválido: strconv.ParseFloat: parsing "1,5": invalid syntax` (espaço único, acento).

Para depois (não bloqueia): itens 4–8 do review 1 (nomes `celsiusPorF`/`baseCalculo`, `sufixo` longe do uso, NaN/Inf, linhas em branco, "método").

Veredito: ✅ aprovado (3 rodadas; dicas N1 e N2 em `erros-como-valores`).
