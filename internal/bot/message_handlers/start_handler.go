package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/keyboards"
)


type StartHandler struct {
	handler *Handler
}

func NewStartHandler(update tgbotapi.Update) *StartHandler {
	return &StartHandler{
		handler: NewHandler(update),
	}
}

func (h StartHandler) Handle() tgbotapi.MessageConfig {
	chatId := h.handler.update.Message.Chat.ID
	text := "Choose action."
	
	return tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              chatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         keyboards.MainKeyboard,
			DisableNotification: false,
		},
		Text:                  text,
		ParseMode:             "",
		DisableWebPagePreview: false,
	}
}

