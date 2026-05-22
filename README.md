# 🍎 Fruit Collector - ECS Game

Um jogo educativo de coletar frutas desenvolvido em **Go** usando a engine **Ebiten** e arquitetura **ECS (Entity Component System)**.

> **Público-alvo:** Crianças a partir de 6 anos  
> **Objetivo:** Aprender ECS, game design e desenvolvimento de jogos 2D

---

## 🎮 Sobre o Jogo

Controle um personagem colorido em um grid, colete todas as frutas para avançar de nível, desvie de obstáculos e pegue power-ups especiais!

### Características
- ✨ Visual colorido e amigável para crianças
- 🎵 Efeitos sonoros divertidos
- 🎯 Múltiplos níveis com dificuldade progressiva
- ⚡ Power-ups que aumentam velocidade temporariamente
- 🏆 Sistema de pontuação
- 🎨 Sprites animados

---

## 🏗️ Arquitetura ECS

O projeto utiliza **Entity Component System**, um padrão de design que separa dados (Components) de lógica (Systems).

### Components (Dados)
- `Position` - Localização no grid
- `Sprite` - Aparência visual
- `Velocity` - Direção de movimento
- `Collectable` - Itens coletáveis com pontos
- `Obstacle` - Objetos sólidos
- `Player` - Marca o jogador
- `Animation` - Frames de animação
- `Audio` - Sons associados

### Systems (Lógica)
- `InputSystem` - Captura teclado
- `MovementSystem` - Move entidades no grid
- `CollisionSystem` - Detecta coletas e colisões
- `RenderSystem` - Desenha na tela
- `AnimationSystem` - Atualiza frames
- `AudioSystem` - Toca efeitos sonoros

### World (Gerenciador)
Armazena todas as entidades e fornece queries para os systems.

---

## 🚀 Como Executar

### Pré-requisitos
- Go 1.21 ou superior
- Sistema operacional: Linux, macOS ou Windows

### Instalação

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/fruit-collector.git
cd fruit-collector

# Baixe as dependências
go mod download

# Execute o jogo
go run main.go
```

### Compilar

```bash
# Build para seu sistema
go build -o fruit-collector

# Executar
./fruit-collector
```

---

## 🎮 Controles

| Tecla | Ação |
|-------|------|
| `↑` ou `W` | Mover para cima |
| `↓` ou `S` | Mover para baixo |
| `←` ou `A` | Mover para esquerda |
| `→` ou `D` | Mover para direita |
| `ESC` | Sair do jogo |

---

## 📁 Estrutura do Projeto

```
fruit-collector/
├── main.go              # Entry point e game loop
├── ecs/
│   ├── components.go    # Definição de todos os components
│   └── world.go         # Gerenciador de entidades
├── systems/
│   ├── input.go         # Sistema de input
│   ├── movement.go      # Sistema de movimento
│   ├── collision.go     # Sistema de colisão
│   ├── render.go        # Sistema de renderização
│   ├── animation.go     # Sistema de animação
│   └── audio.go         # Sistema de áudio
├── assets/
│   ├── sprites/         # Imagens PNG (32x32)
│   │   ├── player.png
│   │   ├── apple.png
│   │   ├── banana.png
│   │   ├── cherry.png
│   │   └── rock.png
│   ├── sounds/          # Efeitos sonoros (.wav)
│   │   ├── collect.wav
│   │   └── victory.wav
│   └── levels/          # Níveis LDtk (futuro)
├── go.mod
├── go.sum
└── README.md
```

---

## 📋 Roadmap de Desenvolvimento

### ✅ Sprint 1: Fundação ECS (5 dias)
- [x] Components definidos
- [x] Sistema de entidades (World)
- [x] Input do jogador
- [x] Movimento em grid
- [x] Renderização básica

### ✅ Sprint 2: Colisão e Game Loop (3 dias)
- [x] Sistema de coleta de itens
- [x] Loop principal integrado
- [x] Primeiro nível jogável

### ✅ Sprint 3: Polimento Visual (3 dias)
- [x] Sistema de animações
- [x] Sprites coloridos
- [x] GUI com placar e vitória

### ✅ Sprint 4: Áudio e Progressão (3 dias)
- [x] Efeitos sonoros
- [x] 3 níveis diferentes
- [x] Power-up de velocidade

### 🔮 Backlog Futuro
- [ ] Integração com LDtk para criar níveis
- [ ] Menu inicial
- [ ] Música de fundo
- [ ] Efeitos de partículas
- [ ] Salvar high score

Veja os tickets detalhados em [`/sprint-tickets`](./sprint-tickets/)

---

## 🎨 Assets

### Sprites
- Dimensões: 32×32 pixels
- Formato: PNG com transparência
- Estilo: Pixel art colorido

### Onde Encontrar Assets Gratuitos
- [itch.io - Game Assets](https://itch.io/game-assets/free)
- [OpenGameArt.org](https://opengameart.org/)
- [Kenney.nl](https://kenney.nl/assets)

### Ferramentas de Criação
- **Sprites:** [Piskel](https://www.piskelapp.com/), [Aseprite](https://www.aseprite.org/)
- **Sons:** [jfxr](https://jfxr.frozenfractal.com/), [Freesound](https://freesound.org/)
- **Níveis:** [LDtk](https://ldtk.io/)

---

## 🧪 Testes

```bash
# Executar testes unitários
go test ./...

# Executar com coverage
go test -cover ./...

# Testes com verbose
go test -v ./...
```

---

## 🛠️ Tecnologias

- **Linguagem:** [Go 1.21+](https://golang.org/)
- **Engine:** [Ebiten 2.6+](https://ebiten.org/)
- **Padrão:** Entity Component System (ECS)
- **Editor de Níveis:** [LDtk](https://ldtk.io/) (futuro)

---

## 📚 Recursos de Aprendizado

### ECS
- [Understanding ECS](https://www.gamedev.net/tutorials/programming/general-and-gameplay-programming/understanding-component-entity-systems-r3013/)
- [ECS Back and Forth](https://skypjack.github.io/2019-02-14-ecs-baf-part-1/)

### Ebiten
- [Documentação Oficial](https://ebiten.org/documents/)
- [Exemplos de Jogos](https://ebiten.org/examples/)
- [Tour of Ebiten](https://ebiten.org/tour/)

### Game Design para Crianças
- [Designing Games for Kids](https://www.gamedeveloper.com/design/designing-games-for-kids-what-you-need-to-know)
- [Color Theory in Games](https://www.gamedeveloper.com/design/color-theory-for-game-design)

---

## 📝 Notas de Desenvolvimento

### Boas Práticas ECS
- Components são dados puros (sem métodos)
- Systems têm responsabilidade única
- World gerencia lifecycle das entidades
- Queries por components para filtrar entidades

---

## 🎯 Objetivos de Aprendizado

Este projeto foi criado para praticar:

- ✅ Arquitetura ECS (Entity Component System)
- ✅ Game loop e delta time
- ✅ Sistema de input
- ✅ Colisão e física simples em grid
- ✅ Renderização 2D com Ebiten
- ✅ Gerenciamento de assets (sprites, sons)
- ✅ Animações frame-by-frame
- ✅ Sistema de áudio
- ✅ Level design e progressão
- ✅ GUI básica (placar, mensagens)
- 🔄 Integração com ferramentas externas (LDtk)

---

### Exemplos:
- **Desafio:** Como fazer colisão eficiente em grid?
- **Solução:** Usar `GetEntitiesAt(x, y)` para query espacial
- **Aprendizado:** ECS facilita adicionar novos tipos de entidades sem modificar systems existentes

---

**Bom aprendizado! 🍎🍌🍓**