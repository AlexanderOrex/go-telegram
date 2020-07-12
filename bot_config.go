package telegram

type BotConfig struct {
	Token   string
	BotName string
	ChatId  int
}

func InitConfig(botName string, token string, chatId int) BotConfig {
	var config BotConfig
	config.Token = token
	config.BotName = botName
	config.ChatId = chatId

	return config
}
