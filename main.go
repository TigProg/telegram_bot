package main

import (
	"regexp"
    "log"
    "os"
    "encoding/json"
    "strings"
	"net/http"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
)

type Config struct {
    TelegramBotToken string
}

// get token from json
func return_bot() string {
    file, _ := os.Open("token.json")
    decoder := json.NewDecoder(file)
    configuration := Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
       log.Panic(err)
    }
    return configuration.TelegramBotToken
}

// authentication with a user token
// token in func not token - it's "/enter_TOKEN" (I hope) e
func token_authentication(token string) string {
	token_old := token
	token = strings.Replace(token, "/email ", "", 1)
	correct_token := "Корректный ввод:" + "\n" + "<\\email YOUR_TOKEN>"

	// text doesn't contain substring like "/enter_"
	if token == token_old {
		return "Прости, но кажется, ты забыл пробел или не ввел токен. " + correct_token
	}
	// text looks like "/enter_" (but message can't ended on <space>) - only for keeping calm
	if token == "" {
		return "Прости, но ты не ввел токен. " + correct_token
	}

	search_url := "https://api.github.com/user/emails?access_token=" + token

	resp, err := http.Get(search_url)
	// only for keeping calm too
	// in my cases github always return response
	if err != nil {
		return "Прости, но с токеном возникли проблемы. " + correct_token
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// damaged response (?)
	if err != nil {
		return "Странно, может ты ошибся? " + correct_token
	}
	json_string :=string(body)

	//any questions about regexp? I'm ready to answer them
	var email_reg = regexp.MustCompile(`"[^"]+@[^"]+"`)

	if email_reg.MatchString(json_string) == false {
		return "В полученном json нет email, попробуй перепроверить. " + correct_token
	}
	email := email_reg.FindString(json_string)
	return strings.Replace(email, "\"", "", 2)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(return_bot())
	if err != nil {
		log.Panic(err)
	}

	// if there are errors, it must be activated
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
    if err != nil {
        log.Panic(err)
    }

	for update := range updates {
		if update.Message == nil {
			continue
		}
		Message 	:= update.Message
		UserName 	:= Message.From.UserName
		Text 		:= Message.Text
		ChatID 		:= Message.Chat.ID

		// for simple logs to the console
		log.Printf("[%s] %s", UserName, Text)

		if Message.IsCommand() {
			send_answer := tgbotapi.NewMessage(ChatID, "")
			switch update.Message.Command() {
			case "start":
				send_answer.Text = "Привет! " +
					"Я - бот-попугай. Я буду передразнивать тебя, но пока умею только печатать. " +
					"Еще я знаю команды /start, /info и /email - последняя самая интересная."
			case "info":
				send_answer.Text = "Я робот - решение ТЗ JetBrains для вступительного испытания. " +
					"Умею повторять за тобой, но только слова и emoji. " +
				    "Можешь попробовать использовать команду /email и добавить к ней свой GitHub token" +
				    	" - посмотри, что получится!"
			case "email":
				send_answer.Text = token_authentication(Text)
			default:
				send_answer.Text = "Попробуй другую команду, я знаю всего 3."
			}
			bot.Send(send_answer)
			continue
		}
		send_message := tgbotapi.NewMessage(ChatID, Text)
		bot.Send(send_message)
	}
}
