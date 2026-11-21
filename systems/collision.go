package systems

import (
	"github.com/afajunior/fruit-collector-game/ecs"
)

type CollisionSystem struct{}

func (s *CollisionSystem) Update(w *ecs.World) error {
	players := w.Query(w.Players, w.Positions)
	collectables := w.Query(w.Collectables, w.Positions)
	for _, p := range players {
		pPos, _ := w.Positions.Get(p)
		dx, dy := pPos.X, pPos.Y

		for _, c := range collectables {
			pCol, _ := w.Positions.Get(c)
			if dx == pCol.X && dy == pCol.Y {
				player, _ := w.Players.Get(p)
				fruit, _ := w.Collectables.Get(c)
				if fruit.Enabled {
					player.Points += fruit.Points
					fruit.Enabled = false
				}
			}
		}
	}
	return nil
}
