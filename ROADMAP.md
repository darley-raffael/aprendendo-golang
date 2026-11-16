# Roadmap: Go

> Gerado em 2026-11-13 para o perfil em `progresso/PERFIL.md` (dev TypeScript, ~5h/semana, modo socrático).
> Toolchain de referência: **Go 1.27**. Plano vivo: ajustamos juntos conforme o ritmo e os objetivos.
> Formato e regras: `METODOLOGIA.md`.

## Visão geral

| Fase | Tema | Projetos | Semanas* |
|---|---|---|---|
| 0 | Primeiros passos: rampa de fundamentos | M01–M06 (micro) → diagnóstico → P00 `gowc` | 3 |
| 1 | Fundamentos pela linha de comando | rampa + P01 `tarefas`, P02 `logstat` | 5 |
| 2 | Sistema de tipos e design idiomático | rampa + P03 `calc` | 5 |
| 3 | Concorrência | rampa + P04 `linkcheck`, P05 `thumbs` | 6 |
| 4 | Backend HTTP e persistência | rampa + P06 `encurta` (API), P07 `encurta` (Postgres + Docker) | 7 |
| 5 | CLIs profissionais e ferramentas DevOps | rampa + P08 `devkit` | 5 |
| 6 | Redes e sistemas distribuídos | rampa + P09 `minikv` (mini Redis) | 7 |
| 7 | Microsserviços e observabilidade | rampa + P10 `encurta` distribuído | 7 |
| 8 | Performance, internals e produção | rampa + P11 otimizar `minikv` | 4 |
| 9 | Carreira e mercado | P12 capstone + entrevistas + open source | 3 + contínuo |

\* A ~5h/semana (≈ 3 sessões de 90 min). **Total ≈ 52 semanas (~12 meses).** Semana corrida: troque a sessão por um `/kata`; o plano absorve isso. Se a rampa de uma fase sair fácil, o diagnóstico libera você direto para o projeto e o prazo encolhe.

**Como a progressão funciona:**
- **Rampa → projeto.** Toda fase abre com **micro-projetos** (uma sessão cada, ~30–60 min) que isolam os conceitos novos, um ou dois por vez. Só depois vem o projeto grande, que **combina** esses conceitos num problema real.
- **Diagnóstico ao fim da rampa.** 3–5 perguntas + análise das dicas usadas. Se uma área estiver fraca, entra um reforço (kata ou micro-projeto extra) **antes** do projeto grande.
- **Introduz → aprofunda.** Cada projeto lista os conceitos que **introduz** e os que **aprofunda**. O mesmo tema volta mais fundo depois (ex.: `slices-basico` na fase 0 → `slices-internals-len-cap` na fase 1 → `escape-analysis` na fase 8).
- **Andaime decrescente.** Nas primeiras micro-projetos o tutor fornece os testes; aos poucos você passa a completá-los e, a partir do P00, escreve todos.

**Fio condutor:** a trilha alterna CLI → concorrência → backend → DevOps → distribuído, e alguns projetos **evoluem** ao longo das fases (o `encurta` nasce como monolito na fase 4 e vira microsserviços na fase 7; o `minikv` nasce na fase 6 e é otimizado na fase 8). Assim você sente na prática o custo e o benefício de cada decisão.

---

## Fase 0: Primeiros passos, rampa de fundamentos (~3 semanas)

**Objetivo:** fixar a sintaxe e os fundamentos de Go em passos pequenos, **um punhado de conceitos por vez**, e descobrir cedo quais áreas precisam de reforço.
**Checkpoint de saída:** explicar a diferença entre módulo e pacote, o que `go build`/`go run`/`go vet`/`gofmt` fazem, zero values, por que erro é valor, e por que Go não tem `while`, `try/catch` nem `class`.

### Rampa: micro-projetos (uma sessão cada)

Cada micro-projeto vive em `projetos/00-rampa/MNN-nome/` e passa por um `/review` curto. Nas micro-projetos M01–M03 os testes são do tutor; M04–M05 vêm com testes parciais para você completar; no M06 você escreve todos.

#### M01. ola: seu primeiro programa
- **O quê:** `ola Maria` → `Olá, Maria!`; sem argumento → `Olá, mundo!`. Gerar o binário e rodá-lo.
- **Conceitos introduzidos:** `toolchain-go`, `modules-e-go-mod`, `pacotes-e-visibilidade`, `fmt-verbos`
- **Ponte com TS:** `go.mod` ≈ `package.json`, sem `node_modules` e sem scripts. `go run` ≈ `ts-node`; `go build` gera **um binário nativo**. Letra maiúscula = `export`.

#### M02. conversor: conversor de unidades
- **O quê:** `conversor 100 c f` → `212.00 °F` (temperatura, distância, peso). Entrada inválida → mensagem clara e exit code ≠ 0.
- **Conceitos introduzidos:** `variaveis-e-zero-values`, `tipos-basicos`, `conversao-de-tipos-explicita`, `constantes`, `strconv`, `erros-como-valores`, `if-com-init`
- **Ponte com TS:** não existe `number`: `int` e `float64` não se misturam sem conversão explícita. `strconv.Atoi` devolve `(valor, err)` em vez de `NaN` ou exceção. Não há `undefined`: toda variável nasce com **zero value**.

#### M03. adivinha: jogo de adivinhar o número
- **O quê:** o programa sorteia de 1 a 100 e lê palpites do stdin, respondendo "maior" ou "menor" até acertar; mostra o número de tentativas e oferece jogar de novo.
- **Conceitos introduzidos:** `for-unico-loop`, `switch`, `funcoes-multiplos-retornos`, `bufio-scanner`, `math-rand`
- **Ponte com TS:** `for` é o único loop (faz o papel de `while` e `for...of`). `switch` não precisa de `break` (não "cai" para o próximo caso) e pode não ter expressão (vira um `if/else` encadeado).

#### M03b. docs: documentando o que você já escreveu
- **O quê:** sem código novo. Escrever a documentação dos pacotes de M01–M03 (`saudacao`, `conversao` e o do `adivinha`): package comments, doc comments de tudo que é exportado, e conferir o resultado com `go doc -all` e com o `pkgsite` rodando localmente. Fechar com um mini-guia pessoal de convenções em `projetos/00-rampa/M03b-docs/NOTAS.md`.
- **Conceitos introduzidos:** `doc-comments`
- **Cobre:** comentário colado na declaração e começando pelo nome (`Package x ...`, `Func ...`); documentar o que **não** sai da assinatura (comportamento em bordas, erros, unidades); sintaxe do Go 1.19+ (parágrafos, listas, blocos de código, links `[strings.TrimSpace]`, `Deprecated:`); `doc.go` para pacotes grandes; regras do `staticcheck`/`revive` sobre comentários.
- **Ponte com TS:** ≈ JSDoc/TSDoc, mas **sem tags** (`@param`, `@returns` não existem): é prosa. A ferramenta (`go doc`, pkg.go.dev) extrai direto do código, sem gerador como o TypeDoc.
- **A partir daqui:** doc comments voltam a ser critério de aceite em todo review. `godoc-e-example-tests` (Fase 2) aprofunda com exemplos executáveis.

#### M04. boletim: estatísticas de notas
- **O quê:** recebe notas por argumento ou stdin e mostra média, mediana, maior, menor e situação (aprovado, recuperação, reprovado). Lógica em funções puras, separadas do `main`.
- **Conceitos introduzidos:** `slices-basico`, `range`, `funcoes-variadicas`, `testing-basico`
- **Testes:** parciais do tutor; você adiciona os casos de borda (lista vazia, nota única, número par de notas).
- **Ponte com TS:** slice ≈ array dinâmico, mas sem `.map`/`.filter` embutidos: você escreve o loop (e é idiomático). `go test` já vem na toolchain, sem Jest.

#### M05. agenda: contatos em memória
- **O quê:** um REPL simples (`add`, `find`, `list`, `rm`) com contatos (nome, telefone, email) guardados em memória, com busca por nome.
- **Conceitos introduzidos:** `maps-basico`, `structs-basico`, `metodos-intro`, `ponteiros-basico`, `table-driven-tests`
- **Testes:** parciais do tutor no formato table-driven; você completa.
- **Ponte com TS:** struct é **valor**: passar para uma função **copia** (diferente de objeto em JS). Ponteiro é a forma explícita de compartilhar. Mapa ≈ `Map`, mas a leitura de chave inexistente retorna o zero value (use `v, ok := m[k]`).

#### M06. senha: gerador e validador de senhas
- **O quê:** `senha gerar -n 16 --simbolos` e `senha validar "abc"` (regras: tamanho, maiúscula, dígito, símbolo; relatório do que falta). Flags com valores padrão e `-h`.
- **Conceitos introduzidos:** `os-args-e-flag`, `pacote-strings`, `runes-e-utf8`, `crypto-rand`
- **Testes:** todos seus.
- **Ponte com TS:** string em Go é uma sequência de **bytes** (UTF-8), não de caracteres: `len("ação")` é 5. `range` numa string itera **runes**.

### Diagnóstico da rampa
Antes do P00, o tutor faz 5 perguntas cobrindo M01–M06 e analisa as dicas e os reviews da rampa. O resultado vai para "Observações do tutor" no `PERFIL.md`. Se um conceito ficou abaixo do nível 2, entra **1 kata de reforço** antes de seguir.

### P00. gowc: clone do `wc` do Unix (projeto de consolidação)
- **Por quê:** junta tudo da rampa num programa de verdade e acrescenta leitura de arquivos e streams.
- **Etapas:**
  1. `gowc arquivo.txt` imprime linhas, palavras e bytes.
  2. Flags `-l -w -c -m` (runes), vários arquivos com total, leitura de stdin quando não há arquivo.
  3. Testes table-driven das funções de contagem, escritos por você.
- **Conceitos introduzidos:** `os-open-e-arquivos`, `defer`, `io-reader-intro`
- **Aprofunda:** `bufio-scanner`, `erros-como-valores`, `os-args-e-flag`, `runes-e-utf8`, `table-driven-tests`
- **Critérios de aceite:** `gofmt -l .` vazio; `go vet ./...` limpo; testes para arquivo vazio, sem newline final e UTF-8 (`-m` ≠ `-c` para "ação"); arquivos sempre fechados.
- **Extensões:** `-L` (linha mais longa); comparar a saída com o `wc` real num script de teste.
- **Ponte com TS:** `defer f.Close()` ≈ `finally`, mas ligado à função. Ler arquivo linha a linha sem carregar tudo na memória é o padrão, não a exceção.

---

## Fase 1: Fundamentos pela linha de comando (~4 semanas)

**Objetivo:** escrever programas Go de verdade com structs, métodos, coleções, JSON, erros bem tratados e testes.
**Checkpoint de saída:** entrevista sobre ponteiro vs valor, slices (len/cap/append), mapas, wrapping de erros e por que interfaces de I/O são tão centrais.

### Rampa
- **R1.1 `json-ida-e-volta`:** serializar e desserializar uma struct aninhada, com tags, campos omitidos e campos desconhecidos (`encoding-json-e-struct-tags`)
- **R1.2 `erros-em-camadas`:** três funções que embrulham erros; no topo, decidir com `errors.Is/As` (`wrapping-de-erros`, `errors-is-as`)
- **R1.3 `ponteiro-ou-valor`:** o mesmo tipo com receivers de valor e de ponteiro; prever o que muda (`metodos-e-receivers`)
- **Diagnóstico** → P01

### P01. tarefas: gerenciador de tarefas em CLI com persistência em JSON
- **Por quê:** é um CRUD que você conhece de cor, o que libera a cabeça para focar **só** em como Go modela dados.
- **Etapas:**
  1. Modelo `Tarefa` + operações em memória (`add`, `list`, `done`, `rm`) com subcomandos.
  2. Persistência em arquivo JSON (`~/.tarefas.json`), com escrita atômica (arquivo temporário + rename).
  3. Filtros (`--pendentes`, `--tag`), ordenação por prioridade/data, saída em tabela alinhada.
  4. Pacote interno `internal/store` com testes, incluindo `t.TempDir()` para os testes de arquivo.
- **Conceitos introduzidos:** `metodos-e-receivers`, `slices-internals-len-cap`, `encoding-json-e-struct-tags`, `time-pacote`, `wrapping-de-erros`, `errors-is-as`, `erros-sentinela-e-tipados`, `layout-cmd-internal`, `escrita-atomica-de-arquivos`
- **Aprofunda:** `structs-basico`, `ponteiros-basico`, `slices-basico`, `maps-basico`, `table-driven-tests`, `defer`
- **Critérios de aceite:** nenhum erro ignorado (`_ =` só com justificativa); receivers consistentes; testes do store com cobertura ≥ 70%; erro amigável para JSON corrompido.
- **Extensões:** `tarefas export --csv`; flag `--json` para saída legível por máquina.
- **Ponte com TS:** struct ≠ class (sem herança, sem construtor; use a função `NewX` por convenção). **Armadilhas:** `append` pode ou não realocar (aliasing de slices); iterar mapa não tem ordem; JSON só serializa campos **exportados**. `errors.Is` ≈ `instanceof`, mas para cadeias de wrap.

### P02. logstat: analisador de logs de acesso (formato nginx/combined)
- **Por quê:** processar arquivos grandes em streaming é onde Go brilha. Força você a pensar em `io.Reader`, alocação e interfaces implícitas.
- **Etapas:**
  1. Ler o log em streaming e fazer o parse de cada linha numa struct; contar status, top-N paths e IPs.
  2. Abstrair a fonte: arquivo, stdin, `.gz` (`compress/gzip`), **tudo como `io.Reader`**.
  3. Formatos de saída plugáveis (texto, JSON, markdown) via interface `Reporter`.
  4. Benchmark do parser: medir alocações e otimizar (ex.: evitar `regexp` no hot path).
- **Conceitos introduzidos:** `interfaces-implicitas`, `io-reader-writer`, `composicao-de-interfaces`, `strings-e-bytes-builder`, `pacote-slices-e-sort`, `closures-e-funcoes-de-primeira-classe`, `benchmarks-basico`, `type-assertion`
- **Aprofunda:** `io-reader-intro`, `bufio-scanner`, `pacote-strings`
- **Critérios de aceite:** processa um log de 1 GB com memória constante (`/usr/bin/time -v`); benchmark com `b.ReportAllocs()`; interfaces pequenas, definidas **no consumidor**.
- **Extensões:** `--follow` (tail -f); histograma de latência no terminal.
- **Ponte com TS:** interface em Go é **estrutural e implícita** (parecido com TS!), mas é checada em runtime para type assertions e não existe `implements`. Prefira interfaces de 1–2 métodos. `io.Reader` ≈ Node streams, porém síncrono e muito mais simples.

### Katas da fase
- `rot13-reader`: implementar um `io.Reader` que embrulha outro (`io-reader-writer`)
- `frequencia-palavras`: top-K palavras com mapa + `slices.SortFunc` (`maps-basico`, `pacote-slices-e-sort`)
- `slice-aliasing`: prever e corrigir bugs de `append` compartilhando array (`slices-internals-len-cap`)
- `validar-cpf`: funções puras + table-driven tests (`table-driven-tests`)

---

## Fase 2: Sistema de tipos e design idiomático (~4 semanas)

**Objetivo:** dominar o sistema de tipos (interfaces a fundo, embedding, generics, iterators) e saber **quando** usar cada coisa.
**Checkpoint de saída:** explicar nil interface vs interface contendo nil, method sets, quando generics vale a pena, e desenhar uma API de pacote idiomática.

### Rampa
- **R2.1 `formas`:** `Forma` com `Area()`/`Perimetro()`, implementada por três tipos; somar áreas de uma `[]Forma` e identificar o tipo com type switch (`type-switch`, `method-sets`)
- **R2.2 `dias-da-semana`:** enum com `iota` + `String()` + parse de string para o enum (`iota-e-enums`, `stringer-interface`)
- **R2.3 `pilha-generica`:** `Stack[T]` com `Push/Pop/Peek` e testes para dois tipos (`generics-type-params`)
- **Diagnóstico** → P03

### P03. calc: interpretador de uma linguagem de expressões
- **Por quê:** lexer → parser → AST → avaliador é o projeto clássico para exercitar tipos, interfaces, type switch e erros ricos. Ele também prepara o terreno para o parser de protocolo do `minikv`.
- **Etapas:**
  1. **Lexer**: tokens com `iota`, `Stringer`, posição (linha/coluna), runes e UTF-8.
  2. **Parser** (descida recursiva/Pratt) gerando a AST; `interface Node` + type switch.
  3. **Avaliador** com variáveis, funções embutidas (`max`, `sqrt`) e erros com posição (tipo de erro customizado).
  4. **REPL** interativo + pacote `calc` com API pública documentada (`go doc`) e exemplos testáveis (`Example...`).
  5. Utilitários genéricos próprios (`Stack[T]`, `Map/Filter`) e um **iterator** (`iter.Seq`) sobre os tokens.
- **Conceitos introduzidos:** `iota-e-enums`, `stringer-interface`, `type-switch`, `embedding`, `method-sets`, `nil-interface-gotcha`, `generics-type-params`, `generics-constraints`, `iterators-range-over-func`, `custom-error-types`, `godoc-e-example-tests`, `panic-recover`, `fuzzing-basico`
- **Aprofunda:** `runes-e-utf8`, `interfaces-implicitas`, `erros-sentinela-e-tipados`, `closures-e-funcoes-de-primeira-classe`
- **Critérios de aceite:** fuzz test do lexer/parser rodando ≥ 1 min sem panic; erros indicam linha/coluna; API pública mínima (o resto é não exportado); `Example` tests passando.
- **Extensões:** `if`/funções definidas pelo usuário; compilar para bytecode e rodar numa VM de pilha.
- **Ponte com TS:** não há union types: use interface + type switch (≈ discriminated union, sem exhaustiveness check). Generics em Go são bem mais limitados que em TS (sem conditional/mapped types) **de propósito**. `panic` não é `throw`: é para bug, não para fluxo de controle.

### Katas da fase
- `lru-generico`: `LRU[K comparable, V any]` com `container/list` (`generics-constraints`)
- `nil-interface`: 5 snippets "isso é nil?" (`nil-interface-gotcha`)
- `iter-fibonacci`: `iter.Seq[int]` infinito + `take` (`iterators-range-over-func`)
- `embedding-logger`: estender um tipo via embedding e entender o shadowing de métodos (`embedding`)

---

## Fase 3: Concorrência (~5 semanas)

**Objetivo:** pensar em goroutines e channels com segurança: cancelamento, backpressure, sem vazamentos e sem races.
**Checkpoint de saída:** desenhar no quadro um worker pool com cancelamento; explicar race vs deadlock vs leak; quando usar channel vs mutex; o que o `context` propaga.

### Rampa
- **R3.1 `ping-pong`:** duas goroutines trocando mensagens por um channel N vezes e encerrando de forma limpa (`goroutines`, `channels`)
- **R3.2 `contador-com-race`:** 1000 goroutines incrementando um contador; ver o race detector acusar e corrigir de duas formas, mutex e atomic (`race-detector`, `sync-mutex`)
- **R3.3 `relogio`:** `select` com `time.Ticker`, timeout e cancelamento por `context` (`select`, `context-cancelamento-timeout`)
- **Diagnóstico** → P04

### P04. linkcheck: verificador de links concorrente (crawler)
- **Por quê:** I/O de rede em massa é o caso de uso perfeito para goroutines, e é fácil de "explodir" sem controle, o que torna os limites necessários.
- **Etapas:**
  1. Versão sequencial: baixa uma página, extrai links (`golang.org/x/net/html`) e checa o status de cada um.
  2. Versão concorrente ingênua (uma goroutine por link) → medir e ver o problema.
  3. **Worker pool** com channels + `sync.WaitGroup`; limite de concorrência configurável.
  4. **Context**: timeout global e por request, cancelamento por Ctrl+C (`signal.NotifyContext`).
  5. Crawl recursivo até a profundidade N, sem visitar duas vezes (mapa protegido) e rate limit por host.
- **Conceitos introduzidos:** `goroutines`, `channels`, `buffered-vs-unbuffered`, `select`, `sync-waitgroup`, `worker-pool`, `context-cancelamento-timeout`, `net-http-client`, `sync-mutex`, `race-detector`, `rate-limiting`, `goroutine-leak`
- **Aprofunda:** `closures-e-funcoes-de-primeira-classe` (captura de variável em goroutines), `maps-basico` (mapas não são thread-safe), `wrapping-de-erros`
- **Critérios de aceite:** `go test -race ./...` limpo; nenhuma goroutine viva após o fim (verificar com `runtime.NumGoroutine` em teste ou com `goleak`); Ctrl+C encerra em menos de 1 s; os testes usam `httptest.Server` (sem internet).
- **Extensões:** relatório em JSON; respeitar `robots.txt`; retry com backoff exponencial + jitter.
- **Ponte com TS:** esqueça o event loop: goroutines são **paralelas de verdade** e bloqueiam sem medo. Não existe `Promise.all`: você compõe com WaitGroup/channels/errgroup. **Armadilha:** em Node, um objeto compartilhado é seguro porque só existe uma thread. Em Go **não é**: race condition é real.

### P05. thumbs: pipeline de processamento de imagens
- **Por quê:** consolida padrões avançados (pipeline, fan-out/fan-in, errgroup, semáforo, shutdown gracioso) num problema CPU-bound + I/O.
- **Etapas:**
  1. Pipeline em estágios (descobrir arquivos → decodificar → redimensionar → salvar) conectados por channels.
  2. Fan-out no estágio CPU-bound (N = `GOMAXPROCS`), fan-in dos resultados, propagação de erro com `errgroup`.
  3. Progresso em tempo real (contadores `sync/atomic`), shutdown gracioso que termina o que está em andamento.
  4. Testes determinísticos de concorrência com `testing/synctest`; benchmark comparando paralelismos.
- **Conceitos introduzidos:** `pipeline-pattern`, `fan-out-fan-in`, `errgroup`, `semaforo-com-channel`, `sync-atomic`, `sync-once`, `sync-rwmutex`, `graceful-shutdown`, `gomaxprocs-e-scheduler`, `testing-synctest`, `memory-model-basico`
- **Critérios de aceite:** um erro num arquivo não derruba o lote (relatório de falhas) **ou** derruba tudo com `--fail-fast`, à sua escolha e justificado; sem races; o benchmark mostra o speedup.
- **Extensões:** cache de resultados por hash do conteúdo; modo watch (fsnotify).
- **Ponte com TS:** isso seria worker_threads + muito boilerplate em Node. Em Go é a linguagem. "Don't communicate by sharing memory; share memory by communicating."

### Katas da fase
- `filosofos`: jantar dos filósofos sem deadlock (`sync-mutex`)
- `pubsub`: broker pub/sub em memória com unsubscribe sem leak (`channels`, `goroutine-leak`)
- `primeiro-que-responder`: consultar 3 réplicas e usar a resposta mais rápida, cancelando as outras (`context-cancelamento-timeout`, `select`)
- `bounded-parallel`: `ParallelMap` genérico com limite de concorrência (`semaforo-com-channel`, `generics-type-params`)

---

## Fase 4: Backend HTTP e persistência (~6 semanas)

**Objetivo:** construir uma API de produção com a stdlib: roteamento, middlewares, validação, logs estruturados, banco, migrations, testes de integração e container.
**Checkpoint de saída:** explicar o ciclo de vida de um request em `net/http`, onde o `context` entra, como estruturar camadas sem "Java em Go" e como testar handlers e repositórios.

### Rampa
- **R4.1 `ola-http`:** servidor com duas rotas JSON (`GET /saude`, `POST /eco`) usando os padrões do `ServeMux` (`net-http-server`, `servemux-patterns`)
- **R4.2 `middleware-de-log`:** um middleware que mede a duração e loga com `slog` (`http-handler-interface`, `middlewares`, `log-slog`)
- **R4.3 `testar-handler`:** testar o R4.1 com `httptest` sem subir porta (`httptest`)
- **Diagnóstico** → P06 (rampa curta antes do P07: **R4.4 `sql-cru`**, CRUD de uma tabela com `pgx` num Postgres em Docker, `database-sql-e-pgx`)

### P06. encurta (v1): API de encurtador de URLs só com a stdlib
- **Por quê:** a stdlib de Go é suficiente para APIs sérias. Aprender sem framework mostra o que os frameworks escondem.
- **Etapas:**
  1. `http.ServeMux` com padrões de método/rota (`POST /links`, `GET /{code}`), JSON in/out e status corretos.
  2. Camadas: handler → service → repository (interface), com storage em memória. Injeção de dependência manual no `main`.
  3. Middlewares: logging com `log/slog` (request id), recover, timeout, CORS. Autenticação por API key/JWT.
  4. Validação, erros de API padronizados (RFC 9457 problem+json) e configuração por env.
  5. Testes com `httptest` (handlers) + servidor com timeouts configurados e shutdown gracioso.
- **Conceitos introduzidos:** `net-http-server`, `servemux-patterns`, `http-handler-interface`, `middlewares`, `log-slog`, `context-em-requests`, `injecao-de-dependencia-manual`, `interfaces-no-consumidor`, `httptest`, `configuracao-por-env`, `server-timeouts`, `jwt-auth`
- **Aprofunda:** `encoding-json-e-struct-tags`, `custom-error-types`, `layout-cmd-internal`, `graceful-shutdown`, `context-cancelamento-timeout`
- **Critérios de aceite:** timeouts de servidor definidos (`ReadHeaderTimeout` etc.); nenhum handler sem teste; logs em JSON com request id; `go vet` + `staticcheck` limpos.
- **Extensões:** contagem de cliques; rate limit por API key; documentação OpenAPI.
- **Ponte com TS:** `http.Handler` ≈ middleware do Express, mas tipado e compondo funções. **Armadilha:** não traga estruturas de NestJS (pastas `controllers/`, `services/` globais, DI container). Em Go, organize **por domínio** e mantenha o `main` explícito.

### P07. encurta (v2): Postgres, migrations, Docker e testes de integração
- **Por quê:** persistência real, transações e um ambiente reproduzível são o dia a dia de um backend Go.
- **Etapas:**
  1. Postgres com `pgx` (pool), repository real, migrations versionadas (`goose` ou `golang-migrate`).
  2. Queries type-safe com `sqlc`; transações (criar link + registrar auditoria) e tratamento de erros do driver (violação de unique).
  3. Testes de integração com `testcontainers-go`; build tags para separar unit e integration.
  4. Dockerfile multi-stage (imagem distroless/scratch < 30 MB), `docker compose` com app + db, health check.
  5. Cache de leitura com Redis (cache-aside) + invalidação.
- **Conceitos introduzidos:** `database-sql-e-pgx`, `connection-pool`, `migrations`, `sqlc`, `transacoes`, `testes-de-integracao`, `build-tags`, `docker-multi-stage`, `cross-compilation`, `cache-aside`
- **Critérios de aceite:** testes de integração passando em container; nenhuma query concatenando strings; pool configurado; imagem final sem toolchain.
- **Extensões:** paginação por cursor; CI no GitHub Actions rodando os testes de integração.
- **Ponte com TS:** não há um "Prisma" dominante: a comunidade prefere SQL explícito (`sqlc`) a ORMs pesados. Binário estático = imagem Docker minúscula, sem `node_modules`.

### Katas da fase
- `middleware-chain`: compor middlewares sem biblioteca (`middlewares`, `closures-e-funcoes-de-primeira-classe`)
- `problem-json`: helper genérico de resposta de erro (`http-handler-interface`)
- `tx-helper`: `WithTx(ctx, func(tx) error)` com rollback em erro e panic (`transacoes`, `panic-recover`)

---

## Fase 5: CLIs profissionais e ferramentas DevOps (~4 semanas)

**Objetivo:** construir e distribuir uma ferramenta de linha de comando nível open source, que conversa com Docker, Kubernetes e APIs.
**Checkpoint de saída:** explicar como uma CLI Go é distribuída (cross-compile, release, versão embutida) e como se integra com APIs de infra.

### Rampa
- **R5.1 `cobra-ola`:** CLI com dois subcomandos e uma flag persistente (`cobra-cli`)
- **R5.2 `versao-embutida`:** binário que imprime a versão injetada via `-ldflags` e um template embutido com `go:embed` (`ldflags-versao`, `go-embed`)
- **R5.3 `rodar-comando`:** executar `git status` via `os/exec` com timeout e tratar o exit code (`os-exec`)
- **Diagnóstico** → P08

### P08. devkit: canivete DevOps em CLI
- **Por quê:** a maior parte do ecossistema cloud-native (Docker, Kubernetes, Terraform, gh) é CLI em Go. Este projeto te coloca nesse mundo.
- **Etapas:**
  1. Estrutura com `cobra` (subcomandos, flags persistentes, autocompletion) e config com arquivo + env + flags (precedência clara).
  2. `devkit docker ps|clean`: SDK do Docker (containers, imagens órfãs), com saída em tabela/JSON/YAML.
  3. `devkit k8s pods`: `client-go` (listar pods com problemas, logs, `--watch`).
  4. `devkit gh prs`: API do GitHub com paginação, retries e token via env/keyring; `os/exec` para integrar com o `git` local.
  5. Release: `-ldflags` com versão, `go:embed` para templates, `goreleaser` (binários para linux/mac/windows) e `go install`.
  6. Opcional: modo TUI interativo com `bubbletea`.
- **Conceitos introduzidos:** `cobra-cli`, `config-precedencia`, `os-exec`, `go-embed`, `ldflags-versao`, `goreleaser`, `sdk-docker`, `client-go`, `paginacao-e-retry`, `go-generate`, `tool-directives-go-mod`
- **Critérios de aceite:** `devkit --help` impecável; códigos de saída corretos; funciona com stdout não-TTY (pipe); release gerada localmente pelo `goreleaser --snapshot`.
- **Extensões:** plugin system (binários `devkit-*` no PATH, como o kubectl).
- **Ponte com TS:** é a mesma ideia de `commander`/`oclif`, mas o resultado é **um binário estático** sem runtime, e é por isso que ferramentas DevOps são escritas em Go.

### Katas da fase
- `table-printer`: `text/tabwriter` + saída `--output json|yaml|table` (`io-reader-writer`)
- `exec-timeout`: rodar um comando externo com timeout e capturar stdout/stderr (`os-exec`, `context-cancelamento-timeout`)

---

## Fase 6: Redes e sistemas distribuídos (~6 semanas)

**Objetivo:** programar no nível de TCP, protocolos e concorrência de servidor: o tipo de código que existe dentro de Redis, NATS e etcd.
**Checkpoint de saída:** explicar como seu servidor lida com 10k conexões, como a expiração de TTL funciona, o que acontece num crash (durabilidade) e o trade-off da replicação assíncrona.

### Rampa
- **R6.1 `eco-tcp`:** servidor TCP de eco com uma goroutine por conexão, testado com `nc` (`net-tcp-server`)
- **R6.2 `deadline`:** derrubar conexões ociosas após N segundos (`deadlines-de-conexao`)
- **R6.3 `protocolo-linha`:** um mini protocolo `CMD arg\r\n` com parser e respostas tipadas (`protocolo-binario-e-parsing`, `bufio-reader-writer-de-rede`)
- **Diagnóstico** → P09

### P09. minikv: um mini Redis
- **Por quê:** é um projeto de portfólio que impressiona em entrevistas e junta tudo: parser (fase 2), concorrência (fase 3) e arquitetura de servidor.
- **Etapas:**
  1. Servidor TCP (`net.Listen`), uma goroutine por conexão, parser do protocolo **RESP**, comandos `PING/SET/GET/DEL`. Compatível com `redis-cli`.
  2. Store concorrente (mutex vs sharding de locks: medir); `EXPIRE/TTL` com expiração preguiçosa + varredura ativa.
  3. Durabilidade: AOF (append-only file) com `fsync` configurável e replay na inicialização; snapshot opcional.
  4. Pub/Sub (`SUBSCRIBE/PUBLISH`) com backpressure para clientes lentos.
  5. Replicação líder → seguidor (stream de comandos), reconexão e comando `INFO`.
  6. Teste de carga com `redis-benchmark` e comparação com o Redis real.
- **Conceitos introduzidos:** `net-tcp-server`, `protocolo-binario-e-parsing`, `bufio-reader-writer-de-rede`, `lock-sharding`, `expiracao-ttl`, `durabilidade-aof-fsync`, `backpressure`, `replicacao-lider-seguidor`, `deadlines-de-conexao`, `testes-de-carga`
- **Critérios de aceite:** `redis-cli` e `redis-benchmark` funcionam; sem races sob carga; o crash recovery funciona (matar com `kill -9` e reiniciar sem perder dados confirmados com fsync=always); conexões ociosas têm deadline.
- **Extensões:** cluster com consenso via `hashicorp/raft`; `MULTI/EXEC`.
- **Ponte com TS:** em Node seria um único event loop. Aqui, cada conexão é uma goroutine bloqueante e o runtime faz a multiplexação (netpoller). O código fica sequencial e simples.

### Katas da fase
- `resp-parser`: parse/serialize de RESP com fuzzing (`protocolo-binario-e-parsing`, `fuzzing-basico`)
- `chat-tcp`: servidor de chat TCP com broadcast (`net-tcp-server`, `channels`)
- `consistent-hashing`: anel de hash consistente com nós virtuais (`maps-basico`, `pacote-slices-e-sort`)

---

## Fase 7: Microsserviços e observabilidade (~6 semanas)

**Objetivo:** decompor um sistema em serviços que se comunicam por gRPC e mensageria, com resiliência e observabilidade de produção.
**Checkpoint de saída:** system design oral: "desenhe o encurtador para 10k req/s", cobrindo consistência, idempotência, outbox, tracing distribuído e falhas parciais.

### Rampa
- **R7.1 `grpc-ola`:** definir um `.proto`, gerar com `buf` e fazer chamada unária com deadline (`grpc-e-protobuf`)
- **R7.2 `nats-ola`:** publisher e consumer com JetStream e ack manual (`mensageria-nats-kafka`)
- **R7.3 `trace-local`:** instrumentar o R4.1 com OpenTelemetry e ver o trace no Jaeger (`opentelemetry-tracing`)
- **Diagnóstico** → P10

### P10. encurta (v3): de monolito a microsserviços
- **Por quê:** você só entende o custo de microsserviços partindo do monolito que **você mesmo** escreveu (P06/P07).
- **Etapas:**
  1. Extrair o serviço `analytics` (cliques) e conectar via **gRPC** + protobuf (`buf`), com deadlines e interceptors.
  2. Eventos assíncronos com **NATS JetStream** (ou Kafka): `link.clicked` → analytics. Consumidor idempotente.
  3. **Outbox pattern** para publicar eventos de forma transacional; retries com backoff; dead letter.
  4. Resiliência: circuit breaker, timeouts em cascata, health/readiness probes, graceful shutdown dos consumidores.
  5. **OpenTelemetry**: traces distribuídos HTTP → gRPC → fila, métricas Prometheus (RED), logs correlacionados por trace id; Grafana/Jaeger via compose.
  6. Deploy em Kubernetes local (kind/k3d) com manifests ou Helm.
- **Conceitos introduzidos:** `grpc-e-protobuf`, `grpc-interceptors`, `mensageria-nats-kafka`, `idempotencia`, `outbox-pattern`, `retry-backoff-jitter`, `circuit-breaker`, `health-checks`, `opentelemetry-tracing`, `metricas-prometheus`, `deploy-kubernetes`
- **Critérios de aceite:** trace único atravessando os 3 saltos; derrubar o analytics não derruba o redirect; mensagens duplicadas não duplicam contagem; dashboards com latência p99.
- **Extensões:** API gateway; saga para um fluxo multi-serviço; testes de contrato.
- **Ponte com TS:** gRPC em Go é cidadão de primeira classe (o ecossistema nasceu junto). **Armadilha:** não crie um "shared lib" gigante entre os serviços. Compartilhe só os `.proto`.

### Katas da fase
- `idempotency-key`: middleware de idempotência com armazenamento (`idempotencia`, `middlewares`)
- `breaker`: circuit breaker com estados e testes com `synctest` (`circuit-breaker`, `testing-synctest`)

---

## Fase 8: Performance, internals e produção (~3 semanas)

**Objetivo:** medir antes de otimizar e entender o runtime (alocação, escape analysis, GC, scheduler) o bastante para tomar decisões com dados.
**Checkpoint de saída:** ler um profile de CPU/heap e explicar a causa; explicar escape analysis, o GC tri-color concorrente, `GOGC`/`GOMEMLIMIT` e o modelo G-M-P.

### Rampa
- **R8.1 `concat-bench`:** comparar `+`, `fmt.Sprintf`, `strings.Builder` e `bytes.Buffer` com benchmarks e explicar os números (`benchmarks-basico`, `escape-analysis`)
- **R8.2 `primeiro-pprof`:** gerar e ler um profile de CPU de um programa pequeno (`pprof`)
- **Diagnóstico** → P11

### P11. Otimizar o `minikv`
- **Por quê:** otimizar o seu próprio servidor sob carga real ensina mais do que qualquer artigo.
- **Etapas:**
  1. Baseline: benchmarks + `redis-benchmark`; perfis com `pprof` (CPU, heap, allocs, mutex, block) e `go tool trace`.
  2. Reduzir alocações no hot path (parser e respostas): escape analysis (`-gcflags=-m`), `sync.Pool`, reuso de buffers.
  3. Ajustar GC (`GOGC`, `GOMEMLIMIT`) e medir p99; experimentar **PGO** com um perfil de produção.
  4. Hardening: `golangci-lint` com config séria, `govulncheck`, fuzzing contínuo do parser e CI completo (lint, test, race, vuln, build multiplataforma).
- **Conceitos introduzidos:** `pprof`, `execution-tracer`, `escape-analysis`, `sync-pool`, `gc-gogc-gomemlimit`, `scheduler-gmp`, `pgo`, `golangci-lint`, `govulncheck`, `ci-github-actions`
- **Critérios de aceite:** relatório antes/depois com números (ns/op, allocs/op, p99); cada otimização justificada por um profile, nenhuma "no achismo".
- **Extensões:** comparar `map` + mutex vs `sync.Map` vs sharding sob diferentes cargas; ler o código do runtime do `map`.

### Katas da fase
- `zero-alloc`: reescrever uma função de formatação para 0 allocs/op (`escape-analysis`, `benchmarks-basico`)
- `profile-detective`: achar o gargalo num programa que o tutor fornece (`pprof`)

---

## Fase 9: Carreira e mercado (~3 semanas + contínuo)

**Objetivo:** transformar o aprendizado em empregabilidade: portfólio, fluência em entrevistas e contribuição open source.
**Checkpoint de saída:** simulado de entrevista completo (Go + concorrência + system design) feito com o tutor.

### P12. Capstone e preparação
- **Etapas:**
  1. **Portfólio:** polir 3 repos (sugestão: `minikv`, `encurta`, `devkit`) com README, arquitetura, benchmarks e CI verde.
  2. **Simulados de entrevista** (`/revisar fase N` em modo entrevista): slices/maps internals, interfaces, concorrência, context, erros, generics, GC, live coding de 30 min.
  3. **Leitura de código de referência:** `net/http` (server), `context`, `sync.Mutex`, e um projeto como `chi`, `nats-server` ou `etcd`. Explicar um trecho ao tutor.
  4. **Open source:** contribuir com 1–2 PRs (good first issue em projetos Go do ecossistema CNCF ou em ferramentas que você usa).
- **Conceitos introduzidos:** `leitura-de-codigo-stdlib`, `system-design-com-go`, `contribuicao-open-source`
- **Critérios de aceite:** 3 repos públicos apresentáveis; 1 PR aberto num projeto open source; simulado de entrevista aprovado.

---

## Recursos oficiais e de referência

- [A Tour of Go](https://go.dev/tour): use como consulta, não como curso
- [Effective Go](https://go.dev/doc/effective_go) e [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments): **base do `/review`**
- [Google Go Style Guide](https://google.github.io/styleguide/go/)
- [The Go Programming Language Specification](https://go.dev/ref/spec) e [The Go Memory Model](https://go.dev/ref/mem)
- [Go Blog](https://go.dev/blog), em especial os artigos sobre slices, errors, context, pipelines, generics, iterators e PGO
- Livros: *Learning Go* (Jon Bodner, 2ª ed.), *Concurrency in Go* (Katherine Cox-Buday), *100 Go Mistakes and How to Avoid Them* (Teiva Harsanyi)
- Prática extra: [Exercism Go track](https://exercism.org/tracks/go), que você já tem instalado em `~/exercism`
