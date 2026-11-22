package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/afajunior/fruit-collector-game/ecs"
	"github.com/afajunior/fruit-collector-game/ldtk"
	"github.com/afajunior/fruit-collector-game/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	world        *ecs.World
	systems      []ecs.System
	renderer     systems.RenderSystem
	ldtk         *ldtk.Project
	currentLevel int
	screenWidth  int
	screenHeight int
}

func NewGame() *Game {
	ldtkSchema := loadLdtk("phase1.ldtk")
	currentLevel := ldtkSchema.Levels[0]
	collisionsLayer, err := currentLevel.GetLayer("Collisions")
	if err != nil {
		panic(err)
	}
	numCols := collisionsLayer.Width
	numRows := collisionsLayer.Height
	cellSize := collisionsLayer.GridSize
	g := &Game{
		world: ecs.NewWorld(),
		systems: []ecs.System{
			&systems.InputSystem{},
			&systems.MovementSystem{Cols: numCols, Rows: numRows},
			&systems.CollisionSystem{},
		},
		renderer: systems.RenderSystem{
			CellSize:  cellSize,
			GridColor: color.RGBA{0xE0, 0xE0, 0xE0, 0xFF},
		},
		ldtk:         ldtkSchema,
		currentLevel: 0,
		screenWidth:  numCols * cellSize,
		screenHeight: numRows * cellSize,
	}
	sprites, err := loadSprites()
	if err != nil {
		log.Fatal(err)
	}
	g.setupLevel(sprites)
	return g
}

func (g *Game) setupLevel(sprites map[string]*ebiten.Image) {
	w := g.world
	currentLevel := g.ldtk.Levels[g.currentLevel]
	for _, layer := range currentLevel.Layers {
		layer.SpawnCollisions(w)
		err := layer.SpawnEntities(w, sprites)
		if err != nil {
			panic(err)
		}
	}
}

func loadLdtk(file string) *ldtk.Project {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	res := ldtk.Project{}
	err = json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	return &res
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xFF, 0xFB, 0xEB, 0xFF})
	g.renderer.Draw(g.world, screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	currentLevel := g.ldtk.Levels[g.currentLevel]
	return currentLevel.ScreenWidth, currentLevel.ScreenHeight
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	for _, s := range g.systems {
		if err := s.Update(g.world); err != nil {
			return err
		}
	}
	return nil
}

func loadSprites() (map[string]*ebiten.Image, error) {
	entries, err := os.ReadDir("./assets/sprites")
	if err != nil {
		return nil, err
	}
	var m = make(map[string]*ebiten.Image)
	for _, e := range entries {
		filename := e.Name()
		img, _, err := ebitenutil.NewImageFromFile("./assets/sprites/" + filename)
		if err != nil {
			return nil, err
		}
		fruitName := strings.TrimSuffix(filename, ".png")
		m[fruitName] = img
	}
	return m, nil
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(g.screenWidth*2, g.screenHeight*2)
	ebiten.SetWindowTitle("Fruit Collector")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
