# telegram_bot
 Телеграм-бот доступен (код размещен локально, может не работать) по ссылке [@tigprog_bot](https://t.me/tigprog_bot).
 
 Знает команды:
 1. /start - приветствие;
 2. /info - вывод информации о боте;
 3. /email YOUR_GitHub_TOKEN - отправляет запрос с GitHub API и возвращает email пользователя.
 
 Помимо этого повторяет любой текст (начинающийся не со "/") и emoji.
 
 # Установка:
 
 1. Необходим стабильный релиз [telegram-bot-api(v4)](https://github.com/go-telegram-bot-api/telegram-bot-api) - требуется Go 1.4 и новее
 ```bash
 go get gopkg.in/telegram-bot-api.v4
 ```
 2. Клонирование [данного репозитория](https://github.com/go-telegram-bot-api/telegram-bot-api)
 ```bash
 go get gopkg.in/telegram-bot-api.v4
 ``` 
 3. Также необходимо переименовать token_example.json в token.json и добавить в него полученный от [@BotFather](https://t.me/tigprog_bot) токен.
 
 # Демонстрация:
 
 ![](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x) ![](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x) ![](https://github.com/TigProg/telegram_bot/blob/master/another/1.jpg =615x)
 
 # Тестирование:
 
 [@tigprog_bot](https://t.me/tigprog_bot) - бот.
 [https://github.com/settings/tokens](Здесь) можно получить персональный токен. Как видно из кода, в консоль исполняемой программы попадают логи, в том числе может попасть ваш токен - пользуйтесь с осторожностью. 
 


 
