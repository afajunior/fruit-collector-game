package ldtk

import (
	"encoding/json"
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
	Identifier string          `json:"__identifier"`
	Type       string          `json:"__type"`
	Value      json.RawMessage `json:"__value"`
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

type Project struct {
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
	if l.Identifier != "Collisions" {
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
	if l.Identifier != "Entities" {
		return nil
	}
	for _, ent := range l.Entities {
		position, err := ent.GetPosition(l.GridSize)
		if err != nil {
			return err
		}

		e := w.NewEntity()

		switch ent.Identifier {
		case "Player":
			w.Players.Add(e, ecs.Player{})
			w.Velocities.Add(e, ecs.Velocity{})
			w.Sprites.Add(e, ecs.Sprite{Color: color.RGBA{0x3B, 0x82, 0xF6, 0xFF}})
		case "Fruit":
			points, err := GetField[int](&ent, "Points")

			if err != nil {
				return err
			}
			if points < 1 {
				return errors.New("Invalid fruit point")
			}

			fruitType, err := GetField[string](&ent, "Type")
			if err != nil {
				return err
			}

			w.Collectables.Add(e, ecs.Collectable{Points: points, Type: fruitType, Enabled: true})
			w.Sprites.Add(e, ecs.Sprite{Color: color.RGBA{0xEF, 0x44, 0x44, 0xFF}})
		}
		w.Positions.Add(e, ecs.Position{X: position[0], Y: position[1]})
	}
	return nil
}

func GetField[T any](e *Entity, name string) (T, error) {
	var zero T
	for _, f := range e.Fields {
		if f.Identifier == name {
			var v T
			err := json.Unmarshal(f.Value, &v)
			if err != nil {
				return zero, err
			}
			return v, err
		}
	}
	return zero, errors.New("Field not found")
}
