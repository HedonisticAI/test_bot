package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	//"database/sql"
	//"github.com/lib/pq"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	TelegramBotToken string
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("For heroku stuff"))
}

func main() {
	file, _ := os.Open("cfg.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	var x int8 = 0
	// инициализируем канал, куда будут прилетать обновления от API
	//var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	//ucfg.Timeout = 60
	//upd, _ := bot.GetUpdatesChan(ucfg)
	upd := bot.ListenForWebhook("/" + configuration.TelegramBotToken)
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	client := ipinit()
	// читаем обновления из канала
	for x != 1 {
		x = 0
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "use commands please"
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "help":
					msg.Text = "I understand ipinfo and userhistory."
				case "ipinfo":
					msg.Text = parseip(update.Message.CommandArguments(), update.Message.From.ID, client)
				case "userhistory":
					msg.Text = userhistory(update.Message.From.ID)
				default:
					msg.Text = "I don't know that command"
				}
				msg.ReplyToMessageID = update.Message.MessageID
			}
			bot.Send(msg)
		default:
			continue
		}

	}
}
