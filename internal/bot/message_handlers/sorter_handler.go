package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/si_project_back/internal/bot/keyboards"
)

type SorterHandler struct {
	handler *Handler
}

func NewSorterHandler(update tgbotapi.Update) *SorterHandler {
	return &SorterHandler{
		handler: NewHandler(update),
	}
}

func (h SorterHandler) Handle() tgbotapi.MessageConfig {
	chatId := h.handler.update.Message.Chat.ID
	text := "Sorter action panel."

	return tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              chatId,
			ChannelUsername:     "",
			ReplyToMessageID:    0,
			ReplyMarkup:         keyboards.SorterKeyboard,
			DisableNotification: false,
		},
		Text:                  text,
		ParseMode:             "",
		DisableWebPagePreview: false,
	}
}

