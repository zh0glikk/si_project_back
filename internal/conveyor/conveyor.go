package conveyor

import (
	"github.com/si_project_back/internal/garbage"
)

type Conveyor struct {
	IsEnabled bool
	Speed float64

	Garbage []garbage.Garbage

	GlassGarbage []garbage.Garbage
	PlasticGarbage []garbage.Garbage
	OrganicGarbage []garbage.Garbage
}

func NewConveyor(Speed float64) *Conveyor {
	return &Conveyor{
		IsEnabled: true,
		Speed: Speed,
	}
}

func (c *Conveyor) Sort() {
	for i, g := range c.Garbage {
		if g.GarbageType == garbage.Glass {
			c.Garbage = append(c.Garbage[:i], c.Garbage[i+1:]...)

			c.GlassGarbage = append(c.GlassGarbage, g)
		} else if g.GarbageType == garbage.Organic {
			c.Garbage = append(c.Garbage[:i], c.Garbage[i+1:]...)

			c.OrganicGarbage = append(c.GlassGarbage, g)
		} else if g.GarbageType == garbage.Plastic {
			c.Garbage = append(c.Garbage[:i], c.Garbage[i+1:]...)

			c.PlasticGarbage = append(c.GlassGarbage, g)
		}
	}
}

func (c *Conveyor) Transform() []garbage.TransformedGarbage {
	var result []garbage.TransformedGarbage

	for i, g := range c.PlasticGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		c.PlasticGarbage = append(c.PlasticGarbage[:i], c.PlasticGarbage[i+1:]...)
	}
	for i, g := range c.GlassGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		c.GlassGarbage = append(c.GlassGarbage[:i], c.GlassGarbage[i+1:]...)
	}
	for i, g := range c.OrganicGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		c.OrganicGarbage = append(c.OrganicGarbage[:i], c.OrganicGarbage[i+1:]...)
	}

	return result
}

