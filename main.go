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
func returnBot() string {
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
func tokenAuthentication(token string) string {
	tokenOld := token
	token = strings.Replace(token, "/email ", "", 1)
	correctToken := "Корректный ввод:" + "\n" + "</email YOUR_TOKEN>"

	// text doesn't contain substring like "/enter_"
	if token == tokenOld {
		return "Прости, но кажется, ты забыл пробел или не ввел токен. " + correctToken
	}
	// text looks like "/enter_" (but message can't ended on <space>) - only for keeping calm
	if token == "" {
		return "Прости, но ты не ввел токен. " + correctToken
	}

	searchUrl := "https://api.github.com/user/emails?access_token=" + token

	resp, err := http.Get(searchUrl)
	// only for keeping calm too
	// in my cases github always return response
	if err != nil {
		return "Прости, но с токеном возникли проблемы. " + correctToken
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// damaged response (?)
	if err != nil {
		return "Странно, может ты ошибся? " + correctToken
	}
	jsonString :=string(body)

	//any questions about regexp? I'm ready to answer them
	var emailReg = regexp.MustCompile(`"[^"@\s]+@[^"@\s]+"`)

	if emailReg.MatchString(jsonString) == false {
		return "В полученном json нет email, попробуй перепроверить. " + correctToken
	}
	email := emailReg.FindString(jsonString)
	return strings.Replace(email, "\"", "", 2)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(returnBot())
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
		//UserName 	:= Message.From.UserName 
		Text 		:= Message.Text
		ChatID 		:= Message.Chat.ID

		// for simple logs to the console - only for debug
		// log.Printf("[%s] %s", UserName, Text)
		log.Printf("another one message")

		if Message.IsCommand() {
			sendAnswer := tgbotapi.NewMessage(ChatID, "")
			switch update.Message.Command() {
			case "start":
				sendAnswer.Text = "Привет! " +
					"Я - бот-попугай. Я буду передразнивать тебя, но пока умею только печатать. " +
					"Еще я знаю команды /start, /info и /email - последняя самая интересная."
			case "info":
				sendAnswer.Text = "Я робот - решение ТЗ JetBrains для вступительного испытания. " +
					"Умею повторять за тобой, но только слова и emoji. " +
				    "Можешь попробовать использовать команду /email и добавить к ней свой GitHub token" +
				    	" - посмотри, что получится!"
			case "email":
				sendAnswer.Text = tokenAuthentication(Text)
			default:
				sendAnswer.Text = "Попробуй другую команду, я знаю всего 3."
			}
			bot.Send(sendAnswer)
			continue
		}
		sendMessage := tgbotapi.NewMessage(ChatID, Text)
		bot.Send(sendMessage)
	}
}
