# Roadmap: Go

> Gerado em 2026-11-13 para o perfil em `progresso/PERFIL.md` (dev TypeScript, ~5h/semana, modo socrático).
> Toolchain de referência: **Go 1.27**. Plano vivo: ajustamos juntos conforme o ritmo e os objetivos.
> Formato e regras: `METODOLOGIA.md`.

## Visão geral

| Fase | Tema | Projetos | Semanas* |
|---|---|---|---|
| 0 | Toolchain e primeiro contato | P00 `gowc` | 1 |
| 1 | Fundamentos pela linha de comando | P01 `tarefas`, P02 `logstat` | 4 |
| 2 | Sistema de tipos e design idiomático | P03 `calc` | 4 |
| 3 | Concorrência | P04 `linkcheck`, P05 `thumbs` | 5 |
| 4 | Backend HTTP e persistência | P06 `encurta` (API), P07 `encurta` (Postgres + Docker) | 6 |
| 5 | CLIs profissionais e ferramentas DevOps | P08 `devkit` | 4 |
| 6 | Redes e sistemas distribuídos | P09 `minikv` (mini Redis) | 6 |
| 7 | Microsserviços e observabilidade | P10 `encurta` distribuído | 6 |
| 8 | Performance, internals e produção | P11 otimizar `minikv` | 3 |
| 9 | Carreira e mercado | P12 capstone + entrevistas + open source | 3 + contínuo |

\* A ~5h/semana (≈ 3 sessões de 90 min). **Total ≈ 42 semanas (~10 meses).** Semana corrida: troque a sessão por um `/kata`; o plano absorve isso.

**Fio condutor:** a trilha alterna CLI → concorrência → backend → DevOps → distribuído, e alguns projetos **evoluem** ao longo das fases (o `encurta` nasce como monolito na fase 4 e vira microsserviços na fase 7; o `minikv` nasce na fase 6 e é otimizado na fase 8). Assim você sente na prática o custo e o benefício de cada decisão.

---

## Fase 0: Toolchain e primeiro contato (~1 semana)

**Objetivo:** sentir-se em casa com o ciclo editar → rodar → testar e com a filosofia "a ferramenta decide o estilo".
**Checkpoint de saída:** explicar a diferença entre módulo e pacote, o que `go build`/`go run`/`go vet`/`gofmt` fazem e por que Go não tem `while`, `try/catch` nem `class`.

### P00. gowc: clone do `wc` do Unix
- **Por quê:** é um programa pequeno, mas toca em I/O, argumentos, erros e saída formatada. Mostra de cara como Go trata erros e strings/bytes.
- **Etapas:**
  1. `gowc arquivo.txt` imprime linhas, palavras e bytes.
  2. Flags `-l -w -c -m` (runes), vários arquivos com total, leitura de stdin quando não há arquivo.
  3. Testes table-driven das funções de contagem.
- **Conceitos introduzidos:** `modules-e-go-mod`, `pacotes-e-visibilidade`, `toolchain-go`, `variaveis-e-zero-values`, `tipos-basicos`, `for-unico-loop`, `if-com-init`, `erros-como-valores`, `os-args-e-flag`, `bufio-scanner`, `fmt-verbos`, `testing-basico`
- **Critérios de aceite:** `gofmt -l .` vazio; `go vet ./...` limpo; testes para arquivo vazio, sem newline final e UTF-8 (`-m` ≠ `-c` para "ação").
- **Extensões:** `-L` (linha mais longa); comparar a saída com o `wc` real num script de teste.
- **Ponte com TS:** não há `undefined`/`null` para tipos básicos, existem **zero values**. Erro não é exceção, é **valor de retorno**: `if err != nil` é o `try/catch` explícito. Letra maiúscula = `export`. `go.mod` ≈ `package.json`, mas sem `node_modules` e sem scripts.

---

## Fase 1: Fundamentos pela linha de comando (~4 semanas)

**Objetivo:** escrever programas Go de verdade com structs, métodos, coleções, JSON, erros bem tratados e testes.
**Checkpoint de saída:** entrevista sobre ponteiro vs valor, slices (len/cap/append), mapas, wrapping de erros e por que interfaces de I/O são tão centrais.

### P01. tarefas: gerenciador de tarefas em CLI com persistência em JSON
- **Por quê:** é um CRUD que você conhece de cor, o que libera a cabeça para focar **só** em como Go modela dados.
- **Etapas:**
  1. Modelo `Tarefa` + operações em memória (`add`, `list`, `done`, `rm`) com subcomandos.
  2. Persistência em arquivo JSON (`~/.tarefas.json`), com escrita atômica (arquivo temporário + rename).
  3. Filtros (`--pendentes`, `--tag`), ordenação por prioridade/data, saída em tabela alinhada.
  4. Pacote interno `internal/store` com testes, incluindo `t.TempDir()` para os testes de arquivo.
- **Conceitos introduzidos:** `structs`, `metodos-e-receivers`, `ponteiros`, `slices`, `maps`, `encoding-json-e-struct-tags`, `time-pacote`, `wrapping-de-erros`, `errors-is-as`, `erros-sentinela-e-tipados`, `layout-cmd-internal`, `table-driven-tests`, `defer`
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
- **Critérios de aceite:** processa um log de 1 GB com memória constante (`/usr/bin/time -v`); benchmark com `b.ReportAllocs()`; interfaces pequenas, definidas **no consumidor**.
- **Extensões:** `--follow` (tail -f); histograma de latência no terminal.
- **Ponte com TS:** interface em Go é **estrutural e implícita** (parecido com TS!), mas é checada em runtime para type assertions e não existe `implements`. Prefira interfaces de 1–2 métodos. `io.Reader` ≈ Node streams, porém síncrono e muito mais simples.

### Katas da fase
- `rot13-reader`: implementar um `io.Reader` que embrulha outro (`io-reader-writer`)
- `frequencia-palavras`: top-K palavras com `maps` + `slices.SortFunc`
- `slice-aliasing`: prever e corrigir bugs de `append` compartilhando array (`slices`)
- `validar-cpf`: funções puras + table-driven tests (`table-driven-tests`)

---

## Fase 2: Sistema de tipos e design idiomático (~4 semanas)

**Objetivo:** dominar o sistema de tipos (interfaces a fundo, embedding, generics, iterators) e saber **quando** usar cada coisa.
**Checkpoint de saída:** explicar nil interface vs interface contendo nil, method sets, quando generics vale a pena, e desenhar uma API de pacote idiomática.

### P03. calc: interpretador de uma linguagem de expressões
- **Por quê:** lexer → parser → AST → avaliador é o projeto clássico para exercitar tipos, interfaces, type switch e erros ricos. Ele também prepara o terreno para o parser de protocolo do `minikv`.
- **Etapas:**
  1. **Lexer**: tokens com `iota`, `Stringer`, posição (linha/coluna), runes e UTF-8.
  2. **Parser** (descida recursiva/Pratt) gerando a AST; `interface Node` + type switch.
  3. **Avaliador** com variáveis, funções embutidas (`max`, `sqrt`) e erros com posição (tipo de erro customizado).
  4. **REPL** interativo + pacote `calc` com API pública documentada (`go doc`) e exemplos testáveis (`Example...`).
  5. Utilitários genéricos próprios (`Stack[T]`, `Map/Filter`) e um **iterator** (`iter.Seq`) sobre os tokens.
- **Conceitos introduzidos:** `iota-e-enums`, `stringer-interface`, `runes-e-utf8`, `type-switch`, `embedding`, `method-sets`, `nil-interface-gotcha`, `generics-type-params`, `generics-constraints`, `iterators-range-over-func`, `custom-error-types`, `godoc-e-example-tests`, `panic-recover`, `fuzzing-basico`
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

### P04. linkcheck: verificador de links concorrente (crawler)
- **Por quê:** I/O de rede em massa é o caso de uso perfeito para goroutines, e é fácil de "explodir" sem controle, o que torna os limites necessários.
- **Etapas:**
  1. Versão sequencial: baixa uma página, extrai links (`golang.org/x/net/html`) e checa o status de cada um.
  2. Versão concorrente ingênua (uma goroutine por link) → medir e ver o problema.
  3. **Worker pool** com channels + `sync.WaitGroup`; limite de concorrência configurável.
  4. **Context**: timeout global e por request, cancelamento por Ctrl+C (`signal.NotifyContext`).
  5. Crawl recursivo até a profundidade N, sem visitar duas vezes (mapa protegido) e rate limit por host.
- **Conceitos introduzidos:** `goroutines`, `channels`, `buffered-vs-unbuffered`, `select`, `sync-waitgroup`, `worker-pool`, `context-cancelamento-timeout`, `net-http-client`, `sync-mutex`, `race-detector`, `rate-limiting`, `goroutine-leak`
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

### P06. encurta (v1): API de encurtador de URLs só com a stdlib
- **Por quê:** a stdlib de Go é suficiente para APIs sérias. Aprender sem framework mostra o que os frameworks escondem.
- **Etapas:**
  1. `http.ServeMux` com padrões de método/rota (`POST /links`, `GET /{code}`), JSON in/out e status corretos.
  2. Camadas: handler → service → repository (interface), com storage em memória. Injeção de dependência manual no `main`.
  3. Middlewares: logging com `log/slog` (request id), recover, timeout, CORS. Autenticação por API key/JWT.
  4. Validação, erros de API padronizados (RFC 9457 problem+json) e configuração por env.
  5. Testes com `httptest` (handlers) + servidor com timeouts configurados e shutdown gracioso.
- **Conceitos introduzidos:** `net-http-server`, `servemux-patterns`, `http-handler-interface`, `middlewares`, `log-slog`, `context-em-requests`, `injecao-de-dependencia-manual`, `interfaces-no-consumidor`, `httptest`, `configuracao-por-env`, `server-timeouts`, `jwt-auth`
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
- `consistent-hashing`: anel de hash consistente com nós virtuais (`maps`, `pacote-slices-e-sort`)

---

## Fase 7: Microsserviços e observabilidade (~6 semanas)

**Objetivo:** decompor um sistema em serviços que se comunicam por gRPC e mensageria, com resiliência e observabilidade de produção.
**Checkpoint de saída:** system design oral: "desenhe o encurtador para 10k req/s", cobrindo consistência, idempotência, outbox, tracing distribuído e falhas parciais.

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
