package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/afajunior/fruit-collector-game/ecs"
	"github.com/afajunior/fruit-collector-game/ldtk"
	"github.com/afajunior/fruit-collector-game/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gridCols = 16
	gridRows = 16
	cellSize = 18

	screenWidth  = gridCols * cellSize
	screenHeight = gridRows * cellSize
)

type Game struct {
	world        *ecs.World
	systems      []ecs.System
	renderer     systems.RenderSystem
	ldtk         *ldtk.Ltdk
	currentLevel int
}

func NewGame() *Game {
	ldtk := loadLtdk("phase1.ldtk")
	currentLevel := ldtk.Levels[0]
	colsLvl1 := currentLevel.Layers[0].Width
	rowsLvl1 := currentLevel.Layers[0].Height
	cellSize := currentLevel.Layers[0].GridSize
	g := &Game{
		world: ecs.NewWorld(),
		systems: []ecs.System{
			&systems.InputSystem{},
			&systems.MovementSystem{Cols: colsLvl1, Rows: rowsLvl1},
			&systems.CollisionSystem{},
		},
		renderer: systems.RenderSystem{
			CellSize:  cellSize,
			GridColor: color.RGBA{0xE0, 0xE0, 0xE0, 0xFF},
		},
		ldtk:         ldtk,
		currentLevel: 0,
	}
	g.setupLevel()
	return g
}

func (g *Game) setupLevel() {
	w := g.world
	currentLevel := g.ldtk.Levels[g.currentLevel]
	for _, layer := range currentLevel.Layers {
		layer.SpawnCollisions(w)
		error := layer.SpawnEntities(w)
		if error != nil {
			panic(error)
		}
	}
}

func loadLtdk(file string) *ldtk.Ltdk {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	res := ldtk.Ltdk{}
	err = json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	return &res
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xFF, 0xFB, 0xEB, 0xFF})
	g.renderer.Draw(g.world, screen)
	g.renderer.DrawGrid(screen, gridCols, gridRows)
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

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Fruir Collector")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
