package systems

import (
	"image/color"
	"sort"

	"github.com/afajunior/fruit-collector-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RenderSystem struct {
	CellSize  int
	GridColor color.RGBA
}

func (s *RenderSystem) Draw(w *ecs.World, screen *ebiten.Image) {
	entities := w.Query(w.Sprites, w.Positions)

	sort.Slice(entities, func(i, j int) bool {
		pi, pj := w.Players.Has(entities[i]), w.Players.Has(entities[j])
		if pi != pj {
			return pj
		}
		return entities[i] < entities[j]
	})

	cell := float32(s.CellSize)
	const padding = 2

	for _, e := range entities {
		p, _ := w.Positions.Get(e)
		sp, _ := w.Sprites.Get(e)
		c, _ := w.Collectables.Get(e)

		if c != nil && !c.Enabled {
			continue
		}
		x := float32(p.X) * cell
		y := float32(p.Y) * cell
		if sp.Image == nil {
			vector.DrawFilledRect(screen, x+padding, y+padding, cell-2*padding, cell-2*padding, sp.Color, false)
		} else {
			opts := &ebiten.DrawImageOptions{}
			dx := float64(p.X * s.CellSize)
			dy := float64(p.Y * s.CellSize)
			opts.GeoM.Translate(dx, dy)
			screen.DrawImage(sp.Image, opts)
		}
	}
}

func (s *RenderSystem) DrawGrid(screen *ebiten.Image, cols, rows int) {
	cell := float32(s.CellSize)
	width, height := float32(cols)*cell, float32(rows)*cell
	for c := 0; c <= cols; c++ {
		x := float32(c) * cell
		vector.StrokeLine(screen, x, 0, x, height, 1, s.GridColor, false)
	}

	for r := 0; r <= rows; r++ {
		y := float32(r) * cell
		vector.StrokeLine(screen, 0, y, width, y, 1, s.GridColor, false)
	}
}
