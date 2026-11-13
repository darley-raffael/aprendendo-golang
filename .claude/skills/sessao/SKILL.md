---
name: sessao
description: Inicia uma sessão de estudo. Retoma de onde o aluno parou, aplica a revisão espaçada dos conceitos vencidos e propõe o plano do dia conforme o tempo disponível. Use no começo de toda sessão de estudo, ou quando o aluno disser "bora estudar", "vamos continuar" ou algo parecido.
---

# /sessao

## 1. Retomada (sem perguntar nada ainda)

Leia `progresso/PROGRESSO.md`, `progresso/CONCEITOS.md`, o diário mais recente em `progresso/diario/` e o `DESAFIO.md` e `REVIEWS.md` do projeto atual, se houver.

Apresente em até 6 linhas:
- dias desde a última sessão;
- fase e etapa atual e o status;
- o que ficou pendente (refatorações do último review, dúvidas registradas no diário);
- quantos conceitos estão com revisão vencida.

## 2. Revisão espaçada (aquecimento, 5–10 min)

- Selecione até **3 conceitos** com `Próxima revisão ≤ hoje` (menor nível primeiro). Se nenhum venceu, pule esta etapa.
- Faça **uma pergunta por vez** e espere a resposta. Varie os formatos:
  - conceitual ("por que X é assim?", "quando você usaria X em vez de Y?");
  - **prever a saída** de um snippet curto (5–15 linhas);
  - **achar o bug** num snippet;
  - comparação com a stack do aluno ("como isso difere de <equivalente na stack dele>?").
- Avalie com franqueza: correto, parcial ou incorreto. Explique em 2–4 linhas o que faltou.
- Atualize `CONCEITOS.md` conforme as regras de `METODOLOGIA.md` (nível, acertos seguidos, datas).
- Se a sessão anterior registrou muitas dicas de nível 3 num conceito, inclua esse conceito na revisão mesmo que não esteja vencido.

## 3. Plano do dia

Pergunte quanto tempo o aluno tem hoje e proponha:
- **< 40 min**: `/kata` focado num conceito fraco, ou terminar uma refatoração pendente;
- **40–90 min**: continuar a etapa atual (ou `/desafio` se a etapa anterior foi aprovada);
- **> 90 min**: etapa atual + início da próxima, ou etapa + kata.

Se houver refatoração pendente de review, ela vem **antes** de qualquer etapa nova.

Se todos os projetos da fase estão aprovados e os conceitos da fase estão em nível ≥ 3, proponha o **checkpoint de fase** (veja `METODOLOGIA.md`): 6–10 perguntas no estilo entrevista técnica, uma por vez. Registre o resultado em `PROGRESSO.md`.
