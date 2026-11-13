# Metodologia

Aprendizado **orientado a projetos**, com **conceitos introduzidos just-in-time**, **tutoria socrática**, **code review a cada entrega** e **revisão espaçada** para fixar a memória de longo prazo.

## Princípios

1. **Projeto primeiro, teoria depois.** Cada fase gira em torno de projetos reais. Um conceito só aparece quando o projeto precisa dele, e é aí que ele faz sentido.
2. **Ponte com o que você já sabe.** Todo conceito novo é comparado com a sua stack atual: o que é igual, o que é diferente e que armadilha vem de trazer hábitos antigos.
3. **Esforço desejável.** Travar faz parte. O tutor não entrega a resposta; dá dicas em níveis crescentes. Lutar 10–15 min com um problema fixa mais do que ler a solução.
4. **Idiomático > funcional.** Uma entrega só é aprovada quando segue as convenções da comunidade da tecnologia.
5. **Recuperação ativa e espaçada.** Revisar puxando da memória (quiz, prever saída, achar bug) em intervalos crescentes vence reler.
6. **Progresso visível.** Tudo é registrado em arquivos versionados: dá para ver a evolução e retomar de qualquer máquina.
7. **Rampa antes do projeto.** Toda fase começa com **micro-projetos** pequenos (uma sessão cada) que isolam 1–4 conceitos novos. O projeto grande vem depois e **combina** o que a rampa ensinou num problema real. Nunca jogue o aluno num projeto que exige muitos conceitos inéditos de uma vez.
8. **Introduz → aprofunda.** Um tema aparece primeiro no básico e volta mais fundo em fases seguintes (ex.: "listas básicas" → "internals de memória" → "otimização de alocação"). Cada projeto declara o que **introduz** e o que **aprofunda**.
9. **Andaime decrescente.** No início, o tutor fornece testes e estrutura; o suporte diminui gradualmente até o aluno conduzir tudo sozinho.

## Estrutura do roadmap

```
Fase
 ├── Objetivo e critério de saída (checkpoint)
 ├── Rampa: 2–6 micro-projetos (1 sessão cada, 1–4 conceitos novos cada)
 │    └── Diagnóstico da rampa → reforço se necessário
 ├── Projetos (1–3 por fase, cada um dividido em ETAPAS entregáveis)
 │    ├── Conceitos introduzidos (só o que a rampa não cobriu)
 │    ├── Conceitos aprofundados (vistos antes, agora mais fundo)
 │    ├── Requisitos MVP + extensões opcionais
 │    └── Critérios de aceite (o que o /review verifica)
 └── Katas (exercícios curtos de 20–40 min)
```

Projetos são **divididos em etapas** de 1 a 3 sessões cada. Cada etapa termina em algo que roda e passa num `/review`. Isso mantém o ritmo mesmo com pouco tempo semanal.

## Rampa, andaime e diagnóstico

**Micro-projeto ≠ kata.** A kata é um exercício de reforço sobre conceitos já vistos. O micro-projeto é um **programa pequeno e completo** (roda, tem entrada e saída) que **introduz** conceitos novos com o mínimo de distração.

**Andaime decrescente** (principalmente na primeira fase, quando a linguagem é toda nova):

| Momento | Testes | Estrutura |
|---|---|---|
| Primeiros micro-projetos | o tutor fornece todos | o tutor diz quais arquivos e funções criar |
| Micro-projetos do meio | o tutor fornece parte; o aluno completa as bordas | o aluno decide, o tutor revisa |
| Últimos micro-projetos e projetos | o aluno escreve todos | o aluno decide |

Nas fases seguintes, a rampa já começa no nível "o aluno escreve os testes", salvo quando o tema é muito novo (ex.: concorrência, rede).

**Diagnóstico da rampa** (ao final de cada rampa, ~10 min):
1. 3–5 perguntas cobrindo os conceitos da rampa (no formato da revisão espaçada).
2. Análise do diário: quantas dicas, e de que nível, em cada conceito; itens 🟡 recorrentes nos reviews.
3. Resultado registrado em "Observações do tutor" no `PERFIL.md`.
4. Conceito abaixo do nível 2 → **reforço** (1 kata ou micro-projeto extra) antes do projeto grande. Rampa muito fácil → seguir direto e considerar uma extensão do projeto.

## Ciclo de uma sessão (60–90 min)

| Momento | Duração | O que acontece |
|---|---|---|
| **Aquecimento** (`/sessao`) | 5–10 min | Retomada do ponto onde parou + 2–3 perguntas de revisão espaçada |
| **Mini-aula** (`/desafio`) | 5–15 min | Só quando a etapa traz conceitos novos: explicação curta, comparação com sua stack, 1 pergunta de checagem |
| **Prática** | 40–60 min | Você codifica. `/dica` quando travar há mais de 10–15 min |
| **Review** (`/review`) | 10–15 min | Quando a etapa estiver pronta: ferramentas automáticas + review humano-idiomático |
| **Encerramento** (`/encerrar`) | 5 min | Diário, atualização do progresso, agenda de revisão, commit |

**Semana corrida, sessão curta (< 40 min)?** Faça um `/kata` em vez de avançar no projeto.

## Dicas progressivas (`/dica`)

| Nível | Forma | Exemplo |
|---|---|---|
| 1. Direção | Uma pergunta que reorienta o raciocínio | "O que acontece com a goroutine se ninguém ler desse channel?" |
| 2. Ferramenta | Indica o conceito/API certo + link da doc | "Veja `context.WithTimeout` e como `select` reage a `ctx.Done()`" |
| 3. Esqueleto | Pseudocódigo ou estrutura parcial, sem a lógica final | Assinaturas das funções e comentários `// TODO` |

Cada dica usada é registrada no diário. Muitas dicas nível 3 num conceito = conceito volta para a revisão mais cedo.

## Code review (`/review`)

1. **Ferramentas automáticas** (formatação, lint, testes, race detector etc.): precisam passar antes do review humano.
2. **Critérios de aceite** da etapa: checklist.
3. **Feedback em categorias**, com severidade:
   - 🔴 **Bloqueante**: bug, race, vazamento, erro ignorado
   - 🟡 **Idiomático**: funciona, mas não é como a comunidade escreve
   - 🔵 **Design**: organização, nomes, responsabilidades, testabilidade
   - ⚪ **Sugestão**: melhoria opcional, curiosidade, otimização
4. O tutor **aponta e pergunta**, não reescreve.
5. **Veredito**: ✅ aprovado ou 🔁 refatorar (com lista do que mudar). Sem 🔴 e sem 🟡 pendente = aprovado.
6. O review fica salvo em `projetos/NN-nome/REVIEWS.md`.

## Domínio de conceitos e revisão espaçada

Cada conceito em `progresso/CONCEITOS.md` tem um **nível de domínio**:

| Nível | Significado |
|---|---|
| 0 | Não visto |
| 1 | Introduzido (mini-aula) |
| 2 | Aplicado em projeto/kata com ajuda |
| 3 | Aplicado sozinho e aprovado em review |
| 4 | Lembrado corretamente em revisão espaçada (≥ 2 vezes) |
| 5 | Dominado: consegue explicar e ensinar, reconhece armadilhas |

**Agenda de revisão** (intervalo até a próxima revisão, conforme o nível após a revisão):

| Nível | 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|
| Intervalo | 1 dia | 3 dias | 7 dias | 14 dias | 30 dias |

- **Acertou** a pergunta de revisão → nível +1 (máx. 5), próxima revisão = hoje + intervalo do novo nível. (Subir de 3 para 4 exige 2 acertos seguidos no nível 3.)
- **Errou ou acertou parcialmente** → nível −1 (mín. 1), próxima revisão = amanhã.
- Perguntas variam: conceitual ("por que…?"), prever a saída de um snippet, achar o bug, escolher entre duas abordagens.

## Checkpoint de fase

Para sair de uma fase, você precisa:
- ter todos os projetos obrigatórios da fase **aprovados** no review;
- ter todos os conceitos da fase em **nível ≥ 3**;
- passar no **checkpoint**: uma conversa de 10–15 min em que o tutor faz perguntas no estilo de entrevista técnica sobre os temas da fase.

## Formato do roadmap (usado pelo `/iniciar`)

Cada fase em `ROADMAP.md` segue:

```markdown
## Fase N: Nome (~X semanas)
**Objetivo:** …
**Checkpoint de saída:** …

### Rampa
- **RN.1 `micro-projeto`:** o que construir em 1 sessão (`conceito-a`, `conceito-b`)
- **RN.2 `micro-projeto`:** … (`conceito-c`)
- **Diagnóstico** → PN

### PN. nome-do-projeto: descrição curta
- **Por quê:** o que esse projeto ensina e por que nesta ordem
- **Etapas:** 1) … 2) … 3) …
- **Conceitos introduzidos:** `conceito-a`, `conceito-b`, … (inclui os da rampa; é a lista que gera o CONCEITOS.md)
- **Aprofunda:** `conceito-de-fase-anterior`, …
- **Critérios de aceite:** …
- **Extensões (opcional):** …
- **Ponte com sua stack:** comparações e armadilhas

### Katas da fase
- `kata`: descrição (conceito-alvo)
```
