package ldtk

import (
	"errors"
	"image/color"

	"github.com/afajunior/fruit-collector-game/ecs"
)

type Entity struct {
	Identifier string         `json:"__identifier"`
	Px         []int          `json:"px"`
	Fields     []EntityFields `json:"fieldInstances"`
}

type EntityFields struct {
	Identifier string `json:"__identifier"`
	Type       string `json:"__type"`
	Value      any    `json:"__value"`
}

type Layer struct {
	Identifier string   `json:"__identifier"`
	Type       string   `json:"__type"`
	GridSize   int      `json:"__gridSize"`
	Entities   []Entity `json:"entityInstances"`
	Width      int      `json:"__cWid"`
	Height     int      `json:"__cHei"`
	GridCsv    []int    `json:"intGridCsv"`
}

type Level struct {
	Identifier   string  `json:"identifier"`
	Layers       []Layer `json:"layerInstances"`
	ScreenWidth  int     `json:"pxWid"`
	ScreenHeight int     `json:"pxHei"`
}

type Ltdk struct {
	Levels []Level `json:"levels"`
}

func (e *Entity) GetPosition(gridSize int) ([2]int, error) {
	if len(e.Px) != 2 {
		return [2]int{}, errors.New("Invalid px size")
	}
	if gridSize <= 0 {
		return [2]int{}, errors.New("Invalid grid size")
	}
	return [2]int{e.Px[0] / gridSize, e.Px[1] / gridSize}, nil
}

func (l *Layer) SpawnCollisions(w *ecs.World) {
	if l.Type == "Collisions" {
		return
	}
	for i, v := range l.GridCsv {
		if v == 0 {
			continue
		}
		e := w.NewEntity()
		w.Obstacles.Add(e, ecs.Obstacle{})
		w.Positions.Add(e, ecs.Position{X: i % l.Width, Y: i / l.Width})
		w.Sprites.Add(e, ecs.Sprite{Color: color.RGBA{0x6B, 0x72, 0x80, 0xFF}})
	}
}

func (l *Layer) SpawnEntities(w *ecs.World) error {
	if l.Type == "entities" {
		return nil
	}
	for _, ent := range l.Entities {
		position, error := ent.GetPosition(l.GridSize)
		if error != nil {
			return error
		}

		e := w.NewEntity()
		w.Positions.Add(e, ecs.Position{X: position[0], Y: position[1]})

		switch ent.Identifier {
		case "Player":
			w.Players.Add(e, ecs.Player{})
			w.Velocities.Add(e, ecs.Velocity{})
			w.Sprites.Add(e, ecs.Sprite{Color: color.RGBA{0x3B, 0x82, 0xF6, 0xFF}})
		case "Fruit":
			points := 0
			for _, field := range ent.Fields {
				switch field.Identifier {
				case "Type":
				case "Points":
					/*points, error =
					if error != nil {
						return nil
					}*/
				}
			}
			w.Collectables.Add(e, ecs.Collectable{Points: points, Enabled: true})
			w.Sprites.Add(e, ecs.Sprite{Color: color.RGBA{0xEF, 0x44, 0x44, 0xFF}})
		}
	}
	return nil
}
