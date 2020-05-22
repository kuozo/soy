package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	prices = map[string]int{
		"i5/2.0G/16G/512G/3733MHz/Bar": 14499,
		"i5/2.0G/16G/1TB/3733MHz/Bar":  15999,
		"i5/1.4G/8G/512G/2133MHz/Bar":  9999,
	}
)

// Apple apple name
type Apple struct {
	Name          string
	Price         int
	OfficialPrice int
}

// Response Data
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
func GetAppleData(keyword string) ([]Apple, error) {
	var name string
	var price string
	var apples = []Apple{}
	cli := &http.Client{}
	req, err := http.NewRequest("GET", ResourceURL, nil)
	if err != nil {
		return apples, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return apples, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return apples, err
	}
	doc.Find(".price-info").Each(
		func(i int, s *goquery.Selection) {
			name = s.Find(".model-name").Text()
			price = s.Find(".price-cell").Find("a").First().Text()
			p, _ := strconv.Atoi(price)
			if strings.Contains(name, keyword) && price != "" {
				apples = append(apples, Apple{
					Name:          name,
					Price:         p,
					OfficialPrice: matchOPrice(name),
				})
			}
		})
	return apples, nil
}

func matchOPrice(name string) int {
	for k, v := range prices {
		if strings.Contains(name, k) {
			return v
		}
	}
	return 0
}

func calDiscount(p, o int) string {
	if o == 0 {
		return "N/A"
	}
	return strconv.Itoa(int(float64(p) / float64(o) * 100))
}

func genMessage(keyword string) (string, error) {

	var b strings.Builder
	b.WriteString("<b>Tody Apple Price</b>\n")
	b.WriteString("------------------\n")
	apples, err := GetAppleData(keyword)
	if err != nil {
		return "", err
	}
	for _, apple := range apples {
		b.WriteString("<strong>")
		b.WriteString(apple.Name)
		b.WriteString("</strong>\n")
		b.WriteString("<i>Price Info: ")
		b.WriteString(strconv.Itoa(apple.Price))
		b.WriteString(" / ")
		b.WriteString(strconv.Itoa(apple.OfficialPrice))
		b.WriteString(" / ")
		b.WriteString(calDiscount(apple.Price, apple.OfficialPrice))
		b.WriteString("%</i>\n")
		b.WriteString("------------------\n")
	}
	b.WriteString(`<pre><code class="language-python">print("work hard")</code></pre>`)
	return b.String(), nil
}

// AppleHandler fetch Apple Price from `appletuan`
func AppleHandler(w http.ResponseWriter, r *http.Request) {
	//auth api
	resp := Response{Code: 200, Message: "success"}
	token := r.Header.Get("x-auth-token")
	if token != getEnvData(AuthToken) {
		w.WriteHeader(401)
		resp.Code = 401
		resp.Message = "not auth"
		json.NewEncoder(w).Encode(resp)
		return
	}

	bot := newBot()
	if bot == nil {
		w.WriteHeader(504)
		resp.Code = 504
		resp.Message = "telegram bot error"
		json.NewEncoder(w).Encode(resp)
		return
	}
	msgText, err := genMessage("13寸")
	if err != nil {
		w.WriteHeader(504)
		resp.Code = 504
		resp.Message = "get apple product price data error"
		json.NewEncoder(w).Encode(resp)
		return
	}
	msg := tgbotapi.NewMessageToChannel(getEnvData(Channel), msgText)
	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		w.WriteHeader(504)
		resp.Code = 504
		resp.Message = "send telegram message error"
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
