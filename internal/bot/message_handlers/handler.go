package message_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	update tgbotapi.Update
}

func NewHandler(update tgbotapi.Update) *Handler {
	return &Handler{update: update}
}

type IHandler interface {
	Handle() tgbotapi.MessageConfig
}
