---
name: kata
description: Cria um exercício curto (20 a 40 min), focado em um conceito, com testes escritos pelo tutor como especificação. Use em dias com pouco tempo, para reforçar conceitos fracos ou vencidos na revisão espaçada, ou quando o aluno pedir "/kata" ou um exercício rápido.
---

# /kata [conceito]

## 1. Escolher o alvo

Em ordem de prioridade:
1. o conceito passado como argumento;
2. um conceito com revisão vencida e nível baixo em `CONCEITOS.md`;
3. um conceito em que o aluno usou dica de nível 3 recentemente (veja o diário);
4. uma kata listada para a fase atual em `ROADMAP.md`.

Use apenas conceitos que o aluno **já viu** (nível ≥ 1). Kata não é lugar para conteúdo novo.

## 2. Criar `katas/NNN-nome/`

- `KATA.md`: enunciado curto, exemplos de entrada → saída, conceito-alvo, tempo sugerido e restrições (ex.: "sem bibliotecas externas", "implemente sem usar mutex").
- **Arquivo de testes escrito pelo tutor**: ele é a especificação. Cubra casos felizes, bordas e erros. Quando fizer sentido, inclua benchmark ou teste com race detector.
- Um arquivo de código com **apenas as assinaturas** necessárias para compilar (corpo com `panic("TODO")` ou equivalente), para que os testes rodem e falhem.
- Inicialize o módulo/projeto se a tecnologia exigir (a kata deve rodar de forma isolada).

Confirme que os testes **compilam e falham** antes de entregar ao aluno.

## 3. Durante e depois

- O aluno faz os testes passarem. Valem as regras de `/dica`.
- Ao terminar, faça um **mini-review** (versão curta do `/review`, foco no idiomático) e salve em `katas/NNN-nome/REVIEW.md`.
- Ofereça uma **variação "hard"** opcional (ex.: mesma kata com generics ou concorrência, ou com restrição de alocação).
- Atualize `CONCEITOS.md` (acertou sem dica → conta como revisão bem-sucedida) e o contador de katas em `PROGRESSO.md`.
