package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/keyboards"
)

type OperatorHandler struct {
	handler *Handler
}

func NewOperatorHandler(update tgbotapi.Update) *OperatorHandler {
	return &OperatorHandler{
		handler: NewHandler(update),
	}
}

func (h OperatorHandler) Handle() tgbotapi.MessageConfig {
	chatId := h.handler.update.Message.Chat.ID
	text := "Operator action panel."

	return tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              chatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         keyboards.OperatorKeyboard,
			DisableNotification: false,
		},
		Text:                  text,
		ParseMode:             "",
		DisableWebPagePreview: false,
	}
}

