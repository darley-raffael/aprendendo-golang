---
name: encerrar
description: Encerra a sessão de estudo. Faz o aluno resumir o que aprendeu (recuperação ativa), escreve o diário, atualiza PROGRESSO e CONCEITOS, ajusta a memória do Claude se as preferências mudaram e propõe o commit. Use quando o aluno disser "/encerrar", "por hoje é só", "vou parar" ou algo parecido.
---

# /encerrar

## 1. Recuperação ativa (pergunte, uma de cada vez)

1. "Em 1–2 frases: o que você aprendeu hoje?"
2. "O que ainda está nebuloso?"

Corrija com gentileza se o resumo tiver algum equívoco. Isso também é aprendizado.

## 2. Diário: `progresso/diario/AAAA-MM-DD.md`

Crie ou complemente o arquivo do dia (se houver mais de uma sessão no dia, adicione uma seção `## Sessão N`):

```markdown
# AAAA-MM-DD

**Duração:** ~X min · **Fase/etapa:** … · **Tipo:** projeto | kata | revisão | checkpoint

## Feito
- …
## Aprendi (nas palavras do aluno)
- …
## Dúvidas em aberto
- …
## Revisão espaçada
- conceito: ✅/🟡/❌
## Dicas
- conceito: nível N, contexto
## Próxima ação
- …
```

## 3. Atualizações

- `PROGRESSO.md`: "Agora" (etapa, status, próxima ação, última sessão) e métricas (sessões +1, horas, dicas, katas).
- `CONCEITOS.md`: garanta que tudo que mudou hoje está registrado. Dúvidas em aberto sobre um conceito → próxima revisão = amanhã.
- `PERFIL.md`: novas observações do tutor, se surgiram.
- **Memória do Claude:** atualize apenas se o aluno expressou uma preferência nova ou mudou um objetivo ou o ritmo. Estado de progresso **não** vai para a memória.

## 4. Commit

Mostre `git status` resumido e proponha a mensagem, por exemplo:
`docs(progresso): sessão AAAA-MM-DD — P03 etapa 2`
(e um commit separado de código, se houver código não commitado).
Faça o commit **só com a confirmação do aluno**. Push só se ele pedir.

## 5. Despedida

Em 2–3 linhas: a próxima ação concreta e quando vence a próxima revisão espaçada.
