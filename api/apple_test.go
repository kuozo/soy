package handler

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestNewBot(t *testing.T) {
	bot := newBot()
	if bot == nil {
		t.Fatalf("bot not success")
	}

	msg := tgbotapi.NewMessageToChannel(getEnvData("Channel"), "Today Mac 13 123332")
	_, err := bot.Send(msg)
	if err != nil {
		t.Fatalf("send message error:%s", err)
	}
}
