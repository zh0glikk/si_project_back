package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Vehicle struct {
	Capacity int `fig:"capacity"`
}

type Vehicler interface {
	Vehicle() *Vehicle
}

func NewVehicler(getter kv.Getter) Vehicler {
	return &vehicler{
		getter: getter,
	}
}

type vehicler struct {
	getter kv.Getter
	once   comfig.Once
}

func (l *vehicler) Vehicle() *Vehicle {
	return l.once.Do(func() interface{} {
		var config Vehicle
		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(l.getter, "vehicle")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out vehicle"))
		}

		return &config
	}).(*Vehicle)
}


