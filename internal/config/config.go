package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
	TelegramBoter
	Vehicler
	Conveyorer
}

type config struct {
	getter kv.Getter

	comfig.Logger
	comfig.Listenerer
	TelegramBoter
	Vehicler
	Conveyorer
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter:           getter,
		Logger:           comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Listenerer:		  comfig.NewListenerer(getter),
		TelegramBoter:	  NewTelegramBoter(getter),
		Vehicler: 		  NewVehicler(getter),
		Conveyorer:		  NewConveyorer(getter),
	}
}

