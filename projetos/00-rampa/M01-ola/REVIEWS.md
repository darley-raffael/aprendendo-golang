# Reviews: M01 ola

## 2026-11-13, review 1: 🔁 refatorar

### Ferramentas
- `gofmt -l .`: vazio ✅
- `go vet ./...`: ok ✅
- `staticcheck`: não instalado ⚠️
- `go test -cover ./...`: ok; `saudacao` 100%, `main` 0% (esperado)

### Critérios de aceite
- ✅ `./ola Maria` → `Olá, Maria!`
- ✅ `./ola` → `Olá, mundo!`
- ✅ `./ola Maria João` → `Olá, Maria!`
- ✅ regra em `saudacao.Ola`, `main` fino
- ❌ caminho do módulo: `go.mod` declara `ola`, não `github.com/darley-raffael/aprendendo-golang/projetos/00-rampa/M01-ola`
- ✅ nada exportado além do necessário

### Itens
1. 🟡 `go.mod:1`: caminho de módulo `ola` (sem ponto no primeiro elemento é reservado para a stdlib; não é importável de fora; diverge do enunciado).
2. 🟡 `saudacao/sudacao.go:8`: `strings.Trim(name, "")` não faz nada (cutset vazio). `./ola "   "` → `Olá,    !`.
3. 🟡 `main.go:10-20`: `fmt.Println(saudacao.Ola(...))` duplicado + early return desnecessário; mistura `args` e `os.Args`.
4. 🟡 `saudacao/sudacao.go:7`: função exportada sem doc comment.
5. 🔵 nome do arquivo `sudacao.go` (typo).
6. ⚪ binário `ola` na pasta do projeto: não versionar.
7. ⚪ import único entre parênteses: o normal é `import "strings"`.

### Pontos fortes
- Uso consciente do zero value de `string` como "sem nome".
- `Ola` é pura e o `main` só faz I/O: boa separação já no primeiro programa.
- Testes verdes de primeira, gofmt limpo.

### Veredito
🔁 Refatorar itens 1–5.

## 2026-11-13, review 2: 🔁 refatorar

Ferramentas: gofmt, vet e testes ✅ (staticcheck ainda ausente).

1. ✅ caminho do módulo corrigido; import ajustado e agrupado à parte da stdlib.
2. 🟡 parcial: `strings.Trim(name, " ")` remove só espaço; `./ola "<TAB>Ana"` → `Olá, <TAB>Ana!`. Existe função específica para whitespace.
3. ✅ `main` com um único `Println`, sem early return.
4. ❌ `Ola` continua sem doc comment (`go doc ./saudacao Ola` não mostra descrição).
5. ✅ arquivo renomeado para `saudacao.go`.

Veredito: 🔁 faltam itens 2 e 4.

## 2026-11-13, review 3: ✅ aprovado

Ferramentas: gofmt, vet e testes ✅ (staticcheck ainda ausente).

- ✅ item 2: `strings.TrimSpace`.
- ✅ item 4: doc comment começa com `Ola` e é exibido no `go doc`.

Para depois (não bloqueia):
- 🔵 o doc de `Ola` não diz o que acontece com nome vazio/só espaços (a regra que o chamador não adivinha).
- 🔵 package comment deve começar com `Package saudacao ...`; ele diz "imprime", mas o pacote **retorna** a string (quem imprime é o `main`); typos "fornce"/"funçào".
- ⚪ `import ( "strings" )` → `import "strings"`.
- ⚪ binário `ola` na pasta: não versionar.

Veredito: ✅ aprovado (3 rodadas, 0 dicas).
