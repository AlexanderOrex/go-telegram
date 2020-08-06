package telegram

type BotConfig struct {
	Token   string
	BotName string
	ChatId  int
}

// Init the config for telegram Bot API
// https://core.telegram.org/bots/api
func InitConfig(botName string, token string, chatId int) BotConfig {
	var config BotConfig
	config.Token = token
	config.BotName = botName
	config.ChatId = chatId

	return config
}
