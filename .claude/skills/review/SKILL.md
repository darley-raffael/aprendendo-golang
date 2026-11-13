---
name: review
description: Faz o code review de uma etapa de projeto ou kata entregue pelo aluno. Roda as ferramentas automáticas, confere os critérios de aceite e dá feedback categorizado (bloqueante, idiomático, design, sugestão) sem reescrever o código. Termina com o veredito aprovado ou refatorar. Use quando o aluno disser "/review", "terminei", "pode revisar" ou entregar uma etapa.
---

# /review

## 1. Ferramentas automáticas primeiro

Rode os comandos da seção "Comandos e ferramentas da tecnologia" do `CLAUDE.md` (formatação, vet/lint, testes, race detector, cobertura) na pasta do projeto. Se algo falhar:
- mostre a saída relevante;
- pergunte ao aluno o que ele acha que significa;
- **pare aqui**: o review humano só acontece com as ferramentas verdes (exceto se ele pedir ajuda para entender a falha, e aí vale `/dica`).

## 2. Critérios de aceite

Leia a etapa no `DESAFIO.md` e marque cada item como ✅ ou ❌, com uma evidência (nome do teste, comportamento observado, comando executado). Teste manualmente os casos de borda que os testes do aluno não cobrem.

## 3. Review humano

Leia **todo** o código da etapa. Classifique cada observação:

- 🔴 **Bloqueante**: bug, race condition, vazamento de recurso, erro ignorado, panic possível, segurança
- 🟡 **Idiomático**: funciona, mas não é como a comunidade escreve. **Atenção especial aos hábitos trazidos da stack antiga do aluno** (veja `PERFIL.md`)
- 🔵 **Design**: nomes, responsabilidades, acoplamento, testabilidade, API pública
- ⚪ **Sugestão**: opcional, para quem quiser ir além

Para cada item: `arquivo:linha`, **o que** está errado, **por que** importa e uma **pergunta** que leve à correção. Não reescreva. No máximo 3–5 linhas ilustrando um *padrão* genérico.

Também aponte **o que ficou bom** (1–3 itens concretos). Reforço positivo específico ajuda a fixar.

## 4. Veredito

- ✅ **Aprovado**: nenhum 🔴 e nenhum 🟡 pendente. (🔵 e ⚪ podem ficar como "para depois".)
- 🔁 **Refatorar**: lista numerada do que precisa mudar. O aluno corrige e roda `/review` de novo; na re-revisão, foque nos itens apontados.

## 5. Registro

- Acrescente em `projetos/NN-nome/REVIEWS.md` (ou `katas/NNN-nome/REVIEW.md`): data, etapa, resultado das ferramentas, critérios, itens do review e veredito.
- Se **aprovado**:
  - `PROGRESSO.md`: marque a etapa como concluída e defina a próxima ação;
  - `CONCEITOS.md`: conceitos da etapa aplicados **sem dica de nível 3** → nível 3; **com** dica de nível 3 → nível 2. Próxima revisão conforme a tabela.
- Se aparecer um padrão recorrente (ex.: sempre ignora erros), anote em "Observações do tutor" no `PERFIL.md`.
- Sugira uma mensagem de commit (`feat(NN-nome): …`).
