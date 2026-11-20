package ecs

import "image/color"

type Player struct {
	Points int
}

type Sprite struct {
	//TODO: Adicionar ebiten.Image para PNGs
	Color color.RGBA
}

type Position struct {
	X, Y int
}

type Velocity struct {
	DX, DY int
}

type Collectable struct {
	Points  int
	Enabled bool
	Type string
}

type Obstacle struct {
}
