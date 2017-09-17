package main

import (
	//"fmt" //for local debug regexp
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "encoding/json"
//    "regexp"
)

type Config struct {
    TelegramBotToken string
}

//get token from json
func return_bot() string {
    file, _ := os.Open("token.json")
    decoder := json.NewDecoder(file)
    configuration := Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
       log.Panic(err)
    }
    return(configuration.TelegramBotToken)
}

//check for commands
//not used, because in the lib there is a check for commands
/*
func some_command(text string) (answer bool) {
	var command_bot = regexp.MustCompile(`^/[a-zA-Z]+$`)
	//fmt.Println(command_bot.MatchString(text))
	return command_bot.MatchString(text)
}
*/

func main() {
	bot, err := tgbotapi.NewBotAPI(return_bot())
	if err != nil {
		log.Panic(err)
	}

	//if there are errors, it must be activated
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

		//for simple logs to the console
		log.Printf("[%s] %s", UserName, Text)

		if Message.IsCommand() {
			send_answer := tgbotapi.NewMessage(ChatID, "")
			switch update.Message.Command() {
			case "info":
				send_answer.Text = "Я робот - решение ТЗ JetBrains для вступительного испытания. " +
					"Умею повторять за тобой, правда стикеры задержались по дороге."
			case "email":
				send_answer.Text = "Мой создатель работает над этим"
			default:
				send_answer.Text = "Попробуй другую команду, я знаю всего 2"
			}
			bot.Send(send_answer)
			continue
		}

		//check for command
		/*
		if some_command(Text) == true { //
			continue
		}
		*/

		send_message := tgbotapi.NewMessage(ChatID, Text)

		bot.Send(send_message)
	}
}

