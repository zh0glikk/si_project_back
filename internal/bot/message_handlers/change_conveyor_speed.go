package message_handlers

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net/url"
	"strings"

	"github.com/si_project_back/internal/api/requests"
	"github.com/si_project_back/internal/bot/keyboards"
	"github.com/si_project_back/pkg/connector"
)

type ChangeConveyorSpeed struct {
	handler *Handler
}

func NewChangeConveyorSpeed(update tgbotapi.Update) *ChangeConveyorSpeed {
	return &ChangeConveyorSpeed{
		handler: NewHandler(update),
	}
}

func (h ChangeConveyorSpeed) Handle() tgbotapi.MessageConfig {
	conn := connector.NewConnector(&url.URL{
		Scheme:      "http",
		Host:        "127.0.0.1:8080",
		ForceQuery:  false,
	})

	tmp := strings.Split(h.handler.update.Message.Text, " ")

	req := requests.ChangeConveyorSpeedRequest{
		Speed: cast.ToInt(tmp[1]),
	}

	postBody, _ := json.Marshal(&req)

	responseBody := bytes.NewBuffer(postBody)

	_, err := conn.Post("/operator/change-conveyor-speed", responseBody)
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



