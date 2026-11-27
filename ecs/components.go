package ecs

import (
	"image"
	"image/color"
)

type Player struct {
	Points int
}

type Sprite struct {
	Color color.RGBA
	Image image.Image
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
	Type    string
}

type Obstacle struct {
}
