package handler

import (
	"fmt"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// TelegramToken telegram token
	TelegramToken = "TelegramToken"
	// Channel telegram channel
	Channel = "Channel"
	// ResourceURL fetch data resouce url
	ResourceURL = "http://www.appletuan.com/price"
	// AuthToken auth token by self
	AuthToken = "AuthToken"
)

var (
	headerMap = map[string]string{
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Safari/605.1.15",
	}
)

func getEnvData(key string) string {
	return os.Getenv(key)
}

func newBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(getEnvData(TelegramToken))
	if err != nil {
		fmt.Println(fmt.Errorf("init tg bot error: %s", err))
		return nil
	}
	return bot
}

// GetAppleData get apple data from appletuan
func GetAppleData() error {
	cli := &http.Client{}
	req, err := http.NewRequest("GET", ResourceURL, nil)
	if err != nil {
		return err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// AppleHandler fetch Apple Price from `appletuan`
func AppleHandler(w http.ResponseWriter, r *http.Request) {
	//auth api
	token := r.Header.Get("x-auth-token")
	if token != getEnvData(AuthToken) {
		w.WriteHeader(401)
		fmt.Fprintf(w, "not auth")
		return
	}

	bot := newBot()
	if bot == nil {
		w.WriteHeader(504)
		fmt.Fprintf(w, "bot can not init.")
		return
	}
	msg := tgbotapi.NewMessageToChannel(getEnvData(Channel), "Today not found apple price")
	if _, err := bot.Send(msg); err != nil {
		w.WriteHeader(504)
		fmt.Fprintf(w, "send telegram message error")
		return
	}
}
