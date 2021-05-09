package message_handlers

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"net/url"

	"github.com/si_project_back/internal/bot/keyboards"
	"github.com/si_project_back/internal/garbage"
	"github.com/si_project_back/pkg/connector"
)

type AddGarbage struct {
	handler *Handler

	garbageType garbage.GarbageType
}

func NewAddGarbage(update tgbotapi.Update, garbageType garbage.GarbageType) *AddGarbage {
	return &AddGarbage{
		handler: NewHandler(update),
		garbageType: garbageType,
	}
}

func (h AddGarbage) Handle() tgbotapi.MessageConfig {
	conn := connector.NewConnector(&url.URL{
		Scheme:      "http",
		Host:        "127.0.0.1:8080",
		ForceQuery:  false,
	})

	garb := garbage.Garbage{
		Weight:      10,
		GarbageType: h.garbageType,
	}

	postBody, _ := json.Marshal(&garb)

	responseBody := bytes.NewBuffer(postBody)

	_, err := conn.Post("/driver/add-garbage", responseBody)
	if err != nil {
		logrus.Info("err")
	}

	text := "Driver action panel."

	chatId := h.handler.update.Message.Chat.ID

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
