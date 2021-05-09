package bot

import (
	"github.com/sirupsen/logrus"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/message_handlers"
	"github.com/si_project_back/internal/config"
	"github.com/si_project_back/internal/garbage"
)

func RunBot(cfg config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.Bot().Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			if !update.Message.IsCommand() {
				continue
			}

			logrus.Info(update.Message.Text)

			var msg tgbotapi.MessageConfig

			switch update.Message.Command() {
			case "start":
				msg = message_handlers.NewStartHandler(update).Handle()
			case "back":
				msg = message_handlers.NewStartHandler(update).Handle()

			case "sorter":
				msg = message_handlers.NewSorterHandler(update).Handle()
			case "sort":
				continue

			case "driver":
				msg = message_handlers.NewDriverHandler(update).Handle()
			case "add_garbage_glass":
				msg = message_handlers.NewAddGarbage(update, garbage.Glass).Handle()
			case "add_garbage_plastic":
				msg = message_handlers.NewAddGarbage(update, garbage.Plastic).Handle()
			case "add_garbage_organic":
				msg = message_handlers.NewAddGarbage(update, garbage.Organic).Handle()


			case "operator":
				continue

			case "loader":
				msg = message_handlers.NewLoaderHandler(update).Handle()
			case "unload_garbage_truck":
				msg = message_handlers.NewUnloadGarbageTruck(update).Handle()
			case "load_truck":
				continue
			case "load_conv":
				continue
			}


			if _, err := bot.Send(msg); err != nil {
				logrus.Info("Can't send message")
			}
		}

	}
}

