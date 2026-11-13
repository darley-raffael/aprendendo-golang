---
name: iniciar
description: Inicializa o aprendizado de uma nova linguagem/tecnologia num repositório criado a partir do learning-template. Entrevista o aluno, gera ROADMAP.md, PERFIL, PROGRESSO, CONCEITOS e a seção de comandos do CLAUDE.md. Use quando o aluno disser "/iniciar <tecnologia>", "inicialize o aprendizado de X" ou quando PERFIL.md ainda tiver placeholders {{...}}.
---

# /iniciar <tecnologia>

Objetivo: transformar o template genérico num plano de estudos **completo, profundo e progressivo** para a tecnologia informada, sob medida para o aluno.

## 1. Verificações

- Leia `CLAUDE.md`, `METODOLOGIA.md` e `progresso/PERFIL.md`.
- Se o `PERFIL.md` **não** tiver mais `{{`, o repo já foi inicializado: pergunte se o aluno quer **replanejar** (mantendo o histórico) ou cancelar.
- Se a tecnologia não veio como argumento, pergunte.
- Verifique se a toolchain está instalada (ex.: `go version`, `rustc --version`, `dotnet --version`). Se não estiver, mostre como instalar e siga mesmo assim.
- Na memória do Claude, procure um perfil de usuário já salvo (stack, ritmo, estilo). Se existir, **confirme** em vez de perguntar de novo.

## 2. Entrevista (AskUserQuestion, no máximo 2 rodadas de até 4 perguntas)

Pergunte apenas o que ainda não se sabe:
1. **Stack atual** (base para as comparações)
2. **Experiência prévia** com a tecnologia (zero, já brinquei, uso no trabalho)
3. **Objetivos** (multiSelect, com opções específicas da tecnologia: ex. para Rust, CLIs, sistemas, WebAssembly, backend, embarcados)
4. **Tempo semanal** (até 5h, 5–10h, mais de 10h)
5. **Estilo de tutoria** (socrático recomendado, equilibrado, guiado)
6. **Elementos de metodologia** (code review, revisão espaçada, katas, diário)
7. **Domínios de interesse** para os projetos (finanças, jogos, DevOps, dados…), para os projetos ficarem motivadores

## 3. Gerar o ROADMAP.md

Siga **exatamente** o formato descrito em `METODOLOGIA.md` ("Formato do roadmap"). Requisitos de qualidade:

- **Completo e não superficial:** do zero até nível de produção/mercado. Inclua os tópicos avançados que um profissional sênior da tecnologia domina (internals, performance, concorrência, tooling, testes, deploy, ecossistema).
- **Progressão real e gradual:**
  - **Fase 0 é sempre uma rampa de fundamentos:** 4–6 micro-projetos de uma sessão (1–4 conceitos básicos cada, com andaime decrescente de testes), depois o diagnóstico, e só então um projeto de consolidação pequeno.
  - **Toda fase seguinte abre com uma rampa** de 2–3 micro-projetos que isolam os conceitos novos mais difíceis, seguida de diagnóstico, antes do projeto grande.
  - Cada projeto adiciona no máximo ~6 conceitos inéditos além dos da rampa e declara **Aprofunda:** com conceitos de fases anteriores, revisitados mais fundo.
  - Nenhum conceito é citado (rampa, kata, aprofunda) antes da fase em que é introduzido. Nada de conceito sem projeto que o use.
- **Projetos com cara de mundo real**, ligados aos objetivos e domínios do aluno. Cada projeto dividido em **etapas** de 1–3 sessões.
- **Katas** por fase, ligadas aos conceitos da fase.
- **Ponte com a stack do aluno** em todo projeto: equivalências e, principalmente, **armadilhas** de trazer hábitos antigos.
- **Duração estimada** por fase, calculada pelo tempo semanal do aluno. Mostre o total.
- Tabela-resumo das fases no início, seção de recursos oficiais no fim (documentação, specs, livros de referência).
- Use a versão **atual** estável da tecnologia (confira a instalada e, se possível, as notas de release recentes).

## 4. Preencher os demais arquivos

- `progresso/PERFIL.md`: todos os placeholders.
- `progresso/PROGRESSO.md`: fase 0/etapa 1, status "não iniciado", próxima ação "rodar /sessao e depois /desafio", checklist de fases e projetos extraído do roadmap.
- `progresso/CONCEITOS.md`: **uma linha por conceito** listado nos "Conceitos introduzidos" do roadmap, com a fase, nível 0 e datas em branco (`-`).
- `CLAUDE.md`: troque `{{TECNOLOGIA}}`; substitua `{{COMANDOS_TECNOLOGIA}}` pelos comandos de build, test, lint, format, ferramentas de análise e docs locais que o `/review` deve usar.
- `ROADMAP.md`: troque `{{TECNOLOGIA}}` e `{{ROADMAP}}`.
- `README.md`: adapte para "Aprendendo <tecnologia>", com link para o ROADMAP e um resumo das fases. Mantenha a menção de que o repo veio do learning-template.

Ao final, `grep -n "{{" README.md ROADMAP.md progresso/*.md` não pode retornar nada, e no `CLAUDE.md` só pode sobrar a linha de aviso sobre inicialização.

## 5. Memória do Claude

- Salve/atualize uma memória **user** com stack, ritmo e estilo de tutoria preferido (vale para todas as tecnologias).
- Salve uma memória **project** apontando para este repositório como fonte da verdade do progresso.

## 6. Fechamento

- Se não for repositório git, rode `git init`. Faça commit: `chore: inicializa aprendizado de <tecnologia>` (peça confirmação antes de dar push).
- Mostre ao aluno: resumo das fases com a duração, o primeiro projeto e como começar (`/sessao` → `/desafio`).
