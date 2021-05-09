package vehicle

import (
	"github.com/si_project_back/internal/garbage"
)

type Truck struct {
	Vehicle

	TransformedGarbage []garbage.TransformedGarbage
}

func NewTruck(capacity int) *Truck {
	return &Truck{
		Vehicle : Vehicle{
			Capacity: capacity,
		},
	}
}

func (v *Truck) AddGarbage(garbage garbage.TransformedGarbage) bool {
	if v.Fullness + garbage.Weight > v.Capacity {
		return false
	}

	v.TransformedGarbage = append(v.TransformedGarbage, garbage)

	v.Fullness += garbage.Weight

	return true
}

func (v *Truck) Dump() []garbage.TransformedGarbage {
	tmp := v.TransformedGarbage

	v.TransformedGarbage = v.TransformedGarbage[:0]
	v.Fullness = 0

	return tmp
}


