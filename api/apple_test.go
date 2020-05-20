package handler

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestNewBot(t *testing.T) {
	bot := newBot()
	if bot == nil {
		t.Fatalf("bot not success")
	}

	// msg := tgbotapi.NewMessageToChannel(getEnvData("Channel"), "Today Mac 13 123332")
	// _, err := bot.Send(msg)
	// if err != nil {
	// 	t.Fatalf("send message error:%s", err)
	// }
}

func TestGetAppleData(t *testing.T) {
	apples, err := GetAppleData("13寸")
	if err != nil {
		t.Fatalf("get apple data error:%s", err)
	}
	for _, apple := range apples {
		t.Logf("Name: %s, Price:%d, OPrice: %d", apple.Name, apple.Price, apple.OfficialPrice)
	}
}

func TestSendMessage(t *testing.T) {
	bot := newBot()
	if bot == nil {
		t.Fatalf("bot not success")
	}
	msgText, err := genMessage("13寸")
	if err != nil {
		t.Fatalf("get message err:%s", err)
	}
	fmt.Println(msgText)
	msg := tgbotapi.NewMessageToChannel(getEnvData(Channel), msgText)
	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("send message error: %s", err)
	}
}
