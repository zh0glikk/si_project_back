package conveyor

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/si_project_back/internal/garbage"
)

type Conveyor struct {
	IsEnabled bool
	Speed int

	Garbage []garbage.Garbage

	GlassGarbage []garbage.Garbage
	PlasticGarbage []garbage.Garbage
	OrganicGarbage []garbage.Garbage
}

func NewConveyor(Speed int) *Conveyor {
	return &Conveyor{
		IsEnabled: true,
		Speed: Speed,
	}
}

func (c *Conveyor) Sort() bool {
	logrus.Info(c.Garbage)
	if len(c.Garbage) == 0 {
		return false
	}

	logrus.Info("Sorting started")

	for _, g := range c.Garbage {
		logrus.Info(g)
		if g.GarbageType == garbage.Glass {
			c.GlassGarbage = append(c.GlassGarbage, g)
			logrus.Info("add to glass")
		} else if g.GarbageType == garbage.Organic {
			c.OrganicGarbage = append(c.OrganicGarbage, g)
			logrus.Info("add to organic")
		} else if g.GarbageType == garbage.Plastic {
			c.PlasticGarbage = append(c.PlasticGarbage, g)
			logrus.Info("add to plastic")
		}
	}

	c.Garbage = c.Garbage[:0]

	return true
}

func (c *Conveyor) Transform() []garbage.TransformedGarbage {
	var result []garbage.TransformedGarbage

	if len(c.OrganicGarbage) == 0 && len(c.GlassGarbage) == 0 && len(c.PlasticGarbage) == 0 {
		return nil
	}

	logrus.Info("Start garbage transforming...")

	logrus.Info(c.PlasticGarbage)
	logrus.Info(c.OrganicGarbage)
	logrus.Info(c.GlassGarbage)

	logrus.Info("plastic")
	for _, g := range c.PlasticGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		dur := time.Duration(int64(g.Weight) / int64(c.Speed))  * time.Second

		logrus.Info(fmt.Sprintf("Transforming of %v during %v", g, dur))
		time.Sleep(dur)
	}
	logrus.Info("glass")
	for _, g := range c.GlassGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		dur := time.Duration(int64(g.Weight) / int64(c.Speed))  * time.Second

		logrus.Info(fmt.Sprintf("Transforming of %v during %v", g, dur))
		time.Sleep(dur)
	}
	logrus.Info("organic")
	for _, g := range c.OrganicGarbage {
		result = append(result, garbage.TransformedGarbage{Weight: g.Weight})

		dur := time.Duration(int64(g.Weight) / int64(c.Speed))  * time.Second

		logrus.Info(fmt.Sprintf("Transforming of %v during %v", g, dur))
		time.Sleep(dur)
	}

	c.GlassGarbage = c.GlassGarbage[:0]
	c.PlasticGarbage = c.PlasticGarbage[:0]
	c.OrganicGarbage = c.OrganicGarbage[:0]

	return result
}

func (c *Conveyor) ChangeSpeed(speed int) {
	c.Speed = speed
}

func (c * Conveyor) ChangeState() {
	c.IsEnabled = !c.IsEnabled
}