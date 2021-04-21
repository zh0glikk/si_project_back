package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramBot struct {
	Token         string   `fig:"token"`
}

type TelegramBoter interface {
	Bot() *TelegramBot
}

func NewTelegramBoter(getter kv.Getter) TelegramBoter {
	return &telegramBoter{
		getter: getter,
	}
}

type telegramBoter struct {
	getter kv.Getter
	once   comfig.Once
}

func (l *telegramBoter) Bot() *TelegramBot {
	return l.once.Do(func() interface{} {
		var config TelegramBot
		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(l.getter, "telegram_bot")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out stripe"))
		}


		return &config
	}).(*TelegramBot)
}

