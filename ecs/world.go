package ecs

type World struct {
	nextID   Entity
	entities map[Entity]struct{}

	Players      *Store[Player]
	Sprites      *Store[Sprite]
	Positions    *Store[Position]
	Velocities   *Store[Velocity]
	Collectables *Store[Collectable]
	Obstacles    *Store[Obstacle]
}

type Filter interface {
	Has(e Entity) bool
	Entities() []Entity
}

func NewWorld() *World {
	return &World{
		nextID:   1,
		entities: make(map[Entity]struct{}),

		Players:      NewStore[Player](),
		Sprites:      NewStore[Sprite](),
		Positions:    NewStore[Position](),
		Velocities:   NewStore[Velocity](),
		Collectables: NewStore[Collectable](),
		Obstacles:    NewStore[Obstacle](),
	}
}

func (w *World) NewEntity() Entity {
	e := w.nextID
	w.nextID++
	w.entities[e] = struct{}{}
	return e
}

func (w *World) Query(first Filter, rest ...Filter) []Entity {
	var out []Entity
	for _, e := range first.Entities() {
		match := true
		for _, f := range rest {
			if !f.Has(e) {
				match = false
				break
			}
		}
		if match {
			out = append(out, e)
		}
	}

	return out
}
