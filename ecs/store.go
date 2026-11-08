package ecs

type Store[T any] struct {
	data map[Entity]*T
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{data: make(map[Entity]*T)}
}

func (s *Store[T]) Add(e Entity, c T) {
	s.data[e] = &c
}

func (s *Store[T]) Get(e Entity) (*T, bool) {
	c, ok := s.data[e]
	return c, ok
}

// Funções da interface Filter
func (s *Store[T]) Has(e Entity) bool {
	_, ok := s.data[e]
	return ok
}

func (s *Store[T]) Del(e Entity) {
	delete(s.data, e)
}

func (s *Store[T]) Entities() []Entity {
	out := make([]Entity, 0, len(s.data))
	for e := range s.data {
		out = append(out, e)
	}
	return out
}
