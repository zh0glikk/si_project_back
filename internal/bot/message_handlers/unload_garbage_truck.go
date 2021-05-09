package message_handlers

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/url"

	"github.com/si_project_back/internal/bot/keyboards"
	"github.com/si_project_back/pkg/connector"
)

type UnloadGarbageTruck struct {
	handler *Handler
}

func NewUnloadGarbageTruck(update tgbotapi.Update) *UnloadGarbageTruck {
	return &UnloadGarbageTruck{
		handler: NewHandler(update),
	}
}

func (h UnloadGarbageTruck) Handle() tgbotapi.MessageConfig {
	conn := connector.NewConnector(&url.URL{
		Scheme:      "http",
		Host:        "127.0.0.1:8080",
		ForceQuery:  false,
	})

	postBody, _ := json.Marshal(0)

	responseBody := bytes.NewBuffer(postBody)

	_, _ = conn.Post("/loader/unload-garbage-truck", responseBody)

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



