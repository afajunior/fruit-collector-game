package systems

import (
	"github.com/afajunior/fruit-collector-game/ecs"
)

type MovementSystem struct {
	Cols, Rows int
}

func (s *MovementSystem) Update(w *ecs.World) error {
	for _, e := range w.Query(w.Velocities, w.Positions) {
		v, _ := w.Velocities.Get(e)
		if v.DX == 0 && v.DY == 0 {
			continue
		}
		p, _ := w.Positions.Get(e)

		nx, ny := p.X+v.DX, p.Y+v.DY

		if s.inBounds(nx, ny) && !blocked(w, nx, ny) {
			p.X, p.Y = nx, ny
		}

		v.DX, v.DY = 0, 0
	}
	return nil
}

func (s *MovementSystem) inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < s.Cols && y < s.Rows
}

func blocked(w *ecs.World, x, y int) bool {
	for _, e := range w.Query(w.Obstacles, w.Positions) {
		p, _ := w.Positions.Get(e)
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}
