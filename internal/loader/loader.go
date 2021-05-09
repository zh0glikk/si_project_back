package loader

import (
	"github.com/si_project_back/internal/garbage"
)

type Loader struct {
	TransformedGarbage  []garbage.TransformedGarbage
	CollectedGarbage 	[]garbage.Garbage
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Dump() []garbage.Garbage {
	res := l.CollectedGarbage

	l.CollectedGarbage = l.CollectedGarbage[:0]

	return res
}

func (l *Loader) DumpTransformedGarbage() []garbage.TransformedGarbage {
	res := l.TransformedGarbage

	l.TransformedGarbage = l.TransformedGarbage[:0]

	return res
}


