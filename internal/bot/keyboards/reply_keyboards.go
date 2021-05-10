package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var MainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/driver",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/loader",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/sorter",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/operator",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
)


var SorterKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/sort_conveyor",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/back",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	)

var DriverKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/add_garbage_glass",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/add_garbage_plastic",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/add_garbage_organic",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/take_away",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/back",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
)

var LoaderKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/unload_garbage_truck",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/load_truck",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/load_conv",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/back",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
)

var OperatorKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/change_conveyor_state",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/change_conveyor_speed 5",
			RequestContact:  false,
			RequestLocation: false,
		},
		tgbotapi.KeyboardButton{
			Text:            "/change_conveyor_speed 10",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/change_ventilation_state",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/transform_garbage",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.KeyboardButton{
			Text:            "/back",
			RequestContact:  false,
			RequestLocation: false,
		},
	),
)


