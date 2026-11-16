# Conceitos e revisão espaçada

> Níveis e intervalos definidos em `METODOLOGIA.md`.
> `/sessao` e `/revisar` escolhem conceitos com **Próxima revisão ≤ hoje**, priorizando o menor nível.
> O `/iniciar` gera a lista inicial a partir do `ROADMAP.md`, com todos os conceitos em nível 0.

| Conceito | Fase | Nível | Acertos seguidos | Última revisão | Próxima revisão | Notas |
|---|---|---|---|---|---|---|
| `toolchain-go` | 0 | 3 | 0 | 2026-11-13 | 2026-11-20 | M01 aprovado sem dicas |
| `modules-e-go-mod` | 0 | 2 | 0 | 2026-11-13 | 2026-11-16 | M01: caminho do módulo só corrigido após review |
| `pacotes-e-visibilidade` | 0 | 2 | 0 | 2026-11-13 | 2026-11-16 | checagem parcial: confundiu pacote x módulo; import sem caminho do módulo; esqueceu campo minúsculo; M01: doc comment só após 2 reviews |
| `fmt-verbos` | 0 | 3 | 0 | 2026-11-16 | 2026-11-23 | M01 só usou Println; verbos ainda não aplicados (M02 aplica); M02: %.2f, %q e Fprintln aplicados |
| `variaveis-e-zero-values` | 0 | 3 | 0 | 2026-11-16 | 2026-11-23 | M02: zero value no retorno de erro |
| `tipos-basicos` | 0 | 3 | 0 | 2026-11-16 | 2026-11-23 | M02 aprovado |
| `conversao-de-tipos-explicita` | 0 | 1 | 0 | 2026-11-16 | 2026-11-17 | M02 não exigiu T(v) na prática; revisar |
| `constantes` | 0 | 3 | 0 | 2026-11-16 | 2026-11-23 | checagem: compila/não certo, não deu saídas nem o porquê (untyped const); M02: nomes de constantes enganosos (celsiusPorF, baseCalculo) |
| `strconv` | 0 | 3 | 0 | 2026-11-16 | 2026-11-23 | M02 aprovado |
| `erros-como-valores` | 0 | 2 | 0 | 2026-11-16 | 2026-11-19 | checagem: chamou retorno múltiplo de "tupla"; não citou zero value no erro; M02: precisou de dicas N1+N2 e 3 rodadas (erro genérico, checagem longe da chamada) |
| `if-com-init` | 0 | 1 | 0 | 2026-11-16 | 2026-11-17 | não aplicado no M02; revisar |
| `for-unico-loop` | 0 | 0 | 0 | - | - | |
| `switch` | 0 | 0 | 0 | - | - | |
| `funcoes-multiplos-retornos` | 0 | 0 | 0 | - | - | |
| `bufio-scanner` | 0 | 0 | 0 | - | - | |
| `math-rand` | 0 | 0 | 0 | - | - | |
| `slices-basico` | 0 | 0 | 0 | - | - | |
| `range` | 0 | 0 | 0 | - | - | |
| `funcoes-variadicas` | 0 | 0 | 0 | - | - | |
| `testing-basico` | 0 | 0 | 0 | - | - | |
| `maps-basico` | 0 | 0 | 0 | - | - | |
| `structs-basico` | 0 | 0 | 0 | - | - | |
| `metodos-intro` | 0 | 0 | 0 | - | - | |
| `ponteiros-basico` | 0 | 0 | 0 | - | - | |
| `table-driven-tests` | 0 | 0 | 0 | - | - | |
| `os-args-e-flag` | 0 | 0 | 0 | - | - | |
| `pacote-strings` | 0 | 0 | 0 | - | - | |
| `runes-e-utf8` | 0 | 0 | 0 | - | - | |
| `crypto-rand` | 0 | 0 | 0 | - | - | |
| `os-open-e-arquivos` | 0 | 0 | 0 | - | - | |
| `defer` | 0 | 0 | 0 | - | - | |
| `io-reader-intro` | 0 | 0 | 0 | - | - | |
| `doc-comments` | 0 | 0 | 0 | - | - | micro-projeto dedicado M03b (pedido do aluno em 2026-11-16) |
| `metodos-e-receivers` | 1 | 0 | 0 | - | - | |
| `slices-internals-len-cap` | 1 | 0 | 0 | - | - | |
| `encoding-json-e-struct-tags` | 1 | 0 | 0 | - | - | |
| `time-pacote` | 1 | 0 | 0 | - | - | |
| `wrapping-de-erros` | 1 | 0 | 0 | - | - | |
| `errors-is-as` | 1 | 0 | 0 | - | - | |
| `erros-sentinela-e-tipados` | 1 | 0 | 0 | - | - | |
| `layout-cmd-internal` | 1 | 0 | 0 | - | - | |
| `escrita-atomica-de-arquivos` | 1 | 0 | 0 | - | - | |
| `interfaces-implicitas` | 1 | 0 | 0 | - | - | |
| `io-reader-writer` | 1 | 0 | 0 | - | - | |
| `composicao-de-interfaces` | 1 | 0 | 0 | - | - | |
| `strings-e-bytes-builder` | 1 | 0 | 0 | - | - | |
| `pacote-slices-e-sort` | 1 | 0 | 0 | - | - | |
| `closures-e-funcoes-de-primeira-classe` | 1 | 0 | 0 | - | - | |
| `benchmarks-basico` | 1 | 0 | 0 | - | - | |
| `type-assertion` | 1 | 0 | 0 | - | - | |
| `iota-e-enums` | 2 | 0 | 0 | - | - | |
| `stringer-interface` | 2 | 0 | 0 | - | - | |
| `type-switch` | 2 | 0 | 0 | - | - | |
| `embedding` | 2 | 0 | 0 | - | - | |
| `method-sets` | 2 | 0 | 0 | - | - | |
| `nil-interface-gotcha` | 2 | 0 | 0 | - | - | |
| `generics-type-params` | 2 | 0 | 0 | - | - | |
| `generics-constraints` | 2 | 0 | 0 | - | - | |
| `iterators-range-over-func` | 2 | 0 | 0 | - | - | |
| `custom-error-types` | 2 | 0 | 0 | - | - | |
| `godoc-e-example-tests` | 2 | 0 | 0 | - | - | |
| `panic-recover` | 2 | 0 | 0 | - | - | |
| `fuzzing-basico` | 2 | 0 | 0 | - | - | |
| `goroutines` | 3 | 0 | 0 | - | - | |
| `channels` | 3 | 0 | 0 | - | - | |
| `buffered-vs-unbuffered` | 3 | 0 | 0 | - | - | |
| `select` | 3 | 0 | 0 | - | - | |
| `sync-waitgroup` | 3 | 0 | 0 | - | - | |
| `worker-pool` | 3 | 0 | 0 | - | - | |
| `context-cancelamento-timeout` | 3 | 0 | 0 | - | - | |
| `net-http-client` | 3 | 0 | 0 | - | - | |
| `sync-mutex` | 3 | 0 | 0 | - | - | |
| `race-detector` | 3 | 0 | 0 | - | - | |
| `rate-limiting` | 3 | 0 | 0 | - | - | |
| `goroutine-leak` | 3 | 0 | 0 | - | - | |
| `pipeline-pattern` | 3 | 0 | 0 | - | - | |
| `fan-out-fan-in` | 3 | 0 | 0 | - | - | |
| `errgroup` | 3 | 0 | 0 | - | - | |
| `semaforo-com-channel` | 3 | 0 | 0 | - | - | |
| `sync-atomic` | 3 | 0 | 0 | - | - | |
| `sync-once` | 3 | 0 | 0 | - | - | |
| `sync-rwmutex` | 3 | 0 | 0 | - | - | |
| `graceful-shutdown` | 3 | 0 | 0 | - | - | |
| `gomaxprocs-e-scheduler` | 3 | 0 | 0 | - | - | |
| `testing-synctest` | 3 | 0 | 0 | - | - | |
| `memory-model-basico` | 3 | 0 | 0 | - | - | |
| `net-http-server` | 4 | 0 | 0 | - | - | |
| `servemux-patterns` | 4 | 0 | 0 | - | - | |
| `http-handler-interface` | 4 | 0 | 0 | - | - | |
| `middlewares` | 4 | 0 | 0 | - | - | |
| `log-slog` | 4 | 0 | 0 | - | - | |
| `context-em-requests` | 4 | 0 | 0 | - | - | |
| `injecao-de-dependencia-manual` | 4 | 0 | 0 | - | - | |
| `interfaces-no-consumidor` | 4 | 0 | 0 | - | - | |
| `httptest` | 4 | 0 | 0 | - | - | |
| `configuracao-por-env` | 4 | 0 | 0 | - | - | |
| `server-timeouts` | 4 | 0 | 0 | - | - | |
| `jwt-auth` | 4 | 0 | 0 | - | - | |
| `database-sql-e-pgx` | 4 | 0 | 0 | - | - | |
| `connection-pool` | 4 | 0 | 0 | - | - | |
| `migrations` | 4 | 0 | 0 | - | - | |
| `sqlc` | 4 | 0 | 0 | - | - | |
| `transacoes` | 4 | 0 | 0 | - | - | |
| `testes-de-integracao` | 4 | 0 | 0 | - | - | |
| `build-tags` | 4 | 0 | 0 | - | - | |
| `docker-multi-stage` | 4 | 0 | 0 | - | - | |
| `cross-compilation` | 4 | 0 | 0 | - | - | |
| `cache-aside` | 4 | 0 | 0 | - | - | |
| `cobra-cli` | 5 | 0 | 0 | - | - | |
| `config-precedencia` | 5 | 0 | 0 | - | - | |
| `os-exec` | 5 | 0 | 0 | - | - | |
| `go-embed` | 5 | 0 | 0 | - | - | |
| `ldflags-versao` | 5 | 0 | 0 | - | - | |
| `goreleaser` | 5 | 0 | 0 | - | - | |
| `sdk-docker` | 5 | 0 | 0 | - | - | |
| `client-go` | 5 | 0 | 0 | - | - | |
| `paginacao-e-retry` | 5 | 0 | 0 | - | - | |
| `go-generate` | 5 | 0 | 0 | - | - | |
| `tool-directives-go-mod` | 5 | 0 | 0 | - | - | |
| `net-tcp-server` | 6 | 0 | 0 | - | - | |
| `protocolo-binario-e-parsing` | 6 | 0 | 0 | - | - | |
| `bufio-reader-writer-de-rede` | 6 | 0 | 0 | - | - | |
| `lock-sharding` | 6 | 0 | 0 | - | - | |
| `expiracao-ttl` | 6 | 0 | 0 | - | - | |
| `durabilidade-aof-fsync` | 6 | 0 | 0 | - | - | |
| `backpressure` | 6 | 0 | 0 | - | - | |
| `replicacao-lider-seguidor` | 6 | 0 | 0 | - | - | |
| `deadlines-de-conexao` | 6 | 0 | 0 | - | - | |
| `testes-de-carga` | 6 | 0 | 0 | - | - | |
| `grpc-e-protobuf` | 7 | 0 | 0 | - | - | |
| `grpc-interceptors` | 7 | 0 | 0 | - | - | |
| `mensageria-nats-kafka` | 7 | 0 | 0 | - | - | |
| `idempotencia` | 7 | 0 | 0 | - | - | |
| `outbox-pattern` | 7 | 0 | 0 | - | - | |
| `retry-backoff-jitter` | 7 | 0 | 0 | - | - | |
| `circuit-breaker` | 7 | 0 | 0 | - | - | |
| `health-checks` | 7 | 0 | 0 | - | - | |
| `opentelemetry-tracing` | 7 | 0 | 0 | - | - | |
| `metricas-prometheus` | 7 | 0 | 0 | - | - | |
| `deploy-kubernetes` | 7 | 0 | 0 | - | - | |
| `pprof` | 8 | 0 | 0 | - | - | |
| `execution-tracer` | 8 | 0 | 0 | - | - | |
| `escape-analysis` | 8 | 0 | 0 | - | - | |
| `sync-pool` | 8 | 0 | 0 | - | - | |
| `gc-gogc-gomemlimit` | 8 | 0 | 0 | - | - | |
| `scheduler-gmp` | 8 | 0 | 0 | - | - | |
| `pgo` | 8 | 0 | 0 | - | - | |
| `golangci-lint` | 8 | 0 | 0 | - | - | |
| `govulncheck` | 8 | 0 | 0 | - | - | |
| `ci-github-actions` | 8 | 0 | 0 | - | - | |
| `leitura-de-codigo-stdlib` | 9 | 0 | 0 | - | - | |
| `system-design-com-go` | 9 | 0 | 0 | - | - | |
| `contribuicao-open-source` | 9 | 0 | 0 | - | - | |
