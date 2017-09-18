# telegram_bot
 Телеграм-бот доступен (код размещен локально, может не работать) по ссылке [@tigprog_bot](https://t.me/tigprog_bot).
 
 Знает команды:
 1. /start - приветствие;
 2. /info - вывод информации о боте;
 3. /email YOUR_GitHub_TOKEN - отправляет запрос к GitHub API и возвращает email пользователя.
 
 Кроме команд,может повторяет любой текст (начинающийся не со "/") и emoji.
 
 ## Установка:
 
 1. Необходим стабильный релиз [telegram-bot-api(v4)](https://github.com/go-telegram-bot-api/telegram-bot-api) - требуется Go 1.4 и новее:
 ```bash
 go get gopkg.in/telegram-bot-api.v4
 ```
 2. Клонирование [данного репозитория](https://github.com/TigProg/telegram_bot.git)
 ```bash
 git clone https://github.com/TigProg/telegram_bot.git
 ``` 
 3. Также необходимо переименовать token_example.json в token.json и добавить в него полученный от [@BotFather](https://t.me/BotFather) токен.
 
 ## Демонстрация:
 
 ![screen1](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x) 
 ![screen2](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x) 
 ![screen3](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x)
 
 ## Тестирование:
 
 [@tigprog_bot](https://t.me/tigprog_bot) - бот.
 [Здесь](https://github.com/settings/tokens) можно получить персональный токен. Как видно из кода, в консоль исполняемой программы попадают логи, в том числе может попасть ваш токен - пользуйтесь с осторожностью. 
 


 
