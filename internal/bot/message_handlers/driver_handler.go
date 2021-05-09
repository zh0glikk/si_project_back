package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/keyboards"
)

type DriverHandler struct {
	handler *Handler
}

func NewDriverHandler(update tgbotapi.Update) *DriverHandler {
	return &DriverHandler{
		handler: NewHandler(update),
	}
}

func (h DriverHandler) Handle() tgbotapi.MessageConfig {
	chatId := h.handler.update.Message.Chat.ID
	text := "Driver action panel."

	return tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              chatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         keyboards.DriverKeyboard,
			DisableNotification: false,
		},
		Text:                  text,
		ParseMode:             "",
		DisableWebPagePreview: false,
	}
}


