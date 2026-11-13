---
name: dica
description: Dá uma dica progressiva (nível 1 direção, nível 2 ferramenta/API, nível 3 esqueleto) quando o aluno trava num projeto ou kata, sem entregar a solução. Use quando o aluno disser "/dica", "travei", "não sei como fazer X" ou pedir ajuda num desafio.
---

# /dica [nível]

## Procedimento

1. Leia o `DESAFIO.md` (ou o `KATA.md`) da etapa atual e o **código atual do aluno**. Se houver erro de compilação ou teste, rode o comando para ver a mensagem real.
2. Se não estiver claro onde ele travou, pergunte **uma coisa**: "o que você tentou e o que aconteceu?"
3. Veja no diário de hoje quantas dicas já foram dadas **para este mesmo ponto**. Dê o **próximo nível**; o aluno pode pedir um nível explicitamente (`/dica 2`).

| Nível | O que dar | O que NÃO dar |
|---|---|---|
| **1. Direção** | 1–2 perguntas socráticas que apontam o problema real; ou "rode X e observe Y" | Nome de API, código |
| **2. Ferramenta** | O conceito ou API certo, o link da doc e um exemplo mínimo **de outro domínio** | Código aplicado ao desafio |
| **3. Esqueleto** | Estrutura parcial: assinaturas, tipos, comentários `// TODO` descrevendo cada passo | A lógica dentro dos TODOs |

4. Depois do nível 3, se o aluno continuar travado, ofereça: (a) uma pausa com um kata menor sobre o mesmo conceito; ou (b) mostrar a solução, **só se ele pedir explicitamente**. Nesse caso, registre e proponha uma variação para ele resolver sozinho.

## Erros de compilação/runtime

Ensine a **ler** a mensagem: peça para ele explicar o que o erro diz antes de você explicar. Mensagens de erro do compilador são uma habilidade a treinar.

## Registro

Anote no diário de hoje (`progresso/diario/AAAA-MM-DD.md`, seção "Dicas"): o conceito, o nível dado e o contexto em uma linha. Some ao contador de dicas em `PROGRESSO.md`.
