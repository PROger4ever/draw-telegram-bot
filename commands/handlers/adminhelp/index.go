package adminhelppkg

import (
	"github.com/PROger4ever-Golang/draw-telegram-bot/bot"
	"github.com/PROger4ever-Golang/draw-telegram-bot/config"
	"github.com/PROger4ever-Golang/draw-telegram-bot/error"
	"github.com/PROger4ever-Golang/draw-telegram-bot/userApi"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const commandsList = `Список админских команд:
adminHelp - список админских команд
стат - количество зарегистрированных

пользователь @username - проверка регистрации по username
пользователь #12345678 - проверка регистрации по userID

звук - состояние уведомлений
звук 0 - отключить уведомления
звук 1 - включить уведомления

розыграй - провести розыгрыш 1 приза
розыграй N - провести розыгрыш N призов
розыграйОнлайн 15m - провести розыгрыш 1 приза среди писавших боту не более 15 минут
розыграйОнлайн 1h15m30s N - провести розыгрыш N призов среди писавших боту не более 1 часа, 15 минут и 30 секунд`

type Handler struct {
	Bot  *botpkg.Bot
	Conf *config.Config
	Tool *userapi.Tool
}

func (h *Handler) GetAliases() []string {
	return []string{"adminHelp"}
}

func (h *Handler) IsForOwnersOnly() bool {
	return true
}

func (h *Handler) GetParamsMinCount() int {
	return 0
}

func (h *Handler) Init(conf *config.Config, tool *userapi.Tool, bot *botpkg.Bot) {
	h.Bot = bot
	h.Conf = conf
	h.Tool = tool
}

func (h *Handler) Execute(msg *tgbotapi.Message, params []string) *eepkg.ExtendedError {
	return h.Bot.SendMessageMarkdown(int64(msg.Chat.ID), commandsList)
}
