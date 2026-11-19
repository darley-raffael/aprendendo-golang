# Perfil do aluno

> Preenchido pelo `/iniciar` em 2026-11-13. Atualize quando objetivos ou preferências mudarem.

- **Tecnologia estudada:** Go (toolchain local: go1.27.1 linux/amd64)
- **Início:** 2026-11-13
- **Stack atual (base para comparações):** JavaScript / TypeScript (Node.js)
- **Experiência prévia com a tecnologia:** iniciante em Go; desenvolvedor experiente em geral
- **Objetivos:** APIs e microsserviços · CLIs e ferramentas DevOps · sistemas concorrentes/distribuídos · carreira/mercado Go
- **Tempo disponível:** até 5h/semana (~3 sessões de 60–90 min)
- **Estilo de tutoria:** socrático: sem solução pronta, dicas progressivas e code review
- **Elementos de metodologia ativos:** code review a cada entrega · revisão espaçada · katas · diário de aprendizado
- **Domínios de interesse para projetos:** backend, DevOps/cloud-native, infraestrutura (a refinar)
- **Preferência de aprendizado:** praticando com projetos; conceitos introduzidos entre projetos e desafios, nunca aulas teóricas longas

## Observações do tutor

<!-- Padrões notados ao longo do tempo: pontos fortes, dificuldades recorrentes, hábitos de TS/Node que aparecem no código Go. -->
- Atenção a hábitos de TS/Node que costumam aparecer: tratar erro como exceção, estruturas no estilo NestJS (DI container, pastas por camada), excesso de abstração/genéricos, assumir execução single-thread (races).
- 2026-11-13 (M01): tende a responder perguntas pela metade (pula subitens) e a tratar doc comment como formalidade; código limpo e boa separação de responsabilidades desde o início. Pedir respostas a todos os subitens.
- 2026-11-16 (M02): resposta parcial voltou a aparecer (checagem de constantes sem saídas/porquê). Tende a corrigir só a metade de um item com duas partes; ao apontar itens compostos, numerar as partes. Pede dica quando trava e aplica bem (N2 de erros foi aplicada certinho). Pediu para adiar doc comments → M03b.
- 2026-11-19 (M03): do parecer parcial, corrigiu os itens pontuais (bug, constantes, rand/v2, switch) e deixou de fora os dois que pediam repensar (mensagens de erro e estrutura dos loops). Marcou todos os critérios do DESAFIO como [x] antes do review. Revisões: bem nas perguntas de escopo e de visibilidade; erra o fluxo de `continue`/`break`. Review 2: corrigiu de novo só 1 de 3 itens e introduziu regressão ao mexer no flag. Sugerir ler o REVIEWS.md item a item antes de pedir re-review. Depois da dica N2, reestruturou bem e corrigiu o EOF em 1 rodada. Quando falta base (ponteiro), pesquisa sozinho e avisa: bom sinal de autonomia.
