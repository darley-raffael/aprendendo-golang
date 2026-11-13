---
name: desafio
description: Apresenta o próximo projeto ou a próxima etapa do roadmap. Faz uma mini-aula dos conceitos novos (just-in-time, comparando com a stack do aluno) e cria o DESAFIO.md com o enunciado, os requisitos e os critérios de aceite, sem escrever a solução. Use quando o aluno pedir o próximo desafio, projeto ou etapa, ou depois de uma etapa aprovada.
---

# /desafio

## 1. Descobrir o que vem a seguir

- Leia `ROADMAP.md`, `progresso/PROGRESSO.md` e `progresso/PERFIL.md`.
- Se houver etapa com review 🔁 pendente, **não avance**: lembre o aluno do que falta refatorar.
- Identifique o próximo item na ordem do roadmap: **micro-projeto da rampa → diagnóstico → etapa de projeto**.
  - **Micro-projeto** (`projetos/NN-rampa/MNN-nome/` ou `RN.x-nome/`): uma sessão, um `DESAFIO.md` enxuto (o quê + exemplos de entrada/saída + critérios). Aplique o **andaime decrescente** de `METODOLOGIA.md`: nos primeiros, crie você mesmo o arquivo de testes (que compila e falha) e diga quais arquivos e funções criar; nos do meio, forneça testes parciais; nos últimos, só o enunciado.
  - **Diagnóstico** (ao terminar a rampa): siga a seção "Diagnóstico da rampa" de `METODOLOGIA.md` **antes** de liberar o projeto grande. Se houver reforço, use `/kata` e só depois avance.
  - **Etapa de projeto**: crie a pasta `projetos/NN-nome/` se for a primeira etapa.
- Na mini-aula de um projeto, os conceitos de **Aprofunda** ganham um "agora mais fundo": 2–4 linhas ligando o que o aluno já sabe ao aspecto novo (ex.: "você já usou slices; agora vamos ver por que `append` às vezes altera outro slice").

## 2. Mini-aula dos conceitos novos (só os desta etapa)

Para cada conceito **novo** desta etapa (nível 0 em `CONCEITOS.md`):
- **O que é e por que existe**, em 3–6 linhas, com foco no *design* da linguagem.
- **Ponte com a stack do aluno:** o equivalente e, principalmente, onde a analogia **quebra**.
- **Exemplo mínimo de outro domínio** (nunca o código do desafio), com até 15 linhas.
- **Armadilha comum**, em 1–2 linhas.
- **Link** da documentação oficial.

Ao final, faça **1 pergunta de checagem** e espere a resposta antes de seguir. Atualize esses conceitos para **nível 1** em `CONCEITOS.md`, com a próxima revisão para amanhã.

Mantenha a mini-aula curta: se houver mais de 4 conceitos novos, distribua-os entre as etapas.

## 3. Criar/atualizar `projetos/NN-nome/DESAFIO.md`

Estrutura:

```markdown
# PNN: nome

## Contexto
Problema do "mundo real" que o projeto resolve (2–4 linhas).

## Etapa X: título  ← (adicione cada etapa conforme for liberada)
### Requisitos
- [ ] comportamento observável 1 (com exemplo de entrada → saída)
- [ ] …
### Critérios de aceite
- [ ] ferramentas automáticas passam (formatação, lint, testes, race…)
- [ ] testes cobrindo: casos felizes, bordas, erros
- [ ] requisitos não-funcionais (ex.: não vaza goroutine, processa arquivo de 1 GB com memória constante)
### Conceitos em jogo
`conceito-a`, `conceito-b`
### Restrições didáticas
Ex.: "use apenas a stdlib", "sem frameworks", "sem mutex nesta etapa"
```

- Os requisitos descrevem **o quê**, nunca **como**.
- Em projetos, **não** crie arquivos de código. Só para o primeiro projeto da trilha, explique o comando de scaffolding (ex.: `go mod init`), mas deixe o aluno rodar.
- Opcional: forneça 1–3 **testes de aceitação** em `projetos/NN-nome/acceptance_test.*` quando a especificação for mais clara em código do que em texto.

## 4. Fechar

Atualize `PROGRESSO.md` (etapa atual, status "em andamento", próxima ação) e diga ao aluno:
- por onde começar (a primeira pergunta que ele deve se fazer, **não** a resposta);
- que `/dica` existe se ele travar por mais de 10–15 min;
- que `/review` é o próximo passo quando a etapa estiver pronta.
