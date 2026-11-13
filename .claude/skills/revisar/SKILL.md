---
name: revisar
description: Faz um quiz avulso de revisão espaçada sobre conceitos já estudados e atualiza os níveis de domínio e a agenda em CONCEITOS.md. Use quando o aluno pedir "/revisar", "me testa", "quiz" ou quiser revisar uma fase ou conceito específico.
---

# /revisar [n | fase N | conceito]

1. Leia `progresso/CONCEITOS.md` e `progresso/PERFIL.md`.
2. Selecione as perguntas:
   - sem argumento: até **5** conceitos com revisão vencida (menor nível primeiro); se não houver vencidos, os 5 de menor nível ≥ 1;
   - `n`: n conceitos;
   - `fase N`: conceitos da fase N, em modo "simulado de entrevista";
   - `conceito`: 3 perguntas de ângulos diferentes sobre ele.
3. **Uma pergunta por vez.** Espere a resposta antes da próxima. Varie entre conceitual, prever a saída, achar o bug, "qual abordagem e por quê" e comparação com a stack do aluno. Snippets curtos e realistas, com armadilhas reais da linguagem.
4. Depois de cada resposta: veredito (✅, 🟡 parcial ou ❌) e uma explicação de 2–5 linhas focada no **porquê**.
5. Aplique as regras de nível e intervalo de `METODOLOGIA.md` e atualize `CONCEITOS.md`.
6. Ao final: placar, conceitos que subiram ou caíram e uma sugestão (ex.: "kata de channels para reforçar").

Nunca pergunte sobre conceitos de nível 0 (ainda não vistos).
