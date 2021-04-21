package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	TelegramBoter
}

type config struct {
	getter kv.Getter

	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	TelegramBoter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter:           getter,
		Logger:           comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser:        pgdb.NewDatabaser(getter),
		Listenerer:		  comfig.NewListenerer(getter),
		TelegramBoter:	  NewTelegramBoter(getter),
	}
}

