package vehicle

import (
	"github.com/si_project_back/internal/garbage"
)

type GarbageTruck struct {
	Vehicle

	CollectedGarbage []garbage.Garbage
}

func NewGarbageTruck(capacity int) *GarbageTruck {
	return &GarbageTruck{
		Vehicle: Vehicle{
			Capacity: capacity,
		},
	}
}

func (v *GarbageTruck) AddGarbage(garbage garbage.Garbage) bool {
	if v.Fullness + garbage.Weight > v.Capacity {
		return false
	}

	v.CollectedGarbage = append(v.CollectedGarbage, garbage)

	v.Fullness += garbage.Weight

	return true
}

func (v *GarbageTruck) Dump() []garbage.Garbage {
	tmp := v.CollectedGarbage

	v.CollectedGarbage = v.CollectedGarbage[:0]
	v.Fullness = 0

	return tmp
}
