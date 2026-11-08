package systems

import (
	"github.com/afajunior/fruit-collector-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputSystem struct{}

func (s *InputSystem) Update(w *ecs.World) error {
	dx, dy := readDirection()

	for _, e := range w.Query(w.Players, w.Velocities) {
		v, _ := w.Velocities.Get(e)
		v.DX, v.DY = dx, dy
	}
	return nil
}

// TODO: Implementar segurar tecla para movimento contínuo
func readDirection() (dx, dy int) {
	switch {
	case justPressed(ebiten.KeyArrowUp, ebiten.KeyW):
		return 0, -1
	case justPressed(ebiten.KeyArrowDown, ebiten.KeyS):
		return 0, 1
	case justPressed(ebiten.KeyArrowLeft, ebiten.KeyA):
		return -1, 0
	case justPressed(ebiten.KeyRight, ebiten.KeyD):
		return 1, 0
	}
	return 0, 0
}

func justPressed(keys ...ebiten.Key) bool {
	for _, k := range keys {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}
