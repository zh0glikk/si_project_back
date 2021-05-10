package message_handlers

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"net/url"

	"github.com/si_project_back/internal/bot/keyboards"
	"github.com/si_project_back/pkg/connector"
)

type ChangeConveyorState struct {
	handler *Handler
}

func NewChangeConveyorState(update tgbotapi.Update) *ChangeConveyorState {
	return &ChangeConveyorState{
		handler: NewHandler(update),
	}
}

func (h ChangeConveyorState) Handle() tgbotapi.MessageConfig {
	conn := connector.NewConnector(&url.URL{
		Scheme:      "http",
		Host:        "127.0.0.1:8080",
		ForceQuery:  false,
	})

	postBody, _ := json.Marshal(0)

	responseBody := bytes.NewBuffer(postBody)

	_, err := conn.Post("/operator/change-conveyor-state", responseBody)
	if err != nil {
		logrus.Info("bad request")
	}


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


