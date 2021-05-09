package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/keyboards"
)

type LoaderHandler struct {
	handler *Handler
}

func NewLoaderHandler(update tgbotapi.Update) *LoaderHandler {
	return &LoaderHandler{
		handler: NewHandler(update),
	}
}

func (h LoaderHandler) Handle() tgbotapi.MessageConfig {
	chatId := h.handler.update.Message.Chat.ID
	text := "Loader action panel."

	return tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              chatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         keyboards.LoaderKeyboard,
			DisableNotification: false,
		},
		Text:                  text,
		ParseMode:             "",
		DisableWebPagePreview: false,
	}
}


