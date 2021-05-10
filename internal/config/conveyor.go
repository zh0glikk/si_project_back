package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Conveyor struct {
	Speed int `fig:"speed"`
}

type Conveyorer interface {
	Conveyor() *Conveyor
}

func NewConveyorer(getter kv.Getter) Conveyorer {
	return &conveyor{
		getter: getter,
	}
}

type conveyor struct {
	getter kv.Getter
	once   comfig.Once
}

func (l *conveyor) Conveyor() *Conveyor {
	return l.once.Do(func() interface{} {
		var config Conveyor
		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(l.getter, "conveyor")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out conveyor"))
		}

		return &config
	}).(*Conveyor)
}



